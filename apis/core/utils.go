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
