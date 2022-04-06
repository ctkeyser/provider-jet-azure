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

type CustomParametersObservation struct {
}

type CustomParametersParameters struct {

	// +kubebuilder:validation:Optional
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceId,omitempty" tf:"machine_learning_workspace_id,omitempty"`

	// +kubebuilder:validation:Optional
	NATGatewayName *string `json:"natGatewayName,omitempty" tf:"nat_gateway_name,omitempty"`

	// +kubebuilder:validation:Optional
	NoPublicIP *bool `json:"noPublicIp,omitempty" tf:"no_public_ip,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +kubebuilder:validation:Optional
	PrivateSubnetName *string `json:"privateSubnetName,omitempty" tf:"private_subnet_name,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateSubnetNameRef *v1.Reference `json:"privateSubnetNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrivateSubnetNameSelector *v1.Selector `json:"privateSubnetNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrivateSubnetNetworkSecurityGroupAssociationID *string `json:"privateSubnetNetworkSecurityGroupAssociationId,omitempty" tf:"private_subnet_network_security_group_association_id,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIPName *string `json:"publicIpName,omitempty" tf:"public_ip_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +kubebuilder:validation:Optional
	PublicSubnetName *string `json:"publicSubnetName,omitempty" tf:"public_subnet_name,omitempty"`

	// +kubebuilder:validation:Optional
	PublicSubnetNameRef *v1.Reference `json:"publicSubnetNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PublicSubnetNameSelector *v1.Selector `json:"publicSubnetNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PublicSubnetNetworkSecurityGroupAssociationID *string `json:"publicSubnetNetworkSecurityGroupAssociationId,omitempty" tf:"public_subnet_network_security_group_association_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountSkuName *string `json:"storageAccountSkuName,omitempty" tf:"storage_account_sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`

	// +kubebuilder:validation:Optional
	VnetAddressPrefix *string `json:"vnetAddressPrefix,omitempty" tf:"vnet_address_prefix,omitempty"`
}

type StorageAccountIdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StorageAccountIdentityParameters struct {
}

type WorkspaceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	StorageAccountIdentity []StorageAccountIdentityObservation `json:"storageAccountIdentity,omitempty" tf:"storage_account_identity,omitempty"`

	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	WorkspaceURL *string `json:"workspaceUrl,omitempty" tf:"workspace_url,omitempty"`
}

type WorkspaceParameters struct {

	// +kubebuilder:validation:Optional
	CustomParameters []CustomParametersParameters `json:"customParameters,omitempty" tf:"custom_parameters,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerManagedKeyEnabled *bool `json:"customerManagedKeyEnabled,omitempty" tf:"customer_managed_key_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	InfrastructureEncryptionEnabled *bool `json:"infrastructureEncryptionEnabled,omitempty" tf:"infrastructure_encryption_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerBackendAddressPoolID *string `json:"loadBalancerBackendAddressPoolId,omitempty" tf:"load_balancer_backend_address_pool_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagedResourceGroupID *string `json:"managedResourceGroupId,omitempty" tf:"managed_resource_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ManagedResourceGroupIDRef *v1.Reference `json:"managedResourceGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ManagedResourceGroupIDSelector *v1.Selector `json:"managedResourceGroupIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty" tf:"managed_resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ManagedResourceGroupNameRef *v1.Reference `json:"managedResourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ManagedResourceGroupNameSelector *v1.Selector `json:"managedResourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ManagedServicesCmkKeyVaultKeyID *string `json:"managedServicesCmkKeyVaultKeyId,omitempty" tf:"managed_services_cmk_key_vault_key_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkSecurityGroupRulesRequired *string `json:"networkSecurityGroupRulesRequired,omitempty" tf:"network_security_group_rules_required,omitempty"`

	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters `json:"forProvider"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workspace is the Schema for the Workspaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Repository type metadata.
var (
	Workspace_Kind             = "Workspace"
	Workspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workspace_Kind}.String()
	Workspace_KindAPIVersion   = Workspace_Kind + "." + CRDGroupVersion.String()
	Workspace_GroupVersionKind = CRDGroupVersion.WithKind(Workspace_Kind)
)

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}