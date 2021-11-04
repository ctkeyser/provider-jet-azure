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

type SystemTopicIdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type SystemTopicIdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SystemTopicObservation struct {
	MetricArmResourceID *string `json:"metricArmResourceId,omitempty" tf:"metric_arm_resource_id,omitempty"`
}

type SystemTopicParameters struct {

	// +kubebuilder:validation:Optional
	Identity []SystemTopicIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SourceArmResourceID *string `json:"sourceArmResourceId" tf:"source_arm_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TopicType *string `json:"topicType" tf:"topic_type,omitempty"`
}

// SystemTopicSpec defines the desired state of SystemTopic
type SystemTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SystemTopicParameters `json:"forProvider"`
}

// SystemTopicStatus defines the observed state of SystemTopic.
type SystemTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SystemTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SystemTopic is the Schema for the SystemTopics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type SystemTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SystemTopicSpec   `json:"spec"`
	Status            SystemTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SystemTopicList contains a list of SystemTopics
type SystemTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SystemTopic `json:"items"`
}

// Repository type metadata.
var (
	SystemTopicKind             = "SystemTopic"
	SystemTopicGroupKind        = schema.GroupKind{Group: Group, Kind: SystemTopicKind}.String()
	SystemTopicKindAPIVersion   = SystemTopicKind + "." + GroupVersion.String()
	SystemTopicGroupVersionKind = GroupVersion.WithKind(SystemTopicKind)
)

func init() {
	SchemeBuilder.Register(&SystemTopic{}, &SystemTopicList{})
}
