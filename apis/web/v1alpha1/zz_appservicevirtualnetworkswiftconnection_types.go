/*
Copyright 2022 The Crossplane Authors.

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

type AppServiceVirtualNetworkSwiftConnectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppServiceVirtualNetworkSwiftConnectionParameters struct {

	// +kubebuilder:validation:Required
	AppServiceID *string `json:"appServiceId" tf:"app_service_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// AppServiceVirtualNetworkSwiftConnectionSpec defines the desired state of AppServiceVirtualNetworkSwiftConnection
type AppServiceVirtualNetworkSwiftConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppServiceVirtualNetworkSwiftConnectionParameters `json:"forProvider"`
}

// AppServiceVirtualNetworkSwiftConnectionStatus defines the observed state of AppServiceVirtualNetworkSwiftConnection.
type AppServiceVirtualNetworkSwiftConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppServiceVirtualNetworkSwiftConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppServiceVirtualNetworkSwiftConnection is the Schema for the AppServiceVirtualNetworkSwiftConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AppServiceVirtualNetworkSwiftConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceVirtualNetworkSwiftConnectionSpec   `json:"spec"`
	Status            AppServiceVirtualNetworkSwiftConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppServiceVirtualNetworkSwiftConnectionList contains a list of AppServiceVirtualNetworkSwiftConnections
type AppServiceVirtualNetworkSwiftConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppServiceVirtualNetworkSwiftConnection `json:"items"`
}

// Repository type metadata.
var (
	AppServiceVirtualNetworkSwiftConnection_Kind             = "AppServiceVirtualNetworkSwiftConnection"
	AppServiceVirtualNetworkSwiftConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppServiceVirtualNetworkSwiftConnection_Kind}.String()
	AppServiceVirtualNetworkSwiftConnection_KindAPIVersion   = AppServiceVirtualNetworkSwiftConnection_Kind + "." + CRDGroupVersion.String()
	AppServiceVirtualNetworkSwiftConnection_GroupVersionKind = CRDGroupVersion.WithKind(AppServiceVirtualNetworkSwiftConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&AppServiceVirtualNetworkSwiftConnection{}, &AppServiceVirtualNetworkSwiftConnectionList{})
}
