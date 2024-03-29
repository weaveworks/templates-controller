//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package core

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chart) DeepCopyInto(out *Chart) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new(HelmReleaseValues)
		(*in).DeepCopyInto(*out)
	}
	in.HelmReleaseTemplate.DeepCopyInto(&out.HelmReleaseTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chart.
func (in *Chart) DeepCopy() *Chart {
	if in == nil {
		return nil
	}
	out := new(Chart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartsSpec) DeepCopyInto(out *ChartsSpec) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Chart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.HelmRepositoryTemplate = in.HelmRepositoryTemplate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartsSpec.
func (in *ChartsSpec) DeepCopy() *ChartsSpec {
	if in == nil {
		return nil
	}
	out := new(ChartsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseTemplate) DeepCopyInto(out *HelmReleaseTemplate) {
	*out = *in
	in.RawExtension.DeepCopyInto(&out.RawExtension)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseTemplate.
func (in *HelmReleaseTemplate) DeepCopy() *HelmReleaseTemplate {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseTemplateSpec) DeepCopyInto(out *HelmReleaseTemplateSpec) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(HelmReleaseTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseTemplateSpec.
func (in *HelmReleaseTemplateSpec) DeepCopy() *HelmReleaseTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseValues) DeepCopyInto(out *HelmReleaseValues) {
	*out = *in
	in.RawExtension.DeepCopyInto(&out.RawExtension)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseValues.
func (in *HelmReleaseValues) DeepCopy() *HelmReleaseValues {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRepositoryTemplateSpec) DeepCopyInto(out *HelmRepositoryTemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRepositoryTemplateSpec.
func (in *HelmRepositoryTemplateSpec) DeepCopy() *HelmRepositoryTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HelmRepositoryTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTemplate) DeepCopyInto(out *ResourceTemplate) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = make([]ResourceTemplateContent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTemplate.
func (in *ResourceTemplate) DeepCopy() *ResourceTemplate {
	if in == nil {
		return nil
	}
	out := new(ResourceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTemplateContent) DeepCopyInto(out *ResourceTemplateContent) {
	*out = *in
	in.RawExtension.DeepCopyInto(&out.RawExtension)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTemplateContent.
func (in *ResourceTemplateContent) DeepCopy() *ResourceTemplateContent {
	if in == nil {
		return nil
	}
	out := new(ResourceTemplateContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateParam) DeepCopyInto(out *TemplateParam) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateParam.
func (in *TemplateParam) DeepCopy() *TemplateParam {
	if in == nil {
		return nil
	}
	out := new(TemplateParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpec) DeepCopyInto(out *TemplateSpec) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]TemplateParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Charts.DeepCopyInto(&out.Charts)
	if in.ResourceTemplates != nil {
		in, out := &in.ResourceTemplates, &out.ResourceTemplates
		*out = make([]ResourceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpec.
func (in *TemplateSpec) DeepCopy() *TemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecV1) DeepCopyInto(out *TemplateSpecV1) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]TemplateParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Charts.DeepCopyInto(&out.Charts)
	if in.ResourceTemplates != nil {
		in, out := &in.ResourceTemplates, &out.ResourceTemplates
		*out = make([]ResourceTemplateContent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecV1.
func (in *TemplateSpecV1) DeepCopy() *TemplateSpecV1 {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecV1)
	in.DeepCopyInto(out)
	return out
}
