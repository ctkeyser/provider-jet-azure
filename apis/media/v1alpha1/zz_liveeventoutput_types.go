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

type LiveEventOutputObservation struct {
}

type LiveEventOutputParameters struct {

	// +kubebuilder:validation:Required
	ArchiveWindowDuration *string `json:"archiveWindowDuration" tf:"archive_window_duration,omitempty"`

	// +kubebuilder:validation:Required
	AssetName *string `json:"assetName" tf:"asset_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	HlsFragmentsPerTSSegment *int64 `json:"hlsFragmentsPerTsSegment,omitempty" tf:"hls_fragments_per_ts_segment,omitempty"`

	// +kubebuilder:validation:Required
	LiveEventID *string `json:"liveEventId" tf:"live_event_id,omitempty"`

	// +kubebuilder:validation:Optional
	ManifestName *string `json:"manifestName,omitempty" tf:"manifest_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OutputSnapTimeInSeconds *int64 `json:"outputSnapTimeInSeconds,omitempty" tf:"output_snap_time_in_seconds,omitempty"`
}

// LiveEventOutputSpec defines the desired state of LiveEventOutput
type LiveEventOutputSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LiveEventOutputParameters `json:"forProvider"`
}

// LiveEventOutputStatus defines the observed state of LiveEventOutput.
type LiveEventOutputStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LiveEventOutputObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LiveEventOutput is the Schema for the LiveEventOutputs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type LiveEventOutput struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LiveEventOutputSpec   `json:"spec"`
	Status            LiveEventOutputStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LiveEventOutputList contains a list of LiveEventOutputs
type LiveEventOutputList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiveEventOutput `json:"items"`
}

// Repository type metadata.
var (
	LiveEventOutputKind             = "LiveEventOutput"
	LiveEventOutputGroupKind        = schema.GroupKind{Group: Group, Kind: LiveEventOutputKind}.String()
	LiveEventOutputKindAPIVersion   = LiveEventOutputKind + "." + GroupVersion.String()
	LiveEventOutputGroupVersionKind = GroupVersion.WithKind(LiveEventOutputKind)
)

func init() {
	SchemeBuilder.Register(&LiveEventOutput{}, &LiveEventOutputList{})
}
