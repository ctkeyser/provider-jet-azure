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

type VideoAnalyzerIdentityObservation struct {
}

type VideoAnalyzerIdentityParameters struct {

	// +kubebuilder:validation:Required
	IdentityIds []*string `json:"identityIds" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type VideoAnalyzerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VideoAnalyzerParameters struct {

	// +kubebuilder:validation:Required
	Identity []VideoAnalyzerIdentityParameters `json:"identity" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	StorageAccount []VideoAnalyzerStorageAccountParameters `json:"storageAccount" tf:"storage_account,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VideoAnalyzerStorageAccountObservation struct {
}

type VideoAnalyzerStorageAccountParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	UserAssignedIdentityID *string `json:"userAssignedIdentityId" tf:"user_assigned_identity_id,omitempty"`
}

// VideoAnalyzerSpec defines the desired state of VideoAnalyzer
type VideoAnalyzerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VideoAnalyzerParameters `json:"forProvider"`
}

// VideoAnalyzerStatus defines the observed state of VideoAnalyzer.
type VideoAnalyzerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VideoAnalyzerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VideoAnalyzer is the Schema for the VideoAnalyzers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VideoAnalyzer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VideoAnalyzerSpec   `json:"spec"`
	Status            VideoAnalyzerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VideoAnalyzerList contains a list of VideoAnalyzers
type VideoAnalyzerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VideoAnalyzer `json:"items"`
}

// Repository type metadata.
var (
	VideoAnalyzer_Kind             = "VideoAnalyzer"
	VideoAnalyzer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VideoAnalyzer_Kind}.String()
	VideoAnalyzer_KindAPIVersion   = VideoAnalyzer_Kind + "." + CRDGroupVersion.String()
	VideoAnalyzer_GroupVersionKind = CRDGroupVersion.WithKind(VideoAnalyzer_Kind)
)

func init() {
	SchemeBuilder.Register(&VideoAnalyzer{}, &VideoAnalyzerList{})
}