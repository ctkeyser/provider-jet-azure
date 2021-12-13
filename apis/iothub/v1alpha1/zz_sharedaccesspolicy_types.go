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

type SharedAccessPolicyObservation_2 struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SharedAccessPolicyParameters_2 struct {

	// +kubebuilder:validation:Optional
	DeviceConnect *bool `json:"deviceConnect,omitempty" tf:"device_connect,omitempty"`

	// +crossplane:generate:reference:type=IOTHub
	// +kubebuilder:validation:Optional
	IothubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Optional
	IothubNameRef *v1.Reference `json:"iothubNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IothubNameSelector *v1.Selector `json:"iothubNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RegistryRead *bool `json:"registryRead,omitempty" tf:"registry_read,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryWrite *bool `json:"registryWrite,omitempty" tf:"registry_write,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceConnect *bool `json:"serviceConnect,omitempty" tf:"service_connect,omitempty"`
}

// SharedAccessPolicySpec defines the desired state of SharedAccessPolicy
type SharedAccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharedAccessPolicyParameters_2 `json:"forProvider"`
}

// SharedAccessPolicyStatus defines the observed state of SharedAccessPolicy.
type SharedAccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharedAccessPolicyObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SharedAccessPolicy is the Schema for the SharedAccessPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SharedAccessPolicySpec   `json:"spec"`
	Status            SharedAccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedAccessPolicyList contains a list of SharedAccessPolicys
type SharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedAccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	SharedAccessPolicy_Kind             = "SharedAccessPolicy"
	SharedAccessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharedAccessPolicy_Kind}.String()
	SharedAccessPolicy_KindAPIVersion   = SharedAccessPolicy_Kind + "." + CRDGroupVersion.String()
	SharedAccessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SharedAccessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SharedAccessPolicy{}, &SharedAccessPolicyList{})
}