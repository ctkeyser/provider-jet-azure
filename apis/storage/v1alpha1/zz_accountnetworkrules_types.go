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

type AccountNetworkRulesObservation struct {
}

type AccountNetworkRulesParameters struct {

	// +kubebuilder:validation:Optional
	Bypass []*string `json:"bypass,omitempty" tf:"bypass,omitempty"`

	// +kubebuilder:validation:Required
	DefaultAction *string `json:"defaultAction" tf:"default_action,omitempty"`

	// +kubebuilder:validation:Optional
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateLinkAccess []AccountNetworkRulesPrivateLinkAccessParameters `json:"privateLinkAccess,omitempty" tf:"private_link_access,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIds []*string `json:"virtualNetworkSubnetIds,omitempty" tf:"virtual_network_subnet_ids,omitempty"`
}

type AccountNetworkRulesPrivateLinkAccessObservation struct {
}

type AccountNetworkRulesPrivateLinkAccessParameters struct {

	// +kubebuilder:validation:Required
	EndpointResourceID *string `json:"endpointResourceId" tf:"endpoint_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointTenantID *string `json:"endpointTenantId,omitempty" tf:"endpoint_tenant_id,omitempty"`
}

// AccountNetworkRulesSpec defines the desired state of AccountNetworkRules
type AccountNetworkRulesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountNetworkRulesParameters `json:"forProvider"`
}

// AccountNetworkRulesStatus defines the observed state of AccountNetworkRules.
type AccountNetworkRulesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountNetworkRulesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccountNetworkRules is the Schema for the AccountNetworkRuless API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type AccountNetworkRules struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountNetworkRulesSpec   `json:"spec"`
	Status            AccountNetworkRulesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountNetworkRulesList contains a list of AccountNetworkRuless
type AccountNetworkRulesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountNetworkRules `json:"items"`
}

// Repository type metadata.
var (
	AccountNetworkRulesKind             = "AccountNetworkRules"
	AccountNetworkRulesGroupKind        = schema.GroupKind{Group: Group, Kind: AccountNetworkRulesKind}.String()
	AccountNetworkRulesKindAPIVersion   = AccountNetworkRulesKind + "." + GroupVersion.String()
	AccountNetworkRulesGroupVersionKind = GroupVersion.WithKind(AccountNetworkRulesKind)
)

func init() {
	SchemeBuilder.Register(&AccountNetworkRules{}, &AccountNetworkRulesList{})
}
