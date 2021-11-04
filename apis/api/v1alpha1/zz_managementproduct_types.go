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

type ManagementProductObservation struct {
}

type ManagementProductParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Optional
	ApprovalRequired *bool `json:"approvalRequired,omitempty" tf:"approval_required,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Required
	ProductID *string `json:"productId" tf:"product_id,omitempty"`

	// +kubebuilder:validation:Required
	Published *bool `json:"published" tf:"published,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SubscriptionRequired *bool `json:"subscriptionRequired" tf:"subscription_required,omitempty"`

	// +kubebuilder:validation:Optional
	SubscriptionsLimit *int64 `json:"subscriptionsLimit,omitempty" tf:"subscriptions_limit,omitempty"`

	// +kubebuilder:validation:Optional
	Terms *string `json:"terms,omitempty" tf:"terms,omitempty"`
}

// ManagementProductSpec defines the desired state of ManagementProduct
type ManagementProductSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementProductParameters `json:"forProvider"`
}

// ManagementProductStatus defines the observed state of ManagementProduct.
type ManagementProductStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementProductObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementProduct is the Schema for the ManagementProducts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type ManagementProduct struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementProductSpec   `json:"spec"`
	Status            ManagementProductStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementProductList contains a list of ManagementProducts
type ManagementProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementProduct `json:"items"`
}

// Repository type metadata.
var (
	ManagementProductKind             = "ManagementProduct"
	ManagementProductGroupKind        = schema.GroupKind{Group: Group, Kind: ManagementProductKind}.String()
	ManagementProductKindAPIVersion   = ManagementProductKind + "." + GroupVersion.String()
	ManagementProductGroupVersionKind = GroupVersion.WithKind(ManagementProductKind)
)

func init() {
	SchemeBuilder.Register(&ManagementProduct{}, &ManagementProductList{})
}
