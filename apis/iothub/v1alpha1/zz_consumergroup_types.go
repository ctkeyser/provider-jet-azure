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

type ConsumerGroupObservation struct {
}

type ConsumerGroupParameters struct {

	// +kubebuilder:validation:Required
	EventhubEndpointName *string `json:"eventhubEndpointName" tf:"eventhub_endpoint_name,omitempty"`

	// +kubebuilder:validation:Required
	IothubName *string `json:"iothubName" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// ConsumerGroupSpec defines the desired state of ConsumerGroup
type ConsumerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConsumerGroupParameters `json:"forProvider"`
}

// ConsumerGroupStatus defines the observed state of ConsumerGroup.
type ConsumerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConsumerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConsumerGroup is the Schema for the ConsumerGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type ConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConsumerGroupSpec   `json:"spec"`
	Status            ConsumerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConsumerGroupList contains a list of ConsumerGroups
type ConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsumerGroup `json:"items"`
}

// Repository type metadata.
var (
	ConsumerGroupKind             = "ConsumerGroup"
	ConsumerGroupGroupKind        = schema.GroupKind{Group: Group, Kind: ConsumerGroupKind}.String()
	ConsumerGroupKindAPIVersion   = ConsumerGroupKind + "." + GroupVersion.String()
	ConsumerGroupGroupVersionKind = GroupVersion.WithKind(ConsumerGroupKind)
)

func init() {
	SchemeBuilder.Register(&ConsumerGroup{}, &ConsumerGroupList{})
}
