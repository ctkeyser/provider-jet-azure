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

type VirtualWanObservation struct {
}

type VirtualWanParameters struct {

	// +kubebuilder:validation:Optional
	AllowBranchToBranchTraffic *bool `json:"allowBranchToBranchTraffic,omitempty" tf:"allow_branch_to_branch_traffic"`

	// +kubebuilder:validation:Optional
	AllowVnetToVnetTraffic *bool `json:"allowVnetToVnetTraffic,omitempty" tf:"allow_vnet_to_vnet_traffic"`

	// +kubebuilder:validation:Optional
	DisableVpnEncryption *bool `json:"disableVpnEncryption,omitempty" tf:"disable_vpn_encryption"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Office365LocalBreakoutCategory *string `json:"office365LocalBreakoutCategory,omitempty" tf:"office365_local_breakout_category"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`
}

// VirtualWanSpec defines the desired state of VirtualWan
type VirtualWanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualWanParameters `json:"forProvider"`
}

// VirtualWanStatus defines the observed state of VirtualWan.
type VirtualWanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualWanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualWan is the Schema for the VirtualWans API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualWan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualWanSpec   `json:"spec"`
	Status            VirtualWanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualWanList contains a list of VirtualWans
type VirtualWanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualWan `json:"items"`
}

// Repository type metadata.
var (
	VirtualWanKind             = "VirtualWan"
	VirtualWanGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualWanKind}.String()
	VirtualWanKindAPIVersion   = VirtualWanKind + "." + GroupVersion.String()
	VirtualWanGroupVersionKind = GroupVersion.WithKind(VirtualWanKind)
)

func init() {
	SchemeBuilder.Register(&VirtualWan{}, &VirtualWanList{})
}