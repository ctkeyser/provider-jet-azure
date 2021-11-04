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

type ShareDirectoryObservation struct {
}

type ShareDirectoryParameters struct {

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ShareName *string `json:"shareName" tf:"share_name,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountName *string `json:"storageAccountName" tf:"storage_account_name,omitempty"`
}

// ShareDirectorySpec defines the desired state of ShareDirectory
type ShareDirectorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ShareDirectoryParameters `json:"forProvider"`
}

// ShareDirectoryStatus defines the observed state of ShareDirectory.
type ShareDirectoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ShareDirectoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ShareDirectory is the Schema for the ShareDirectorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type ShareDirectory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ShareDirectorySpec   `json:"spec"`
	Status            ShareDirectoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShareDirectoryList contains a list of ShareDirectorys
type ShareDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShareDirectory `json:"items"`
}

// Repository type metadata.
var (
	ShareDirectoryKind             = "ShareDirectory"
	ShareDirectoryGroupKind        = schema.GroupKind{Group: Group, Kind: ShareDirectoryKind}.String()
	ShareDirectoryKindAPIVersion   = ShareDirectoryKind + "." + GroupVersion.String()
	ShareDirectoryGroupVersionKind = GroupVersion.WithKind(ShareDirectoryKind)
)

func init() {
	SchemeBuilder.Register(&ShareDirectory{}, &ShareDirectoryList{})
}
