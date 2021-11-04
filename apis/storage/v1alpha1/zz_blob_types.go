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

type BlobObservation struct {
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type BlobParameters struct {

	// +kubebuilder:validation:Optional
	AccessTier *string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`

	// +kubebuilder:validation:Optional
	ContentMd5 *string `json:"contentMd5,omitempty" tf:"content_md5,omitempty"`

	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parallelism *int64 `json:"parallelism,omitempty" tf:"parallelism,omitempty"`

	// +kubebuilder:validation:Optional
	Size *int64 `json:"size,omitempty" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	SourceContent *string `json:"sourceContent,omitempty" tf:"source_content,omitempty"`

	// +kubebuilder:validation:Optional
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountName *string `json:"storageAccountName" tf:"storage_account_name,omitempty"`

	// +kubebuilder:validation:Required
	StorageContainerName *string `json:"storageContainerName" tf:"storage_container_name,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// BlobSpec defines the desired state of Blob
type BlobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BlobParameters `json:"forProvider"`
}

// BlobStatus defines the observed state of Blob.
type BlobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BlobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Blob is the Schema for the Blobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type Blob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BlobSpec   `json:"spec"`
	Status            BlobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlobList contains a list of Blobs
type BlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Blob `json:"items"`
}

// Repository type metadata.
var (
	BlobKind             = "Blob"
	BlobGroupKind        = schema.GroupKind{Group: Group, Kind: BlobKind}.String()
	BlobKindAPIVersion   = BlobKind + "." + GroupVersion.String()
	BlobGroupVersionKind = GroupVersion.WithKind(BlobKind)
)

func init() {
	SchemeBuilder.Register(&Blob{}, &BlobList{})
}
