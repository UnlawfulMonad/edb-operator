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

// ExternalDatabaseSpec defines the desired state of ExternalDatabase
// +k8s:openapi-gen=true
type ExternalDatabaseSpec struct {
	Host              string                   `json:"host"`
	AdminUser         string                   `json:"adminUser"`
	AdminPasswordRef  corev1.SecretKeySelector `json:"adminPasswordRef"`
	Type              DatabaseType             `json:"type"`
	NamespaceSelector *metav1.LabelSelector    `json:"namespaceSelector,omitempty"`
}

// ExternalDatabaseRef talks about an external database.
// +k8s:openapi-gen=true
type ExternalDatabaseRef struct {
	Name string `json:"name,omitempty"`
}

// ExternalDatabaseStatus defines the observed state of ExternalDatabase
// +k8s:openapi-gen=true
type ExternalDatabaseStatus struct {
	Reachable bool   `json:"reachable,omitempty"`
	Error     string `json:"error,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExternalDatabase is the Schema for the externaldatabases API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=externaldatabases,scope=Cluster
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
