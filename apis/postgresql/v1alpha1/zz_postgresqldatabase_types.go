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

type PostgresqlDatabaseObservation struct {
}

type PostgresqlDatabaseParameters struct {

	// +kubebuilder:validation:Required
	Charset *string `json:"charset" tf:"charset"`

	// +kubebuilder:validation:Required
	Collation *string `json:"collation" tf:"collation"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=PostgresqlServer
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name"`

	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`
}

// PostgresqlDatabaseSpec defines the desired state of PostgresqlDatabase
type PostgresqlDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlDatabaseParameters `json:"forProvider"`
}

// PostgresqlDatabaseStatus defines the observed state of PostgresqlDatabase.
type PostgresqlDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlDatabase is the Schema for the PostgresqlDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PostgresqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlDatabaseSpec   `json:"spec"`
	Status            PostgresqlDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlDatabaseList contains a list of PostgresqlDatabases
type PostgresqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlDatabase `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlDatabaseKind             = "PostgresqlDatabase"
	PostgresqlDatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresqlDatabaseKind}.String()
	PostgresqlDatabaseKindAPIVersion   = PostgresqlDatabaseKind + "." + GroupVersion.String()
	PostgresqlDatabaseGroupVersionKind = GroupVersion.WithKind(PostgresqlDatabaseKind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlDatabase{}, &PostgresqlDatabaseList{})
}