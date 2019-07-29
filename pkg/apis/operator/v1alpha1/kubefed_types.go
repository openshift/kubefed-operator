package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

/// Type and variable declarations common to all CRs that this operator is responsible for

// InstallationScope defines the scope of the resource being installed
// +k8s:openapi-gen=true
type InstallationScope string

var (
	// InstallationScopeNamespaceScoped defines Namespace scoped installation scope for a resource
	InstallationScopeNamespaceScoped InstallationScope = "Namespaced"
	// InstallationScopeClusterScoped defines Cluster scoped installation scope for a resource
	InstallationScopeClusterScoped InstallationScope = "Cluster"
)

/// KubeFed resource related type declarations

// KubeFedSpec defines the desired state of KubeFed
// +k8s:openapi-gen=true
type KubeFedSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	Scope InstallationScope `json:"scope"`
}

// KubeFedStatus defines the observed state of KubeFed
// +k8s:openapi-gen=true
type KubeFedStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// The version of the installed release
	// +optional
	Version string `json:"version,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeFed is the Schema for the installs API
// +k8s:openapi-gen=true
type KubeFed struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeFedSpec   `json:"spec,omitempty"`
	Status KubeFedStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeFedList contains a list of KubeFed
type KubeFedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeFed `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeFed{}, &KubeFedList{})
}
