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

type AuthorizationObservation struct {
}

type AuthorizationParameters struct {

	// +kubebuilder:validation:Optional
	Parameter *string `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`
}

type BackendObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackendParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Optional
	Credentials []CredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Proxy []BackendProxyParameters `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceFabricCluster []ServiceFabricClusterParameters `json:"serviceFabricCluster,omitempty" tf:"service_fabric_cluster,omitempty"`

	// +kubebuilder:validation:Optional
	TLS []TLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type BackendProxyObservation struct {
}

type BackendProxyParameters struct {

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type CredentialsObservation struct {
}

type CredentialsParameters struct {

	// +kubebuilder:validation:Optional
	Authorization []AuthorizationParameters `json:"authorization,omitempty" tf:"authorization,omitempty"`

	// +kubebuilder:validation:Optional
	Certificate []*string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	Header map[string]*string `json:"header,omitempty" tf:"header,omitempty"`

	// +kubebuilder:validation:Optional
	Query map[string]*string `json:"query,omitempty" tf:"query,omitempty"`
}

type ServerX509NameObservation struct {
}

type ServerX509NameParameters struct {

	// +kubebuilder:validation:Required
	IssuerCertificateThumbprint *string `json:"issuerCertificateThumbprint" tf:"issuer_certificate_thumbprint,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type ServiceFabricClusterObservation struct {
}

type ServiceFabricClusterParameters struct {

	// +kubebuilder:validation:Optional
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClientCertificateThumbprint *string `json:"clientCertificateThumbprint,omitempty" tf:"client_certificate_thumbprint,omitempty"`

	// +kubebuilder:validation:Required
	ManagementEndpoints []*string `json:"managementEndpoints" tf:"management_endpoints,omitempty"`

	// +kubebuilder:validation:Required
	MaxPartitionResolutionRetries *float64 `json:"maxPartitionResolutionRetries" tf:"max_partition_resolution_retries,omitempty"`

	// +kubebuilder:validation:Optional
	ServerCertificateThumbprints []*string `json:"serverCertificateThumbprints,omitempty" tf:"server_certificate_thumbprints,omitempty"`

	// +kubebuilder:validation:Optional
	ServerX509Name []ServerX509NameParameters `json:"serverX509Name,omitempty" tf:"server_x509_name,omitempty"`
}

type TLSObservation struct {
}

type TLSParameters struct {

	// +kubebuilder:validation:Optional
	ValidateCertificateChain *bool `json:"validateCertificateChain,omitempty" tf:"validate_certificate_chain,omitempty"`

	// +kubebuilder:validation:Optional
	ValidateCertificateName *bool `json:"validateCertificateName,omitempty" tf:"validate_certificate_name,omitempty"`
}

// BackendSpec defines the desired state of Backend
type BackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendParameters `json:"forProvider"`
}

// BackendStatus defines the observed state of Backend.
type BackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Backend is the Schema for the Backends API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Backend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackendSpec   `json:"spec"`
	Status            BackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendList contains a list of Backends
type BackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backend `json:"items"`
}

// Repository type metadata.
var (
	Backend_Kind             = "Backend"
	Backend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Backend_Kind}.String()
	Backend_KindAPIVersion   = Backend_Kind + "." + CRDGroupVersion.String()
	Backend_GroupVersionKind = CRDGroupVersion.WithKind(Backend_Kind)
)

func init() {
	SchemeBuilder.Register(&Backend{}, &BackendList{})
}
