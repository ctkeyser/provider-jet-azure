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

type DataConnectorAwsCloudTrailObservation struct {
}

type DataConnectorAwsCloudTrailParameters struct {

	// +kubebuilder:validation:Required
	AwsRoleArn *string `json:"awsRoleArn" tf:"aws_role_arn,omitempty"`

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// DataConnectorAwsCloudTrailSpec defines the desired state of DataConnectorAwsCloudTrail
type DataConnectorAwsCloudTrailSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataConnectorAwsCloudTrailParameters `json:"forProvider"`
}

// DataConnectorAwsCloudTrailStatus defines the observed state of DataConnectorAwsCloudTrail.
type DataConnectorAwsCloudTrailStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataConnectorAwsCloudTrailObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataConnectorAwsCloudTrail is the Schema for the DataConnectorAwsCloudTrails API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type DataConnectorAwsCloudTrail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataConnectorAwsCloudTrailSpec   `json:"spec"`
	Status            DataConnectorAwsCloudTrailStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataConnectorAwsCloudTrailList contains a list of DataConnectorAwsCloudTrails
type DataConnectorAwsCloudTrailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataConnectorAwsCloudTrail `json:"items"`
}

// Repository type metadata.
var (
	DataConnectorAwsCloudTrailKind             = "DataConnectorAwsCloudTrail"
	DataConnectorAwsCloudTrailGroupKind        = schema.GroupKind{Group: Group, Kind: DataConnectorAwsCloudTrailKind}.String()
	DataConnectorAwsCloudTrailKindAPIVersion   = DataConnectorAwsCloudTrailKind + "." + GroupVersion.String()
	DataConnectorAwsCloudTrailGroupVersionKind = GroupVersion.WithKind(DataConnectorAwsCloudTrailKind)
)

func init() {
	SchemeBuilder.Register(&DataConnectorAwsCloudTrail{}, &DataConnectorAwsCloudTrailList{})
}
