package mysqldatabase

import (
	"context"
	"time"

	apiv1alpha1 "github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1"
	"github.com/UnlawfulMonad/edb-operator/pkg/edb"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_mysqldatabase")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new MySqlDatabase Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMySQLDatabase{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("mysqldatabase-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource MySqlDatabase
	err = c.Watch(&source.Kind{Type: &apiv1alpha1.MySQLDatabase{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner MySqlDatabase
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &apiv1alpha1.MySQLDatabase{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileMySqlDatabase implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileMySQLDatabase{}

// ReconcileMySQLDatabase reconciles a MySqlDatabase object
type ReconcileMySQLDatabase struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a MySqlDatabase object and makes changes based on the state read
// and what is in the MySqlDatabase.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMySQLDatabase) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling MySQLDatabase")
	defer reqLogger.Info("Finished reconciling MySQLDatabase")

	// Fetch the MySqlDatabase instance
	instance := &apiv1alpha1.MySQLDatabase{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	database := edb.LookupExternalDatabase(instance.Spec.ExternalDatabaseRef.Name)
	if database == nil {
		instance.Status.Error = "unable to find ExternalDatabase instance"
		//err = r.client.Update(context.TODO(), instance)
		err = r.client.Status().Update(context.TODO(), instance)
		return reconcile.Result{RequeueAfter: 10 * time.Second}, err
	}

	err = database.CreateDB(instance.Spec.Name)
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}
