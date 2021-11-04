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

type ManagementApiOperationPolicyObservation struct {
}

type ManagementApiOperationPolicyParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	APIName *string `json:"apiName" tf:"api_name,omitempty"`

	// +kubebuilder:validation:Required
	OperationID *string `json:"operationId" tf:"operation_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// +kubebuilder:validation:Optional
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

// ManagementApiOperationPolicySpec defines the desired state of ManagementApiOperationPolicy
type ManagementApiOperationPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementApiOperationPolicyParameters `json:"forProvider"`
}

// ManagementApiOperationPolicyStatus defines the observed state of ManagementApiOperationPolicy.
type ManagementApiOperationPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementApiOperationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementApiOperationPolicy is the Schema for the ManagementApiOperationPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type ManagementApiOperationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementApiOperationPolicySpec   `json:"spec"`
	Status            ManagementApiOperationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementApiOperationPolicyList contains a list of ManagementApiOperationPolicys
type ManagementApiOperationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementApiOperationPolicy `json:"items"`
}

// Repository type metadata.
var (
	ManagementApiOperationPolicyKind             = "ManagementApiOperationPolicy"
	ManagementApiOperationPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: ManagementApiOperationPolicyKind}.String()
	ManagementApiOperationPolicyKindAPIVersion   = ManagementApiOperationPolicyKind + "." + GroupVersion.String()
	ManagementApiOperationPolicyGroupVersionKind = GroupVersion.WithKind(ManagementApiOperationPolicyKind)
)

func init() {
	SchemeBuilder.Register(&ManagementApiOperationPolicy{}, &ManagementApiOperationPolicyList{})
}
