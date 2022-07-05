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

type BotChannelSlackObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BotChannelSlackParameters struct {

	// +kubebuilder:validation:Required
	BotName *string `json:"botName" tf:"bot_name,omitempty"`

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	LandingPageURL *string `json:"landingPageUrl,omitempty" tf:"landing_page_url,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SigningSecretSecretRef *v1.SecretKeySelector `json:"signingSecretSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	VerificationTokenSecretRef v1.SecretKeySelector `json:"verificationTokenSecretRef" tf:"-"`
}

// BotChannelSlackSpec defines the desired state of BotChannelSlack
type BotChannelSlackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotChannelSlackParameters `json:"forProvider"`
}

// BotChannelSlackStatus defines the observed state of BotChannelSlack.
type BotChannelSlackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotChannelSlackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelSlack is the Schema for the BotChannelSlacks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type BotChannelSlack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotChannelSlackSpec   `json:"spec"`
	Status            BotChannelSlackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelSlackList contains a list of BotChannelSlacks
type BotChannelSlackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelSlack `json:"items"`
}

// Repository type metadata.
var (
	BotChannelSlack_Kind             = "BotChannelSlack"
	BotChannelSlack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotChannelSlack_Kind}.String()
	BotChannelSlack_KindAPIVersion   = BotChannelSlack_Kind + "." + CRDGroupVersion.String()
	BotChannelSlack_GroupVersionKind = CRDGroupVersion.WithKind(BotChannelSlack_Kind)
)

func init() {
	SchemeBuilder.Register(&BotChannelSlack{}, &BotChannelSlackList{})
}
