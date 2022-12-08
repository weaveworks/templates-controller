package core

// ConvertV1SpecToSpec converts between the two specs.
func ConvertV1SpecToSpec(s TemplateSpecV1) TemplateSpec {
	return TemplateSpec{
		Description: s.Description,
		RenderType:  s.RenderType,
		Params:      s.Params,
		Charts:      s.Charts,
		ResourceTemplates: []ResourceTemplate{
			{Content: s.ResourceTemplates},
		},
	}
}

func ConvertV2SpecToV1Spec(s TemplateSpec) TemplateSpecV1 {
	resourceTemplates := []ResourceTemplateContent{}
	for _, rt := range s.ResourceTemplates {
		resourceTemplates = append(resourceTemplates, rt.Content...)
	}

	return TemplateSpecV1{
		Description:       s.Description,
		RenderType:        s.RenderType,
		Params:            s.Params,
		Charts:            s.Charts,
		ResourceTemplates: resourceTemplates,
	}
}
