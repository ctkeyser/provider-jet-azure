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

type SharedImageGalleryObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UniqueName *string `json:"uniqueName,omitempty" tf:"unique_name,omitempty"`
}

type SharedImageGalleryParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

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

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SharedImageGallerySpec defines the desired state of SharedImageGallery
type SharedImageGallerySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharedImageGalleryParameters `json:"forProvider"`
}

// SharedImageGalleryStatus defines the observed state of SharedImageGallery.
type SharedImageGalleryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharedImageGalleryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SharedImageGallery is the Schema for the SharedImageGallerys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SharedImageGallery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SharedImageGallerySpec   `json:"spec"`
	Status            SharedImageGalleryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedImageGalleryList contains a list of SharedImageGallerys
type SharedImageGalleryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedImageGallery `json:"items"`
}

// Repository type metadata.
var (
	SharedImageGallery_Kind             = "SharedImageGallery"
	SharedImageGallery_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharedImageGallery_Kind}.String()
	SharedImageGallery_KindAPIVersion   = SharedImageGallery_Kind + "." + CRDGroupVersion.String()
	SharedImageGallery_GroupVersionKind = CRDGroupVersion.WithKind(SharedImageGallery_Kind)
)

func init() {
	SchemeBuilder.Register(&SharedImageGallery{}, &SharedImageGalleryList{})
}
