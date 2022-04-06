/*
Copyright 2022 The Crossplane Authors.

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

type CircuitObservation struct {
	ExpressRouteID *string `json:"expressRouteId,omitempty" tf:"express_route_id,omitempty"`

	ExpressRoutePrivatePeeringID *string `json:"expressRoutePrivatePeeringId,omitempty" tf:"express_route_private_peering_id,omitempty"`

	PrimarySubnetCidr *string `json:"primarySubnetCidr,omitempty" tf:"primary_subnet_cidr,omitempty"`

	SecondarySubnetCidr *string `json:"secondarySubnetCidr,omitempty" tf:"secondary_subnet_cidr,omitempty"`
}

type CircuitParameters struct {
}

type ManagementClusterObservation struct {
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagementClusterParameters struct {

	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`
}

type VMwarePrivateCloudObservation struct {
	Circuit []CircuitObservation `json:"circuit,omitempty" tf:"circuit,omitempty"`

	HcxCloudManagerEndpoint *string `json:"hcxCloudManagerEndpoint,omitempty" tf:"hcx_cloud_manager_endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ManagementSubnetCidr *string `json:"managementSubnetCidr,omitempty" tf:"management_subnet_cidr,omitempty"`

	NsxtCertificateThumbprint *string `json:"nsxtCertificateThumbprint,omitempty" tf:"nsxt_certificate_thumbprint,omitempty"`

	NsxtManagerEndpoint *string `json:"nsxtManagerEndpoint,omitempty" tf:"nsxt_manager_endpoint,omitempty"`

	ProvisioningSubnetCidr *string `json:"provisioningSubnetCidr,omitempty" tf:"provisioning_subnet_cidr,omitempty"`

	VcenterCertificateThumbprint *string `json:"vcenterCertificateThumbprint,omitempty" tf:"vcenter_certificate_thumbprint,omitempty"`

	VcsaEndpoint *string `json:"vcsaEndpoint,omitempty" tf:"vcsa_endpoint,omitempty"`

	VmotionSubnetCidr *string `json:"vmotionSubnetCidr,omitempty" tf:"vmotion_subnet_cidr,omitempty"`
}

type VMwarePrivateCloudParameters struct {

	// +kubebuilder:validation:Optional
	InternetConnectionEnabled *bool `json:"internetConnectionEnabled,omitempty" tf:"internet_connection_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	ManagementCluster []ManagementClusterParameters `json:"managementCluster" tf:"management_cluster,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkSubnetCidr *string `json:"networkSubnetCidr" tf:"network_subnet_cidr,omitempty"`

	// +kubebuilder:validation:Optional
	NsxtPasswordSecretRef *v1.SecretKeySelector `json:"nsxtPasswordSecretRef,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VcenterPasswordSecretRef *v1.SecretKeySelector `json:"vcenterPasswordSecretRef,omitempty" tf:"-"`
}

// VMwarePrivateCloudSpec defines the desired state of VMwarePrivateCloud
type VMwarePrivateCloudSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VMwarePrivateCloudParameters `json:"forProvider"`
}

// VMwarePrivateCloudStatus defines the observed state of VMwarePrivateCloud.
type VMwarePrivateCloudStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VMwarePrivateCloudObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VMwarePrivateCloud is the Schema for the VMwarePrivateClouds API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VMwarePrivateCloud struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VMwarePrivateCloudSpec   `json:"spec"`
	Status            VMwarePrivateCloudStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VMwarePrivateCloudList contains a list of VMwarePrivateClouds
type VMwarePrivateCloudList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMwarePrivateCloud `json:"items"`
}

// Repository type metadata.
var (
	VMwarePrivateCloud_Kind             = "VMwarePrivateCloud"
	VMwarePrivateCloud_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VMwarePrivateCloud_Kind}.String()
	VMwarePrivateCloud_KindAPIVersion   = VMwarePrivateCloud_Kind + "." + CRDGroupVersion.String()
	VMwarePrivateCloud_GroupVersionKind = CRDGroupVersion.WithKind(VMwarePrivateCloud_Kind)
)

func init() {
	SchemeBuilder.Register(&VMwarePrivateCloud{}, &VMwarePrivateCloudList{})
}