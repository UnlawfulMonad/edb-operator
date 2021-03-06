package mysqlgrant

import (
	"context"
	"strings"
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

var log = logf.Log.WithName("controller_mysqlgrant")

// Add creates a new MySQLGrant Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMySQLGrant{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("mysqlgrant-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource MySQLGrant
	err = c.Watch(&source.Kind{Type: &apiv1alpha1.MySQLGrant{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &apiv1alpha1.MySQLGrant{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileMySQLGrant implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileMySQLGrant{}

// ReconcileMySQLGrant reconciles a MySQLGrant object
type ReconcileMySQLGrant struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a MySQLGrant object and makes changes based on the state read
// and what is in the MySQLGrant.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMySQLGrant) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling MySQLGrant")

	// Fetch the MySQLGrant instance
	instance := &apiv1alpha1.MySQLGrant{}
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

	permission := instance.Spec.Permission
	if !isValidGrant(permission) {
		reqLogger.Error(edb.ErrUnsupportedPermission, "provided permission was invalid")
		instance.Status.Granted = false
		instance.Status.Error = "provided permission is invalid"
		r.client.Status().Update(context.TODO(), instance)
		return reconcile.Result{}, edb.ErrUnsupportedPermission
	}

	// TODO
	// Ensure the database and user are created (and in the same namespace)

	database := edb.LookupExternalDatabase(instance.Spec.ExternalDatabaseRef.Name)
	if database == nil {
		return reconcile.Result{RequeueAfter: time.Second * 5}, nil
	}

	err = database.Grant(instance.Spec.Permission, instance.Spec.Database, instance.Spec.User)
	if err != nil {
		return reconcile.Result{}, err
	}

	instance.Status.Granted = true
	instance.Status.Error = ""
	r.client.Status().Update(context.TODO(), instance)

	return reconcile.Result{}, nil
}

var validGrants = []string{"all", "select", "insert", "update", "delete"}

func isValidGrant(permission string) bool {
	permission = strings.ToLower(permission)
	for _, grant := range validGrants {
		if permission == grant {
			return true
		}
	}

	return false
}
