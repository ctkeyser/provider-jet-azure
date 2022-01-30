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

type FunctionJavascriptUDFObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FunctionJavascriptUDFParameters struct {

	// +kubebuilder:validation:Required
	Input []InputParameters `json:"input" tf:"input,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Output []OutputParameters `json:"output" tf:"output,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Script *string `json:"script" tf:"script,omitempty"`

	// +kubebuilder:validation:Required
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name,omitempty"`
}

type InputObservation struct {
}

type InputParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type OutputObservation struct {
}

type OutputParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// FunctionJavascriptUDFSpec defines the desired state of FunctionJavascriptUDF
type FunctionJavascriptUDFSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionJavascriptUDFParameters `json:"forProvider"`
}

// FunctionJavascriptUDFStatus defines the observed state of FunctionJavascriptUDF.
type FunctionJavascriptUDFStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionJavascriptUDFObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionJavascriptUDF is the Schema for the FunctionJavascriptUDFs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FunctionJavascriptUDF struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionJavascriptUDFSpec   `json:"spec"`
	Status            FunctionJavascriptUDFStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionJavascriptUDFList contains a list of FunctionJavascriptUDFs
type FunctionJavascriptUDFList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionJavascriptUDF `json:"items"`
}

// Repository type metadata.
var (
	FunctionJavascriptUDF_Kind             = "FunctionJavascriptUDF"
	FunctionJavascriptUDF_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionJavascriptUDF_Kind}.String()
	FunctionJavascriptUDF_KindAPIVersion   = FunctionJavascriptUDF_Kind + "." + CRDGroupVersion.String()
	FunctionJavascriptUDF_GroupVersionKind = CRDGroupVersion.WithKind(FunctionJavascriptUDF_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionJavascriptUDF{}, &FunctionJavascriptUDFList{})
}