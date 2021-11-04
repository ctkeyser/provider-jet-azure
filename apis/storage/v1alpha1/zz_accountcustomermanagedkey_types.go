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

type AccountCustomerManagedKeyObservation struct {
}

type AccountCustomerManagedKeyParameters struct {

	// +kubebuilder:validation:Required
	KeyName *string `json:"keyName" tf:"key_name,omitempty"`

	// +kubebuilder:validation:Required
	KeyVaultID *string `json:"keyVaultId" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVersion *string `json:"keyVersion,omitempty" tf:"key_version,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

// AccountCustomerManagedKeySpec defines the desired state of AccountCustomerManagedKey
type AccountCustomerManagedKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountCustomerManagedKeyParameters `json:"forProvider"`
}

// AccountCustomerManagedKeyStatus defines the observed state of AccountCustomerManagedKey.
type AccountCustomerManagedKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountCustomerManagedKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccountCustomerManagedKey is the Schema for the AccountCustomerManagedKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type AccountCustomerManagedKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountCustomerManagedKeySpec   `json:"spec"`
	Status            AccountCustomerManagedKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountCustomerManagedKeyList contains a list of AccountCustomerManagedKeys
type AccountCustomerManagedKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountCustomerManagedKey `json:"items"`
}

// Repository type metadata.
var (
	AccountCustomerManagedKeyKind             = "AccountCustomerManagedKey"
	AccountCustomerManagedKeyGroupKind        = schema.GroupKind{Group: Group, Kind: AccountCustomerManagedKeyKind}.String()
	AccountCustomerManagedKeyKindAPIVersion   = AccountCustomerManagedKeyKind + "." + GroupVersion.String()
	AccountCustomerManagedKeyGroupVersionKind = GroupVersion.WithKind(AccountCustomerManagedKeyKind)
)

func init() {
	SchemeBuilder.Register(&AccountCustomerManagedKey{}, &AccountCustomerManagedKeyList{})
}
