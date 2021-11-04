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

type SqlTriggerObservation struct {
}

type SqlTriggerParameters struct {

	// +kubebuilder:validation:Required
	Body *string `json:"body" tf:"body,omitempty"`

	// +kubebuilder:validation:Required
	ContainerID *string `json:"containerId" tf:"container_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// SqlTriggerSpec defines the desired state of SqlTrigger
type SqlTriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SqlTriggerParameters `json:"forProvider"`
}

// SqlTriggerStatus defines the observed state of SqlTrigger.
type SqlTriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SqlTriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SqlTrigger is the Schema for the SqlTriggers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type SqlTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlTriggerSpec   `json:"spec"`
	Status            SqlTriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlTriggerList contains a list of SqlTriggers
type SqlTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlTrigger `json:"items"`
}

// Repository type metadata.
var (
	SqlTriggerKind             = "SqlTrigger"
	SqlTriggerGroupKind        = schema.GroupKind{Group: Group, Kind: SqlTriggerKind}.String()
	SqlTriggerKindAPIVersion   = SqlTriggerKind + "." + GroupVersion.String()
	SqlTriggerGroupVersionKind = GroupVersion.WithKind(SqlTriggerKind)
)

func init() {
	SchemeBuilder.Register(&SqlTrigger{}, &SqlTriggerList{})
}
