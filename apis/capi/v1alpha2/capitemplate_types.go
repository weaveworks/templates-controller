package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/weaveworks/templates-controller/apis/core"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// CAPITemplate is the Schema for the capitemplates API
type CAPITemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec core.TemplateSpec `json:"spec,omitempty"`
}

// Hub marks this as a conversion hub.
//
// https://book.kubebuilder.io/multiversion-tutorial/conversion.html
func (*CAPITemplate) Hub() {}

//+kubebuilder:object:root=true

// CAPITemplateList contains a list of CAPITemplate
type CAPITemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CAPITemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CAPITemplate{}, &CAPITemplateList{})
}
