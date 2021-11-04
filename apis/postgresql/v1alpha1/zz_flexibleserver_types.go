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

type FlexibleServerObservation struct {
	CmkEnabled *string `json:"cmkEnabled,omitempty" tf:"cmk_enabled,omitempty"`

	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`
}

type FlexibleServerParameters struct {

	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// +kubebuilder:validation:Optional
	AdministratorPasswordSecretRef v1.SecretKeySelector `json:"administratorPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BackupRetentionDays *int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`

	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// +kubebuilder:validation:Optional
	DelegatedSubnetID *string `json:"delegatedSubnetId,omitempty" tf:"delegated_subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	HighAvailability []HighAvailabilityParameters `json:"highAvailability,omitempty" tf:"high_availability,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PointInTimeRestoreTimeInUtc *string `json:"pointInTimeRestoreTimeInUtc,omitempty" tf:"point_in_time_restore_time_in_utc,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateDNSZoneID *string `json:"privateDnsZoneId,omitempty" tf:"private_dns_zone_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	SourceServerID *string `json:"sourceServerId,omitempty" tf:"source_server_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageMb *int64 `json:"storageMb,omitempty" tf:"storage_mb,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type HighAvailabilityObservation struct {
}

type HighAvailabilityParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty" tf:"standby_availability_zone,omitempty"`
}

type MaintenanceWindowObservation struct {
}

type MaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	DayOfWeek *int64 `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// +kubebuilder:validation:Optional
	StartHour *int64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	// +kubebuilder:validation:Optional
	StartMinute *int64 `json:"startMinute,omitempty" tf:"start_minute,omitempty"`
}

// FlexibleServerSpec defines the desired state of FlexibleServer
type FlexibleServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerParameters `json:"forProvider"`
}

// FlexibleServerStatus defines the observed state of FlexibleServer.
type FlexibleServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServer is the Schema for the FlexibleServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type FlexibleServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServerSpec   `json:"spec"`
	Status            FlexibleServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerList contains a list of FlexibleServers
type FlexibleServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServer `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServerKind             = "FlexibleServer"
	FlexibleServerGroupKind        = schema.GroupKind{Group: Group, Kind: FlexibleServerKind}.String()
	FlexibleServerKindAPIVersion   = FlexibleServerKind + "." + GroupVersion.String()
	FlexibleServerGroupVersionKind = GroupVersion.WithKind(FlexibleServerKind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServer{}, &FlexibleServerList{})
}
