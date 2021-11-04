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

type ApplicationInsightsObservation struct {
}

type ApplicationInsightsParameters struct {

	// +kubebuilder:validation:Required
	InstrumentationKeySecretRef v1.SecretKeySelector `json:"instrumentationKeySecretRef" tf:"-"`
}

type EventhubObservation struct {
}

type EventhubParameters struct {

	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type ManagementLoggerObservation struct {
}

type ManagementLoggerParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Optional
	ApplicationInsights []ApplicationInsightsParameters `json:"applicationInsights,omitempty" tf:"application_insights,omitempty"`

	// +kubebuilder:validation:Optional
	Buffered *bool `json:"buffered,omitempty" tf:"buffered,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Eventhub []EventhubParameters `json:"eventhub,omitempty" tf:"eventhub,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

// ManagementLoggerSpec defines the desired state of ManagementLogger
type ManagementLoggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementLoggerParameters `json:"forProvider"`
}

// ManagementLoggerStatus defines the observed state of ManagementLogger.
type ManagementLoggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementLoggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementLogger is the Schema for the ManagementLoggers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type ManagementLogger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementLoggerSpec   `json:"spec"`
	Status            ManagementLoggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementLoggerList contains a list of ManagementLoggers
type ManagementLoggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementLogger `json:"items"`
}

// Repository type metadata.
var (
	ManagementLoggerKind             = "ManagementLogger"
	ManagementLoggerGroupKind        = schema.GroupKind{Group: Group, Kind: ManagementLoggerKind}.String()
	ManagementLoggerKindAPIVersion   = ManagementLoggerKind + "." + GroupVersion.String()
	ManagementLoggerGroupVersionKind = GroupVersion.WithKind(ManagementLoggerKind)
)

func init() {
	SchemeBuilder.Register(&ManagementLogger{}, &ManagementLoggerList{})
}
