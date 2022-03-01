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

type RolesWorkerNodeAutoscaleRecurrenceObservation struct {
}

type RolesWorkerNodeAutoscaleRecurrenceParameters struct {

	// +kubebuilder:validation:Required
	Schedule []WorkerNodeAutoscaleRecurrenceScheduleParameters `json:"schedule" tf:"schedule,omitempty"`

	// +kubebuilder:validation:Required
	Timezone *string `json:"timezone" tf:"timezone,omitempty"`
}

type SparkClusterComponentVersionObservation struct {
}

type SparkClusterComponentVersionParameters struct {

	// +kubebuilder:validation:Required
	Spark *string `json:"spark" tf:"spark,omitempty"`
}

type SparkClusterGatewayObservation struct {
}

type SparkClusterGatewayParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type SparkClusterMetastoresAmbariObservation struct {
}

type SparkClusterMetastoresAmbariParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type SparkClusterMetastoresHiveObservation struct {
}

type SparkClusterMetastoresHiveParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type SparkClusterMetastoresObservation struct {
}

type SparkClusterMetastoresOozieObservation struct {
}

type SparkClusterMetastoresOozieParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type SparkClusterMetastoresParameters struct {

	// +kubebuilder:validation:Optional
	Ambari []SparkClusterMetastoresAmbariParameters `json:"ambari,omitempty" tf:"ambari,omitempty"`

	// +kubebuilder:validation:Optional
	Hive []SparkClusterMetastoresHiveParameters `json:"hive,omitempty" tf:"hive,omitempty"`

	// +kubebuilder:validation:Optional
	Oozie []SparkClusterMetastoresOozieParameters `json:"oozie,omitempty" tf:"oozie,omitempty"`
}

type SparkClusterMonitorObservation struct {
}

type SparkClusterMonitorParameters struct {

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	PrimaryKeySecretRef v1.SecretKeySelector `json:"primaryKeySecretRef" tf:"-"`
}

type SparkClusterNetworkObservation struct {
}

type SparkClusterNetworkParameters struct {

	// +kubebuilder:validation:Optional
	ConnectionDirection *string `json:"connectionDirection,omitempty" tf:"connection_direction,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateLinkEnabled *bool `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled,omitempty"`
}

type SparkClusterObservation struct {
	HTTPSEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SSHEndpoint *string `json:"sshEndpoint,omitempty" tf:"ssh_endpoint,omitempty"`
}

