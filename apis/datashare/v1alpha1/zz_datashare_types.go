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

type DataShareObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataShareParameters struct {

	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotSchedule []SnapshotScheduleParameters `json:"snapshotSchedule,omitempty" tf:"snapshot_schedule,omitempty"`

	// +kubebuilder:validation:Optional
	Terms *string `json:"terms,omitempty" tf:"terms,omitempty"`
}

type SnapshotScheduleObservation struct {
}

type SnapshotScheduleParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Recurrence *string `json:"recurrence" tf:"recurrence,omitempty"`

	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

// DataShareSpec defines the desired state of DataShare
type DataShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataShareParameters `json:"forProvider"`
}

// DataShareStatus defines the observed state of DataShare.
type DataShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataShare is the Schema for the DataShares API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DataShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataShareSpec   `json:"spec"`
	Status            DataShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataShareList contains a list of DataShares
type DataShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataShare `json:"items"`
}

// Repository type metadata.
var (
	DataShare_Kind             = "DataShare"
	DataShare_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataShare_Kind}.String()
	DataShare_KindAPIVersion   = DataShare_Kind + "." + CRDGroupVersion.String()
	DataShare_GroupVersionKind = CRDGroupVersion.WithKind(DataShare_Kind)
)

func init() {
	SchemeBuilder.Register(&DataShare{}, &DataShareList{})
}
