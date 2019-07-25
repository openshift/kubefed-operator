/*
Copyright 2019 The kubefed-operator Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package openshift

import (
	"context"
	"fmt"
	"strings"

	b64 "encoding/base64"

	mf "github.com/jcrossley3/manifestival"
	configv1 "github.com/openshift/api/config/v1"
	kubefedv1alpha1 "github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1"
	"github.com/openshift/kubefed-operator/pkg/controller/common"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

const (
	caBundleConfigMapName = "config-service-ca"
	serviceSecretKey      = "service-ca.crt"
	serviceAccountName    = "system:serviceaccount:kubefed-admission-webhook"
)

var (
	extension = common.Extension{
		Transformers:  []mf.Transformer{annotateWebhookService},
		PreInstalls:   []common.Extender{caBundleConfigMap},
		PostInstalls:  []common.Extender{},
		Transclosures: []common.Transclosure{addCaBundleToWebhookConfiguration},
		PreConditions: []common.PreCondition{ensureCaBundleIsPresent},
	}
	log    = logf.Log.WithName("openshift")
	api    client.Client
	scheme *runtime.Scheme
)

func IsRunningOpenshift(c client.Client) (bool, error) {
	if routeExists, err := anyKindExists(c, "", schema.GroupVersionKind{Group: "route.openshift.io", Version: "v1", Kind: "route"}); err != nil {
		return false, err
	} else if !routeExists {
		// Not running in OpenShift
		return false, nil
	} else {
		return true, nil
	}
}

// Configure OpenShift if we're soaking in it
func Configure(c client.Client, s *runtime.Scheme) (*common.Extension, error) {
	if routeExists, err := anyKindExists(c, "", schema.GroupVersionKind{Group: "route.openshift.io", Version: "v1", Kind: "route"}); err != nil {
		return nil, err
	} else if !routeExists {
		// Not running in OpenShift
		return nil, nil
	}

	// Register scheme
	if err := configv1.Install(s); err != nil {
		log.Error(err, "Unable to register scheme")
		return nil, err
	}

	api = c
	scheme = s
	return &extension, nil
}

func annotateWebhookService(u *unstructured.Unstructured) error {
	if u.GetKind() == "Service" && u.GetName() == "kubefed-admission-webhook" {
		log.Info("Trying to annotate webhook service")
		annotations := u.GetAnnotations()
		if annotations == nil {
			annotations = make(map[string]string)
		}
		annotations["service.beta.openshift.io/serving-cert-secret-name"] = "kubefed-admission-webhook-serving-cert"
		u.SetAnnotations(annotations)
		log.Info("Added webhook service annotation")
	}
	return nil
}

func caBundleConfigMap(instance *kubefedv1alpha1.KubeFedWebHook) error {
	cm := &v1.ConfigMap{}
	log.Info("Creating ca config map")
	if err := api.Get(context.TODO(), types.NamespacedName{Name: caBundleConfigMapName, Namespace: instance.GetNamespace()}, cm); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("Populating fields of ca config map")
			// Define a new configmap
			cm.Name = caBundleConfigMapName
			cm.Annotations = make(map[string]string)
			cm.Annotations["service.alpha.openshift.io/inject-cabundle"] = "true"
			cm.Namespace = instance.GetNamespace()
			cm.SetOwnerReferences([]metav1.OwnerReference{*metav1.NewControllerRef(instance, instance.GroupVersionKind())})
			err2 := api.Create(context.TODO(), cm)
			if err2 != nil {
				log.Info(fmt.Sprintf("error2 on api create: %v", err2))
				return err2
			}
			// ConfigMap created successfully
			log.Info(fmt.Sprintf("Created the config map %v", caBundleConfigMapName))
			return nil
		}
		return err
	} else {
		log.Info("Config map already exists")
	}

	return nil
}

// anyKindExists returns true if any of the gvks (GroupVersionKind) exist
func anyKindExists(c client.Client, namespace string, gvks ...schema.GroupVersionKind) (bool, error) {
	for _, gvk := range gvks {
		list := &unstructured.UnstructuredList{}
		list.SetGroupVersionKind(gvk)
		if err := c.List(context.TODO(), &client.ListOptions{Namespace: namespace}, list); err != nil {
			if !meta.IsNoMatchError(err) {
				return false, err
			}
		} else {
			log.Info("Detected", "gvk", gvk.String())
			return true, nil
		}
	}
	return false, nil
}

func itemsExist(c client.Client, kind string, apiVersion string, namespace string) (bool, error) {
	list := &unstructured.UnstructuredList{}
	list.SetKind(kind)
	list.SetAPIVersion(apiVersion)
	if err := c.List(context.TODO(), &client.ListOptions{Namespace: namespace}, list); err != nil {
		return false, err
	}
	return len(list.Items) > 0, nil
}

func ensureCaBundleIsPresent(instance *kubefedv1alpha1.KubeFedWebHook) error {
	cm := &v1.ConfigMap{}
	if err := api.Get(context.TODO(), types.NamespacedName{Name: caBundleConfigMapName, Namespace: instance.GetNamespace()}, cm); err == nil {
		data := cm.Data
		if _, ok := data[serviceSecretKey]; ok {
			return nil
		} else {
			return apierrors.NewNotFound(schema.GroupResource{Group: "ConfigMap", Resource: "Key"}, serviceSecretKey)
		}
	} else {
		return err
	}
}

func getCaBundleFromConfigMap(instance *kubefedv1alpha1.KubeFedWebHook) (error, string) {
	cm := &v1.ConfigMap{}
	if err := api.Get(context.TODO(), types.NamespacedName{Name: caBundleConfigMapName, Namespace: instance.GetNamespace()}, cm); err != nil {
		log.Info(fmt.Sprintf("Couldn't find the configmap %v", caBundleConfigMapName))
		return err, ""
	}
	data := cm.Data
	if v, ok := data[serviceSecretKey]; ok {
		log.Info(fmt.Sprintf("Found the key %v and value %v in configmap %v", serviceSecretKey, v, caBundleConfigMapName))
		return nil, v
	} else {
		log.Info(fmt.Sprintf("Could not find the key %v in configmap %v", serviceSecretKey, caBundleConfigMapName))
		return apierrors.NewNotFound(schema.GroupResource{Group: "ConfigMap", Resource: "Key"}, serviceSecretKey), ""
	}
}

func addCaBundleToWebhookConfiguration(instance *kubefedv1alpha1.KubeFedWebHook) mf.Transformer {
	return func(u *unstructured.Unstructured) error {
		switch strings.ToLower(u.GetKind()) {
		case "validatingwebhookconfiguration":
			fallthrough
		case "mutatingwebhookconfiguration":
			log.Info("Trying to get to the caBundle for validatingwebhook configuration transformer")
			if err, caBundle := getCaBundleFromConfigMap(instance); err != nil {
				log.Info("Unable to get to the caBundle, but that's ok...")
				return nil
			} else {
				log.Info("Trying to set the caBundle parameter on validating webhook")
				caEnc := b64.StdEncoding.EncodeToString([]byte(caBundle))
				if webhooks, ok, webhooksErr := unstructured.NestedSlice(u.Object, "webhooks"); ok {
					for _, webhook := range webhooks {
						err := unstructured.SetNestedField(webhook.(map[string]interface{}), caEnc, "clientConfig", "caBundle")
						if err != nil {
							return err
						}
					}
					err := unstructured.SetNestedSlice(u.Object, webhooks, "webhooks")
					if err != nil {
						return err
					}
				} else {
					return webhooksErr
				}
				log.Info("Set the caBundle successfully on the webhooks")
			}
		}
		return nil
	}
}
