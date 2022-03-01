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

type LinkedServiceDataLakeStorageGen2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LinkedServiceDataLakeStorageGen2Parameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

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

	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServicePrincipalKey *string `json:"servicePrincipalKey,omitempty" tf:"service_principal_key,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountKey *string `json:"storageAccountKey,omitempty" tf:"storage_account_key,omitempty"`

	// +kubebuilder:validation:Optional
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// +kubebuilder:validation:Optional
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

// LinkedServiceDataLakeStorageGen2Spec defines the desired state of LinkedServiceDataLakeStorageGen2
type LinkedServiceDataLakeStorageGen2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceDataLakeStorageGen2Parameters `json:"forProvider"`
}

// LinkedServiceDataLakeStorageGen2Status defines the observed state of LinkedServiceDataLakeStorageGen2.
type LinkedServiceDataLakeStorageGen2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceDataLakeStorageGen2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceDataLakeStorageGen2 is the Schema for the LinkedServiceDataLakeStorageGen2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LinkedServiceDataLakeStorageGen2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceDataLakeStorageGen2Spec   `json:"spec"`
	Status            LinkedServiceDataLakeStorageGen2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceDataLakeStorageGen2List contains a list of LinkedServiceDataLakeStorageGen2s
type LinkedServiceDataLakeStorageGen2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceDataLakeStorageGen2 `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceDataLakeStorageGen2_Kind             = "LinkedServiceDataLakeStorageGen2"
	LinkedServiceDataLakeStorageGen2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceDataLakeStorageGen2_Kind}.String()
	LinkedServiceDataLakeStorageGen2_KindAPIVersion   = LinkedServiceDataLakeStorageGen2_Kind + "." + CRDGroupVersion.String()
	LinkedServiceDataLakeStorageGen2_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceDataLakeStorageGen2_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceDataLakeStorageGen2{}, &LinkedServiceDataLakeStorageGen2List{})
}
