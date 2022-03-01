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

type MSSQLServerSecurityAlertPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MSSQLServerSecurityAlertPolicyParameters struct {

	// +kubebuilder:validation:Optional
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAccountAdmins *bool `json:"emailAccountAdmins,omitempty" tf:"email_account_admins,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// +kubebuilder:validation:Required
	ServerName *string `json:"serverName" tf:"server_name,omitempty"`

	// +kubebuilder:validation:Required
	State *string `json:"state" tf:"state,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

// MSSQLServerSecurityAlertPolicySpec defines the desired state of MSSQLServerSecurityAlertPolicy
type MSSQLServerSecurityAlertPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLServerSecurityAlertPolicyParameters `json:"forProvider"`
}

// MSSQLServerSecurityAlertPolicyStatus defines the observed state of MSSQLServerSecurityAlertPolicy.
type MSSQLServerSecurityAlertPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLServerSecurityAlertPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerSecurityAlertPolicy is the Schema for the MSSQLServerSecurityAlertPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type MSSQLServerSecurityAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLServerSecurityAlertPolicySpec   `json:"spec"`
	Status            MSSQLServerSecurityAlertPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerSecurityAlertPolicyList contains a list of MSSQLServerSecurityAlertPolicys
type MSSQLServerSecurityAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLServerSecurityAlertPolicy `json:"items"`
}

// Repository type metadata.
var (
	MSSQLServerSecurityAlertPolicy_Kind             = "MSSQLServerSecurityAlertPolicy"
	MSSQLServerSecurityAlertPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLServerSecurityAlertPolicy_Kind}.String()
	MSSQLServerSecurityAlertPolicy_KindAPIVersion   = MSSQLServerSecurityAlertPolicy_Kind + "." + CRDGroupVersion.String()
	MSSQLServerSecurityAlertPolicy_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLServerSecurityAlertPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLServerSecurityAlertPolicy{}, &MSSQLServerSecurityAlertPolicyList{})
}
