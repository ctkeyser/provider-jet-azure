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

type IOTSecuritySolutionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTSecuritySolutionParameters struct {

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EventsToExport []*string `json:"eventsToExport,omitempty" tf:"events_to_export,omitempty"`

	// +kubebuilder:validation:Required
	IOTHubIds []*string `json:"iothubIds" tf:"iothub_ids,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Optional
	LogUnmaskedIpsEnabled *bool `json:"logUnmaskedIpsEnabled,omitempty" tf:"log_unmasked_ips_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	QueryForResources *string `json:"queryForResources,omitempty" tf:"query_for_resources,omitempty"`

	// +kubebuilder:validation:Optional
	QuerySubscriptionIds []*string `json:"querySubscriptionIds,omitempty" tf:"query_subscription_ids,omitempty"`

	// +kubebuilder:validation:Optional
	RecommendationsEnabled []RecommendationsEnabledParameters `json:"recommendationsEnabled,omitempty" tf:"recommendations_enabled,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RecommendationsEnabledObservation struct {
}

type RecommendationsEnabledParameters struct {

	// +kubebuilder:validation:Optional
	AcrAuthentication *bool `json:"acrAuthentication,omitempty" tf:"acr_authentication,omitempty"`

	// +kubebuilder:validation:Optional
	AgentSendUnutilizedMsg *bool `json:"agentSendUnutilizedMsg,omitempty" tf:"agent_send_unutilized_msg,omitempty"`

	// +kubebuilder:validation:Optional
	Baseline *bool `json:"baseline,omitempty" tf:"baseline,omitempty"`

	// +kubebuilder:validation:Optional
	EdgeHubMemOptimize *bool `json:"edgeHubMemOptimize,omitempty" tf:"edge_hub_mem_optimize,omitempty"`

	// +kubebuilder:validation:Optional
	EdgeLoggingOption *bool `json:"edgeLoggingOption,omitempty" tf:"edge_logging_option,omitempty"`

	// +kubebuilder:validation:Optional
	IPFilterDenyAll *bool `json:"ipFilterDenyAll,omitempty" tf:"ip_filter_deny_all,omitempty"`

	// +kubebuilder:validation:Optional
	IPFilterPermissiveRule *bool `json:"ipFilterPermissiveRule,omitempty" tf:"ip_filter_permissive_rule,omitempty"`

	// +kubebuilder:validation:Optional
	InconsistentModuleSettings *bool `json:"inconsistentModuleSettings,omitempty" tf:"inconsistent_module_settings,omitempty"`

	// +kubebuilder:validation:Optional
	InstallAgent *bool `json:"installAgent,omitempty" tf:"install_agent,omitempty"`

	// +kubebuilder:validation:Optional
	OpenPorts *bool `json:"openPorts,omitempty" tf:"open_ports,omitempty"`

	// +kubebuilder:validation:Optional
	PermissiveFirewallPolicy *bool `json:"permissiveFirewallPolicy,omitempty" tf:"permissive_firewall_policy,omitempty"`

	// +kubebuilder:validation:Optional
	PermissiveInputFirewallRules *bool `json:"permissiveInputFirewallRules,omitempty" tf:"permissive_input_firewall_rules,omitempty"`

	// +kubebuilder:validation:Optional
	PermissiveOutputFirewallRules *bool `json:"permissiveOutputFirewallRules,omitempty" tf:"permissive_output_firewall_rules,omitempty"`

	// +kubebuilder:validation:Optional
	PrivilegedDockerOptions *bool `json:"privilegedDockerOptions,omitempty" tf:"privileged_docker_options,omitempty"`

	// +kubebuilder:validation:Optional
	SharedCredentials *bool `json:"sharedCredentials,omitempty" tf:"shared_credentials,omitempty"`

	// +kubebuilder:validation:Optional
	VulnerableTLSCipherSuite *bool `json:"vulnerableTlsCipherSuite,omitempty" tf:"vulnerable_tls_cipher_suite,omitempty"`
}

// IOTSecuritySolutionSpec defines the desired state of IOTSecuritySolution
type IOTSecuritySolutionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTSecuritySolutionParameters `json:"forProvider"`
}

// IOTSecuritySolutionStatus defines the observed state of IOTSecuritySolution.
type IOTSecuritySolutionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTSecuritySolutionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTSecuritySolution is the Schema for the IOTSecuritySolutions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type IOTSecuritySolution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTSecuritySolutionSpec   `json:"spec"`
	Status            IOTSecuritySolutionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTSecuritySolutionList contains a list of IOTSecuritySolutions
type IOTSecuritySolutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTSecuritySolution `json:"items"`
}

// Repository type metadata.
var (
	IOTSecuritySolution_Kind             = "IOTSecuritySolution"
	IOTSecuritySolution_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTSecuritySolution_Kind}.String()
	IOTSecuritySolution_KindAPIVersion   = IOTSecuritySolution_Kind + "." + CRDGroupVersion.String()
	IOTSecuritySolution_GroupVersionKind = CRDGroupVersion.WithKind(IOTSecuritySolution_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTSecuritySolution{}, &IOTSecuritySolutionList{})
}
