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

type PacketCaptureFilterObservation struct {
}

type PacketCaptureFilterParameters struct {

	// +kubebuilder:validation:Optional
	LocalIPAddress *string `json:"localIpAddress,omitempty" tf:"local_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	LocalPort *string `json:"localPort,omitempty" tf:"local_port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteIPAddress *string `json:"remoteIpAddress,omitempty" tf:"remote_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	RemotePort *string `json:"remotePort,omitempty" tf:"remote_port,omitempty"`
}

type PacketCaptureObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PacketCaptureParameters struct {

	// +kubebuilder:validation:Optional
	Filter []PacketCaptureFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumBytesPerPacket *int64 `json:"maximumBytesPerPacket,omitempty" tf:"maximum_bytes_per_packet,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumBytesPerSession *int64 `json:"maximumBytesPerSession,omitempty" tf:"maximum_bytes_per_session,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumCaptureDuration *int64 `json:"maximumCaptureDuration,omitempty" tf:"maximum_capture_duration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkWatcherName *string `json:"networkWatcherName" tf:"network_watcher_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	StorageLocation []PacketCaptureStorageLocationParameters `json:"storageLocation" tf:"storage_location,omitempty"`

	// +kubebuilder:validation:Required
	TargetResourceID *string `json:"targetResourceId" tf:"target_resource_id,omitempty"`
}

type PacketCaptureStorageLocationObservation struct {
	StoragePath *string `json:"storagePath,omitempty" tf:"storage_path,omitempty"`
}

type PacketCaptureStorageLocationParameters struct {

	// +kubebuilder:validation:Optional
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

// PacketCaptureSpec defines the desired state of PacketCapture
type PacketCaptureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PacketCaptureParameters `json:"forProvider"`
}

// PacketCaptureStatus defines the observed state of PacketCapture.
type PacketCaptureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PacketCaptureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PacketCapture is the Schema for the PacketCaptures API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type PacketCapture struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PacketCaptureSpec   `json:"spec"`
	Status            PacketCaptureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PacketCaptureList contains a list of PacketCaptures
type PacketCaptureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PacketCapture `json:"items"`
}

// Repository type metadata.
var (
	PacketCapture_Kind             = "PacketCapture"
	PacketCapture_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PacketCapture_Kind}.String()
	PacketCapture_KindAPIVersion   = PacketCapture_Kind + "." + CRDGroupVersion.String()
	PacketCapture_GroupVersionKind = CRDGroupVersion.WithKind(PacketCapture_Kind)
)

func init() {
	SchemeBuilder.Register(&PacketCapture{}, &PacketCaptureList{})
}
