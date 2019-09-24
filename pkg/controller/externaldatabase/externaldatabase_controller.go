package externaldatabase

import (
	"context"

	"os"

	apiv1alpha1 "github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1"
	"github.com/UnlawfulMonad/edb-operator/pkg/edb"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_externaldatabase")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ExternalDatabase Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileExternalDatabase{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("externaldatabase-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ExternalDatabase
	err = c.Watch(&source.Kind{Type: &apiv1alpha1.ExternalDatabase{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &apiv1alpha1.ExternalDatabase{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileExternalDatabase implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileExternalDatabase{}

// ReconcileExternalDatabase reconciles a ExternalDatabase object
type ReconcileExternalDatabase struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a ExternalDatabase object and makes changes based on the state read
// and what is in the ExternalDatabase.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileExternalDatabase) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ExternalDatabase")

	// Fetch the ExternalDatabase instance
	instance := &apiv1alpha1.ExternalDatabase{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			edb.RemoveExternalDatabase(request.Name)
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Try connect to DB and add to in-memory list
	switch instance.Spec.Type {
	case apiv1alpha1.MySQL:
		err = r.connectMySQL(instance)
		if err != nil {
			reqLogger.Error(err, "failed to connect to mysql")
			return reconcile.Result{}, err
		}
	case apiv1alpha1.PostgreSQL:
		panic("unimplemented")
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileExternalDatabase) getPasswordRef(db *apiv1alpha1.ExternalDatabase) (*corev1.Secret, error) {
	spec := db.Spec

	secretNamespacedName := types.NamespacedName{
		Name:      spec.AdminPasswordRef.Name,
		Namespace: os.Getenv("POD_NAMESPACE"),
	}

	passwordSecret := &corev1.Secret{}
	err := r.client.Get(context.TODO(), secretNamespacedName, passwordSecret)
	if err != nil {
		log.Error(err, "Secret.Name", secretNamespacedName.Name, "Secret.Namespace", secretNamespacedName.Namespace)
		return nil, err
	}

	return passwordSecret, nil
}

func (r *ReconcileExternalDatabase) connectMySQL(db *apiv1alpha1.ExternalDatabase) error {
	spec := db.Spec
	passwordSecret, err := r.getPasswordRef(db)
	if err != nil {
		db.Status.Reachable = false
		db.Status.Error = err.Error()
		r.client.Update(context.TODO(), db)
		log.Info("Failed to get password secret")
		return err
	}

	password := string(passwordSecret.Data[spec.AdminPasswordRef.Key])
	mysql, err := edb.NewMySQL(spec.AdminUser, password, spec.Host)
	if err != nil {
		return err
	}

	edb.AddOrUpdateExternalDatabase(db.Name, mysql)
	return nil
}
