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

type DatabaseMigrationServiceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DatabaseMigrationServiceParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DatabaseMigrationServiceSpec defines the desired state of DatabaseMigrationService
type DatabaseMigrationServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseMigrationServiceParameters `json:"forProvider"`
}

// DatabaseMigrationServiceStatus defines the observed state of DatabaseMigrationService.
type DatabaseMigrationServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseMigrationServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseMigrationService is the Schema for the DatabaseMigrationServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DatabaseMigrationService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseMigrationServiceSpec   `json:"spec"`
	Status            DatabaseMigrationServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseMigrationServiceList contains a list of DatabaseMigrationServices
type DatabaseMigrationServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseMigrationService `json:"items"`
}

// Repository type metadata.
var (
	DatabaseMigrationService_Kind             = "DatabaseMigrationService"
	DatabaseMigrationService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseMigrationService_Kind}.String()
	DatabaseMigrationService_KindAPIVersion   = DatabaseMigrationService_Kind + "." + CRDGroupVersion.String()
	DatabaseMigrationService_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseMigrationService_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseMigrationService{}, &DatabaseMigrationServiceList{})
}
