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

type RoleAssignmentObservation struct {
}

type RoleAssignmentParameters struct {

	// +kubebuilder:validation:Required
	PrincipalID *string `json:"principalId" tf:"principal_id,omitempty"`

	// +kubebuilder:validation:Required
	RoleName *string `json:"roleName" tf:"role_name,omitempty"`

	// +kubebuilder:validation:Optional
	SynapseSparkPoolID *string `json:"synapseSparkPoolId,omitempty" tf:"synapse_spark_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`
}

// RoleAssignmentSpec defines the desired state of RoleAssignment
type RoleAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleAssignmentParameters `json:"forProvider"`
}

// RoleAssignmentStatus defines the observed state of RoleAssignment.
type RoleAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignment is the Schema for the RoleAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignmentSpec   `json:"spec"`
	Status            RoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignmentList contains a list of RoleAssignments
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	RoleAssignmentKind             = "RoleAssignment"
	RoleAssignmentGroupKind        = schema.GroupKind{Group: Group, Kind: RoleAssignmentKind}.String()
	RoleAssignmentKindAPIVersion   = RoleAssignmentKind + "." + GroupVersion.String()
	RoleAssignmentGroupVersionKind = GroupVersion.WithKind(RoleAssignmentKind)
)

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
