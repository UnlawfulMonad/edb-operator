package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DatabaseType string

const (
	MySQL      DatabaseType = "mysql"
	PostgreSQL DatabaseType = "postgresql"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExternalDatabaseSpec defines the desired state of ExternalDatabase
// +k8s:openapi-gen=true
type ExternalDatabaseSpec struct {
	Host          string                 `json:"host"`
	AdminUser     string                 `json:"adminUser"`
	AdminPassword corev1.SecretReference `json:"adminPasswordSecretRef"`
	Type          DatabaseType           `json:"type"`
	Selector      *metav1.LabelSelector  `json:"namespaceSelector"`
}

type ExternalDatabaseRef struct {
	Name string `json:"name"`
}

// ExternalDatabaseStatus defines the observed state of ExternalDatabase
// +k8s:openapi-gen=true
type ExternalDatabaseStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExternalDatabase is the Schema for the externaldatabases API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=externaldatabases,scope=Namespaced
type ExternalDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalDatabaseSpec   `json:"spec,omitempty"`
	Status ExternalDatabaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExternalDatabaseList contains a list of ExternalDatabase
type ExternalDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExternalDatabase{}, &ExternalDatabaseList{})
}
