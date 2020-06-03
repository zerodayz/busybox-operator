package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BusyboxSpec defines the desired state of Busybox
type BusyboxSpec struct {
	// Size is the size of the busybox deployment
	Size int32 `json:"size"`
}

// BusyboxStatus defines the observed state of Busybox
type BusyboxStatus struct {
	// Nodes are the names of the busybox pods
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Busybox is the Schema for the busyboxes API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=busyboxes,scope=Namespaced
type Busybox struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BusyboxSpec   `json:"spec,omitempty"`
	Status BusyboxStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BusyboxList contains a list of Busybox
type BusyboxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Busybox `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Busybox{}, &BusyboxList{})
}
