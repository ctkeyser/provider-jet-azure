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

type ManagementUserObservation struct {
}

type ManagementUserParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Optional
	Confirmation *string `json:"confirmation,omitempty" tf:"confirmation,omitempty"`

	// +kubebuilder:validation:Required
	Email *string `json:"email" tf:"email,omitempty"`

	// +kubebuilder:validation:Required
	FirstName *string `json:"firstName" tf:"first_name,omitempty"`

	// +kubebuilder:validation:Required
	LastName *string `json:"lastName" tf:"last_name,omitempty"`

	// +kubebuilder:validation:Optional
	Note *string `json:"note,omitempty" tf:"note,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +kubebuilder:validation:Required
	UserID *string `json:"userId" tf:"user_id,omitempty"`
}

// ManagementUserSpec defines the desired state of ManagementUser
type ManagementUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementUserParameters `json:"forProvider"`
}

// ManagementUserStatus defines the observed state of ManagementUser.
type ManagementUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementUser is the Schema for the ManagementUsers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type ManagementUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementUserSpec   `json:"spec"`
	Status            ManagementUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementUserList contains a list of ManagementUsers
type ManagementUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementUser `json:"items"`
}

// Repository type metadata.
var (
	ManagementUserKind             = "ManagementUser"
	ManagementUserGroupKind        = schema.GroupKind{Group: Group, Kind: ManagementUserKind}.String()
	ManagementUserKindAPIVersion   = ManagementUserKind + "." + GroupVersion.String()
	ManagementUserGroupVersionKind = GroupVersion.WithKind(ManagementUserKind)
)

func init() {
	SchemeBuilder.Register(&ManagementUser{}, &ManagementUserList{})
}
