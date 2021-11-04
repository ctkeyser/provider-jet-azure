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

type PrivateLinkHubObservation struct {
}

type PrivateLinkHubParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PrivateLinkHubSpec defines the desired state of PrivateLinkHub
type PrivateLinkHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateLinkHubParameters `json:"forProvider"`
}

// PrivateLinkHubStatus defines the observed state of PrivateLinkHub.
type PrivateLinkHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateLinkHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLinkHub is the Schema for the PrivateLinkHubs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type PrivateLinkHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateLinkHubSpec   `json:"spec"`
	Status            PrivateLinkHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLinkHubList contains a list of PrivateLinkHubs
type PrivateLinkHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateLinkHub `json:"items"`
}

// Repository type metadata.
var (
	PrivateLinkHubKind             = "PrivateLinkHub"
	PrivateLinkHubGroupKind        = schema.GroupKind{Group: Group, Kind: PrivateLinkHubKind}.String()
	PrivateLinkHubKindAPIVersion   = PrivateLinkHubKind + "." + GroupVersion.String()
	PrivateLinkHubGroupVersionKind = GroupVersion.WithKind(PrivateLinkHubKind)
)

func init() {
	SchemeBuilder.Register(&PrivateLinkHub{}, &PrivateLinkHubList{})
}
