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

// GetTerraformResourceType returns Terraform resource type for this SentinelAlertRuleFusion
func (mg *SentinelAlertRuleFusion) GetTerraformResourceType() string {
	return "azurerm_sentinel_alert_rule_fusion"
}

// GetConnectionDetailsMapping for this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelAlertRuleFusion using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelAlertRuleFusion) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelAlertRuleFusionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelAlertRuleFusion) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetTerraformResourceType() string {
	return "azurerm_sentinel_alert_rule_machine_learning_behavior_analytics"
}

// GetConnectionDetailsMapping for this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelAlertRuleMachineLearningBehaviorAnalytics using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelAlertRuleMachineLearningBehaviorAnalyticsParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelAlertRuleMSSecurityIncident
func (mg *SentinelAlertRuleMSSecurityIncident) GetTerraformResourceType() string {
	return "azurerm_sentinel_alert_rule_ms_security_incident"
}

// GetConnectionDetailsMapping for this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelAlertRuleMSSecurityIncident using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelAlertRuleMSSecurityIncident) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelAlertRuleMSSecurityIncidentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelAlertRuleMSSecurityIncident) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelAlertRuleScheduled
func (mg *SentinelAlertRuleScheduled) GetTerraformResourceType() string {
	return "azurerm_sentinel_alert_rule_scheduled"
}

// GetConnectionDetailsMapping for this SentinelAlertRuleScheduled
func (tr *SentinelAlertRuleScheduled) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelAlertRuleScheduled
func (tr *SentinelAlertRuleScheduled) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelAlertRuleScheduled
func (tr *SentinelAlertRuleScheduled) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelAlertRuleScheduled
func (tr *SentinelAlertRuleScheduled) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelAlertRuleScheduled
func (tr *SentinelAlertRuleScheduled) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelAlertRuleScheduled
func (tr *SentinelAlertRuleScheduled) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelAlertRuleScheduled using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelAlertRuleScheduled) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelAlertRuleScheduledParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelAlertRuleScheduled) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelDataConnectorAWSCloudTrail
func (mg *SentinelDataConnectorAWSCloudTrail) GetTerraformResourceType() string {
	return "azurerm_sentinel_data_connector_aws_cloud_trail"
}

// GetConnectionDetailsMapping for this SentinelDataConnectorAWSCloudTrail
func (tr *SentinelDataConnectorAWSCloudTrail) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelDataConnectorAWSCloudTrail
func (tr *SentinelDataConnectorAWSCloudTrail) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelDataConnectorAWSCloudTrail
func (tr *SentinelDataConnectorAWSCloudTrail) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelDataConnectorAWSCloudTrail
func (tr *SentinelDataConnectorAWSCloudTrail) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelDataConnectorAWSCloudTrail
func (tr *SentinelDataConnectorAWSCloudTrail) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelDataConnectorAWSCloudTrail
func (tr *SentinelDataConnectorAWSCloudTrail) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelDataConnectorAWSCloudTrail using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelDataConnectorAWSCloudTrail) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelDataConnectorAWSCloudTrailParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelDataConnectorAWSCloudTrail) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelDataConnectorAzureActiveDirectory
func (mg *SentinelDataConnectorAzureActiveDirectory) GetTerraformResourceType() string {
	return "azurerm_sentinel_data_connector_azure_active_directory"
}

// GetConnectionDetailsMapping for this SentinelDataConnectorAzureActiveDirectory
func (tr *SentinelDataConnectorAzureActiveDirectory) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelDataConnectorAzureActiveDirectory
func (tr *SentinelDataConnectorAzureActiveDirectory) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelDataConnectorAzureActiveDirectory
func (tr *SentinelDataConnectorAzureActiveDirectory) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelDataConnectorAzureActiveDirectory
func (tr *SentinelDataConnectorAzureActiveDirectory) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelDataConnectorAzureActiveDirectory
func (tr *SentinelDataConnectorAzureActiveDirectory) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelDataConnectorAzureActiveDirectory
func (tr *SentinelDataConnectorAzureActiveDirectory) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelDataConnectorAzureActiveDirectory using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelDataConnectorAzureActiveDirectory) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelDataConnectorAzureActiveDirectoryParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelDataConnectorAzureActiveDirectory) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelDataConnectorAzureAdvancedThreatProtection
func (mg *SentinelDataConnectorAzureAdvancedThreatProtection) GetTerraformResourceType() string {
	return "azurerm_sentinel_data_connector_azure_advanced_threat_protection"
}

