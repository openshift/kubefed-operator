package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

/// KubeFedWebHook related type declarations

// KubeFedWebHookSpec defines the desired state of KubeFedWebHook
// +k8s:openapi-gen=true
type KubeFedWebHookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

}

// KubeFedWebHookStatus defines the observed state of KubeFedWebHook
// +k8s:openapi-gen=true
type KubeFedWebHookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// The version of the installed release
	// +optional
	Version string `json:"version,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeFedWebHook is the Schema for the installs API
// +k8s:openapi-gen=true
type KubeFedWebHook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeFedWebHookSpec   `json:"spec"`
	Status KubeFedWebHookStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeFedWebHookList contains a list of KubeFedWebHook
type KubeFedWebHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeFedWebHook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeFedWebHook{}, &KubeFedWebHookList{})
}
