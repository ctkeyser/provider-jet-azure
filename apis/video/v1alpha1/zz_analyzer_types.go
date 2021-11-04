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

type AnalyzerObservation struct {
}

type AnalyzerParameters struct {

	// +kubebuilder:validation:Required
	Identity []IdentityParameters `json:"identity" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccount []StorageAccountParameters `json:"storageAccount" tf:"storage_account,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityObservation struct {
}

type IdentityParameters struct {

	// +kubebuilder:validation:Required
	IdentityIds []*string `json:"identityIds" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type StorageAccountObservation struct {
}

type StorageAccountParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	UserAssignedIdentityID *string `json:"userAssignedIdentityId" tf:"user_assigned_identity_id,omitempty"`
}

// AnalyzerSpec defines the desired state of Analyzer
type AnalyzerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyzerParameters `json:"forProvider"`
}

// AnalyzerStatus defines the observed state of Analyzer.
type AnalyzerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyzerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Analyzer is the Schema for the Analyzers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type Analyzer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyzerSpec   `json:"spec"`
	Status            AnalyzerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyzerList contains a list of Analyzers
type AnalyzerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Analyzer `json:"items"`
}

// Repository type metadata.
var (
	AnalyzerKind             = "Analyzer"
	AnalyzerGroupKind        = schema.GroupKind{Group: Group, Kind: AnalyzerKind}.String()
	AnalyzerKindAPIVersion   = AnalyzerKind + "." + GroupVersion.String()
	AnalyzerGroupVersionKind = GroupVersion.WithKind(AnalyzerKind)
)

func init() {
	SchemeBuilder.Register(&Analyzer{}, &AnalyzerList{})
}
