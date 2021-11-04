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

type GatewayConnectionObservation struct {
}

type GatewayConnectionParameters struct {

	// +kubebuilder:validation:Optional
	InternetSecurityEnabled *bool `json:"internetSecurityEnabled,omitempty" tf:"internet_security_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RemoteVpnSiteID *string `json:"remoteVpnSiteId" tf:"remote_vpn_site_id,omitempty"`

	// +kubebuilder:validation:Optional
	Routing []RoutingParameters `json:"routing,omitempty" tf:"routing,omitempty"`

	// +kubebuilder:validation:Required
	VpnGatewayID *string `json:"vpnGatewayId" tf:"vpn_gateway_id,omitempty"`

	// +kubebuilder:validation:Required
	VpnLink []VpnLinkParameters `json:"vpnLink" tf:"vpn_link,omitempty"`
}

type IpsecPolicyObservation struct {
}

type IpsecPolicyParameters struct {

	// +kubebuilder:validation:Required
	DhGroup *string `json:"dhGroup" tf:"dh_group,omitempty"`

	// +kubebuilder:validation:Required
	EncryptionAlgorithm *string `json:"encryptionAlgorithm" tf:"encryption_algorithm,omitempty"`

	// +kubebuilder:validation:Required
	IkeEncryptionAlgorithm *string `json:"ikeEncryptionAlgorithm" tf:"ike_encryption_algorithm,omitempty"`

	// +kubebuilder:validation:Required
	IkeIntegrityAlgorithm *string `json:"ikeIntegrityAlgorithm" tf:"ike_integrity_algorithm,omitempty"`

	// +kubebuilder:validation:Required
	IntegrityAlgorithm *string `json:"integrityAlgorithm" tf:"integrity_algorithm,omitempty"`

	// +kubebuilder:validation:Required
	PfsGroup *string `json:"pfsGroup" tf:"pfs_group,omitempty"`

	// +kubebuilder:validation:Required
	SaDataSizeKb *int64 `json:"saDataSizeKb" tf:"sa_data_size_kb,omitempty"`

	// +kubebuilder:validation:Required
	SaLifetimeSec *int64 `json:"saLifetimeSec" tf:"sa_lifetime_sec,omitempty"`
}

type RoutingObservation struct {
}

type RoutingParameters struct {

	// +kubebuilder:validation:Required
	AssociatedRouteTable *string `json:"associatedRouteTable" tf:"associated_route_table,omitempty"`

	// +kubebuilder:validation:Required
	PropagatedRouteTables []*string `json:"propagatedRouteTables" tf:"propagated_route_tables,omitempty"`
}

type VpnLinkObservation struct {
}

type VpnLinkParameters struct {

	// +kubebuilder:validation:Optional
	BandwidthMbps *int64 `json:"bandwidthMbps,omitempty" tf:"bandwidth_mbps,omitempty"`

	// +kubebuilder:validation:Optional
	BgpEnabled *bool `json:"bgpEnabled,omitempty" tf:"bgp_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IpsecPolicy []IpsecPolicyParameters `json:"ipsecPolicy,omitempty" tf:"ipsec_policy,omitempty"`

	// +kubebuilder:validation:Optional
	LocalAzureIPAddressEnabled *bool `json:"localAzureIpAddressEnabled,omitempty" tf:"local_azure_ip_address_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyBasedTrafficSelectorEnabled *bool `json:"policyBasedTrafficSelectorEnabled,omitempty" tf:"policy_based_traffic_selector_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	RatelimitEnabled *bool `json:"ratelimitEnabled,omitempty" tf:"ratelimit_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RouteWeight *int64 `json:"routeWeight,omitempty" tf:"route_weight,omitempty"`

	// +kubebuilder:validation:Optional
	SharedKey *string `json:"sharedKey,omitempty" tf:"shared_key,omitempty"`

	// +kubebuilder:validation:Required
	VpnSiteLinkID *string `json:"vpnSiteLinkId" tf:"vpn_site_link_id,omitempty"`
}

// GatewayConnectionSpec defines the desired state of GatewayConnection
type GatewayConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayConnectionParameters `json:"forProvider"`
}

// GatewayConnectionStatus defines the observed state of GatewayConnection.
type GatewayConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayConnection is the Schema for the GatewayConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type GatewayConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayConnectionSpec   `json:"spec"`
	Status            GatewayConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayConnectionList contains a list of GatewayConnections
type GatewayConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayConnection `json:"items"`
}

// Repository type metadata.
var (
	GatewayConnectionKind             = "GatewayConnection"
	GatewayConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: GatewayConnectionKind}.String()
	GatewayConnectionKindAPIVersion   = GatewayConnectionKind + "." + GroupVersion.String()
	GatewayConnectionGroupVersionKind = GroupVersion.WithKind(GatewayConnectionKind)
)

func init() {
	SchemeBuilder.Register(&GatewayConnection{}, &GatewayConnectionList{})
}
