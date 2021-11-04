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

type MxRecordObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
}

type MxRecordParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Record []MxRecordRecordParameters `json:"record" tf:"record,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	TTL *int64 `json:"ttl" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	ZoneName *string `json:"zoneName" tf:"zone_name,omitempty"`
}

type MxRecordRecordObservation struct {
}

type MxRecordRecordParameters struct {

	// +kubebuilder:validation:Required
	Exchange *string `json:"exchange" tf:"exchange,omitempty"`

	// +kubebuilder:validation:Required
	Preference *string `json:"preference" tf:"preference,omitempty"`
}

// MxRecordSpec defines the desired state of MxRecord
type MxRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MxRecordParameters `json:"forProvider"`
}

// MxRecordStatus defines the observed state of MxRecord.
type MxRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MxRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MxRecord is the Schema for the MxRecords API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type MxRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MxRecordSpec   `json:"spec"`
	Status            MxRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MxRecordList contains a list of MxRecords
type MxRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MxRecord `json:"items"`
}

// Repository type metadata.
var (
	MxRecordKind             = "MxRecord"
	MxRecordGroupKind        = schema.GroupKind{Group: Group, Kind: MxRecordKind}.String()
	MxRecordKindAPIVersion   = MxRecordKind + "." + GroupVersion.String()
	MxRecordGroupVersionKind = GroupVersion.WithKind(MxRecordKind)
)

func init() {
	SchemeBuilder.Register(&MxRecord{}, &MxRecordList{})
}