// GetConnectionDetailsMapping for this SentinelDataConnectorAzureAdvancedThreatProtection
func (tr *SentinelDataConnectorAzureAdvancedThreatProtection) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelDataConnectorAzureAdvancedThreatProtection
func (tr *SentinelDataConnectorAzureAdvancedThreatProtection) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelDataConnectorAzureAdvancedThreatProtection
func (tr *SentinelDataConnectorAzureAdvancedThreatProtection) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelDataConnectorAzureAdvancedThreatProtection
func (tr *SentinelDataConnectorAzureAdvancedThreatProtection) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelDataConnectorAzureAdvancedThreatProtection
func (tr *SentinelDataConnectorAzureAdvancedThreatProtection) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelDataConnectorAzureAdvancedThreatProtection
func (tr *SentinelDataConnectorAzureAdvancedThreatProtection) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelDataConnectorAzureAdvancedThreatProtection using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelDataConnectorAzureAdvancedThreatProtection) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelDataConnectorAzureAdvancedThreatProtectionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelDataConnectorAzureAdvancedThreatProtection) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelDataConnectorAzureSecurityCenter
func (mg *SentinelDataConnectorAzureSecurityCenter) GetTerraformResourceType() string {
	return "azurerm_sentinel_data_connector_azure_security_center"
}

// GetConnectionDetailsMapping for this SentinelDataConnectorAzureSecurityCenter
func (tr *SentinelDataConnectorAzureSecurityCenter) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelDataConnectorAzureSecurityCenter
func (tr *SentinelDataConnectorAzureSecurityCenter) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelDataConnectorAzureSecurityCenter
func (tr *SentinelDataConnectorAzureSecurityCenter) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelDataConnectorAzureSecurityCenter
func (tr *SentinelDataConnectorAzureSecurityCenter) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelDataConnectorAzureSecurityCenter
func (tr *SentinelDataConnectorAzureSecurityCenter) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelDataConnectorAzureSecurityCenter
func (tr *SentinelDataConnectorAzureSecurityCenter) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelDataConnectorAzureSecurityCenter using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelDataConnectorAzureSecurityCenter) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelDataConnectorAzureSecurityCenterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelDataConnectorAzureSecurityCenter) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelDataConnectorMicrosoftCloudAppSecurity
func (mg *SentinelDataConnectorMicrosoftCloudAppSecurity) GetTerraformResourceType() string {
	return "azurerm_sentinel_data_connector_microsoft_cloud_app_security"
}

// GetConnectionDetailsMapping for this SentinelDataConnectorMicrosoftCloudAppSecurity
func (tr *SentinelDataConnectorMicrosoftCloudAppSecurity) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelDataConnectorMicrosoftCloudAppSecurity
func (tr *SentinelDataConnectorMicrosoftCloudAppSecurity) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelDataConnectorMicrosoftCloudAppSecurity
func (tr *SentinelDataConnectorMicrosoftCloudAppSecurity) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelDataConnectorMicrosoftCloudAppSecurity
func (tr *SentinelDataConnectorMicrosoftCloudAppSecurity) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelDataConnectorMicrosoftCloudAppSecurity
func (tr *SentinelDataConnectorMicrosoftCloudAppSecurity) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelDataConnectorMicrosoftCloudAppSecurity
func (tr *SentinelDataConnectorMicrosoftCloudAppSecurity) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelDataConnectorMicrosoftCloudAppSecurity using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelDataConnectorMicrosoftCloudAppSecurity) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelDataConnectorMicrosoftCloudAppSecurityParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelDataConnectorMicrosoftCloudAppSecurity) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelDataConnectorOffice365
func (mg *SentinelDataConnectorOffice365) GetTerraformResourceType() string {
	return "azurerm_sentinel_data_connector_office_365"
}

// GetConnectionDetailsMapping for this SentinelDataConnectorOffice365
func (tr *SentinelDataConnectorOffice365) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelDataConnectorOffice365
func (tr *SentinelDataConnectorOffice365) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelDataConnectorOffice365
func (tr *SentinelDataConnectorOffice365) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelDataConnectorOffice365
func (tr *SentinelDataConnectorOffice365) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelDataConnectorOffice365
func (tr *SentinelDataConnectorOffice365) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelDataConnectorOffice365
func (tr *SentinelDataConnectorOffice365) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelDataConnectorOffice365 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelDataConnectorOffice365) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelDataConnectorOffice365Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelDataConnectorOffice365) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelDataConnectorThreatIntelligence
func (mg *SentinelDataConnectorThreatIntelligence) GetTerraformResourceType() string {
	return "azurerm_sentinel_data_connector_threat_intelligence"
}

// GetConnectionDetailsMapping for this SentinelDataConnectorThreatIntelligence
func (tr *SentinelDataConnectorThreatIntelligence) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelDataConnectorThreatIntelligence
func (tr *SentinelDataConnectorThreatIntelligence) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelDataConnectorThreatIntelligence
func (tr *SentinelDataConnectorThreatIntelligence) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelDataConnectorThreatIntelligence
func (tr *SentinelDataConnectorThreatIntelligence) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelDataConnectorThreatIntelligence
func (tr *SentinelDataConnectorThreatIntelligence) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelDataConnectorThreatIntelligence
func (tr *SentinelDataConnectorThreatIntelligence) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelDataConnectorThreatIntelligence using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelDataConnectorThreatIntelligence) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelDataConnectorThreatIntelligenceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelDataConnectorThreatIntelligence) GetTerraformSchemaVersion() int {
	return 0
}
