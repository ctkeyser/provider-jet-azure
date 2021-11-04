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

type AdvancedFilterBoolEqualsObservation struct {
}

type AdvancedFilterBoolEqualsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *bool `json:"value" tf:"value,omitempty"`
}

type AdvancedFilterIsNotNullObservation struct {
}

type AdvancedFilterIsNotNullParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type AdvancedFilterIsNullOrUndefinedObservation struct {
}

type AdvancedFilterIsNullOrUndefinedParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type AdvancedFilterNumberGreaterThanObservation struct {
}

type AdvancedFilterNumberGreaterThanOrEqualsObservation struct {
}

type AdvancedFilterNumberGreaterThanOrEqualsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type AdvancedFilterNumberGreaterThanParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type AdvancedFilterNumberInObservation struct {
}

type AdvancedFilterNumberInParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*float64 `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterNumberInRangeObservation struct {
}

type AdvancedFilterNumberInRangeParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values [][]*float64 `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterNumberLessThanObservation struct {
}

type AdvancedFilterNumberLessThanOrEqualsObservation struct {
}

type AdvancedFilterNumberLessThanOrEqualsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type AdvancedFilterNumberLessThanParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type AdvancedFilterNumberNotInObservation struct {
}

type AdvancedFilterNumberNotInParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*float64 `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterNumberNotInRangeObservation struct {
}

type AdvancedFilterNumberNotInRangeParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values [][]*float64 `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterStringBeginsWithObservation struct {
}

type AdvancedFilterStringBeginsWithParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterStringContainsObservation struct {
}

type AdvancedFilterStringContainsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterStringEndsWithObservation struct {
}

type AdvancedFilterStringEndsWithParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterStringInObservation struct {
}

type AdvancedFilterStringInParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterStringNotBeginsWithObservation struct {
}

type AdvancedFilterStringNotBeginsWithParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterStringNotContainsObservation struct {
}

type AdvancedFilterStringNotContainsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterStringNotEndsWithObservation struct {
}

type AdvancedFilterStringNotEndsWithParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AdvancedFilterStringNotInObservation struct {
}

type AdvancedFilterStringNotInParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type SystemTopicEventSubscriptionAdvancedFilterObservation struct {
}

type SystemTopicEventSubscriptionAdvancedFilterParameters struct {

	// +kubebuilder:validation:Optional
	BoolEquals []AdvancedFilterBoolEqualsParameters `json:"boolEquals,omitempty" tf:"bool_equals,omitempty"`

	// +kubebuilder:validation:Optional
	IsNotNull []AdvancedFilterIsNotNullParameters `json:"isNotNull,omitempty" tf:"is_not_null,omitempty"`

	// +kubebuilder:validation:Optional
	IsNullOrUndefined []AdvancedFilterIsNullOrUndefinedParameters `json:"isNullOrUndefined,omitempty" tf:"is_null_or_undefined,omitempty"`

	// +kubebuilder:validation:Optional
	NumberGreaterThan []AdvancedFilterNumberGreaterThanParameters `json:"numberGreaterThan,omitempty" tf:"number_greater_than,omitempty"`

	// +kubebuilder:validation:Optional
	NumberGreaterThanOrEquals []AdvancedFilterNumberGreaterThanOrEqualsParameters `json:"numberGreaterThanOrEquals,omitempty" tf:"number_greater_than_or_equals,omitempty"`

	// +kubebuilder:validation:Optional
	NumberIn []AdvancedFilterNumberInParameters `json:"numberIn,omitempty" tf:"number_in,omitempty"`

	// +kubebuilder:validation:Optional
	NumberInRange []AdvancedFilterNumberInRangeParameters `json:"numberInRange,omitempty" tf:"number_in_range,omitempty"`

	// +kubebuilder:validation:Optional
	NumberLessThan []AdvancedFilterNumberLessThanParameters `json:"numberLessThan,omitempty" tf:"number_less_than,omitempty"`

	// +kubebuilder:validation:Optional
	NumberLessThanOrEquals []AdvancedFilterNumberLessThanOrEqualsParameters `json:"numberLessThanOrEquals,omitempty" tf:"number_less_than_or_equals,omitempty"`

	// +kubebuilder:validation:Optional
	NumberNotIn []AdvancedFilterNumberNotInParameters `json:"numberNotIn,omitempty" tf:"number_not_in,omitempty"`

	// +kubebuilder:validation:Optional
	NumberNotInRange []AdvancedFilterNumberNotInRangeParameters `json:"numberNotInRange,omitempty" tf:"number_not_in_range,omitempty"`

	// +kubebuilder:validation:Optional
	StringBeginsWith []AdvancedFilterStringBeginsWithParameters `json:"stringBeginsWith,omitempty" tf:"string_begins_with,omitempty"`

	// +kubebuilder:validation:Optional
	StringContains []AdvancedFilterStringContainsParameters `json:"stringContains,omitempty" tf:"string_contains,omitempty"`

	// +kubebuilder:validation:Optional
	StringEndsWith []AdvancedFilterStringEndsWithParameters `json:"stringEndsWith,omitempty" tf:"string_ends_with,omitempty"`

	// +kubebuilder:validation:Optional
	StringIn []AdvancedFilterStringInParameters `json:"stringIn,omitempty" tf:"string_in,omitempty"`

	// +kubebuilder:validation:Optional
	StringNotBeginsWith []AdvancedFilterStringNotBeginsWithParameters `json:"stringNotBeginsWith,omitempty" tf:"string_not_begins_with,omitempty"`

	// +kubebuilder:validation:Optional
	StringNotContains []AdvancedFilterStringNotContainsParameters `json:"stringNotContains,omitempty" tf:"string_not_contains,omitempty"`

	// +kubebuilder:validation:Optional
	StringNotEndsWith []AdvancedFilterStringNotEndsWithParameters `json:"stringNotEndsWith,omitempty" tf:"string_not_ends_with,omitempty"`

	// +kubebuilder:validation:Optional
	StringNotIn []AdvancedFilterStringNotInParameters `json:"stringNotIn,omitempty" tf:"string_not_in,omitempty"`
}

