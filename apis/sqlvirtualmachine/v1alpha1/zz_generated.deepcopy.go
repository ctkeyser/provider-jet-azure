//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoBackupObservation) DeepCopyInto(out *AutoBackupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoBackupObservation.
func (in *AutoBackupObservation) DeepCopy() *AutoBackupObservation {
	if in == nil {
		return nil
	}
	out := new(AutoBackupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoBackupParameters) DeepCopyInto(out *AutoBackupParameters) {
	*out = *in
	if in.EncryptionEnabled != nil {
		in, out := &in.EncryptionEnabled, &out.EncryptionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionPasswordSecretRef != nil {
		in, out := &in.EncryptionPasswordSecretRef, &out.EncryptionPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ManualSchedule != nil {
		in, out := &in.ManualSchedule, &out.ManualSchedule
		*out = make([]ManualScheduleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RetentionPeriodInDays != nil {
		in, out := &in.RetentionPeriodInDays, &out.RetentionPeriodInDays
		*out = new(float64)
		**out = **in
	}
	if in.StorageAccountAccessKey != nil {
		in, out := &in.StorageAccountAccessKey, &out.StorageAccountAccessKey
		*out = new(string)
		**out = **in
	}
	if in.StorageBlobEndpoint != nil {
		in, out := &in.StorageBlobEndpoint, &out.StorageBlobEndpoint
		*out = new(string)
		**out = **in
	}
	if in.SystemDatabasesBackupEnabled != nil {
		in, out := &in.SystemDatabasesBackupEnabled, &out.SystemDatabasesBackupEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoBackupParameters.
func (in *AutoBackupParameters) DeepCopy() *AutoBackupParameters {
	if in == nil {
		return nil
	}
	out := new(AutoBackupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoPatchingObservation) DeepCopyInto(out *AutoPatchingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoPatchingObservation.
func (in *AutoPatchingObservation) DeepCopy() *AutoPatchingObservation {
	if in == nil {
		return nil
	}
	out := new(AutoPatchingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoPatchingParameters) DeepCopyInto(out *AutoPatchingParameters) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindowDurationInMinutes != nil {
		in, out := &in.MaintenanceWindowDurationInMinutes, &out.MaintenanceWindowDurationInMinutes
		*out = new(float64)
		**out = **in
	}
	if in.MaintenanceWindowStartingHour != nil {
		in, out := &in.MaintenanceWindowStartingHour, &out.MaintenanceWindowStartingHour
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoPatchingParameters.
func (in *AutoPatchingParameters) DeepCopy() *AutoPatchingParameters {
	if in == nil {
		return nil
	}
	out := new(AutoPatchingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSettingsObservation) DeepCopyInto(out *DataSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSettingsObservation.
func (in *DataSettingsObservation) DeepCopy() *DataSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(DataSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSettingsParameters) DeepCopyInto(out *DataSettingsParameters) {
	*out = *in
	if in.DefaultFilePath != nil {
		in, out := &in.DefaultFilePath, &out.DefaultFilePath
		*out = new(string)
		**out = **in
	}
	if in.Luns != nil {
		in, out := &in.Luns, &out.Luns
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSettingsParameters.
func (in *DataSettingsParameters) DeepCopy() *DataSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(DataSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultCredentialObservation) DeepCopyInto(out *KeyVaultCredentialObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultCredentialObservation.
func (in *KeyVaultCredentialObservation) DeepCopy() *KeyVaultCredentialObservation {
	if in == nil {
		return nil
	}
	out := new(KeyVaultCredentialObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultCredentialParameters) DeepCopyInto(out *KeyVaultCredentialParameters) {
	*out = *in
	out.KeyVaultURLSecretRef = in.KeyVaultURLSecretRef
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	out.ServicePrincipalNameSecretRef = in.ServicePrincipalNameSecretRef
	out.ServicePrincipalSecretSecretRef = in.ServicePrincipalSecretSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultCredentialParameters.
func (in *KeyVaultCredentialParameters) DeepCopy() *KeyVaultCredentialParameters {
	if in == nil {
		return nil
	}
	out := new(KeyVaultCredentialParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSettingsObservation) DeepCopyInto(out *LogSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSettingsObservation.
func (in *LogSettingsObservation) DeepCopy() *LogSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(LogSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSettingsParameters) DeepCopyInto(out *LogSettingsParameters) {
	*out = *in
	if in.DefaultFilePath != nil {
		in, out := &in.DefaultFilePath, &out.DefaultFilePath
		*out = new(string)
		**out = **in
	}
	if in.Luns != nil {
		in, out := &in.Luns, &out.Luns
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSettingsParameters.
func (in *LogSettingsParameters) DeepCopy() *LogSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(LogSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLVirtualMachine) DeepCopyInto(out *MSSQLVirtualMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLVirtualMachine.
func (in *MSSQLVirtualMachine) DeepCopy() *MSSQLVirtualMachine {
	if in == nil {
		return nil
	}
	out := new(MSSQLVirtualMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MSSQLVirtualMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLVirtualMachineList) DeepCopyInto(out *MSSQLVirtualMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MSSQLVirtualMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLVirtualMachineList.
func (in *MSSQLVirtualMachineList) DeepCopy() *MSSQLVirtualMachineList {
	if in == nil {
		return nil
	}
	out := new(MSSQLVirtualMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MSSQLVirtualMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLVirtualMachineObservation) DeepCopyInto(out *MSSQLVirtualMachineObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLVirtualMachineObservation.
func (in *MSSQLVirtualMachineObservation) DeepCopy() *MSSQLVirtualMachineObservation {
	if in == nil {
		return nil
	}
	out := new(MSSQLVirtualMachineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLVirtualMachineParameters) DeepCopyInto(out *MSSQLVirtualMachineParameters) {
	*out = *in
	if in.AutoBackup != nil {
		in, out := &in.AutoBackup, &out.AutoBackup
		*out = make([]AutoBackupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoPatching != nil {
		in, out := &in.AutoPatching, &out.AutoPatching
		*out = make([]AutoPatchingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KeyVaultCredential != nil {
		in, out := &in.KeyVaultCredential, &out.KeyVaultCredential
		*out = make([]KeyVaultCredentialParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RServicesEnabled != nil {
		in, out := &in.RServicesEnabled, &out.RServicesEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SQLConnectivityPort != nil {
		in, out := &in.SQLConnectivityPort, &out.SQLConnectivityPort
		*out = new(float64)
		**out = **in
	}
	if in.SQLConnectivityType != nil {
		in, out := &in.SQLConnectivityType, &out.SQLConnectivityType
		*out = new(string)
		**out = **in
	}
	if in.SQLConnectivityUpdatePasswordSecretRef != nil {
		in, out := &in.SQLConnectivityUpdatePasswordSecretRef, &out.SQLConnectivityUpdatePasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SQLConnectivityUpdateUsernameSecretRef != nil {
		in, out := &in.SQLConnectivityUpdateUsernameSecretRef, &out.SQLConnectivityUpdateUsernameSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SQLLicenseType != nil {
		in, out := &in.SQLLicenseType, &out.SQLLicenseType
		*out = new(string)
		**out = **in
	}
	if in.StorageConfiguration != nil {
		in, out := &in.StorageConfiguration, &out.StorageConfiguration
		*out = make([]StorageConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VirtualMachineID != nil {
		in, out := &in.VirtualMachineID, &out.VirtualMachineID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLVirtualMachineParameters.
func (in *MSSQLVirtualMachineParameters) DeepCopy() *MSSQLVirtualMachineParameters {
	if in == nil {
		return nil
	}
	out := new(MSSQLVirtualMachineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLVirtualMachineSpec) DeepCopyInto(out *MSSQLVirtualMachineSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLVirtualMachineSpec.
func (in *MSSQLVirtualMachineSpec) DeepCopy() *MSSQLVirtualMachineSpec {
	if in == nil {
		return nil
	}
	out := new(MSSQLVirtualMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLVirtualMachineStatus) DeepCopyInto(out *MSSQLVirtualMachineStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLVirtualMachineStatus.
func (in *MSSQLVirtualMachineStatus) DeepCopy() *MSSQLVirtualMachineStatus {
	if in == nil {
		return nil
	}
	out := new(MSSQLVirtualMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualScheduleObservation) DeepCopyInto(out *ManualScheduleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualScheduleObservation.
func (in *ManualScheduleObservation) DeepCopy() *ManualScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(ManualScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualScheduleParameters) DeepCopyInto(out *ManualScheduleParameters) {
	*out = *in
	if in.FullBackupFrequency != nil {
		in, out := &in.FullBackupFrequency, &out.FullBackupFrequency
		*out = new(string)
		**out = **in
	}
	if in.FullBackupStartHour != nil {
		in, out := &in.FullBackupStartHour, &out.FullBackupStartHour
		*out = new(float64)
		**out = **in
	}
	if in.FullBackupWindowInHours != nil {
		in, out := &in.FullBackupWindowInHours, &out.FullBackupWindowInHours
		*out = new(float64)
		**out = **in
	}
	if in.LogBackupFrequencyInMinutes != nil {
		in, out := &in.LogBackupFrequencyInMinutes, &out.LogBackupFrequencyInMinutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualScheduleParameters.
func (in *ManualScheduleParameters) DeepCopy() *ManualScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(ManualScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfigurationObservation) DeepCopyInto(out *StorageConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfigurationObservation.
func (in *StorageConfigurationObservation) DeepCopy() *StorageConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(StorageConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfigurationParameters) DeepCopyInto(out *StorageConfigurationParameters) {
	*out = *in
	if in.DataSettings != nil {
		in, out := &in.DataSettings, &out.DataSettings
		*out = make([]DataSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.LogSettings != nil {
		in, out := &in.LogSettings, &out.LogSettings
		*out = make([]LogSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageWorkloadType != nil {
		in, out := &in.StorageWorkloadType, &out.StorageWorkloadType
		*out = new(string)
		**out = **in
	}
	if in.TempDBSettings != nil {
		in, out := &in.TempDBSettings, &out.TempDBSettings
		*out = make([]TempDBSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfigurationParameters.
func (in *StorageConfigurationParameters) DeepCopy() *StorageConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(StorageConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempDBSettingsObservation) DeepCopyInto(out *TempDBSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempDBSettingsObservation.
func (in *TempDBSettingsObservation) DeepCopy() *TempDBSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(TempDBSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempDBSettingsParameters) DeepCopyInto(out *TempDBSettingsParameters) {
	*out = *in
	if in.DefaultFilePath != nil {
		in, out := &in.DefaultFilePath, &out.DefaultFilePath
		*out = new(string)
		**out = **in
	}
	if in.Luns != nil {
		in, out := &in.Luns, &out.Luns
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempDBSettingsParameters.
func (in *TempDBSettingsParameters) DeepCopy() *TempDBSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(TempDBSettingsParameters)
	in.DeepCopyInto(out)
	return out
}
