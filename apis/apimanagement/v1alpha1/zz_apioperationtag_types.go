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

type APIOperationTagObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APIOperationTagParameters struct {

	// +kubebuilder:validation:Required
	APIOperationID *string `json:"apiOperationId" tf:"api_operation_id,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// APIOperationTagSpec defines the desired state of APIOperationTag
type APIOperationTagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIOperationTagParameters `json:"forProvider"`
}

// APIOperationTagStatus defines the observed state of APIOperationTag.
type APIOperationTagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIOperationTagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIOperationTag is the Schema for the APIOperationTags API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type APIOperationTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIOperationTagSpec   `json:"spec"`
	Status            APIOperationTagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIOperationTagList contains a list of APIOperationTags
type APIOperationTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIOperationTag `json:"items"`
}

// Repository type metadata.
var (
	APIOperationTag_Kind             = "APIOperationTag"
	APIOperationTag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIOperationTag_Kind}.String()
	APIOperationTag_KindAPIVersion   = APIOperationTag_Kind + "." + CRDGroupVersion.String()
	APIOperationTag_GroupVersionKind = CRDGroupVersion.WithKind(APIOperationTag_Kind)
)

func init() {
	SchemeBuilder.Register(&APIOperationTag{}, &APIOperationTagList{})
}