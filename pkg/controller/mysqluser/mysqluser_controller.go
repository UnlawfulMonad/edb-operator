package mysqluser

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

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

var log = logf.Log.WithName("controller_mysqluser")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new MySqlUser Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMySQLUser{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("mysqluser-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource MySqlUser
	err = c.Watch(&source.Kind{Type: &apiv1alpha1.MySQLUser{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner MySqlUser
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &apiv1alpha1.MySQLUser{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileMySqlUser implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileMySQLUser{}

// ReconcileMySQLUser reconciles a MySqlUser object
type ReconcileMySQLUser struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a MySqlUser object and makes changes based on the state read
// and what is in the MySqlUser.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMySQLUser) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling MySqlUser")

	// Fetch the MySqlUser instance
	instance := &apiv1alpha1.MySQLUser{}
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

	spec := instance.Spec

	db := edb.LookupExternalDatabase(spec.ExternalDatabaseRef.Name)
	if db == nil {
		return reconcile.Result{RequeueAfter: 5 * time.Second}, nil
	}

	secret, err := r.getSecretForUser(instance)
	if err != nil {

		// If the secret doesn't exist then we create it.
		if errors.IsNotFound(err) {
			reqLogger.Info("Creating new secret")
			secret = newSecretForMySQLUser(instance)
			err := r.client.Create(context.TODO(), secret)
			if err != nil {
				reqLogger.Error(err, "failed to create secret")
				return reconcile.Result{}, err
			}
		} else {
			return reconcile.Result{}, err
		}
	}

	err = db.CreateUser(instance.Name, string(secret.Data["mysql-password"]))
	if err != nil {
		return reconcile.Result{}, err
	}

	instance.Status.Created = true
	err = r.client.Status().Update(context.TODO(), instance)
	return reconcile.Result{}, err
}

// getSecretForUser gets the secret if it already exists in the given namespace
func (r *ReconcileMySQLUser) getSecretForUser(instance *apiv1alpha1.MySQLUser) (*corev1.Secret, error) {
	nn := types.NamespacedName{Name: instance.Spec.PasswordSecretName, Namespace: instance.Namespace}

	secret := &corev1.Secret{}
	err := r.client.Get(context.TODO(), nn, secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

var (
	defaultSecretAnnotations = map[string]string{
		"api.edb-operator.com/generated": "true",
	}
)

func newSecretForMySQLUser(instance *apiv1alpha1.MySQLUser) *corev1.Secret {
	password := edb.GenPassword()

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:        instance.Spec.PasswordSecretName,
			Namespace:   instance.Namespace,
			Annotations: defaultSecretAnnotations,
		},

		Data: map[string][]byte{
			"mysql-password": []byte(password),
		},
	}

	return secret
}
