package v1alpha1

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/weaveworks/templates-controller/apis/core"
	gitopsv1alpha2 "github.com/weaveworks/templates-controller/apis/gitops/v1alpha2"
)

// ConvertTo converts this GitOpsTemplate to the Hub version (v1alpha2).
func (src *GitOpsTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*gitopsv1alpha2.GitOpsTemplate)
	dst.ObjectMeta = src.ObjectMeta
	dst.Spec = src.GetSpec()

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha2) to this version (v1alpha1).
func (dst *GitOpsTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*gitopsv1alpha2.GitOpsTemplate)
	dst.ObjectMeta = src.ObjectMeta
	dst.Spec = core.ConvertV2SpecToV1Spec(src.Spec)

	return nil
}
