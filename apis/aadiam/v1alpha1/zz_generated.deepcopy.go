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
func (in *LogObservation) DeepCopyInto(out *LogObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogObservation.
func (in *LogObservation) DeepCopy() *LogObservation {
	if in == nil {
		return nil
	}
	out := new(LogObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogParameters) DeepCopyInto(out *LogParameters) {
	*out = *in
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = make([]RetentionPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogParameters.
func (in *LogParameters) DeepCopy() *LogParameters {
	if in == nil {
		return nil
	}
	out := new(LogParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorAADDiagnosticSetting) DeepCopyInto(out *MonitorAADDiagnosticSetting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorAADDiagnosticSetting.
func (in *MonitorAADDiagnosticSetting) DeepCopy() *MonitorAADDiagnosticSetting {
	if in == nil {
		return nil
	}
	out := new(MonitorAADDiagnosticSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorAADDiagnosticSetting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorAADDiagnosticSettingList) DeepCopyInto(out *MonitorAADDiagnosticSettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitorAADDiagnosticSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorAADDiagnosticSettingList.
func (in *MonitorAADDiagnosticSettingList) DeepCopy() *MonitorAADDiagnosticSettingList {
	if in == nil {
		return nil
	}
	out := new(MonitorAADDiagnosticSettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorAADDiagnosticSettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorAADDiagnosticSettingObservation) DeepCopyInto(out *MonitorAADDiagnosticSettingObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorAADDiagnosticSettingObservation.
func (in *MonitorAADDiagnosticSettingObservation) DeepCopy() *MonitorAADDiagnosticSettingObservation {
	if in == nil {
		return nil
	}
	out := new(MonitorAADDiagnosticSettingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorAADDiagnosticSettingParameters) DeepCopyInto(out *MonitorAADDiagnosticSettingParameters) {
	*out = *in
	if in.EventHubAuthorizationRuleID != nil {
		in, out := &in.EventHubAuthorizationRuleID, &out.EventHubAuthorizationRuleID
		*out = new(string)
		**out = **in
	}
	if in.EventHubName != nil {
		in, out := &in.EventHubName, &out.EventHubName
		*out = new(string)
		**out = **in
	}
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = make([]LogParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LogAnalyticsWorkspaceID != nil {
		in, out := &in.LogAnalyticsWorkspaceID, &out.LogAnalyticsWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorAADDiagnosticSettingParameters.
func (in *MonitorAADDiagnosticSettingParameters) DeepCopy() *MonitorAADDiagnosticSettingParameters {
	if in == nil {
		return nil
	}
	out := new(MonitorAADDiagnosticSettingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorAADDiagnosticSettingSpec) DeepCopyInto(out *MonitorAADDiagnosticSettingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorAADDiagnosticSettingSpec.
func (in *MonitorAADDiagnosticSettingSpec) DeepCopy() *MonitorAADDiagnosticSettingSpec {
	if in == nil {
		return nil
	}
	out := new(MonitorAADDiagnosticSettingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorAADDiagnosticSettingStatus) DeepCopyInto(out *MonitorAADDiagnosticSettingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorAADDiagnosticSettingStatus.
func (in *MonitorAADDiagnosticSettingStatus) DeepCopy() *MonitorAADDiagnosticSettingStatus {
	if in == nil {
		return nil
	}
	out := new(MonitorAADDiagnosticSettingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicyObservation) DeepCopyInto(out *RetentionPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicyObservation.
func (in *RetentionPolicyObservation) DeepCopy() *RetentionPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicyParameters) DeepCopyInto(out *RetentionPolicyParameters) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicyParameters.
func (in *RetentionPolicyParameters) DeepCopy() *RetentionPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicyParameters)
	in.DeepCopyInto(out)
	return out
}