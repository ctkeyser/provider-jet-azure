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

type DpsCertificateObservation struct {
}

type DpsCertificateParameters struct {

	// +kubebuilder:validation:Required
	CertificateContentSecretRef v1.SecretKeySelector `json:"certificateContentSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	IotDpsName *string `json:"iotDpsName" tf:"iot_dps_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// DpsCertificateSpec defines the desired state of DpsCertificate
type DpsCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DpsCertificateParameters `json:"forProvider"`
}

// DpsCertificateStatus defines the observed state of DpsCertificate.
type DpsCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DpsCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DpsCertificate is the Schema for the DpsCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type DpsCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DpsCertificateSpec   `json:"spec"`
	Status            DpsCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DpsCertificateList contains a list of DpsCertificates
type DpsCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DpsCertificate `json:"items"`
}

// Repository type metadata.
var (
	DpsCertificateKind             = "DpsCertificate"
	DpsCertificateGroupKind        = schema.GroupKind{Group: Group, Kind: DpsCertificateKind}.String()
	DpsCertificateKindAPIVersion   = DpsCertificateKind + "." + GroupVersion.String()
	DpsCertificateGroupVersionKind = GroupVersion.WithKind(DpsCertificateKind)
)

func init() {
	SchemeBuilder.Register(&DpsCertificate{}, &DpsCertificateList{})
}
