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

type FlexibleServerConfigurationObservation struct {
}

type FlexibleServerConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ServerID *string `json:"serverId" tf:"server_id,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// FlexibleServerConfigurationSpec defines the desired state of FlexibleServerConfiguration
type FlexibleServerConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerConfigurationParameters `json:"forProvider"`
}

// FlexibleServerConfigurationStatus defines the observed state of FlexibleServerConfiguration.
type FlexibleServerConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerConfiguration is the Schema for the FlexibleServerConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type FlexibleServerConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServerConfigurationSpec   `json:"spec"`
	Status            FlexibleServerConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerConfigurationList contains a list of FlexibleServerConfigurations
type FlexibleServerConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServerConfiguration `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServerConfigurationKind             = "FlexibleServerConfiguration"
	FlexibleServerConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: FlexibleServerConfigurationKind}.String()
	FlexibleServerConfigurationKindAPIVersion   = FlexibleServerConfigurationKind + "." + GroupVersion.String()
	FlexibleServerConfigurationGroupVersionKind = GroupVersion.WithKind(FlexibleServerConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServerConfiguration{}, &FlexibleServerConfigurationList{})
}
