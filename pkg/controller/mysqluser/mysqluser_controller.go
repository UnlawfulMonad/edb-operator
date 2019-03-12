package mysqluser

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/UnlawfulMonad/edb-operator/pkg/edb"
	"k8s.io/apimachinery/pkg/types"

	apiv1alpha1 "github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_mysqluser")

// Add creates a new MySQLUser Controller and adds it to the Manager. The Manager will set fields on the Controller
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

	// Watch for changes to primary resource MySQLUser
	err = c.Watch(&source.Kind{Type: &apiv1alpha1.MySQLUser{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner MySQLUser
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &apiv1alpha1.MySQLUser{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileMySQLUser{}

// ReconcileMySQLUser reconciles a MySQLUser object
type ReconcileMySQLUser struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a MySQLUser object and makes changes based on the state read
// and what is in the MySQLUser.Spec
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMySQLUser) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling MySQLUser")

	// Fetch the MySQLUser instance
	user := &apiv1alpha1.MySQLUser{}
	err := r.client.Get(context.TODO(), request.NamespacedName, user)
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

	db := edb.LookupExternalDatabase(user.Spec.ExternalDatabaseRef.Name)
	if db == nil {
		return reconcile.Result{RequeueAfter: time.Second * 30}, errors.NewBadRequest("external database specified doesn't exist")
	}

	ext := &apiv1alpha1.ExternalDatabase{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: user.Spec.ExternalDatabaseRef.Name}, ext)
	if err != nil {
		return reconcile.Result{}, err
	}

	validNs, err := edb.CanUseDB(r.client, user.Namespace, ext)
	if err != nil {
		return reconcile.Result{}, err
	}

	if !validNs {
		reqLogger.Info("Could not find valid ExternalDatabase")
		return reconcile.Result{Requeue: true, RequeueAfter: time.Second * 15}, nil
	}

	// Generate the password if it doesn't already exist
	if user.Spec.Password == nil && user.Status.PasswordSecretName == nil {
		err = r.generatePassword(user)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if user.Spec.Password != nil {
		user.Status.PasswordSecretName = user.Spec.Password.DeepCopy()
		if err := r.client.Update(context.TODO(), user); err != nil {
			return reconcile.Result{}, err
		}
	}

	secret := &corev1.Secret{}
	if err := r.client.Get(context.TODO(),
		types.NamespacedName{Name: user.Status.PasswordSecretName.Name, Namespace: user.Namespace},
		secret); err != nil {
		return reconcile.Result{}, err
	}

	password := string(secret.Data[user.Status.PasswordSecretName.Key])
	if err := db.CreateUser(user.Spec.Name, password); err != nil {
		return reconcile.Result{}, err
	}

	reqLogger.Info("Done reconciling user")
	return reconcile.Result{}, nil
}

func (r *ReconcileMySQLUser) generatePassword(user *apiv1alpha1.MySQLUser) error {
	s := &corev1.Secret{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: user.Name + "-db-secret", Namespace: user.Namespace}, s)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	length := 32
	charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		return err
	}

	for i := 0; i < length; i++ {
		buf[i] = charset[int(buf[i])%len(charset)]
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      user.Name + "-db-secret",
			Namespace: user.Namespace,
			Labels:    user.ObjectMeta.Labels,
		},

		Data: map[string][]byte{
			"mysql-password": buf,
		},
	}

	if err := r.client.Create(context.TODO(), secret); err != nil {
		// FIX ME
		// This should instead try to get the secret
		if !errors.IsAlreadyExists(err) {
			return err
		}
	}

	user.Status.PasswordSecretName = &corev1.SecretKeySelector{
		LocalObjectReference: corev1.LocalObjectReference{
			Name: user.Name + "-db-secret",
		},
		Key: "mysql-password",
	}

	if err := r.client.Update(context.TODO(), user); err != nil {
		return err
	}

	return nil
}
