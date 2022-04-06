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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this HPCCache.
func (mg *HPCCache) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HPCCache.
func (mg *HPCCache) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HPCCache.
func (mg *HPCCache) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HPCCache.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HPCCache) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HPCCache.
func (mg *HPCCache) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HPCCache.
func (mg *HPCCache) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HPCCache.
func (mg *HPCCache) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HPCCache.
func (mg *HPCCache) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HPCCache.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HPCCache) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HPCCache.
func (mg *HPCCache) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HPCCacheAccessPolicy.
func (mg *HPCCacheAccessPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HPCCacheAccessPolicy.
func (mg *HPCCacheAccessPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HPCCacheAccessPolicy.
func (mg *HPCCacheAccessPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HPCCacheAccessPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HPCCacheAccessPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HPCCacheAccessPolicy.
func (mg *HPCCacheAccessPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HPCCacheAccessPolicy.
func (mg *HPCCacheAccessPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HPCCacheAccessPolicy.
func (mg *HPCCacheAccessPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HPCCacheAccessPolicy.
func (mg *HPCCacheAccessPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HPCCacheAccessPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HPCCacheAccessPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HPCCacheAccessPolicy.
func (mg *HPCCacheAccessPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HPCCacheBlobNFSTarget.
func (mg *HPCCacheBlobNFSTarget) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HPCCacheBlobNFSTarget.
func (mg *HPCCacheBlobNFSTarget) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HPCCacheBlobNFSTarget.
func (mg *HPCCacheBlobNFSTarget) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HPCCacheBlobNFSTarget.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HPCCacheBlobNFSTarget) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HPCCacheBlobNFSTarget.
func (mg *HPCCacheBlobNFSTarget) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HPCCacheBlobNFSTarget.
func (mg *HPCCacheBlobNFSTarget) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HPCCacheBlobNFSTarget.
func (mg *HPCCacheBlobNFSTarget) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HPCCacheBlobNFSTarget.
func (mg *HPCCacheBlobNFSTarget) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HPCCacheBlobNFSTarget.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HPCCacheBlobNFSTarget) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HPCCacheBlobNFSTarget.
func (mg *HPCCacheBlobNFSTarget) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HPCCacheBlobTarget.
func (mg *HPCCacheBlobTarget) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HPCCacheBlobTarget.
func (mg *HPCCacheBlobTarget) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HPCCacheBlobTarget.
func (mg *HPCCacheBlobTarget) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HPCCacheBlobTarget.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HPCCacheBlobTarget) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HPCCacheBlobTarget.
func (mg *HPCCacheBlobTarget) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HPCCacheBlobTarget.
func (mg *HPCCacheBlobTarget) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HPCCacheBlobTarget.
func (mg *HPCCacheBlobTarget) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HPCCacheBlobTarget.
func (mg *HPCCacheBlobTarget) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HPCCacheBlobTarget.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HPCCacheBlobTarget) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HPCCacheBlobTarget.
func (mg *HPCCacheBlobTarget) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HPCCacheNFSTarget.
func (mg *HPCCacheNFSTarget) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HPCCacheNFSTarget.
func (mg *HPCCacheNFSTarget) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HPCCacheNFSTarget.
func (mg *HPCCacheNFSTarget) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HPCCacheNFSTarget.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HPCCacheNFSTarget) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HPCCacheNFSTarget.
func (mg *HPCCacheNFSTarget) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HPCCacheNFSTarget.
func (mg *HPCCacheNFSTarget) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HPCCacheNFSTarget.
func (mg *HPCCacheNFSTarget) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HPCCacheNFSTarget.
func (mg *HPCCacheNFSTarget) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HPCCacheNFSTarget.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HPCCacheNFSTarget) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HPCCacheNFSTarget.
func (mg *HPCCacheNFSTarget) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}