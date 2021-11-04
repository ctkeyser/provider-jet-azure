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

type ConnectionCertificateObservation struct {
}

type ConnectionCertificateParameters struct {

	// +kubebuilder:validation:Required
	AutomationAccountName *string `json:"automationAccountName" tf:"automation_account_name,omitempty"`

	// +kubebuilder:validation:Required
	AutomationCertificateName *string `json:"automationCertificateName" tf:"automation_certificate_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SubscriptionID *string `json:"subscriptionId" tf:"subscription_id,omitempty"`
}

// ConnectionCertificateSpec defines the desired state of ConnectionCertificate
type ConnectionCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionCertificateParameters `json:"forProvider"`
}

// ConnectionCertificateStatus defines the observed state of ConnectionCertificate.
type ConnectionCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionCertificate is the Schema for the ConnectionCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type ConnectionCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionCertificateSpec   `json:"spec"`
	Status            ConnectionCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionCertificateList contains a list of ConnectionCertificates
type ConnectionCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectionCertificate `json:"items"`
}

// Repository type metadata.
var (
	ConnectionCertificateKind             = "ConnectionCertificate"
	ConnectionCertificateGroupKind        = schema.GroupKind{Group: Group, Kind: ConnectionCertificateKind}.String()
	ConnectionCertificateKindAPIVersion   = ConnectionCertificateKind + "." + GroupVersion.String()
	ConnectionCertificateGroupVersionKind = GroupVersion.WithKind(ConnectionCertificateKind)
)

func init() {
	SchemeBuilder.Register(&ConnectionCertificate{}, &ConnectionCertificateList{})
}
