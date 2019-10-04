package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MySQLDatabaseSpec defines the desired state of MySqlDatabase
// +k8s:openapi-gen=true
type MySQLDatabaseSpec struct {
	Name                string              `json:"name"`
	ExternalDatabaseRef ExternalDatabaseRef `json:"externalDatabaseRef"`
}

// MySQLDatabaseStatus defines the observed state of MySQLDatabase
// +k8s:openapi-gen=true
type MySQLDatabaseStatus struct {
	Error   string `json:"error,omitempty"`
	Created bool   `json:"created,omitempty"`

	SecretCreated  bool   `json:"secretCreated"`
	ExistingSecret string `json:"existingSecret,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLDatabase is the Schema for the mysqldatabases API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mysqldatabases,scope=Namespaced
type MySQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLDatabaseSpec   `json:"spec,omitempty"`
	Status MySQLDatabaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLDatabaseList contains a list of MySQLDatabase
type MySQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLDatabase{}, &MySQLDatabaseList{})
}
