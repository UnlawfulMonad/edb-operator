package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MySqlUserSpec defines the desired state of MySqlUser
// +k8s:openapi-gen=true
type MySqlUserSpec struct {
	Name                string                    `json:"name"`
	Host                string                    `json:"host,omitempty"`
	ExternalDatabaseRef *ExternalDatabaseRef      `json:"externalDatabaseRef"`
	Password            *corev1.SecretKeySelector `json:"existingPasswordSecretRef"`
}

// MySqlUserStatus defines the observed state of MySqlUser
// +k8s:openapi-gen=true
type MySqlUserStatus struct {
	Created bool `json:"created"`
	Error   bool `json:"hasError"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySqlUser is the Schema for the mysqlusers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mysqlusers,scope=Namespaced
type MySqlUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySqlUserSpec   `json:"spec,omitempty"`
	Status MySqlUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySqlUserList contains a list of MySqlUser
type MySqlUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySqlUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySqlUser{}, &MySqlUserList{})
}
