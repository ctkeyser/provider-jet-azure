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

type MaintenanceAssignmentVirtualMachineObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MaintenanceAssignmentVirtualMachineParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	MaintenanceConfigurationID *string `json:"maintenanceConfigurationId" tf:"maintenance_configuration_id,omitempty"`

	// +kubebuilder:validation:Required
	VirtualMachineID *string `json:"virtualMachineId" tf:"virtual_machine_id,omitempty"`
}

// MaintenanceAssignmentVirtualMachineSpec defines the desired state of MaintenanceAssignmentVirtualMachine
type MaintenanceAssignmentVirtualMachineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MaintenanceAssignmentVirtualMachineParameters `json:"forProvider"`
}

// MaintenanceAssignmentVirtualMachineStatus defines the observed state of MaintenanceAssignmentVirtualMachine.
type MaintenanceAssignmentVirtualMachineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MaintenanceAssignmentVirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceAssignmentVirtualMachine is the Schema for the MaintenanceAssignmentVirtualMachines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type MaintenanceAssignmentVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MaintenanceAssignmentVirtualMachineSpec   `json:"spec"`
	Status            MaintenanceAssignmentVirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceAssignmentVirtualMachineList contains a list of MaintenanceAssignmentVirtualMachines
type MaintenanceAssignmentVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaintenanceAssignmentVirtualMachine `json:"items"`
}

// Repository type metadata.
var (
	MaintenanceAssignmentVirtualMachine_Kind             = "MaintenanceAssignmentVirtualMachine"
	MaintenanceAssignmentVirtualMachine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MaintenanceAssignmentVirtualMachine_Kind}.String()
	MaintenanceAssignmentVirtualMachine_KindAPIVersion   = MaintenanceAssignmentVirtualMachine_Kind + "." + CRDGroupVersion.String()
	MaintenanceAssignmentVirtualMachine_GroupVersionKind = CRDGroupVersion.WithKind(MaintenanceAssignmentVirtualMachine_Kind)
)

func init() {
	SchemeBuilder.Register(&MaintenanceAssignmentVirtualMachine{}, &MaintenanceAssignmentVirtualMachineList{})
}
