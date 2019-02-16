package mysqldatabase

import (
	"context"
	"github.com/UnlawfulMonad/edb-operator/pkg/edb"
	"k8s.io/apimachinery/pkg/types"
	"time"

	apiv1alpha1 "github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
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

var log = logf.Log.WithName("controller_mysqldatabase")

// Add creates a new MySQLDatabase Controller and adds it to the Manager. The Manager will set fields on the Controller
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

	// Watch for changes to primary resource MySQLDatabase
	err = c.Watch(&source.Kind{Type: &apiv1alpha1.MySQLDatabase{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner MySQLDatabase
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &apiv1alpha1.MySQLDatabase{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileMySQLDatabase{}

// ReconcileMySQLDatabase reconciles a MySQLDatabase object
type ReconcileMySQLDatabase struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a MySQLDatabase object and makes changes based on the state read
// and what is in the MySQLDatabase.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMySQLDatabase) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling MySQLDatabase")

	// Fetch the MySQLDatabase instance
	db := &apiv1alpha1.MySQLDatabase{}
	err := r.client.Get(context.TODO(), request.NamespacedName, db)
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


	dbi := edb.LookupExternalDatabase(db.Spec.ExternalDatabaseRef.Name)
	if dbi == nil {
		return reconcile.Result{RequeueAfter: time.Second * 30}, errors.NewBadRequest("external database specified doesn't exist")
	}

	ext := &apiv1alpha1.ExternalDatabase{}
	err = r.client.Get(context.TODO(), types.NamespacedName{ Name: db.Spec.ExternalDatabaseRef.Name }, ext)
	if err != nil {
		return reconcile.Result{}, err
	}

	valid, err := edb.CanUseDB(r.client, db.Namespace, ext)
	if err != nil {
		return reconcile.Result{RequeueAfter: time.Second * 30}, err
	}

	if valid {
		err = dbi.CreateDB(db.Spec.Name, db.Spec.Owner)
		if err != nil {
			return reconcile.Result{RequeueAfter: time.Second * 30}, err
		}
	}

	return reconcile.Result{}, nil
}
