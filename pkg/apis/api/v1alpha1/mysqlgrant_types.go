package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MySQLGrantSpec defines the desired state of MySQLGrant
// +k8s:openapi-gen=true
type MySQLGrantSpec struct {
	To string `json:"to"`
	On string `json:"on"`

	Permission string `json:"permission"`
}

// MySQLGrantStatus defines the observed state of MySQLGrant
// +k8s:openapi-gen=true
type MySQLGrantStatus struct {
	Granted bool `json:"granted"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLGrant is the Schema for the mysqlgrants API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mysqlgrants,scope=Namespaced
type MySQLGrant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLGrantSpec   `json:"spec,omitempty"`
	Status MySQLGrantStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLGrantList contains a list of MySQLGrant
type MySQLGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLGrant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLGrant{}, &MySQLGrantList{})
}