type SystemTopicEventSubscriptionAzureFunctionEndpointObservation struct {
}

type SystemTopicEventSubscriptionAzureFunctionEndpointParameters struct {

	// +kubebuilder:validation:Required
	FunctionID *string `json:"functionId" tf:"function_id,omitempty"`

	// +kubebuilder:validation:Optional
	MaxEventsPerBatch *int64 `json:"maxEventsPerBatch,omitempty" tf:"max_events_per_batch,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredBatchSizeInKilobytes *int64 `json:"preferredBatchSizeInKilobytes,omitempty" tf:"preferred_batch_size_in_kilobytes,omitempty"`
}

type SystemTopicEventSubscriptionDeadLetterIdentityObservation struct {
}

type SystemTopicEventSubscriptionDeadLetterIdentityParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SystemTopicEventSubscriptionDeliveryIdentityObservation struct {
}

type SystemTopicEventSubscriptionDeliveryIdentityParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SystemTopicEventSubscriptionObservation struct {
}

type SystemTopicEventSubscriptionParameters struct {

	// +kubebuilder:validation:Optional
	AdvancedFilter []SystemTopicEventSubscriptionAdvancedFilterParameters `json:"advancedFilter,omitempty" tf:"advanced_filter,omitempty"`

	// +kubebuilder:validation:Optional
	AdvancedFilteringOnArraysEnabled *bool `json:"advancedFilteringOnArraysEnabled,omitempty" tf:"advanced_filtering_on_arrays_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	AzureFunctionEndpoint []SystemTopicEventSubscriptionAzureFunctionEndpointParameters `json:"azureFunctionEndpoint,omitempty" tf:"azure_function_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	DeadLetterIdentity []SystemTopicEventSubscriptionDeadLetterIdentityParameters `json:"deadLetterIdentity,omitempty" tf:"dead_letter_identity,omitempty"`

	// +kubebuilder:validation:Optional
	DeliveryIdentity []SystemTopicEventSubscriptionDeliveryIdentityParameters `json:"deliveryIdentity,omitempty" tf:"delivery_identity,omitempty"`

	// +kubebuilder:validation:Optional
	EventDeliverySchema *string `json:"eventDeliverySchema,omitempty" tf:"event_delivery_schema,omitempty"`

	// +kubebuilder:validation:Optional
	EventhubEndpointID *string `json:"eventhubEndpointId,omitempty" tf:"eventhub_endpoint_id,omitempty"`

	// +kubebuilder:validation:Optional
	ExpirationTimeUtc *string `json:"expirationTimeUtc,omitempty" tf:"expiration_time_utc,omitempty"`

	// +kubebuilder:validation:Optional
	HybridConnectionEndpointID *string `json:"hybridConnectionEndpointId,omitempty" tf:"hybrid_connection_endpoint_id,omitempty"`

	// +kubebuilder:validation:Optional
	IncludedEventTypes []*string `json:"includedEventTypes,omitempty" tf:"included_event_types,omitempty"`

	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	RetryPolicy []SystemTopicEventSubscriptionRetryPolicyParameters `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceBusQueueEndpointID *string `json:"serviceBusQueueEndpointId,omitempty" tf:"service_bus_queue_endpoint_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceBusTopicEndpointID *string `json:"serviceBusTopicEndpointId,omitempty" tf:"service_bus_topic_endpoint_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageBlobDeadLetterDestination []SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationParameters `json:"storageBlobDeadLetterDestination,omitempty" tf:"storage_blob_dead_letter_destination,omitempty"`

	// +kubebuilder:validation:Optional
	StorageQueueEndpoint []SystemTopicEventSubscriptionStorageQueueEndpointParameters `json:"storageQueueEndpoint,omitempty" tf:"storage_queue_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	SubjectFilter []SystemTopicEventSubscriptionSubjectFilterParameters `json:"subjectFilter,omitempty" tf:"subject_filter,omitempty"`

	// +kubebuilder:validation:Required
	SystemTopic *string `json:"systemTopic" tf:"system_topic,omitempty"`

	// +kubebuilder:validation:Optional
	WebhookEndpoint []SystemTopicEventSubscriptionWebhookEndpointParameters `json:"webhookEndpoint,omitempty" tf:"webhook_endpoint,omitempty"`
}

type SystemTopicEventSubscriptionRetryPolicyObservation struct {
}

type SystemTopicEventSubscriptionRetryPolicyParameters struct {

	// +kubebuilder:validation:Required
	EventTimeToLive *int64 `json:"eventTimeToLive" tf:"event_time_to_live,omitempty"`

	// +kubebuilder:validation:Required
	MaxDeliveryAttempts *int64 `json:"maxDeliveryAttempts" tf:"max_delivery_attempts,omitempty"`
}

type SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationObservation struct {
}

type SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationParameters struct {

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Required
	StorageBlobContainerName *string `json:"storageBlobContainerName" tf:"storage_blob_container_name,omitempty"`
}

type SystemTopicEventSubscriptionStorageQueueEndpointObservation struct {
}

type SystemTopicEventSubscriptionStorageQueueEndpointParameters struct {

	// +kubebuilder:validation:Required
	QueueName *string `json:"queueName" tf:"queue_name,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

type SystemTopicEventSubscriptionSubjectFilterObservation struct {
}

type SystemTopicEventSubscriptionSubjectFilterParameters struct {

	// +kubebuilder:validation:Optional
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// +kubebuilder:validation:Optional
	SubjectBeginsWith *string `json:"subjectBeginsWith,omitempty" tf:"subject_begins_with,omitempty"`

	// +kubebuilder:validation:Optional
	SubjectEndsWith *string `json:"subjectEndsWith,omitempty" tf:"subject_ends_with,omitempty"`
}

type SystemTopicEventSubscriptionWebhookEndpointObservation struct {
	BaseURL *string `json:"baseUrl,omitempty" tf:"base_url,omitempty"`
}

type SystemTopicEventSubscriptionWebhookEndpointParameters struct {

	// +kubebuilder:validation:Optional
	ActiveDirectoryAppIDOrURI *string `json:"activeDirectoryAppIdOrUri,omitempty" tf:"active_directory_app_id_or_uri,omitempty"`

	// +kubebuilder:validation:Optional
	ActiveDirectoryTenantID *string `json:"activeDirectoryTenantId,omitempty" tf:"active_directory_tenant_id,omitempty"`

	// +kubebuilder:validation:Optional
	MaxEventsPerBatch *int64 `json:"maxEventsPerBatch,omitempty" tf:"max_events_per_batch,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredBatchSizeInKilobytes *int64 `json:"preferredBatchSizeInKilobytes,omitempty" tf:"preferred_batch_size_in_kilobytes,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

// SystemTopicEventSubscriptionSpec defines the desired state of SystemTopicEventSubscription
type SystemTopicEventSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SystemTopicEventSubscriptionParameters `json:"forProvider"`
}

// SystemTopicEventSubscriptionStatus defines the observed state of SystemTopicEventSubscription.
type SystemTopicEventSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SystemTopicEventSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SystemTopicEventSubscription is the Schema for the SystemTopicEventSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type SystemTopicEventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SystemTopicEventSubscriptionSpec   `json:"spec"`
	Status            SystemTopicEventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SystemTopicEventSubscriptionList contains a list of SystemTopicEventSubscriptions
type SystemTopicEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SystemTopicEventSubscription `json:"items"`
}

// Repository type metadata.
var (
	SystemTopicEventSubscriptionKind             = "SystemTopicEventSubscription"
	SystemTopicEventSubscriptionGroupKind        = schema.GroupKind{Group: Group, Kind: SystemTopicEventSubscriptionKind}.String()
	SystemTopicEventSubscriptionKindAPIVersion   = SystemTopicEventSubscriptionKind + "." + GroupVersion.String()
	SystemTopicEventSubscriptionGroupVersionKind = GroupVersion.WithKind(SystemTopicEventSubscriptionKind)
)

func init() {
	SchemeBuilder.Register(&SystemTopicEventSubscription{}, &SystemTopicEventSubscriptionList{})
}
