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

type NATGatewayPublicIPPrefixAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NATGatewayPublicIPPrefixAssociationParameters struct {

	// +kubebuilder:validation:Required
	NATGatewayID *string `json:"natGatewayId" tf:"nat_gateway_id,omitempty"`

	// +kubebuilder:validation:Required
	PublicIPPrefixID *string `json:"publicIpPrefixId" tf:"public_ip_prefix_id,omitempty"`
}

// NATGatewayPublicIPPrefixAssociationSpec defines the desired state of NATGatewayPublicIPPrefixAssociation
type NATGatewayPublicIPPrefixAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NATGatewayPublicIPPrefixAssociationParameters `json:"forProvider"`
}

// NATGatewayPublicIPPrefixAssociationStatus defines the observed state of NATGatewayPublicIPPrefixAssociation.
type NATGatewayPublicIPPrefixAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NATGatewayPublicIPPrefixAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NATGatewayPublicIPPrefixAssociation is the Schema for the NATGatewayPublicIPPrefixAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type NATGatewayPublicIPPrefixAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NATGatewayPublicIPPrefixAssociationSpec   `json:"spec"`
	Status            NATGatewayPublicIPPrefixAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NATGatewayPublicIPPrefixAssociationList contains a list of NATGatewayPublicIPPrefixAssociations
type NATGatewayPublicIPPrefixAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NATGatewayPublicIPPrefixAssociation `json:"items"`
}

// Repository type metadata.
var (
	NATGatewayPublicIPPrefixAssociation_Kind             = "NATGatewayPublicIPPrefixAssociation"
	NATGatewayPublicIPPrefixAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NATGatewayPublicIPPrefixAssociation_Kind}.String()
	NATGatewayPublicIPPrefixAssociation_KindAPIVersion   = NATGatewayPublicIPPrefixAssociation_Kind + "." + CRDGroupVersion.String()
	NATGatewayPublicIPPrefixAssociation_GroupVersionKind = CRDGroupVersion.WithKind(NATGatewayPublicIPPrefixAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&NATGatewayPublicIPPrefixAssociation{}, &NATGatewayPublicIPPrefixAssociationList{})
}
