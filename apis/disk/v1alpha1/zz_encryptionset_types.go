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

type EncryptionSetObservation struct {
}

type EncryptionSetParameters struct {

	// +kubebuilder:validation:Required
	Identity []IdentityParameters `json:"identity" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	KeyVaultKeyID *string `json:"keyVaultKeyId" tf:"key_vault_key_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// EncryptionSetSpec defines the desired state of EncryptionSet
type EncryptionSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EncryptionSetParameters `json:"forProvider"`
}

// EncryptionSetStatus defines the observed state of EncryptionSet.
type EncryptionSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EncryptionSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EncryptionSet is the Schema for the EncryptionSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type EncryptionSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EncryptionSetSpec   `json:"spec"`
	Status            EncryptionSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EncryptionSetList contains a list of EncryptionSets
type EncryptionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EncryptionSet `json:"items"`
}

// Repository type metadata.
var (
	EncryptionSetKind             = "EncryptionSet"
	EncryptionSetGroupKind        = schema.GroupKind{Group: Group, Kind: EncryptionSetKind}.String()
	EncryptionSetKindAPIVersion   = EncryptionSetKind + "." + GroupVersion.String()
	EncryptionSetGroupVersionKind = GroupVersion.WithKind(EncryptionSetKind)
)

func init() {
	SchemeBuilder.Register(&EncryptionSet{}, &EncryptionSetList{})
}
