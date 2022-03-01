//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediation) DeepCopyInto(out *PolicyRemediation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediation.
func (in *PolicyRemediation) DeepCopy() *PolicyRemediation {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyRemediation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediationList) DeepCopyInto(out *PolicyRemediationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyRemediation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediationList.
func (in *PolicyRemediationList) DeepCopy() *PolicyRemediationList {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyRemediationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediationObservation) DeepCopyInto(out *PolicyRemediationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediationObservation.
func (in *PolicyRemediationObservation) DeepCopy() *PolicyRemediationObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediationParameters) DeepCopyInto(out *PolicyRemediationParameters) {
	*out = *in
	if in.LocationFilters != nil {
		in, out := &in.LocationFilters, &out.LocationFilters
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyAssignmentID != nil {
		in, out := &in.PolicyAssignmentID, &out.PolicyAssignmentID
		*out = new(string)
		**out = **in
	}
	if in.PolicyDefinitionReferenceID != nil {
		in, out := &in.PolicyDefinitionReferenceID, &out.PolicyDefinitionReferenceID
		*out = new(string)
		**out = **in
	}
	if in.ResourceDiscoveryMode != nil {
		in, out := &in.ResourceDiscoveryMode, &out.ResourceDiscoveryMode
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediationParameters.
func (in *PolicyRemediationParameters) DeepCopy() *PolicyRemediationParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediationSpec) DeepCopyInto(out *PolicyRemediationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediationSpec.
func (in *PolicyRemediationSpec) DeepCopy() *PolicyRemediationSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediationStatus) DeepCopyInto(out *PolicyRemediationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediationStatus.
func (in *PolicyRemediationStatus) DeepCopy() *PolicyRemediationStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediationStatus)
	in.DeepCopyInto(out)
	return out
}
