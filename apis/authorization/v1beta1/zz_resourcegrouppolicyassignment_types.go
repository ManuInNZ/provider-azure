/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IdentityInitParameters struct {

	// A list of User Managed Identity IDs which should be assigned to the Policy Definition.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Type of Managed Identity which should be added to this Policy Definition. Possible values are SystemAssigned and UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// A list of User Managed Identity IDs which should be assigned to the Policy Definition.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID of the Policy Assignment for this Resource Group.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID of the Policy Assignment for this Resource Group.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The Type of Managed Identity which should be added to this Policy Definition. Possible values are SystemAssigned and UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// A list of User Managed Identity IDs which should be assigned to the Policy Definition.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Type of Managed Identity which should be added to this Policy Definition. Possible values are SystemAssigned and UserAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type NonComplianceMessageInitParameters struct {

	// The non-compliance message text. When assigning policy sets (initiatives), unless policy_definition_reference_id is specified then this message will be the default for all policies.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`
}

type NonComplianceMessageObservation struct {

	// The non-compliance message text. When assigning policy sets (initiatives), unless policy_definition_reference_id is specified then this message will be the default for all policies.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`
}

type NonComplianceMessageParameters struct {

	// The non-compliance message text. When assigning policy sets (initiatives), unless policy_definition_reference_id is specified then this message will be the default for all policies.
	// +kubebuilder:validation:Optional
	Content *string `json:"content" tf:"content,omitempty"`

	// When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
	// +kubebuilder:validation:Optional
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`
}

type OverridesInitParameters struct {

	// One or more override_selector as defined below.
	Selectors []SelectorsInitParameters `json:"selectors,omitempty" tf:"selectors,omitempty"`

	// Specifies the value to override the policy property. Possible values for policyEffect override listed policy effects.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type OverridesObservation struct {

	// One or more override_selector as defined below.
	Selectors []SelectorsObservation `json:"selectors,omitempty" tf:"selectors,omitempty"`

	// Specifies the value to override the policy property. Possible values for policyEffect override listed policy effects.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type OverridesParameters struct {

	// One or more override_selector as defined below.
	// +kubebuilder:validation:Optional
	Selectors []SelectorsParameters `json:"selectors,omitempty" tf:"selectors,omitempty"`

	// Specifies the value to override the policy property. Possible values for policyEffect override listed policy effects.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type ResourceGroupPolicyAssignmentInitParameters struct {

	// A description which should be used for this Policy Assignment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Display Name for this Policy Assignment.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies if this Policy should be enforced or not? Defaults to true.
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A JSON mapping of any Metadata for this Policy.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// One or more non_compliance_message blocks as defined below.
	NonComplianceMessage []NonComplianceMessageInitParameters `json:"nonComplianceMessage,omitempty" tf:"non_compliance_message,omitempty"`

	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes []*string `json:"notScopes,omitempty" tf:"not_scopes,omitempty"`

	// One or more overrides blocks as defined below. More detail about overrides and resource_selectors see policy assignment structure
	Overrides []OverridesInitParameters `json:"overrides,omitempty" tf:"overrides,omitempty"`

	// A JSON mapping of any Parameters for this Policy.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// One or more resource_selectors blocks as defined below to filter polices by resource properties.
	ResourceSelectors []ResourceSelectorsInitParameters `json:"resourceSelectors,omitempty" tf:"resource_selectors,omitempty"`
}

type ResourceGroupPolicyAssignmentObservation struct {

	// A description which should be used for this Policy Assignment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Display Name for this Policy Assignment.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies if this Policy should be enforced or not? Defaults to true.
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// The ID of the Resource Group Policy Assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A JSON mapping of any Metadata for this Policy.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// One or more non_compliance_message blocks as defined below.
	NonComplianceMessage []NonComplianceMessageObservation `json:"nonComplianceMessage,omitempty" tf:"non_compliance_message,omitempty"`

	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes []*string `json:"notScopes,omitempty" tf:"not_scopes,omitempty"`

	// One or more overrides blocks as defined below. More detail about overrides and resource_selectors see policy assignment structure
	Overrides []OverridesObservation `json:"overrides,omitempty" tf:"overrides,omitempty"`

	// A JSON mapping of any Parameters for this Policy.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// The ID of the Resource Group where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// One or more resource_selectors blocks as defined below to filter polices by resource properties.
	ResourceSelectors []ResourceSelectorsObservation `json:"resourceSelectors,omitempty" tf:"resource_selectors,omitempty"`
}

type ResourceGroupPolicyAssignmentParameters struct {

	// A description which should be used for this Policy Assignment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Display Name for this Policy Assignment.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies if this Policy should be enforced or not? Defaults to true.
	// +kubebuilder:validation:Optional
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A JSON mapping of any Metadata for this Policy.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// One or more non_compliance_message blocks as defined below.
	// +kubebuilder:validation:Optional
	NonComplianceMessage []NonComplianceMessageParameters `json:"nonComplianceMessage,omitempty" tf:"non_compliance_message,omitempty"`

	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	// +kubebuilder:validation:Optional
	NotScopes []*string `json:"notScopes,omitempty" tf:"not_scopes,omitempty"`

	// One or more overrides blocks as defined below. More detail about overrides and resource_selectors see policy assignment structure
	// +kubebuilder:validation:Optional
	Overrides []OverridesParameters `json:"overrides,omitempty" tf:"overrides,omitempty"`

	// A JSON mapping of any Parameters for this Policy.
	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/authorization/v1beta1.PolicyDefinition
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// Reference to a PolicyDefinition in authorization to populate policyDefinitionId.
	// +kubebuilder:validation:Optional
	PolicyDefinitionIDRef *v1.Reference `json:"policyDefinitionIdRef,omitempty" tf:"-"`

	// Selector for a PolicyDefinition in authorization to populate policyDefinitionId.
	// +kubebuilder:validation:Optional
	PolicyDefinitionIDSelector *v1.Selector `json:"policyDefinitionIdSelector,omitempty" tf:"-"`

	// The ID of the Resource Group where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupId.
	// +kubebuilder:validation:Optional
	ResourceGroupIDRef *v1.Reference `json:"resourceGroupIdRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupId.
	// +kubebuilder:validation:Optional
	ResourceGroupIDSelector *v1.Selector `json:"resourceGroupIdSelector,omitempty" tf:"-"`

	// One or more resource_selectors blocks as defined below to filter polices by resource properties.
	// +kubebuilder:validation:Optional
	ResourceSelectors []ResourceSelectorsParameters `json:"resourceSelectors,omitempty" tf:"resource_selectors,omitempty"`
}

