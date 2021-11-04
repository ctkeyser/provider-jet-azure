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
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/terrajet/pkg/resource"
	"github.com/crossplane-contrib/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ProtectionBackupInstancePostgresql
func (mg *ProtectionBackupInstancePostgresql) GetTerraformResourceType() string {
	return "azurerm_data_protection_backup_instance_postgresql"
}

// GetTerraformResourceIDField returns Terraform identifier field for this ProtectionBackupInstancePostgresql
func (tr *ProtectionBackupInstancePostgresql) GetTerraformResourceIDField() string {
	return "id"
}

// GetConnectionDetailsMapping for this ProtectionBackupInstancePostgresql
func (tr *ProtectionBackupInstancePostgresql) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ProtectionBackupInstancePostgresql
func (tr *ProtectionBackupInstancePostgresql) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ProtectionBackupInstancePostgresql
func (tr *ProtectionBackupInstancePostgresql) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetParameters of this ProtectionBackupInstancePostgresql
func (tr *ProtectionBackupInstancePostgresql) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ProtectionBackupInstancePostgresql
func (tr *ProtectionBackupInstancePostgresql) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ProtectionBackupInstancePostgresql using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ProtectionBackupInstancePostgresql) LateInitialize(attrs []byte) (bool, error) {
	params := &ProtectionBackupInstancePostgresqlParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ProtectionBackupInstancePostgresql) GetTerraformSchemaVersion() int {
	return 0
}
