/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ResourceGroupPolicyAssignment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ResourceGroupPolicyAssignment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ResourceGroupPolicyAssignment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ResourceGroupPolicyAssignment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
