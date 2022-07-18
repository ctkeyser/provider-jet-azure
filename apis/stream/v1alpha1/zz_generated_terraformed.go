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

// GetTerraformResourceType returns Terraform resource type for this AnalyticsOutputTable
func (mg *AnalyticsOutputTable) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_output_table"
}

// GetConnectionDetailsMapping for this AnalyticsOutputTable
func (tr *AnalyticsOutputTable) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"storage_account_key": "spec.forProvider.storageAccountKeySecretRef"}
}

// GetObservation of this AnalyticsOutputTable
func (tr *AnalyticsOutputTable) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AnalyticsOutputTable
func (tr *AnalyticsOutputTable) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AnalyticsOutputTable
func (tr *AnalyticsOutputTable) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AnalyticsOutputTable
func (tr *AnalyticsOutputTable) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AnalyticsOutputTable
func (tr *AnalyticsOutputTable) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AnalyticsOutputTable using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AnalyticsOutputTable) LateInitialize(attrs []byte) (bool, error) {
	params := &AnalyticsOutputTableParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AnalyticsOutputTable) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AnalyticsReferenceInputMSSQL
func (mg *AnalyticsReferenceInputMSSQL) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_reference_input_mssql"
}

// GetConnectionDetailsMapping for this AnalyticsReferenceInputMSSQL
func (tr *AnalyticsReferenceInputMSSQL) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"password": "spec.forProvider.passwordSecretRef"}
}

// GetObservation of this AnalyticsReferenceInputMSSQL
func (tr *AnalyticsReferenceInputMSSQL) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AnalyticsReferenceInputMSSQL
func (tr *AnalyticsReferenceInputMSSQL) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AnalyticsReferenceInputMSSQL
func (tr *AnalyticsReferenceInputMSSQL) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AnalyticsReferenceInputMSSQL
func (tr *AnalyticsReferenceInputMSSQL) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AnalyticsReferenceInputMSSQL
func (tr *AnalyticsReferenceInputMSSQL) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AnalyticsReferenceInputMSSQL using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AnalyticsReferenceInputMSSQL) LateInitialize(attrs []byte) (bool, error) {
	params := &AnalyticsReferenceInputMSSQLParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AnalyticsReferenceInputMSSQL) GetTerraformSchemaVersion() int {
	return 0
}
