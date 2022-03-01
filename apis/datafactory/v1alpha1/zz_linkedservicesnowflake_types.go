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

type LinkedServiceSnowflakeKeyVaultPasswordObservation struct {
}

type LinkedServiceSnowflakeKeyVaultPasswordParameters struct {

	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type LinkedServiceSnowflakeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LinkedServiceSnowflakeParameters struct {

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
	KeyVaultPassword []LinkedServiceSnowflakeKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// LinkedServiceSnowflakeSpec defines the desired state of LinkedServiceSnowflake
type LinkedServiceSnowflakeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceSnowflakeParameters `json:"forProvider"`
}

// LinkedServiceSnowflakeStatus defines the observed state of LinkedServiceSnowflake.
type LinkedServiceSnowflakeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceSnowflakeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceSnowflake is the Schema for the LinkedServiceSnowflakes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LinkedServiceSnowflake struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceSnowflakeSpec   `json:"spec"`
	Status            LinkedServiceSnowflakeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceSnowflakeList contains a list of LinkedServiceSnowflakes
type LinkedServiceSnowflakeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceSnowflake `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceSnowflake_Kind             = "LinkedServiceSnowflake"
	LinkedServiceSnowflake_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceSnowflake_Kind}.String()
	LinkedServiceSnowflake_KindAPIVersion   = LinkedServiceSnowflake_Kind + "." + CRDGroupVersion.String()
	LinkedServiceSnowflake_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceSnowflake_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceSnowflake{}, &LinkedServiceSnowflakeList{})
}
