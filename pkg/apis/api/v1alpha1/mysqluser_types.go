package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MySQLUserSpec defines the desired state of MySQLUser
type MySQLUserSpec struct {
	Name                string                     `json:"name"`
	Host                string                     `json:"host,omitempty"`
	ExternalDatabaseRef *ExternalDatabaseReference `json:"externalDatabaseRef"`
	ForcePasswordUpdate bool                       `json:"forcePasswordUpdate"`

	// If this is empty a password will be generated
	Password *corev1.SecretKeySelector `json:"existingPasswordSecretRef,omitempty"`
}

// MySQLUserStatus defines the observed state of MySQLUser
type MySQLUserStatus struct {
	Created            bool                      `json:"created"`
	PasswordSecretName *corev1.SecretKeySelector `json:"passwordSecretName"`
	Message            string                    `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLUser is the Schema for the mysqlusers API
// +k8s:openapi-gen=true
type MySQLUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLUserSpec   `json:"spec,omitempty"`
	Status MySQLUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLUserList contains a list of MySQLUser
type MySQLUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLUser{}, &MySQLUserList{})
}
