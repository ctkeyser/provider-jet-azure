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

type DiagnosticSettingLogObservation struct {
}

type DiagnosticSettingLogParameters struct {

	// +kubebuilder:validation:Required
	Category *string `json:"category" tf:"category,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionPolicy []LogRetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type DiagnosticSettingObservation struct {
}

type DiagnosticSettingParameters struct {

	// +kubebuilder:validation:Optional
	EventhubAuthorizationRuleID *string `json:"eventhubAuthorizationRuleId,omitempty" tf:"eventhub_authorization_rule_id,omitempty"`

	// +kubebuilder:validation:Optional
	EventhubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// +kubebuilder:validation:Optional
	Log []DiagnosticSettingLogParameters `json:"log,omitempty" tf:"log,omitempty"`

	// +kubebuilder:validation:Optional
	LogAnalyticsDestinationType *string `json:"logAnalyticsDestinationType,omitempty" tf:"log_analytics_destination_type,omitempty"`

	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Optional
	Metric []MetricParameters `json:"metric,omitempty" tf:"metric,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Required
	TargetResourceID *string `json:"targetResourceId" tf:"target_resource_id,omitempty"`
}

type LogRetentionPolicyObservation struct {
}

type LogRetentionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type MetricObservation struct {
}

type MetricParameters struct {

	// +kubebuilder:validation:Required
	Category *string `json:"category" tf:"category,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionPolicy []MetricRetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type MetricRetentionPolicyObservation struct {
}

type MetricRetentionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

// DiagnosticSettingSpec defines the desired state of DiagnosticSetting
type DiagnosticSettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiagnosticSettingParameters `json:"forProvider"`
}

// DiagnosticSettingStatus defines the observed state of DiagnosticSetting.
type DiagnosticSettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiagnosticSettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiagnosticSetting is the Schema for the DiagnosticSettings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type DiagnosticSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiagnosticSettingSpec   `json:"spec"`
	Status            DiagnosticSettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiagnosticSettingList contains a list of DiagnosticSettings
type DiagnosticSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiagnosticSetting `json:"items"`
}

// Repository type metadata.
var (
	DiagnosticSettingKind             = "DiagnosticSetting"
	DiagnosticSettingGroupKind        = schema.GroupKind{Group: Group, Kind: DiagnosticSettingKind}.String()
	DiagnosticSettingKindAPIVersion   = DiagnosticSettingKind + "." + GroupVersion.String()
	DiagnosticSettingGroupVersionKind = GroupVersion.WithKind(DiagnosticSettingKind)
)

func init() {
	SchemeBuilder.Register(&DiagnosticSetting{}, &DiagnosticSettingList{})
}
