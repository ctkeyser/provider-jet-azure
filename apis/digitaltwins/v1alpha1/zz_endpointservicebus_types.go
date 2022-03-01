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

type EndpointServiceBusObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EndpointServiceBusParameters struct {

	// +kubebuilder:validation:Optional
	DeadLetterStorageSecretSecretRef *v1.SecretKeySelector `json:"deadLetterStorageSecretSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	DigitalTwinsID *string `json:"digitalTwinsId" tf:"digital_twins_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ServiceBusPrimaryConnectionStringSecretRef v1.SecretKeySelector `json:"serviceBusPrimaryConnectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	ServiceBusSecondaryConnectionStringSecretRef v1.SecretKeySelector `json:"serviceBusSecondaryConnectionStringSecretRef" tf:"-"`
}

// EndpointServiceBusSpec defines the desired state of EndpointServiceBus
type EndpointServiceBusSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointServiceBusParameters `json:"forProvider"`
}

// EndpointServiceBusStatus defines the observed state of EndpointServiceBus.
type EndpointServiceBusStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointServiceBusObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointServiceBus is the Schema for the EndpointServiceBuss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type EndpointServiceBus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointServiceBusSpec   `json:"spec"`
	Status            EndpointServiceBusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointServiceBusList contains a list of EndpointServiceBuss
type EndpointServiceBusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointServiceBus `json:"items"`
}

// Repository type metadata.
var (
	EndpointServiceBus_Kind             = "EndpointServiceBus"
	EndpointServiceBus_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointServiceBus_Kind}.String()
	EndpointServiceBus_KindAPIVersion   = EndpointServiceBus_Kind + "." + CRDGroupVersion.String()
	EndpointServiceBus_GroupVersionKind = CRDGroupVersion.WithKind(EndpointServiceBus_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointServiceBus{}, &EndpointServiceBusList{})
}
