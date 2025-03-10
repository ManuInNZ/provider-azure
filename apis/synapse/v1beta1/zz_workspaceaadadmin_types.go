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

type WorkspaceAADAdminInitParameters struct {

	// The login name of the Azure AD Administrator of this Synapse Workspace.
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The object id of the Azure AD Administrator of this Synapse Workspace.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type WorkspaceAADAdminObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The login name of the Azure AD Administrator of this Synapse Workspace.
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The object id of the Azure AD Administrator of this Synapse Workspace.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type WorkspaceAADAdminParameters struct {

	// The login name of the Azure AD Administrator of this Synapse Workspace.
	// +kubebuilder:validation:Optional
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The object id of the Azure AD Administrator of this Synapse Workspace.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Reference to a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDRef *v1.Reference `json:"synapseWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDSelector *v1.Selector `json:"synapseWorkspaceIdSelector,omitempty" tf:"-"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// WorkspaceAADAdminSpec defines the desired state of WorkspaceAADAdmin
type WorkspaceAADAdminSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceAADAdminParameters `json:"forProvider"`
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
	InitProvider WorkspaceAADAdminInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceAADAdminStatus defines the observed state of WorkspaceAADAdmin.
type WorkspaceAADAdminStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceAADAdminObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceAADAdmin is the Schema for the WorkspaceAADAdmins API. Manages Synapse Workspace AAD Admin
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WorkspaceAADAdmin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.login) || has(self.initProvider.login)",message="login is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectId) || has(self.initProvider.objectId)",message="objectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId) || has(self.initProvider.tenantId)",message="tenantId is a required parameter"
	Spec   WorkspaceAADAdminSpec   `json:"spec"`
	Status WorkspaceAADAdminStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceAADAdminList contains a list of WorkspaceAADAdmins
type WorkspaceAADAdminList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceAADAdmin `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceAADAdmin_Kind             = "WorkspaceAADAdmin"
	WorkspaceAADAdmin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceAADAdmin_Kind}.String()
	WorkspaceAADAdmin_KindAPIVersion   = WorkspaceAADAdmin_Kind + "." + CRDGroupVersion.String()
	WorkspaceAADAdmin_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceAADAdmin_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceAADAdmin{}, &WorkspaceAADAdminList{})
}
