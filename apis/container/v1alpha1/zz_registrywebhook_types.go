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

type RegistryWebhookObservation struct {
}

type RegistryWebhookParameters struct {

	// +kubebuilder:validation:Required
	Actions []*string `json:"actions" tf:"actions,omitempty"`

	// +kubebuilder:validation:Optional
	CustomHeaders map[string]*string `json:"customHeaders,omitempty" tf:"custom_headers,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RegistryName *string `json:"registryName" tf:"registry_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Required
	ServiceURI *string `json:"serviceUri" tf:"service_uri,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RegistryWebhookSpec defines the desired state of RegistryWebhook
type RegistryWebhookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryWebhookParameters `json:"forProvider"`
}

// RegistryWebhookStatus defines the observed state of RegistryWebhook.
type RegistryWebhookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryWebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryWebhook is the Schema for the RegistryWebhooks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type RegistryWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegistryWebhookSpec   `json:"spec"`
	Status            RegistryWebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryWebhookList contains a list of RegistryWebhooks
type RegistryWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryWebhook `json:"items"`
}

// Repository type metadata.
var (
	RegistryWebhookKind             = "RegistryWebhook"
	RegistryWebhookGroupKind        = schema.GroupKind{Group: Group, Kind: RegistryWebhookKind}.String()
	RegistryWebhookKindAPIVersion   = RegistryWebhookKind + "." + GroupVersion.String()
	RegistryWebhookGroupVersionKind = GroupVersion.WithKind(RegistryWebhookKind)
)

func init() {
	SchemeBuilder.Register(&RegistryWebhook{}, &RegistryWebhookList{})
}
