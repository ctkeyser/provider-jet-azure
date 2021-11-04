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

type AppIntegrationAccountMapObservation struct {
}

type AppIntegrationAccountMapParameters struct {

	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// +kubebuilder:validation:Required
	IntegrationAccountName *string `json:"integrationAccountName" tf:"integration_account_name,omitempty"`

	// +kubebuilder:validation:Required
	MapType *string `json:"mapType" tf:"map_type,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// AppIntegrationAccountMapSpec defines the desired state of AppIntegrationAccountMap
type AppIntegrationAccountMapSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppIntegrationAccountMapParameters `json:"forProvider"`
}

// AppIntegrationAccountMapStatus defines the observed state of AppIntegrationAccountMap.
type AppIntegrationAccountMapStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppIntegrationAccountMapObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountMap is the Schema for the AppIntegrationAccountMaps API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type AppIntegrationAccountMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppIntegrationAccountMapSpec   `json:"spec"`
	Status            AppIntegrationAccountMapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountMapList contains a list of AppIntegrationAccountMaps
type AppIntegrationAccountMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppIntegrationAccountMap `json:"items"`
}

// Repository type metadata.
var (
	AppIntegrationAccountMapKind             = "AppIntegrationAccountMap"
	AppIntegrationAccountMapGroupKind        = schema.GroupKind{Group: Group, Kind: AppIntegrationAccountMapKind}.String()
	AppIntegrationAccountMapKindAPIVersion   = AppIntegrationAccountMapKind + "." + GroupVersion.String()
	AppIntegrationAccountMapGroupVersionKind = GroupVersion.WithKind(AppIntegrationAccountMapKind)
)

func init() {
	SchemeBuilder.Register(&AppIntegrationAccountMap{}, &AppIntegrationAccountMapList{})
}
