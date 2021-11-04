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

type ShareDatasetDataLakeGen1Observation struct {
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}

type ShareDatasetDataLakeGen1Parameters struct {

	// +kubebuilder:validation:Required
	DataLakeStoreID *string `json:"dataLakeStoreId" tf:"data_lake_store_id,omitempty"`

	// +kubebuilder:validation:Required
	DataShareID *string `json:"dataShareId" tf:"data_share_id,omitempty"`

	// +kubebuilder:validation:Optional
	FileName *string `json:"fileName,omitempty" tf:"file_name,omitempty"`

	// +kubebuilder:validation:Required
	FolderPath *string `json:"folderPath" tf:"folder_path,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// ShareDatasetDataLakeGen1Spec defines the desired state of ShareDatasetDataLakeGen1
type ShareDatasetDataLakeGen1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ShareDatasetDataLakeGen1Parameters `json:"forProvider"`
}

// ShareDatasetDataLakeGen1Status defines the observed state of ShareDatasetDataLakeGen1.
type ShareDatasetDataLakeGen1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ShareDatasetDataLakeGen1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ShareDatasetDataLakeGen1 is the Schema for the ShareDatasetDataLakeGen1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type ShareDatasetDataLakeGen1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ShareDatasetDataLakeGen1Spec   `json:"spec"`
	Status            ShareDatasetDataLakeGen1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShareDatasetDataLakeGen1List contains a list of ShareDatasetDataLakeGen1s
type ShareDatasetDataLakeGen1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShareDatasetDataLakeGen1 `json:"items"`
}

// Repository type metadata.
var (
	ShareDatasetDataLakeGen1Kind             = "ShareDatasetDataLakeGen1"
	ShareDatasetDataLakeGen1GroupKind        = schema.GroupKind{Group: Group, Kind: ShareDatasetDataLakeGen1Kind}.String()
	ShareDatasetDataLakeGen1KindAPIVersion   = ShareDatasetDataLakeGen1Kind + "." + GroupVersion.String()
	ShareDatasetDataLakeGen1GroupVersionKind = GroupVersion.WithKind(ShareDatasetDataLakeGen1Kind)
)

func init() {
	SchemeBuilder.Register(&ShareDatasetDataLakeGen1{}, &ShareDatasetDataLakeGen1List{})
}
