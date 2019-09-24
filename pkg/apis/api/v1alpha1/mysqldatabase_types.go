package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MySqlDatabaseSpec defines the desired state of MySqlDatabase
// +k8s:openapi-gen=true
type MySqlDatabaseSpec struct {
	Name                string              `json:"name"`
	Owner               *string             `json:"owner"`
	ExternalDatabaseRef ExternalDatabaseRef `json:"externalDatabaseRef"`
}

// MySqlDatabaseStatus defines the observed state of MySqlDatabase
// +k8s:openapi-gen=true
type MySqlDatabaseStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySqlDatabase is the Schema for the mysqldatabases API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mysqldatabases,scope=Namespaced
type MySqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySqlDatabaseSpec   `json:"spec,omitempty"`
	Status MySqlDatabaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySqlDatabaseList contains a list of MySqlDatabase
type MySqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySqlDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySqlDatabase{}, &MySqlDatabaseList{})
}
