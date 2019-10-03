package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MySQLUserSpec defines the desired state of MySqlUser
// +k8s:openapi-gen=true
type MySQLUserSpec struct {
	Name                string               `json:"name"`
	Host                string               `json:"host,omitempty"`
	ExternalDatabaseRef *ExternalDatabaseRef `json:"externalDatabaseRef"`
	PasswordSecretName  string               `json:"passwordSecretName"`
	// ExistingPasswordSecretRef *corev1.SecretKeySelector `json:"existingPasswordSecretRef"`
}

// MySQLUserStatus defines the observed state of MySqlUser
// +k8s:openapi-gen=true
type MySQLUserStatus struct {
	Created bool   `json:"created"`
	Error   string `json:"error"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLUser is the Schema for the mysqlusers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mysqlusers,scope=Namespaced
type MySQLUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLUserSpec   `json:"spec,omitempty"`
	Status MySQLUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLUserList contains a list of MySqlUser
type MySQLUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLUser{}, &MySQLUserList{})
}
