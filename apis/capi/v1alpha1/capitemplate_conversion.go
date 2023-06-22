package v1alpha1

import (
	capiv1alpha2 "github.com/weaveworks/templates-controller/apis/capi/v1alpha2"
	"github.com/weaveworks/templates-controller/apis/core"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this CAPITemplate to the Hub version (v1alpha2).
func (src *CAPITemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*capiv1alpha2.CAPITemplate)
	dst.ObjectMeta = src.ObjectMeta
	dst.Spec = src.GetSpec()

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha2) to this version (v1alpha1).
func (dst *CAPITemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*capiv1alpha2.CAPITemplate)
	dst.ObjectMeta = src.ObjectMeta
	dst.Spec = core.ConvertV2SpecToV1Spec(src.Spec)

	return nil
}
