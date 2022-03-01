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

type PublicIPObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
}

type PublicIPParameters struct {

	// +kubebuilder:validation:Required
	AllocationMethod *string `json:"allocationMethod" tf:"allocation_method,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	DomainNameLabel *string `json:"domainNameLabel,omitempty" tf:"domain_name_label,omitempty"`

	// +kubebuilder:validation:Optional
	IPTags map[string]*string `json:"ipTags,omitempty" tf:"ip_tags,omitempty"`

	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ReverseFqdn *string `json:"reverseFqdn,omitempty" tf:"reverse_fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

// PublicIPSpec defines the desired state of PublicIP
type PublicIPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicIPParameters `json:"forProvider"`
}

// PublicIPStatus defines the observed state of PublicIP.
type PublicIPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicIPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIP is the Schema for the PublicIPs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type PublicIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicIPSpec   `json:"spec"`
	Status            PublicIPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIPList contains a list of PublicIPs
type PublicIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicIP `json:"items"`
}

// Repository type metadata.
var (
	PublicIP_Kind             = "PublicIP"
	PublicIP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicIP_Kind}.String()
	PublicIP_KindAPIVersion   = PublicIP_Kind + "." + CRDGroupVersion.String()
	PublicIP_GroupVersionKind = CRDGroupVersion.WithKind(PublicIP_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicIP{}, &PublicIPList{})
}
