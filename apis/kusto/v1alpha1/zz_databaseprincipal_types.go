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

type DatabasePrincipalObservation struct {
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty" tf:"fully_qualified_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DatabasePrincipalParameters struct {

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	ObjectID *string `json:"objectId" tf:"object_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// DatabasePrincipalSpec defines the desired state of DatabasePrincipal
type DatabasePrincipalSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabasePrincipalParameters `json:"forProvider"`
}

// DatabasePrincipalStatus defines the observed state of DatabasePrincipal.
type DatabasePrincipalStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabasePrincipalObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatabasePrincipal is the Schema for the DatabasePrincipals API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DatabasePrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabasePrincipalSpec   `json:"spec"`
	Status            DatabasePrincipalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabasePrincipalList contains a list of DatabasePrincipals
type DatabasePrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabasePrincipal `json:"items"`
}

// Repository type metadata.
var (
	DatabasePrincipal_Kind             = "DatabasePrincipal"
	DatabasePrincipal_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabasePrincipal_Kind}.String()
	DatabasePrincipal_KindAPIVersion   = DatabasePrincipal_Kind + "." + CRDGroupVersion.String()
	DatabasePrincipal_GroupVersionKind = CRDGroupVersion.WithKind(DatabasePrincipal_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabasePrincipal{}, &DatabasePrincipalList{})
}
