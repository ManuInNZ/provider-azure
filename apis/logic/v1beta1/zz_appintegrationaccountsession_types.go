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

type AppIntegrationAccountSessionInitParameters struct {

	// The content of the Logic App Integration Account Session.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`
}

type AppIntegrationAccountSessionObservation struct {

	// The content of the Logic App Integration Account Session.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The ID of the Logic App Integration Account Session.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Logic App Integration Account. Changing this forces a new Logic App Integration Account Session to be created.
	IntegrationAccountName *string `json:"integrationAccountName,omitempty" tf:"integration_account_name,omitempty"`

	// The name of the Resource Group where the Logic App Integration Account Session should exist. Changing this forces a new Logic App Integration Account Session to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type AppIntegrationAccountSessionParameters struct {

	// The content of the Logic App Integration Account Session.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The name of the Logic App Integration Account. Changing this forces a new Logic App Integration Account Session to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/logic/v1beta1.AppIntegrationAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	IntegrationAccountName *string `json:"integrationAccountName,omitempty" tf:"integration_account_name,omitempty"`

	// Reference to a AppIntegrationAccount in logic to populate integrationAccountName.
	// +kubebuilder:validation:Optional
	IntegrationAccountNameRef *v1.Reference `json:"integrationAccountNameRef,omitempty" tf:"-"`

	// Selector for a AppIntegrationAccount in logic to populate integrationAccountName.
	// +kubebuilder:validation:Optional
	IntegrationAccountNameSelector *v1.Selector `json:"integrationAccountNameSelector,omitempty" tf:"-"`

	// The name of the Resource Group where the Logic App Integration Account Session should exist. Changing this forces a new Logic App Integration Account Session to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// AppIntegrationAccountSessionSpec defines the desired state of AppIntegrationAccountSession
type AppIntegrationAccountSessionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppIntegrationAccountSessionParameters `json:"forProvider"`
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
	InitProvider AppIntegrationAccountSessionInitParameters `json:"initProvider,omitempty"`
}

// AppIntegrationAccountSessionStatus defines the observed state of AppIntegrationAccountSession.
type AppIntegrationAccountSessionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppIntegrationAccountSessionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountSession is the Schema for the AppIntegrationAccountSessions API. Manages a Logic App Integration Account Session.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppIntegrationAccountSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.content) || has(self.initProvider.content)",message="content is a required parameter"
	Spec   AppIntegrationAccountSessionSpec   `json:"spec"`
	Status AppIntegrationAccountSessionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountSessionList contains a list of AppIntegrationAccountSessions
type AppIntegrationAccountSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppIntegrationAccountSession `json:"items"`
}

// Repository type metadata.
var (
	AppIntegrationAccountSession_Kind             = "AppIntegrationAccountSession"
	AppIntegrationAccountSession_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppIntegrationAccountSession_Kind}.String()
	AppIntegrationAccountSession_KindAPIVersion   = AppIntegrationAccountSession_Kind + "." + CRDGroupVersion.String()
	AppIntegrationAccountSession_GroupVersionKind = CRDGroupVersion.WithKind(AppIntegrationAccountSession_Kind)
)

func init() {
	SchemeBuilder.Register(&AppIntegrationAccountSession{}, &AppIntegrationAccountSessionList{})
}
