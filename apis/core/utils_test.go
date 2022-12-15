package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestConvertV1SpectToSpec(t *testing.T) {
	v1Spec := TemplateSpecV1{
		Description: "description",
		RenderType:  "renderType",
		Params: []TemplateParam{
			{
				Name:        "name",
				Description: "description",
			},
		},
		Charts: ChartsSpec{
			Items: []Chart{
				{
					Chart: "chart",
				},
			},
		},
		ResourceTemplates: []ResourceTemplateContent{
			makeTemplateResource("raw"),
		},
	}

	spec := ConvertV1SpecToSpec(v1Spec)

	// All the bits of the spec made it
	assert.Equal(t, v1Spec.Description, spec.Description)
	assert.Equal(t, v1Spec.RenderType, spec.RenderType)
	assert.Equal(t, v1Spec.Params, spec.Params)
	assert.Equal(t, v1Spec.Charts, spec.Charts)

	// What we actually convert, wrap the old Content under `Content`
	assert.Equal(t, v1Spec.ResourceTemplates, spec.ResourceTemplates[0].Content)
}

func TestConvertV2SpecToV1Spec(t *testing.T) {
	v2spec := TemplateSpec{
		Description: "description",
		RenderType:  "renderType",
		Params: []TemplateParam{
			{
				Name:        "name",
				Description: "description",
			},
		},
		Charts: ChartsSpec{
			Items: []Chart{
				{
					Chart: "chart",
				},
			},
		},
		ResourceTemplates: []ResourceTemplate{
			{
				Content: []ResourceTemplateContent{
					makeTemplateResource("foo"),
					makeTemplateResource("foo2"),
				},
				Path: "foos.yaml",
			},
			{
				Content: []ResourceTemplateContent{
					makeTemplateResource("bar"),
					makeTemplateResource("bar2"),
				},
				Path: "bars.yaml",
			},
		},
	}

	v1Spec := ConvertV2SpecToV1Spec(v2spec)

	// All the bits of the spec made it
	assert.Equal(t, v2spec.Description, v1Spec.Description)
	assert.Equal(t, v2spec.RenderType, v1Spec.RenderType)
	assert.Equal(t, v2spec.Params, v1Spec.Params)
	assert.Equal(t, v2spec.Charts, v1Spec.Charts)

	// everything concatenated, paths lost
	expectedTemplateResources := []ResourceTemplateContent{
		makeTemplateResource("foo"),
		makeTemplateResource("foo2"),
		makeTemplateResource("bar"),
		makeTemplateResource("bar2"),
	}

	assert.Equal(t, expectedTemplateResources, v1Spec.ResourceTemplates)
}

func makeTemplateResource(data string) ResourceTemplateContent {
	return ResourceTemplateContent{
		RawExtension: runtime.RawExtension{
			Raw: []byte(data),
		},
	}
}