type ResourceSelectorsInitParameters struct {

	// Specifies a name for the resource selector.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more resource_selector block as defined below.
	Selectors []ResourceSelectorsSelectorsInitParameters `json:"selectors,omitempty" tf:"selectors,omitempty"`
}

type ResourceSelectorsObservation struct {

	// Specifies a name for the resource selector.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more resource_selector block as defined below.
	Selectors []ResourceSelectorsSelectorsObservation `json:"selectors,omitempty" tf:"selectors,omitempty"`
}

type ResourceSelectorsParameters struct {

	// Specifies a name for the resource selector.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more resource_selector block as defined below.
	// +kubebuilder:validation:Optional
	Selectors []ResourceSelectorsSelectorsParameters `json:"selectors" tf:"selectors,omitempty"`
}

type ResourceSelectorsSelectorsInitParameters struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// Specifies which characteristic will narrow down the set of evaluated resources. Possible values are resourceLocation,  resourceType and resourceWithoutLocation.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

type ResourceSelectorsSelectorsObservation struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// Specifies which characteristic will narrow down the set of evaluated resources. Possible values are resourceLocation,  resourceType and resourceWithoutLocation.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

type ResourceSelectorsSelectorsParameters struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	// +kubebuilder:validation:Optional
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// Specifies which characteristic will narrow down the set of evaluated resources. Possible values are resourceLocation,  resourceType and resourceWithoutLocation.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	// +kubebuilder:validation:Optional
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

type SelectorsInitParameters struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

type SelectorsObservation struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// Specifies which characteristic will narrow down the set of evaluated resources. Possible values are resourceLocation,  resourceType and resourceWithoutLocation.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

type SelectorsParameters struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	// +kubebuilder:validation:Optional
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	// +kubebuilder:validation:Optional
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

// ResourceGroupPolicyAssignmentSpec defines the desired state of ResourceGroupPolicyAssignment
type ResourceGroupPolicyAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGroupPolicyAssignmentParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ResourceGroupPolicyAssignmentInitParameters `json:"initProvider,omitempty"`
}

// ResourceGroupPolicyAssignmentStatus defines the observed state of ResourceGroupPolicyAssignment.
type ResourceGroupPolicyAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGroupPolicyAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupPolicyAssignment is the Schema for the ResourceGroupPolicyAssignments API. Manages a Resource Group Policy Assignment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceGroupPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroupPolicyAssignmentSpec   `json:"spec"`
	Status            ResourceGroupPolicyAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupPolicyAssignmentList contains a list of ResourceGroupPolicyAssignments
type ResourceGroupPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroupPolicyAssignment `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroupPolicyAssignment_Kind             = "ResourceGroupPolicyAssignment"
	ResourceGroupPolicyAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGroupPolicyAssignment_Kind}.String()
	ResourceGroupPolicyAssignment_KindAPIVersion   = ResourceGroupPolicyAssignment_Kind + "." + CRDGroupVersion.String()
	ResourceGroupPolicyAssignment_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGroupPolicyAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroupPolicyAssignment{}, &ResourceGroupPolicyAssignmentList{})
}
