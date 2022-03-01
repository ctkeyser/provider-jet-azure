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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type HPCCacheAccessPolicyAccessRuleObservation struct {
}

type HPCCacheAccessPolicyAccessRuleParameters struct {

	// +kubebuilder:validation:Required
	Access *string `json:"access" tf:"access,omitempty"`

	// +kubebuilder:validation:Optional
	AnonymousGID *int64 `json:"anonymousGid,omitempty" tf:"anonymous_gid,omitempty"`

	// +kubebuilder:validation:Optional
	AnonymousUID *int64 `json:"anonymousUid,omitempty" tf:"anonymous_uid,omitempty"`

	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Optional
	RootSquashEnabled *bool `json:"rootSquashEnabled,omitempty" tf:"root_squash_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	SubmountAccessEnabled *bool `json:"submountAccessEnabled,omitempty" tf:"submount_access_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	SuidEnabled *bool `json:"suidEnabled,omitempty" tf:"suid_enabled,omitempty"`
}

type HPCCacheAccessPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HPCCacheAccessPolicyParameters struct {

	// +kubebuilder:validation:Required
	AccessRule []HPCCacheAccessPolicyAccessRuleParameters `json:"accessRule" tf:"access_rule,omitempty"`

	// +kubebuilder:validation:Required
	HPCCacheID *string `json:"hpcCacheId" tf:"hpc_cache_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// HPCCacheAccessPolicySpec defines the desired state of HPCCacheAccessPolicy
type HPCCacheAccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HPCCacheAccessPolicyParameters `json:"forProvider"`
}

// HPCCacheAccessPolicyStatus defines the observed state of HPCCacheAccessPolicy.
type HPCCacheAccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HPCCacheAccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheAccessPolicy is the Schema for the HPCCacheAccessPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type HPCCacheAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HPCCacheAccessPolicySpec   `json:"spec"`
	Status            HPCCacheAccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheAccessPolicyList contains a list of HPCCacheAccessPolicys
type HPCCacheAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HPCCacheAccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	HPCCacheAccessPolicy_Kind             = "HPCCacheAccessPolicy"
	HPCCacheAccessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HPCCacheAccessPolicy_Kind}.String()
	HPCCacheAccessPolicy_KindAPIVersion   = HPCCacheAccessPolicy_Kind + "." + CRDGroupVersion.String()
	HPCCacheAccessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(HPCCacheAccessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&HPCCacheAccessPolicy{}, &HPCCacheAccessPolicyList{})
}
