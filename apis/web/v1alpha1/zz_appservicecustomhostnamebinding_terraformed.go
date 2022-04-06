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

// GetTerraformResourceType returns Terraform resource type for this AppServiceCustomHostNameBinding
func (mg *AppServiceCustomHostNameBinding) GetTerraformResourceType() string {
	return "azurerm_app_service_custom_hostname_binding"
}

// GetConnectionDetailsMapping for this AppServiceCustomHostNameBinding
func (tr *AppServiceCustomHostNameBinding) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppServiceCustomHostNameBinding
func (tr *AppServiceCustomHostNameBinding) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppServiceCustomHostNameBinding
func (tr *AppServiceCustomHostNameBinding) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppServiceCustomHostNameBinding
func (tr *AppServiceCustomHostNameBinding) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppServiceCustomHostNameBinding
func (tr *AppServiceCustomHostNameBinding) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppServiceCustomHostNameBinding
func (tr *AppServiceCustomHostNameBinding) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AppServiceCustomHostNameBinding using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppServiceCustomHostNameBinding) LateInitialize(attrs []byte) (bool, error) {
	params := &AppServiceCustomHostNameBindingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppServiceCustomHostNameBinding) GetTerraformSchemaVersion() int {
	return 0
}