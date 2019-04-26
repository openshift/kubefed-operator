package install

import (
	"context"
	"flag"

	mf "github.com/jcrossley3/manifestival"
	configv1 "github.com/openshift/api/config/v1"
	servingv1alpha1 "github.com/pmorie/kubefed-operator/pkg/apis/operator/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
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
	autoinstall = flag.Bool("install", false,
		"Automatically creates an Install resource if none exist")
	namespace = flag.String("namespace", "",
		"Overrides namespace in manifest (env vars resolved in-container)")
	log = logf.Log.WithName("controller_install")
)

// Add creates a new Install Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	manifest, err := mf.NewYamlManifest(*filename, *recursive, mgr.GetClient())
	if err != nil {
		return err
	}
	return add(mgr, newReconciler(mgr, manifest))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, man mf.Manifest) reconcile.Reconciler {
	return &ReconcileInstall{client: mgr.GetClient(), scheme: mgr.GetScheme(), config: man}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Register scheme
	if err := configv1.Install(mgr.GetScheme()); err != nil {
		log.Error(err, "Unable to register scheme")
	}

	// Create a new controller
	c, err := controller.New("install-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Install
	err = c.Watch(&source.Kind{Type: &servingv1alpha1.Install{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileInstall{}

// ReconcileInstall reconciles a Install object
type ReconcileInstall struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
	config mf.Manifest
}

// Reconcile reads that state of the cluster for a Install object and makes changes based on the state read
// and what is in the Install.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileInstall) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Install")

	// Fetch the Install instance
	instance := &servingv1alpha1.Install{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			r.config.DeleteAll()
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	stages := []func(*servingv1alpha1.Install) error{
		r.install,
		r.configure,
	}

	for _, stage := range stages {
		if err := stage(instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

// Apply the embedded resources
func (r *ReconcileInstall) install(instance *servingv1alpha1.Install) error {
	// Filter resources as appropriate
	filters := []mf.FilterFn{mf.ByOwner(instance)}
	if len(*namespace) > 0 {
		filters = append(filters, mf.ByNamespace(*namespace))
	}
	r.config.Filter(filters...)

	// if instance.Status.Version == version.Version {
	// 	// we've already successfully applied our YAML
	// 	return nil
	// }
	// // Apply the resources in the YAML file
	// if err := r.config.ApplyAll(); err != nil {
	// 	return err
	// }

	// // Update status
	// instance.Status.Resources = r.config.ResourceNames()
	// instance.Status.Version = version.Version
	// if err := r.client.Status().Update(context.TODO(), instance); err != nil {
	// 	return err
	// }
	return nil
}

// Set ConfigMap values from Install spec
func (r *ReconcileInstall) configure(instance *servingv1alpha1.Install) error {
	return nil
}
