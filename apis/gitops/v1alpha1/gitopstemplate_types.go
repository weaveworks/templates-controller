package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/weaveworks/templates-controller/apis/core"
)

const Kind = "GitOpsTemplate"

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GitOpsTemplate is the Schema for the gitopstemplates API
type GitOpsTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec core.TemplateSpecV1 `json:"spec,omitempty"`
}

func (t GitOpsTemplate) GetSpec() core.TemplateSpec {
	return core.ConvertV1SpecToSpec(t.Spec)
}

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
