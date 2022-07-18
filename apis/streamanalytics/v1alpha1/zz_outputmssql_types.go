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

type OutputMSSQLObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OutputMSSQLParameters struct {

	// +kubebuilder:validation:Required
	Database *string `json:"database" tf:"database,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Required
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name,omitempty"`

	// +kubebuilder:validation:Required
	Table *string `json:"table" tf:"table,omitempty"`

	// +kubebuilder:validation:Required
	User *string `json:"user" tf:"user,omitempty"`
}

// OutputMSSQLSpec defines the desired state of OutputMSSQL
type OutputMSSQLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputMSSQLParameters `json:"forProvider"`
}

// OutputMSSQLStatus defines the observed state of OutputMSSQL.
type OutputMSSQLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputMSSQLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputMSSQL is the Schema for the OutputMSSQLs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type OutputMSSQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OutputMSSQLSpec   `json:"spec"`
	Status            OutputMSSQLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputMSSQLList contains a list of OutputMSSQLs
type OutputMSSQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputMSSQL `json:"items"`
}

// Repository type metadata.
var (
	OutputMSSQL_Kind             = "OutputMSSQL"
	OutputMSSQL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputMSSQL_Kind}.String()
	OutputMSSQL_KindAPIVersion   = OutputMSSQL_Kind + "." + CRDGroupVersion.String()
	OutputMSSQL_GroupVersionKind = CRDGroupVersion.WithKind(OutputMSSQL_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputMSSQL{}, &OutputMSSQLList{})
}
