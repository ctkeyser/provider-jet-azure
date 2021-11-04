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

type AccessObservation struct {
}

type AccessParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AccessSpec defines the desired state of Access
type AccessSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessParameters `json:"forProvider"`
}

// AccessStatus defines the observed state of Access.
type AccessStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Access is the Schema for the Accesss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type Access struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessSpec   `json:"spec"`
	Status            AccessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessList contains a list of Accesss
type AccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Access `json:"items"`
}

// Repository type metadata.
var (
	AccessKind             = "Access"
	AccessGroupKind        = schema.GroupKind{Group: Group, Kind: AccessKind}.String()
	AccessKindAPIVersion   = AccessKind + "." + GroupVersion.String()
	AccessGroupVersionKind = GroupVersion.WithKind(AccessKind)
)

func init() {
	SchemeBuilder.Register(&Access{}, &AccessList{})
}