type SparkClusterParameters struct {

	// +kubebuilder:validation:Required
	ClusterVersion *string `json:"clusterVersion" tf:"cluster_version,omitempty"`

	// +kubebuilder:validation:Required
	ComponentVersion []SparkClusterComponentVersionParameters `json:"componentVersion" tf:"component_version,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionInTransitEnabled *bool `json:"encryptionInTransitEnabled,omitempty" tf:"encryption_in_transit_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Gateway []SparkClusterGatewayParameters `json:"gateway" tf:"gateway,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Metastores []SparkClusterMetastoresParameters `json:"metastores,omitempty" tf:"metastores,omitempty"`

	// +kubebuilder:validation:Optional
	Monitor []SparkClusterMonitorParameters `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Network []SparkClusterNetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Roles []SparkClusterRolesParameters `json:"roles" tf:"roles,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityProfile []SparkClusterSecurityProfileParameters `json:"securityProfile,omitempty" tf:"security_profile,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccount []SparkClusterStorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountGen2 []SparkClusterStorageAccountGen2Parameters `json:"storageAccountGen2,omitempty" tf:"storage_account_gen2,omitempty"`

	// +kubebuilder:validation:Optional
	TLSMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

type SparkClusterRolesHeadNodeObservation struct {
}

type SparkClusterRolesHeadNodeParameters struct {

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`

	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`
}

type SparkClusterRolesObservation struct {
}

type SparkClusterRolesParameters struct {

	// +kubebuilder:validation:Required
	HeadNode []SparkClusterRolesHeadNodeParameters `json:"headNode" tf:"head_node,omitempty"`

	// +kubebuilder:validation:Required
	WorkerNode []SparkClusterRolesWorkerNodeParameters `json:"workerNode" tf:"worker_node,omitempty"`

	// +kubebuilder:validation:Required
	ZookeeperNode []SparkClusterRolesZookeeperNodeParameters `json:"zookeeperNode" tf:"zookeeper_node,omitempty"`
}

type SparkClusterRolesWorkerNodeAutoscaleObservation struct {
}

type SparkClusterRolesWorkerNodeAutoscaleParameters struct {

	// +kubebuilder:validation:Optional
	Capacity []WorkerNodeAutoscaleCapacityParameters `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Recurrence []RolesWorkerNodeAutoscaleRecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`
}

type SparkClusterRolesWorkerNodeObservation struct {
}

type SparkClusterRolesWorkerNodeParameters struct {

	// +kubebuilder:validation:Optional
	Autoscale []SparkClusterRolesWorkerNodeAutoscaleParameters `json:"autoscale,omitempty" tf:"autoscale,omitempty"`

	// +kubebuilder:validation:Optional
	MinInstanceCount *int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	TargetInstanceCount *int64 `json:"targetInstanceCount" tf:"target_instance_count,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`

	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`
}

type SparkClusterRolesZookeeperNodeObservation struct {
}

type SparkClusterRolesZookeeperNodeParameters struct {

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`

	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`
}

type SparkClusterSecurityProfileObservation struct {
}

type SparkClusterSecurityProfileParameters struct {

	// +kubebuilder:validation:Required
	AaddsResourceID *string `json:"aaddsResourceId" tf:"aadds_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterUsersGroupDNS []*string `json:"clusterUsersGroupDns,omitempty" tf:"cluster_users_group_dns,omitempty"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Required
	DomainUserPasswordSecretRef v1.SecretKeySelector `json:"domainUserPasswordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	DomainUsername *string `json:"domainUsername" tf:"domain_username,omitempty"`

	// +kubebuilder:validation:Required
	LdapsUrls []*string `json:"ldapsUrls" tf:"ldaps_urls,omitempty"`

	// +kubebuilder:validation:Required
	MsiResourceID *string `json:"msiResourceId" tf:"msi_resource_id,omitempty"`
}

type SparkClusterStorageAccountGen2Observation struct {
}

type SparkClusterStorageAccountGen2Parameters struct {

	// +kubebuilder:validation:Required
	FileSystemID *string `json:"filesystemId" tf:"filesystem_id,omitempty"`

	// +kubebuilder:validation:Required
	IsDefault *bool `json:"isDefault" tf:"is_default,omitempty"`

	// +kubebuilder:validation:Required
	ManagedIdentityResourceID *string `json:"managedIdentityResourceId" tf:"managed_identity_resource_id,omitempty"`

	// +kubebuilder:validation:Required
	StorageResourceID *string `json:"storageResourceId" tf:"storage_resource_id,omitempty"`
}

type SparkClusterStorageAccountObservation struct {
}

type SparkClusterStorageAccountParameters struct {

	// +kubebuilder:validation:Required
	IsDefault *bool `json:"isDefault" tf:"is_default,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountKeySecretRef v1.SecretKeySelector `json:"storageAccountKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	StorageContainerID *string `json:"storageContainerId" tf:"storage_container_id,omitempty"`
}

type WorkerNodeAutoscaleCapacityObservation struct {
}

type WorkerNodeAutoscaleCapacityParameters struct {

	// +kubebuilder:validation:Required
	MaxInstanceCount *int64 `json:"maxInstanceCount" tf:"max_instance_count,omitempty"`

	// +kubebuilder:validation:Required
	MinInstanceCount *int64 `json:"minInstanceCount" tf:"min_instance_count,omitempty"`
}

type WorkerNodeAutoscaleRecurrenceScheduleObservation struct {
}

type WorkerNodeAutoscaleRecurrenceScheduleParameters struct {

	// +kubebuilder:validation:Required
	Days []*string `json:"days" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	TargetInstanceCount *int64 `json:"targetInstanceCount" tf:"target_instance_count,omitempty"`

	// +kubebuilder:validation:Required
	Time *string `json:"time" tf:"time,omitempty"`
}

// SparkClusterSpec defines the desired state of SparkCluster
type SparkClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SparkClusterParameters `json:"forProvider"`
}

// SparkClusterStatus defines the observed state of SparkCluster.
type SparkClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SparkClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SparkCluster is the Schema for the SparkClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SparkCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SparkClusterSpec   `json:"spec"`
	Status            SparkClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SparkClusterList contains a list of SparkClusters
type SparkClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SparkCluster `json:"items"`
}

// Repository type metadata.
var (
	SparkCluster_Kind             = "SparkCluster"
	SparkCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SparkCluster_Kind}.String()
	SparkCluster_KindAPIVersion   = SparkCluster_Kind + "." + CRDGroupVersion.String()
	SparkCluster_GroupVersionKind = CRDGroupVersion.WithKind(SparkCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&SparkCluster{}, &SparkClusterList{})
}
