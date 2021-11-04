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

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LearningComputeClusterObservation struct {
}

type LearningComputeClusterParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceId" tf:"machine_learning_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	SSH []SSHParameters `json:"ssh,omitempty" tf:"ssh,omitempty"`

	// +kubebuilder:validation:Optional
	SSHPublicAccessEnabled *bool `json:"sshPublicAccessEnabled,omitempty" tf:"ssh_public_access_enabled,omitempty"`

	// +kubebuilder:validation:Required
	ScaleSettings []ScaleSettingsParameters `json:"scaleSettings" tf:"scale_settings,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetResourceID *string `json:"subnetResourceId,omitempty" tf:"subnet_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VMPriority *string `json:"vmPriority" tf:"vm_priority,omitempty"`

	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`
}

type SSHObservation struct {
}

type SSHParameters struct {

	// +kubebuilder:validation:Optional
	AdminPassword *string `json:"adminPassword,omitempty" tf:"admin_password,omitempty"`

	// +kubebuilder:validation:Required
	AdminUsername *string `json:"adminUsername" tf:"admin_username,omitempty"`

	// +kubebuilder:validation:Optional
	KeyValue *string `json:"keyValue,omitempty" tf:"key_value,omitempty"`
}

type ScaleSettingsObservation struct {
}

type ScaleSettingsParameters struct {

	// +kubebuilder:validation:Required
	MaxNodeCount *int64 `json:"maxNodeCount" tf:"max_node_count,omitempty"`

	// +kubebuilder:validation:Required
	MinNodeCount *int64 `json:"minNodeCount" tf:"min_node_count,omitempty"`

	// +kubebuilder:validation:Required
	ScaleDownNodesAfterIdleDuration *string `json:"scaleDownNodesAfterIdleDuration" tf:"scale_down_nodes_after_idle_duration,omitempty"`
}

// LearningComputeClusterSpec defines the desired state of LearningComputeCluster
type LearningComputeClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LearningComputeClusterParameters `json:"forProvider"`
}

// LearningComputeClusterStatus defines the observed state of LearningComputeCluster.
type LearningComputeClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LearningComputeClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LearningComputeCluster is the Schema for the LearningComputeClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type LearningComputeCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LearningComputeClusterSpec   `json:"spec"`
	Status            LearningComputeClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LearningComputeClusterList contains a list of LearningComputeClusters
type LearningComputeClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LearningComputeCluster `json:"items"`
}

// Repository type metadata.
var (
	LearningComputeClusterKind             = "LearningComputeCluster"
	LearningComputeClusterGroupKind        = schema.GroupKind{Group: Group, Kind: LearningComputeClusterKind}.String()
	LearningComputeClusterKindAPIVersion   = LearningComputeClusterKind + "." + GroupVersion.String()
	LearningComputeClusterGroupVersionKind = GroupVersion.WithKind(LearningComputeClusterKind)
)

func init() {
	SchemeBuilder.Register(&LearningComputeCluster{}, &LearningComputeClusterList{})
}
