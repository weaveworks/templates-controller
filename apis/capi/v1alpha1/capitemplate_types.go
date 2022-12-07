package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/weaveworks/templates-controller/apis/core"
)

const Kind = "CAPITemplate"

//+kubebuilder:object:root=true

// CAPITemplate is the Schema for the capitemplates API
type CAPITemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec core.TemplateSpecV1 `json:"spec,omitempty"`
}

func (t CAPITemplate) GetSpec() core.TemplateSpec {
	return core.ConvertV1SpecToSpec(t.Spec)
}

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
