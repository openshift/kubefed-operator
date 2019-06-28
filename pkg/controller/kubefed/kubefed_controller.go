package kubefed

import (
	"context"
	"flag"
	"strings"

	mf "github.com/jcrossley3/manifestival"
	kubefedv1alpha1 "github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1"
	"github.com/openshift/kubefed-operator/version"
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
	filename = flag.String("filename", "deploy/resources",
		"The filename containing the YAML resources to apply")
	recursive = flag.Bool("recursive", false,
		"If filename is a directory, process all manifests recursively")
	namespace = flag.String("namespace", "",
		"Overrides namespace in manifest (env vars resolved in-container)")
	log = logf.Log.WithName("controller_kubefed")
)

// Add creates a new KubeFed Controller and adds it to the Manager. The Manager will set fields on the Controller
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
	return &ReconcileKubeFed{client: mgr.GetClient(), scheme: mgr.GetScheme(), config: man}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("kubefed-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource KubeFed
	err = c.Watch(&source.Kind{Type: &kubefedv1alpha1.KubeFed{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileKubeFed{}

// ReconcileKubeFed reconciles a KubeFed object
type ReconcileKubeFed struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
	config mf.Manifest
}

// Reconcile reads that state of the cluster for a KubeFed object and makes changes based on the state read
// and what is in the KubeFed.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileKubeFed) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling KubeFed")

	// Fetch the KubeFed instance
	instance := &kubefedv1alpha1.KubeFed{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			r.config.DeleteAll()
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	stages := []func(*kubefedv1alpha1.KubeFed) error{
		r.install,
		r.configure,
	}

	for _, stage := range stages {
		if err := stage(instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	reqLogger.Info("Finished reconciling kubefed")

	return reconcile.Result{}, nil
}

// This is a transform method that ignores clusterrole and clusterrolebinding
// resources for namespace scoped deployment of kubefed-operator
func resourceScopeFilter(resources []unstructured.Unstructured, scope kubefedv1alpha1.InstallationScope) []unstructured.Unstructured {
	if scope != kubefedv1alpha1.InstallationScopeNamespaceScoped {
		return resources
	}

	filtered := []unstructured.Unstructured{}

	for i := range resources {
		switch strings.ToLower(resources[i].GetKind()) {
		case "clusterrole":
			fallthrough
		case "clusterrolebinding":
			continue
		default:
			filtered = append(filtered, resources[i])
		}
	}

	return filtered
}

// This is a transform method that updates the deployment resource's environment variables
// by adding the federation scope env. variable for namespace scoped deployments
func resourceEnvUpdate(scope kubefedv1alpha1.InstallationScope, ns, name string) mf.Transformer {
	return func(u *unstructured.Unstructured) error {
		reqLogger := log.WithValues("Instance.Namespace", ns, "Instance.Name", name)
		kubefedScope := "Namespaced"
		if scope == kubefedv1alpha1.InstallationScopeClusterScoped {
			kubefedScope = "Cluster"

		}

		switch strings.ToLower(u.GetKind()) {
		case "deployment":
			if containers, ok, err := unstructured.NestedSlice(u.Object,
				"spec", "template", "spec", "containers"); ok {
				if envs, envOk, envErr := unstructured.NestedSlice(containers[0].(map[string]interface{}), "env"); envOk {
					if !checkEnvExists(envs, "name", "DEFAULT_KUBEFED_SCOPE") {
						fse := map[string]interface{}{"name": "DEFAULT_KUBEFED_SCOPE", "value": kubefedScope}
						envs = append(envs, fse)
					}
					reqLogger.Info("Transforming deployment resource for environment update - env; ", "envs", envs)
					envErr = unstructured.SetNestedSlice(containers[0].(map[string]interface{}), envs, "env")
					if envErr != nil {
						reqLogger.Info("Failed to update the environment")
					}
					err = unstructured.SetNestedSlice(u.Object, containers, "spec", "template", "spec", "containers")
					if err != nil {
						reqLogger.Info("Failed to update the container array")
					}
				} else {
					reqLogger.Info("Failed to get the env array")
				}
			} else {
				fooLogger := log.WithValues("namespace", ns, "name", name, "scope", scope, "error", err)
				fooLogger.Info("Cannot get the nested slice for env")
			}
		}
		return nil
	}
}

// This function checks if the fedearation scope environment variable exists in the env array
func checkEnvExists(envs []interface{}, envKey, envName string) bool {
	for _, envInterface := range envs {
		env := envInterface.(map[string]interface{})
		if val, ok := env[envKey]; ok {
			if val == envName {
				return true
			}
		}
	}
	return false
}

// This is a transform method that updates the namespace field of the clusterrolebinding resource
// for cluster scoped deployment
func resourceNamespaceUpdate(scope kubefedv1alpha1.InstallationScope, ns, name string) mf.Transformer {
	return func(u *unstructured.Unstructured) error {
		reqLogger := log.WithValues("Instance.Namespace", ns, "Instance.Name", name)
		if scope == kubefedv1alpha1.InstallationScopeClusterScoped {
			switch strings.ToLower(u.GetKind()) {
			case "clusterrolebinding":
				if subjects, ok, err := unstructured.NestedSlice(u.Object, "subjects"); ok {
					err = unstructured.SetNestedField(subjects[0].(map[string]interface{}), ns, "namespace")
					if err != nil {
						reqLogger.Info("Failed to set the namespace nested field")
					} else {
						reqLogger.Info("Added the namespace to the clusterrolebinding subjects element")
						err = unstructured.SetNestedSlice(u.Object, subjects, "subjects")
						if err != nil {
							reqLogger.Info("Failed to update the subjects slice")
						}
					}
				} else {
					reqLogger.Info("Failed to get subjects slice")
				}
			}
		}
		return nil
	}

}

// Apply the embedded resources
func (r *ReconcileKubeFed) install(instance *kubefedv1alpha1.KubeFed) error {
	// filter out resources we don't want
	filteredResources := resourceScopeFilter(r.config.Resources, instance.Spec.Scope)
	r.config.Resources = filteredResources
	// Transform resources as appropriate
	fns := []mf.Transformer{mf.InjectOwner(instance)}
	fns = append(fns, mf.InjectNamespace(instance.Namespace))
	fns = append(fns, resourceEnvUpdate(instance.Spec.Scope, instance.Namespace, instance.Name))
	fns = append(fns, resourceNamespaceUpdate(instance.Spec.Scope, instance.Namespace, instance.Name))
	r.config.Transform(fns...)

	// Apply the resources in the YAML file
	if err := r.config.ApplyAll(); err != nil {
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

// Set ConfigMap values from KubeFed spec
func (r *ReconcileKubeFed) configure(instance *kubefedv1alpha1.KubeFed) error {
	return nil
}
