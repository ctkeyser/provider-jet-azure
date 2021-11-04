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

type MongoCollectionAutoscaleSettingsObservation struct {
}

type MongoCollectionAutoscaleSettingsParameters struct {

	// +kubebuilder:validation:Optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type MongoCollectionIndexObservation struct {
}

type MongoCollectionIndexParameters struct {

	// +kubebuilder:validation:Required
	Keys []*string `json:"keys" tf:"keys,omitempty"`

	// +kubebuilder:validation:Optional
	Unique *bool `json:"unique,omitempty" tf:"unique,omitempty"`
}

type MongoCollectionObservation struct {
	SystemIndexes []SystemIndexesObservation `json:"systemIndexes,omitempty" tf:"system_indexes,omitempty"`
}

type MongoCollectionParameters struct {

	// +kubebuilder:validation:Required
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	AnalyticalStorageTTL *int64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscaleSettings []MongoCollectionAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultTTLSeconds *int64 `json:"defaultTtlSeconds,omitempty" tf:"default_ttl_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	Index []MongoCollectionIndexParameters `json:"index,omitempty" tf:"index,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ShardKey *string `json:"shardKey,omitempty" tf:"shard_key,omitempty"`

	// +kubebuilder:validation:Optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type SystemIndexesObservation struct {
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`

	Unique *bool `json:"unique,omitempty" tf:"unique,omitempty"`
}

type SystemIndexesParameters struct {
}

// MongoCollectionSpec defines the desired state of MongoCollection
type MongoCollectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MongoCollectionParameters `json:"forProvider"`
}

// MongoCollectionStatus defines the observed state of MongoCollection.
type MongoCollectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MongoCollectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MongoCollection is the Schema for the MongoCollections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type MongoCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongoCollectionSpec   `json:"spec"`
	Status            MongoCollectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MongoCollectionList contains a list of MongoCollections
type MongoCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoCollection `json:"items"`
}

// Repository type metadata.
var (
	MongoCollectionKind             = "MongoCollection"
	MongoCollectionGroupKind        = schema.GroupKind{Group: Group, Kind: MongoCollectionKind}.String()
	MongoCollectionKindAPIVersion   = MongoCollectionKind + "." + GroupVersion.String()
	MongoCollectionGroupVersionKind = GroupVersion.WithKind(MongoCollectionKind)
)

func init() {
	SchemeBuilder.Register(&MongoCollection{}, &MongoCollectionList{})
}
