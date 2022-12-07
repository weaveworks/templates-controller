package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/weaveworks/templates-controller/apis/core"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// GitOpsTemplate is the Schema for the gitopstemplates API
type GitOpsTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec core.TemplateSpec `json:"spec,omitempty"`
}

// Hub marks this as a conversion hub.
//
// https://book.kubebuilder.io/multiversion-tutorial/conversion.html
func (*GitOpsTemplate) Hub() {}

//+kubebuilder:object:root=true

// GitOpsTemplateList contains a list of GitOpsTemplate
type GitOpsTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitOpsTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitOpsTemplate{}, &GitOpsTemplateList{})
}
