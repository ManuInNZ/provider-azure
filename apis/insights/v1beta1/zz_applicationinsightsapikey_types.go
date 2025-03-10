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

type ApplicationInsightsAPIKeyInitParameters struct {

	// Specifies the name of the Application Insights API key. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the list of read permissions granted to the API key. Valid values are agentconfig, aggregate, api, draft, extendqueries, search. Please note these values are case sensitive. Changing this forces a new resource to be created.
	ReadPermissions []*string `json:"readPermissions,omitempty" tf:"read_permissions,omitempty"`

	// Specifies the list of write permissions granted to the API key. Valid values are annotations. Please note these values are case sensitive. Changing this forces a new resource to be created.
	WritePermissions []*string `json:"writePermissions,omitempty" tf:"write_permissions,omitempty"`
}

type ApplicationInsightsAPIKeyObservation struct {

	// The ID of the Application Insights component on which the API key operates. Changing this forces a new resource to be created.
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// The ID of the Application Insights API key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Application Insights API key. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the list of read permissions granted to the API key. Valid values are agentconfig, aggregate, api, draft, extendqueries, search. Please note these values are case sensitive. Changing this forces a new resource to be created.
	ReadPermissions []*string `json:"readPermissions,omitempty" tf:"read_permissions,omitempty"`

	// Specifies the list of write permissions granted to the API key. Valid values are annotations. Please note these values are case sensitive. Changing this forces a new resource to be created.
	WritePermissions []*string `json:"writePermissions,omitempty" tf:"write_permissions,omitempty"`
}

type ApplicationInsightsAPIKeyParameters struct {

	// The ID of the Application Insights component on which the API key operates. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// Reference to a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDRef *v1.Reference `json:"applicationInsightsIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDSelector *v1.Selector `json:"applicationInsightsIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Application Insights API key. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the list of read permissions granted to the API key. Valid values are agentconfig, aggregate, api, draft, extendqueries, search. Please note these values are case sensitive. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ReadPermissions []*string `json:"readPermissions,omitempty" tf:"read_permissions,omitempty"`

	// Specifies the list of write permissions granted to the API key. Valid values are annotations. Please note these values are case sensitive. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	WritePermissions []*string `json:"writePermissions,omitempty" tf:"write_permissions,omitempty"`
}

// ApplicationInsightsAPIKeySpec defines the desired state of ApplicationInsightsAPIKey
type ApplicationInsightsAPIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationInsightsAPIKeyParameters `json:"forProvider"`
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
	InitProvider ApplicationInsightsAPIKeyInitParameters `json:"initProvider,omitempty"`
}

// ApplicationInsightsAPIKeyStatus defines the observed state of ApplicationInsightsAPIKey.
type ApplicationInsightsAPIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationInsightsAPIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsAPIKey is the Schema for the ApplicationInsightsAPIKeys API. Manages an Application Insights API key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationInsightsAPIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ApplicationInsightsAPIKeySpec   `json:"spec"`
	Status ApplicationInsightsAPIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsAPIKeyList contains a list of ApplicationInsightsAPIKeys
type ApplicationInsightsAPIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationInsightsAPIKey `json:"items"`
}

// Repository type metadata.
var (
	ApplicationInsightsAPIKey_Kind             = "ApplicationInsightsAPIKey"
	ApplicationInsightsAPIKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationInsightsAPIKey_Kind}.String()
	ApplicationInsightsAPIKey_KindAPIVersion   = ApplicationInsightsAPIKey_Kind + "." + CRDGroupVersion.String()
	ApplicationInsightsAPIKey_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationInsightsAPIKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationInsightsAPIKey{}, &ApplicationInsightsAPIKeyList{})
}
