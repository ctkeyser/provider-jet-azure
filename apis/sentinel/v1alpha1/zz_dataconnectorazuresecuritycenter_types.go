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

type DataConnectorAzureSecurityCenterObservation struct {
}

type DataConnectorAzureSecurityCenterParameters struct {

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

// DataConnectorAzureSecurityCenterSpec defines the desired state of DataConnectorAzureSecurityCenter
type DataConnectorAzureSecurityCenterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataConnectorAzureSecurityCenterParameters `json:"forProvider"`
}

// DataConnectorAzureSecurityCenterStatus defines the observed state of DataConnectorAzureSecurityCenter.
type DataConnectorAzureSecurityCenterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataConnectorAzureSecurityCenterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataConnectorAzureSecurityCenter is the Schema for the DataConnectorAzureSecurityCenters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type DataConnectorAzureSecurityCenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataConnectorAzureSecurityCenterSpec   `json:"spec"`
	Status            DataConnectorAzureSecurityCenterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataConnectorAzureSecurityCenterList contains a list of DataConnectorAzureSecurityCenters
type DataConnectorAzureSecurityCenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataConnectorAzureSecurityCenter `json:"items"`
}

// Repository type metadata.
var (
	DataConnectorAzureSecurityCenterKind             = "DataConnectorAzureSecurityCenter"
	DataConnectorAzureSecurityCenterGroupKind        = schema.GroupKind{Group: Group, Kind: DataConnectorAzureSecurityCenterKind}.String()
	DataConnectorAzureSecurityCenterKindAPIVersion   = DataConnectorAzureSecurityCenterKind + "." + GroupVersion.String()
	DataConnectorAzureSecurityCenterGroupVersionKind = GroupVersion.WithKind(DataConnectorAzureSecurityCenterKind)
)

func init() {
	SchemeBuilder.Register(&DataConnectorAzureSecurityCenter{}, &DataConnectorAzureSecurityCenterList{})
}
