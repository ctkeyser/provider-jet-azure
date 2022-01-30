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

type KafkaClusterComponentVersionObservation struct {
}

type KafkaClusterComponentVersionParameters struct {

	// +kubebuilder:validation:Required
	Kafka *string `json:"kafka" tf:"kafka,omitempty"`
}

type KafkaClusterGatewayObservation struct {
}

type KafkaClusterGatewayParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type KafkaClusterMetastoresAmbariObservation struct {
}

type KafkaClusterMetastoresAmbariParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type KafkaClusterMetastoresHiveObservation struct {
}

type KafkaClusterMetastoresHiveParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type KafkaClusterMetastoresObservation struct {
}

type KafkaClusterMetastoresOozieObservation struct {
}

type KafkaClusterMetastoresOozieParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type KafkaClusterMetastoresParameters struct {

	// +kubebuilder:validation:Optional
	Ambari []KafkaClusterMetastoresAmbariParameters `json:"ambari,omitempty" tf:"ambari,omitempty"`

	// +kubebuilder:validation:Optional
	Hive []KafkaClusterMetastoresHiveParameters `json:"hive,omitempty" tf:"hive,omitempty"`

	// +kubebuilder:validation:Optional
	Oozie []KafkaClusterMetastoresOozieParameters `json:"oozie,omitempty" tf:"oozie,omitempty"`
}

type KafkaClusterMonitorObservation struct {
}

type KafkaClusterMonitorParameters struct {

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	PrimaryKeySecretRef v1.SecretKeySelector `json:"primaryKeySecretRef" tf:"-"`
}

type KafkaClusterObservation struct {
	HTTPSEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KafkaRestProxyEndpoint *string `json:"kafkaRestProxyEndpoint,omitempty" tf:"kafka_rest_proxy_endpoint,omitempty"`

	SSHEndpoint *string `json:"sshEndpoint,omitempty" tf:"ssh_endpoint,omitempty"`
}

type KafkaClusterParameters struct {

	// +kubebuilder:validation:Required
	ClusterVersion *string `json:"clusterVersion" tf:"cluster_version,omitempty"`

	// +kubebuilder:validation:Required
	ComponentVersion []KafkaClusterComponentVersionParameters `json:"componentVersion" tf:"component_version,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionInTransitEnabled *bool `json:"encryptionInTransitEnabled,omitempty" tf:"encryption_in_transit_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Gateway []KafkaClusterGatewayParameters `json:"gateway" tf:"gateway,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Metastores []KafkaClusterMetastoresParameters `json:"metastores,omitempty" tf:"metastores,omitempty"`

	// +kubebuilder:validation:Optional
	Monitor []KafkaClusterMonitorParameters `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RestProxy []RestProxyParameters `json:"restProxy,omitempty" tf:"rest_proxy,omitempty"`

	// +kubebuilder:validation:Required
	Roles []KafkaClusterRolesParameters `json:"roles" tf:"roles,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityProfile []KafkaClusterSecurityProfileParameters `json:"securityProfile,omitempty" tf:"security_profile,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccount []KafkaClusterStorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountGen2 []KafkaClusterStorageAccountGen2Parameters `json:"storageAccountGen2,omitempty" tf:"storage_account_gen2,omitempty"`

	// +kubebuilder:validation:Optional
	TLSMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

type KafkaClusterRolesHeadNodeObservation struct {
}

type KafkaClusterRolesHeadNodeParameters struct {

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

type KafkaClusterRolesObservation struct {
}

type KafkaClusterRolesParameters struct {

	// +kubebuilder:validation:Required
	HeadNode []KafkaClusterRolesHeadNodeParameters `json:"headNode" tf:"head_node,omitempty"`

	// +kubebuilder:validation:Optional
	KafkaManagementNode []KafkaManagementNodeParameters `json:"kafkaManagementNode,omitempty" tf:"kafka_management_node,omitempty"`

	// +kubebuilder:validation:Required
	WorkerNode []KafkaClusterRolesWorkerNodeParameters `json:"workerNode" tf:"worker_node,omitempty"`

	// +kubebuilder:validation:Required
	ZookeeperNode []KafkaClusterRolesZookeeperNodeParameters `json:"zookeeperNode" tf:"zookeeper_node,omitempty"`
}

type KafkaClusterRolesWorkerNodeObservation struct {
}

type KafkaClusterRolesWorkerNodeParameters struct {

	// +kubebuilder:validation:Optional
	MinInstanceCount *int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`

	// +kubebuilder:validation:Required
	NumberOfDisksPerNode *int64 `json:"numberOfDisksPerNode" tf:"number_of_disks_per_node,omitempty"`

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

type KafkaClusterRolesZookeeperNodeObservation struct {
}

type KafkaClusterRolesZookeeperNodeParameters struct {

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

type KafkaClusterSecurityProfileObservation struct {
}

type KafkaClusterSecurityProfileParameters struct {

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

type KafkaClusterStorageAccountGen2Observation struct {
}

type KafkaClusterStorageAccountGen2Parameters struct {

	// +kubebuilder:validation:Required
	FileSystemID *string `json:"filesystemId" tf:"filesystem_id,omitempty"`

	// +kubebuilder:validation:Required
	IsDefault *bool `json:"isDefault" tf:"is_default,omitempty"`

	// +kubebuilder:validation:Required
	ManagedIdentityResourceID *string `json:"managedIdentityResourceId" tf:"managed_identity_resource_id,omitempty"`

	// +kubebuilder:validation:Required
	StorageResourceID *string `json:"storageResourceId" tf:"storage_resource_id,omitempty"`
}

type KafkaClusterStorageAccountObservation struct {
}

type KafkaClusterStorageAccountParameters struct {

	// +kubebuilder:validation:Required
	IsDefault *bool `json:"isDefault" tf:"is_default,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountKeySecretRef v1.SecretKeySelector `json:"storageAccountKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	StorageContainerID *string `json:"storageContainerId" tf:"storage_container_id,omitempty"`
}

type KafkaManagementNodeObservation struct {
}

type KafkaManagementNodeParameters struct {

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

type RestProxyObservation struct {
}

type RestProxyParameters struct {

	// +kubebuilder:validation:Required
	SecurityGroupID *string `json:"securityGroupId" tf:"security_group_id,omitempty"`
}

// KafkaClusterSpec defines the desired state of KafkaCluster
type KafkaClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KafkaClusterParameters `json:"forProvider"`
}

// KafkaClusterStatus defines the observed state of KafkaCluster.
type KafkaClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KafkaClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaCluster is the Schema for the KafkaClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type KafkaCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KafkaClusterSpec   `json:"spec"`
	Status            KafkaClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaClusterList contains a list of KafkaClusters
type KafkaClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaCluster `json:"items"`
}

// Repository type metadata.
var (
	KafkaCluster_Kind             = "KafkaCluster"
	KafkaCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KafkaCluster_Kind}.String()
	KafkaCluster_KindAPIVersion   = KafkaCluster_Kind + "." + CRDGroupVersion.String()
	KafkaCluster_GroupVersionKind = CRDGroupVersion.WithKind(KafkaCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&KafkaCluster{}, &KafkaClusterList{})
}