package kubefedwebhook

import (
	"context"
	"flag"
	"fmt"
	"strings"

	mf "github.com/jcrossley3/manifestival"
	kubefedv1alpha1 "github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1"
	"github.com/openshift/kubefed-operator/pkg/controller/common"
	"github.com/openshift/kubefed-operator/version"
	"github.com/operator-framework/operator-sdk/pkg/predicate"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var (
	filename = flag.String("webhookfilename", "deploy/resources/webhook",
		"The filename containing the YAML resources to apply")
	recursive = flag.Bool("recursivewebhook", false,
		"If filename is a directory, process all manifests recursively")
	namespace = flag.String("webhooknamespace", "",
		"Overrides namespace in manifest (env vars resolved in-container)")
	log       = logf.Log.WithName("controller_kubefedwebhook")
	platforms common.Platforms
)

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new KubeFedWebHook Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	manifest, err := mf.NewManifest(*filename, *recursive, mgr.GetClient())
	if err != nil {
		return err
	}
	return add(mgr, newReconciler(mgr, manifest))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, man mf.Manifest) reconcile.Reconciler {
	return &ReconcileKubeFedWebHook{client: mgr.GetClient(), scheme: mgr.GetScheme(),
		config: man}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("kubefedwebhook-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource KubeFedWebHook
	err = c.Watch(&source.Kind{Type: &kubefedv1alpha1.KubeFedWebHook{}}, &handler.EnqueueRequestForObject{}, predicate.GenerationChangedPredicate{})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileKubeFedWebHook implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileKubeFedWebHook{}

// ReconcileKubeFedWebHook reconciles a KubeFedWebHook object
type ReconcileKubeFedWebHook struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
	config mf.Manifest
}

// Reconcile reads that state of the cluster for a KubeFedWebHook object and makes changes based on the state read
// and what is in the KubeFedWebHook.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileKubeFedWebHook) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling KubeFedWebHook")

	// Fetch the KubeFedWebHook instance
	instance := &kubefedv1alpha1.KubeFedWebHook{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			r.config.DeleteAll()
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	stages := []func(*kubefedv1alpha1.KubeFedWebHook) error{
		r.install,
		r.configure,
	}

	for _, stage := range stages {
		if err := stage(instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	reqLogger.Info("Finished reconciling kubefedwebhook")

	return reconcile.Result{}, nil
}

// This is a transform method that updates the namespace field of the clusterrolebinding resource
// for cluster scoped deployment
func resourceNamespaceUpdate(ns, name string, scheme *runtime.Scheme) mf.Transformer {
	return func(u *unstructured.Unstructured) error {
		reqLogger := log.WithValues("Instance.Namespace", ns, "Instance.Name", name)
		switch strings.ToLower(u.GetKind()) {
		case "clusterrolebinding":
			fallthrough
		case "rolebinding":
			if subjects, ok, _ := unstructured.NestedSlice(u.Object, "subjects"); ok {
				if kind, ok, err := unstructured.NestedString(subjects[0].(map[string]interface{}), "kind"); ok && kind == "ServiceAccount" {
					err = unstructured.SetNestedField(subjects[0].(map[string]interface{}), ns, "namespace")
					if err != nil {
						reqLogger.Info("Failed to set the namespace nested field")
					} else {
						reqLogger.Info(fmt.Sprintf("Added the namespace to the rolebinding %v subjects element", u.GetName()))
						err = unstructured.SetNestedSlice(u.Object, subjects, "subjects")
						if err != nil {
							reqLogger.Info("Failed to update the subjects slice")
						}
					}
				} else {
					reqLogger.Info("Either not able to get the kind field of the element or it is of no interest")
				}
			} else {
				reqLogger.Info("Failed to get subjects slice")
			}
			if u.GetName() == "kubefed-admission-webhook:apiextension-viewer" {
				reqLogger.Info("Setting the namespace for the api extension-viewer role binding")
				u.SetNamespace("kube-system")

				u.SetOwnerReferences(nil)

				roleBinding := &rbacv1.RoleBinding{}
				if err := scheme.Convert(u, roleBinding, nil); err != nil {
					return err
				} else {
					reqLogger.Info(fmt.Sprintf("Dumping rolebinding yaml: %+v", roleBinding))
				}
			}

		case "validatingwebhookconfiguration":
			fallthrough
		case "mutatingwebhookconfiguration":
			if webhooks, ok, _ := unstructured.NestedSlice(u.Object, "webhooks"); ok {
				for _, webhook := range webhooks {

					err := unstructured.SetNestedField(webhook.(map[string]interface{}), ns, "clientConfig", "service", "namespace")
					webhookname, _, _ := unstructured.NestedString(webhook.(map[string]interface{}), "clientConfig", "service", "name")
					if err != nil {
						reqLogger.Info("Failed to set the namespace nested field")
					} else {
						reqLogger.Info(fmt.Sprintf("Added the namespace to the webhook: %s", webhookname))
					}
				}
				err := unstructured.SetNestedSlice(u.Object, webhooks, "webhooks")
				if err != nil {
					reqLogger.Info("Failed to update the subjects slice")
				}
			} else {
				reqLogger.Info("Failed to get webhooks slice")
			}
		}
		return nil
	}

}

func (r *ReconcileKubeFedWebHook) updateStatus(instance *kubefedv1alpha1.KubeFedWebHook) error {

	// Account for https://github.com/kubernetes-sigs/controller-runtime/issues/406
	gvk := instance.GroupVersionKind()
	defer instance.SetGroupVersionKind(gvk)

	if err := r.client.Status().Update(context.TODO(), instance); err != nil {
		return err
	}
	return nil
}

// Apply the embedded resources
func (r *ReconcileKubeFedWebHook) install(instance *kubefedv1alpha1.KubeFedWebHook) error {
	defer r.updateStatus(instance)
	// Transform resources as appropriate

	extensions, err := platforms.Extend(r.client, r.scheme)
	if err != nil {
		return err
	}
	fns := extensions.Transform(instance)
	fns = append(fns, resourceNamespaceUpdate(instance.Namespace, instance.Name, r.scheme))
	fns = append(fns, common.ResourceImageReplace(instance.Namespace, instance.Name))

	preConditionsErr := extensions.PreConditions(instance)

	err = r.config.Transform(fns...)
	if err == nil {
		err = extensions.PreInstall(instance)
		if err == nil {
			err = r.config.ApplyAll()
			if err == nil {
				err = extensions.PostInstall(instance)
			}
		}
	}

	if preConditionsErr != nil {
		return preConditionsErr
	}

	if err != nil {
		return err
	}

	// Update status
	instance.Status.Version = version.Version
	if err := r.client.Status().Update(context.TODO(), instance); err != nil {
		return err
	}
	log.Info("Install succeeded", "version", version.Version)
	return nil
}

// Set ConfigMap values from KubeFedWebHook spec
func (r *ReconcileKubeFedWebHook) configure(instance *kubefedv1alpha1.KubeFedWebHook) error {
	return nil
}
