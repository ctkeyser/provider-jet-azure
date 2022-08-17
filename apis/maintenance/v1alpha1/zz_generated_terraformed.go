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
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this MaintenanceAssignmentDedicatedHost
func (mg *MaintenanceAssignmentDedicatedHost) GetTerraformResourceType() string {
	return "azurerm_maintenance_assignment_dedicated_host"
}

// GetConnectionDetailsMapping for this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MaintenanceAssignmentDedicatedHost using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MaintenanceAssignmentDedicatedHost) LateInitialize(attrs []byte) (bool, error) {
	params := &MaintenanceAssignmentDedicatedHostParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MaintenanceAssignmentDedicatedHost) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MaintenanceAssignmentVirtualMachine
func (mg *MaintenanceAssignmentVirtualMachine) GetTerraformResourceType() string {
	return "azurerm_maintenance_assignment_virtual_machine"
}

// GetConnectionDetailsMapping for this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MaintenanceAssignmentVirtualMachine using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MaintenanceAssignmentVirtualMachine) LateInitialize(attrs []byte) (bool, error) {
	params := &MaintenanceAssignmentVirtualMachineParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MaintenanceAssignmentVirtualMachine) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MaintenanceAssignmentVirtualMachineScaleSet
func (mg *MaintenanceAssignmentVirtualMachineScaleSet) GetTerraformResourceType() string {
	return "azurerm_maintenance_assignment_virtual_machine_scale_set"
}

// GetConnectionDetailsMapping for this MaintenanceAssignmentVirtualMachineScaleSet
func (tr *MaintenanceAssignmentVirtualMachineScaleSet) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MaintenanceAssignmentVirtualMachineScaleSet
func (tr *MaintenanceAssignmentVirtualMachineScaleSet) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MaintenanceAssignmentVirtualMachineScaleSet
func (tr *MaintenanceAssignmentVirtualMachineScaleSet) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MaintenanceAssignmentVirtualMachineScaleSet
func (tr *MaintenanceAssignmentVirtualMachineScaleSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MaintenanceAssignmentVirtualMachineScaleSet
func (tr *MaintenanceAssignmentVirtualMachineScaleSet) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MaintenanceAssignmentVirtualMachineScaleSet
func (tr *MaintenanceAssignmentVirtualMachineScaleSet) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MaintenanceAssignmentVirtualMachineScaleSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MaintenanceAssignmentVirtualMachineScaleSet) LateInitialize(attrs []byte) (bool, error) {
	params := &MaintenanceAssignmentVirtualMachineScaleSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MaintenanceAssignmentVirtualMachineScaleSet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MaintenanceConfiguration
func (mg *MaintenanceConfiguration) GetTerraformResourceType() string {
	return "azurerm_maintenance_configuration"
}

// GetConnectionDetailsMapping for this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MaintenanceConfiguration using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MaintenanceConfiguration) LateInitialize(attrs []byte) (bool, error) {
	params := &MaintenanceConfigurationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MaintenanceConfiguration) GetTerraformSchemaVersion() int {
	return 1
}