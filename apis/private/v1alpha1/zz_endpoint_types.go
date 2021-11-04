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

type CustomDNSConfigsObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`
}

type CustomDNSConfigsParameters struct {
}

type EndpointObservation struct {
	CustomDNSConfigs []CustomDNSConfigsObservation `json:"customDnsConfigs,omitempty" tf:"custom_dns_configs,omitempty"`

	NetworkInterface []NetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	PrivateDNSZoneConfigs []PrivateDNSZoneConfigsObservation `json:"privateDnsZoneConfigs,omitempty" tf:"private_dns_zone_configs,omitempty"`
}

type EndpointParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateDNSZoneGroup []PrivateDNSZoneGroupParameters `json:"privateDnsZoneGroup,omitempty" tf:"private_dns_zone_group,omitempty"`

	// +kubebuilder:validation:Required
	PrivateServiceConnection []PrivateServiceConnectionParameters `json:"privateServiceConnection" tf:"private_service_connection,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetworkInterfaceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type NetworkInterfaceParameters struct {
}

type PrivateDNSZoneConfigsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	PrivateDNSZoneID *string `json:"privateDnsZoneId,omitempty" tf:"private_dns_zone_id,omitempty"`

	RecordSets []RecordSetsObservation `json:"recordSets,omitempty" tf:"record_sets,omitempty"`
}

type PrivateDNSZoneConfigsParameters struct {
}

type PrivateDNSZoneGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateDNSZoneGroupParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PrivateDNSZoneIds []*string `json:"privateDnsZoneIds" tf:"private_dns_zone_ids,omitempty"`
}

type PrivateServiceConnectionObservation struct {
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`
}

type PrivateServiceConnectionParameters struct {

	// +kubebuilder:validation:Required
	IsManualConnection *bool `json:"isManualConnection" tf:"is_manual_connection,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateConnectionResourceAlias *string `json:"privateConnectionResourceAlias,omitempty" tf:"private_connection_resource_alias,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateConnectionResourceID *string `json:"privateConnectionResourceId,omitempty" tf:"private_connection_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// +kubebuilder:validation:Optional
	SubresourceNames []*string `json:"subresourceNames,omitempty" tf:"subresource_names,omitempty"`
}

type RecordSetsObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TTL *int64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RecordSetsParameters struct {
}

// EndpointSpec defines the desired state of Endpoint
type EndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointParameters `json:"forProvider"`
}

// EndpointStatus defines the observed state of Endpoint.
type EndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Endpoint is the Schema for the Endpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointSpec   `json:"spec"`
	Status            EndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointList contains a list of Endpoints
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

// Repository type metadata.
var (
	EndpointKind             = "Endpoint"
	EndpointGroupKind        = schema.GroupKind{Group: Group, Kind: EndpointKind}.String()
	EndpointKindAPIVersion   = EndpointKind + "." + GroupVersion.String()
	EndpointGroupVersionKind = GroupVersion.WithKind(EndpointKind)
)

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
