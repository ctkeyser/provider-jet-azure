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

type IdentityProviderAADB2CObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IdentityProviderAADB2CParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	AllowedTenant *string `json:"allowedTenant" tf:"allowed_tenant,omitempty"`

	// +kubebuilder:validation:Required
	Authority *string `json:"authority" tf:"authority,omitempty"`

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	PasswordResetPolicy *string `json:"passwordResetPolicy,omitempty" tf:"password_reset_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ProfileEditingPolicy *string `json:"profileEditingPolicy,omitempty" tf:"profile_editing_policy,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SigninPolicy *string `json:"signinPolicy" tf:"signin_policy,omitempty"`

	// +kubebuilder:validation:Required
	SigninTenant *string `json:"signinTenant" tf:"signin_tenant,omitempty"`

	// +kubebuilder:validation:Required
	SignupPolicy *string `json:"signupPolicy" tf:"signup_policy,omitempty"`
}

// IdentityProviderAADB2CSpec defines the desired state of IdentityProviderAADB2C
type IdentityProviderAADB2CSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderAADB2CParameters `json:"forProvider"`
}

// IdentityProviderAADB2CStatus defines the observed state of IdentityProviderAADB2C.
type IdentityProviderAADB2CStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderAADB2CObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderAADB2C is the Schema for the IdentityProviderAADB2Cs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type IdentityProviderAADB2C struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IdentityProviderAADB2CSpec   `json:"spec"`
	Status            IdentityProviderAADB2CStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderAADB2CList contains a list of IdentityProviderAADB2Cs
type IdentityProviderAADB2CList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderAADB2C `json:"items"`
}

// Repository type metadata.
var (
	IdentityProviderAADB2C_Kind             = "IdentityProviderAADB2C"
	IdentityProviderAADB2C_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProviderAADB2C_Kind}.String()
	IdentityProviderAADB2C_KindAPIVersion   = IdentityProviderAADB2C_Kind + "." + CRDGroupVersion.String()
	IdentityProviderAADB2C_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProviderAADB2C_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProviderAADB2C{}, &IdentityProviderAADB2CList{})
}
