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

type EndpointEventGridObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EndpointEventGridParameters struct {

	// +kubebuilder:validation:Optional
	DeadLetterStorageSecret *string `json:"deadLetterStorageSecret,omitempty" tf:"dead_letter_storage_secret,omitempty"`

	// +kubebuilder:validation:Required
	DigitalTwinsID *string `json:"digitalTwinsId" tf:"digital_twins_id,omitempty"`

	// +kubebuilder:validation:Required
	EventGridTopicEndpoint *string `json:"eventgridTopicEndpoint" tf:"eventgrid_topic_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	EventGridTopicPrimaryAccessKey *string `json:"eventgridTopicPrimaryAccessKey" tf:"eventgrid_topic_primary_access_key,omitempty"`

	// +kubebuilder:validation:Required
	EventGridTopicSecondaryAccessKey *string `json:"eventgridTopicSecondaryAccessKey" tf:"eventgrid_topic_secondary_access_key,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// EndpointEventGridSpec defines the desired state of EndpointEventGrid
type EndpointEventGridSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointEventGridParameters `json:"forProvider"`
}

// EndpointEventGridStatus defines the observed state of EndpointEventGrid.
type EndpointEventGridStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointEventGridObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointEventGrid is the Schema for the EndpointEventGrids API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type EndpointEventGrid struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointEventGridSpec   `json:"spec"`
	Status            EndpointEventGridStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointEventGridList contains a list of EndpointEventGrids
type EndpointEventGridList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointEventGrid `json:"items"`
}

// Repository type metadata.
var (
	EndpointEventGrid_Kind             = "EndpointEventGrid"
	EndpointEventGrid_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointEventGrid_Kind}.String()
	EndpointEventGrid_KindAPIVersion   = EndpointEventGrid_Kind + "." + CRDGroupVersion.String()
	EndpointEventGrid_GroupVersionKind = CRDGroupVersion.WithKind(EndpointEventGrid_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointEventGrid{}, &EndpointEventGridList{})
}
