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

type FactoryLinkedServiceSnowflakeKeyVaultPasswordObservation struct {
}

type FactoryLinkedServiceSnowflakeKeyVaultPasswordParameters struct {

	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type FactoryLinkedServiceSnowflakeObservation struct {
}

type FactoryLinkedServiceSnowflakeParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionString *string `json:"connectionString" tf:"connection_string,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultPassword []FactoryLinkedServiceSnowflakeKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// FactoryLinkedServiceSnowflakeSpec defines the desired state of FactoryLinkedServiceSnowflake
type FactoryLinkedServiceSnowflakeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FactoryLinkedServiceSnowflakeParameters `json:"forProvider"`
}

// FactoryLinkedServiceSnowflakeStatus defines the observed state of FactoryLinkedServiceSnowflake.
type FactoryLinkedServiceSnowflakeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FactoryLinkedServiceSnowflakeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceSnowflake is the Schema for the FactoryLinkedServiceSnowflakes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type FactoryLinkedServiceSnowflake struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryLinkedServiceSnowflakeSpec   `json:"spec"`
	Status            FactoryLinkedServiceSnowflakeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceSnowflakeList contains a list of FactoryLinkedServiceSnowflakes
type FactoryLinkedServiceSnowflakeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FactoryLinkedServiceSnowflake `json:"items"`
}

// Repository type metadata.
var (
	FactoryLinkedServiceSnowflakeKind             = "FactoryLinkedServiceSnowflake"
	FactoryLinkedServiceSnowflakeGroupKind        = schema.GroupKind{Group: Group, Kind: FactoryLinkedServiceSnowflakeKind}.String()
	FactoryLinkedServiceSnowflakeKindAPIVersion   = FactoryLinkedServiceSnowflakeKind + "." + GroupVersion.String()
	FactoryLinkedServiceSnowflakeGroupVersionKind = GroupVersion.WithKind(FactoryLinkedServiceSnowflakeKind)
)

func init() {
	SchemeBuilder.Register(&FactoryLinkedServiceSnowflake{}, &FactoryLinkedServiceSnowflakeList{})
}
