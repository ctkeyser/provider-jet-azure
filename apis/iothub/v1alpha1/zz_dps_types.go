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

type DpsObservation struct {
	DeviceProvisioningHostName *string `json:"deviceProvisioningHostName,omitempty" tf:"device_provisioning_host_name,omitempty"`

	IDScope *string `json:"idScope,omitempty" tf:"id_scope,omitempty"`

	ServiceOperationsHostName *string `json:"serviceOperationsHostName,omitempty" tf:"service_operations_host_name,omitempty"`
}

type DpsParameters struct {

	// +kubebuilder:validation:Optional
	AllocationPolicy *string `json:"allocationPolicy,omitempty" tf:"allocation_policy,omitempty"`

	// +kubebuilder:validation:Optional
	LinkedHub []LinkedHubParameters `json:"linkedHub,omitempty" tf:"linked_hub,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	Sku []SkuParameters `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LinkedHubObservation struct {
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`
}

type LinkedHubParameters struct {

	// +kubebuilder:validation:Optional
	AllocationWeight *int64 `json:"allocationWeight,omitempty" tf:"allocation_weight,omitempty"`

	// +kubebuilder:validation:Optional
	ApplyAllocationPolicy *bool `json:"applyAllocationPolicy,omitempty" tf:"apply_allocation_policy,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`
}

type SkuObservation struct {
}

type SkuParameters struct {

	// +kubebuilder:validation:Required
	Capacity *int64 `json:"capacity" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// DpsSpec defines the desired state of Dps
type DpsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DpsParameters `json:"forProvider"`
}

// DpsStatus defines the observed state of Dps.
type DpsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DpsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Dps is the Schema for the Dpss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type Dps struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DpsSpec   `json:"spec"`
	Status            DpsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DpsList contains a list of Dpss
type DpsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dps `json:"items"`
}

// Repository type metadata.
var (
	DpsKind             = "Dps"
	DpsGroupKind        = schema.GroupKind{Group: Group, Kind: DpsKind}.String()
	DpsKindAPIVersion   = DpsKind + "." + GroupVersion.String()
	DpsGroupVersionKind = GroupVersion.WithKind(DpsKind)
)

func init() {
	SchemeBuilder.Register(&Dps{}, &DpsList{})
}
