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

type AmbariObservation struct {
}

type AmbariParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type AutoscaleObservation struct {
}

type AutoscaleParameters struct {

	// +kubebuilder:validation:Optional
	Capacity []CapacityParameters `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Recurrence []RecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`
}

type CapacityObservation struct {
}

type CapacityParameters struct {

	// +kubebuilder:validation:Required
	MaxInstanceCount *int64 `json:"maxInstanceCount" tf:"max_instance_count,omitempty"`

	// +kubebuilder:validation:Required
	MinInstanceCount *int64 `json:"minInstanceCount" tf:"min_instance_count,omitempty"`
}

type ComponentVersionObservation struct {
}

type ComponentVersionParameters struct {

	// +kubebuilder:validation:Required
	Hadoop *string `json:"hadoop" tf:"hadoop,omitempty"`
}

type EdgeNodeObservation struct {
}

type EdgeNodeParameters struct {

	// +kubebuilder:validation:Required
	InstallScriptAction []InstallScriptActionParameters `json:"installScriptAction" tf:"install_script_action,omitempty"`

	// +kubebuilder:validation:Required
	TargetInstanceCount *int64 `json:"targetInstanceCount" tf:"target_instance_count,omitempty"`

	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`
}

type GatewayObservation struct {
}

type GatewayParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type HadoopClusterObservation struct {
	HTTPSEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SSHEndpoint *string `json:"sshEndpoint,omitempty" tf:"ssh_endpoint,omitempty"`
}

type HadoopClusterParameters struct {

	// +kubebuilder:validation:Required
	ClusterVersion *string `json:"clusterVersion" tf:"cluster_version,omitempty"`

	// +kubebuilder:validation:Required
	ComponentVersion []ComponentVersionParameters `json:"componentVersion" tf:"component_version,omitempty"`

	// +kubebuilder:validation:Required
	Gateway []GatewayParameters `json:"gateway" tf:"gateway,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Metastores []MetastoresParameters `json:"metastores,omitempty" tf:"metastores,omitempty"`

	// +kubebuilder:validation:Optional
	Monitor []MonitorParameters `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Network []NetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Roles []RolesParameters `json:"roles" tf:"roles,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityProfile []SecurityProfileParameters `json:"securityProfile,omitempty" tf:"security_profile,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccount []StorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountGen2 []StorageAccountGen2Parameters `json:"storageAccountGen2,omitempty" tf:"storage_account_gen2,omitempty"`

	// +kubebuilder:validation:Optional
	TLSMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

type HeadNodeObservation struct {
}

type HeadNodeParameters struct {

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

type HiveObservation struct {
}

type HiveParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type InstallScriptActionObservation struct {
}

type InstallScriptActionParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type MetastoresObservation struct {
}

type MetastoresParameters struct {

	// +kubebuilder:validation:Optional
	Ambari []AmbariParameters `json:"ambari,omitempty" tf:"ambari,omitempty"`

	// +kubebuilder:validation:Optional
	Hive []HiveParameters `json:"hive,omitempty" tf:"hive,omitempty"`

	// +kubebuilder:validation:Optional
	Oozie []OozieParameters `json:"oozie,omitempty" tf:"oozie,omitempty"`
}

type MonitorObservation struct {
}

type MonitorParameters struct {

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	PrimaryKeySecretRef v1.SecretKeySelector `json:"primaryKeySecretRef" tf:"-"`
}

type NetworkObservation struct {
}

type NetworkParameters struct {

	// +kubebuilder:validation:Optional
	ConnectionDirection *string `json:"connectionDirection,omitempty" tf:"connection_direction,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateLinkEnabled *bool `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled,omitempty"`
}

type OozieObservation struct {
}

type OozieParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type RecurrenceObservation struct {
}

type RecurrenceParameters struct {

	// +kubebuilder:validation:Required
	Schedule []ScheduleParameters `json:"schedule" tf:"schedule,omitempty"`

	// +kubebuilder:validation:Required
	Timezone *string `json:"timezone" tf:"timezone,omitempty"`
}

type RolesObservation struct {
}

type RolesParameters struct {

	// +kubebuilder:validation:Optional
	EdgeNode []EdgeNodeParameters `json:"edgeNode,omitempty" tf:"edge_node,omitempty"`

	// +kubebuilder:validation:Required
	HeadNode []HeadNodeParameters `json:"headNode" tf:"head_node,omitempty"`

	// +kubebuilder:validation:Required
	WorkerNode []WorkerNodeParameters `json:"workerNode" tf:"worker_node,omitempty"`

	// +kubebuilder:validation:Required
	ZookeeperNode []ZookeeperNodeParameters `json:"zookeeperNode" tf:"zookeeper_node,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Required
	Days []*string `json:"days" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	TargetInstanceCount *int64 `json:"targetInstanceCount" tf:"target_instance_count,omitempty"`

	// +kubebuilder:validation:Required
	Time *string `json:"time" tf:"time,omitempty"`
}

type SecurityProfileObservation struct {
}

type SecurityProfileParameters struct {

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

type StorageAccountGen2Observation struct {
}

type StorageAccountGen2Parameters struct {

	// +kubebuilder:validation:Required
	FileSystemID *string `json:"filesystemId" tf:"filesystem_id,omitempty"`

	// +kubebuilder:validation:Required
	IsDefault *bool `json:"isDefault" tf:"is_default,omitempty"`

	// +kubebuilder:validation:Required
	ManagedIdentityResourceID *string `json:"managedIdentityResourceId" tf:"managed_identity_resource_id,omitempty"`

	// +kubebuilder:validation:Required
	StorageResourceID *string `json:"storageResourceId" tf:"storage_resource_id,omitempty"`
}

type StorageAccountObservation struct {
}

type StorageAccountParameters struct {

	// +kubebuilder:validation:Required
	IsDefault *bool `json:"isDefault" tf:"is_default,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountKeySecretRef v1.SecretKeySelector `json:"storageAccountKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	StorageContainerID *string `json:"storageContainerId" tf:"storage_container_id,omitempty"`
}

type WorkerNodeObservation struct {
}

type WorkerNodeParameters struct {

	// +kubebuilder:validation:Optional
	Autoscale []AutoscaleParameters `json:"autoscale,omitempty" tf:"autoscale,omitempty"`

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

type ZookeeperNodeObservation struct {
}

type ZookeeperNodeParameters struct {

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

// HadoopClusterSpec defines the desired state of HadoopCluster
type HadoopClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HadoopClusterParameters `json:"forProvider"`
}

// HadoopClusterStatus defines the observed state of HadoopCluster.
type HadoopClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HadoopClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HadoopCluster is the Schema for the HadoopClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type HadoopCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HadoopClusterSpec   `json:"spec"`
	Status            HadoopClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HadoopClusterList contains a list of HadoopClusters
type HadoopClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HadoopCluster `json:"items"`
}

// Repository type metadata.
var (
	HadoopCluster_Kind             = "HadoopCluster"
	HadoopCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HadoopCluster_Kind}.String()
	HadoopCluster_KindAPIVersion   = HadoopCluster_Kind + "." + CRDGroupVersion.String()
	HadoopCluster_GroupVersionKind = CRDGroupVersion.WithKind(HadoopCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&HadoopCluster{}, &HadoopClusterList{})
}
