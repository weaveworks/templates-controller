package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestConvertV1SpectToSpec(t *testing.T) {
	expected := TemplateSpec{
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
		// only diff is here
		ResourceTemplates: []ResourceTemplate{
			{
				Content: []ResourceTemplateContent{
					makeTemplateResource("raw"),
				},
			},
		},
	}

	v1Spec := makeV1Alpha1TemplateSpec()
	spec := ConvertV1SpecToSpec(v1Spec)

	// What we actually convert, wrap the old Content under `Content`
	assert.Equal(t, expected, spec)
}

func makeTemplateSpec() TemplateSpec {
	return TemplateSpec{
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
}

func makeV1Alpha1TemplateSpec() TemplateSpecV1 {
	return TemplateSpecV1{
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
}

func makeTemplateResource(data string) ResourceTemplateContent {
	return ResourceTemplateContent{
		RawExtension: runtime.RawExtension{
			Raw: []byte(data),
		},
	}
}
