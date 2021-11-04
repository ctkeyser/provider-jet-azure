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

type ApplicationRuleCollectionObservation struct {
}

type ApplicationRuleCollectionParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	AzureFirewallName *string `json:"azureFirewallName" tf:"azure_firewall_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`
}

type ProtocolObservation struct {
}

type ProtocolParameters struct {

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FqdnTags []*string `json:"fqdnTags,omitempty" tf:"fqdn_tags,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol []ProtocolParameters `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	SourceIPGroups []*string `json:"sourceIpGroups,omitempty" tf:"source_ip_groups,omitempty"`

	// +kubebuilder:validation:Optional
	TargetFqdns []*string `json:"targetFqdns,omitempty" tf:"target_fqdns,omitempty"`
}

// ApplicationRuleCollectionSpec defines the desired state of ApplicationRuleCollection
type ApplicationRuleCollectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationRuleCollectionParameters `json:"forProvider"`
}

// ApplicationRuleCollectionStatus defines the observed state of ApplicationRuleCollection.
type ApplicationRuleCollectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationRuleCollectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationRuleCollection is the Schema for the ApplicationRuleCollections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type ApplicationRuleCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationRuleCollectionSpec   `json:"spec"`
	Status            ApplicationRuleCollectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationRuleCollectionList contains a list of ApplicationRuleCollections
type ApplicationRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationRuleCollection `json:"items"`
}

// Repository type metadata.
var (
	ApplicationRuleCollectionKind             = "ApplicationRuleCollection"
	ApplicationRuleCollectionGroupKind        = schema.GroupKind{Group: Group, Kind: ApplicationRuleCollectionKind}.String()
	ApplicationRuleCollectionKindAPIVersion   = ApplicationRuleCollectionKind + "." + GroupVersion.String()
	ApplicationRuleCollectionGroupVersionKind = GroupVersion.WithKind(ApplicationRuleCollectionKind)
)

func init() {
	SchemeBuilder.Register(&ApplicationRuleCollection{}, &ApplicationRuleCollectionList{})
}
