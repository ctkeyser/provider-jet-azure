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

type ContactObservation struct {
}

type ContactParameters struct {

	// +kubebuilder:validation:Required
	CompanyName *string `json:"companyName" tf:"company_name,omitempty"`

	// +kubebuilder:validation:Required
	Emails []*string `json:"emails" tf:"emails,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PhoneNumber *string `json:"phoneNumber" tf:"phone_number,omitempty"`
}

type OrderObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ReturnTracking []ReturnTrackingObservation `json:"returnTracking,omitempty" tf:"return_tracking,omitempty"`

	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	ShipmentHistory []ShipmentHistoryObservation `json:"shipmentHistory,omitempty" tf:"shipment_history,omitempty"`

	ShipmentTracking []ShipmentTrackingObservation `json:"shipmentTracking,omitempty" tf:"shipment_tracking,omitempty"`

	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`
}

type OrderParameters struct {

	// +kubebuilder:validation:Required
	Contact []ContactParameters `json:"contact" tf:"contact,omitempty"`

	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ShipmentAddress []ShipmentAddressParameters `json:"shipmentAddress" tf:"shipment_address,omitempty"`
}

type ReturnTrackingObservation struct {
	CarrierName *string `json:"carrierName,omitempty" tf:"carrier_name,omitempty"`

	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	TrackingID *string `json:"trackingId,omitempty" tf:"tracking_id,omitempty"`

	TrackingURL *string `json:"trackingUrl,omitempty" tf:"tracking_url,omitempty"`
}

type ReturnTrackingParameters struct {
}

type ShipmentAddressObservation struct {
}

type ShipmentAddressParameters struct {

	// +kubebuilder:validation:Required
	Address []*string `json:"address" tf:"address,omitempty"`

	// +kubebuilder:validation:Required
	City *string `json:"city" tf:"city,omitempty"`

	// +kubebuilder:validation:Required
	Country *string `json:"country" tf:"country,omitempty"`

	// +kubebuilder:validation:Required
	PostalCode *string `json:"postalCode" tf:"postal_code,omitempty"`

	// +kubebuilder:validation:Required
	State *string `json:"state" tf:"state,omitempty"`
}

type ShipmentHistoryObservation struct {
	AdditionalDetails map[string]*string `json:"additionalDetails,omitempty" tf:"additional_details,omitempty"`

	Comments *string `json:"comments,omitempty" tf:"comments,omitempty"`

	LastUpdate *string `json:"lastUpdate,omitempty" tf:"last_update,omitempty"`
}

type ShipmentHistoryParameters struct {
}

type ShipmentTrackingObservation struct {
	CarrierName *string `json:"carrierName,omitempty" tf:"carrier_name,omitempty"`

	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	TrackingID *string `json:"trackingId,omitempty" tf:"tracking_id,omitempty"`

	TrackingURL *string `json:"trackingUrl,omitempty" tf:"tracking_url,omitempty"`
}

type ShipmentTrackingParameters struct {
}

type StatusObservation struct {
	AdditionalDetails map[string]*string `json:"additionalDetails,omitempty" tf:"additional_details,omitempty"`

	Comments *string `json:"comments,omitempty" tf:"comments,omitempty"`

	Info *string `json:"info,omitempty" tf:"info,omitempty"`

	LastUpdate *string `json:"lastUpdate,omitempty" tf:"last_update,omitempty"`
}

type StatusParameters struct {
}

// OrderSpec defines the desired state of Order
type OrderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrderParameters `json:"forProvider"`
}

// OrderStatus defines the observed state of Order.
type OrderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Order is the Schema for the Orders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Order struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrderSpec   `json:"spec"`
	Status            OrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrderList contains a list of Orders
type OrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Order `json:"items"`
}

// Repository type metadata.
var (
	Order_Kind             = "Order"
	Order_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Order_Kind}.String()
	Order_KindAPIVersion   = Order_Kind + "." + CRDGroupVersion.String()
	Order_GroupVersionKind = CRDGroupVersion.WithKind(Order_Kind)
)

func init() {
	SchemeBuilder.Register(&Order{}, &OrderList{})
}