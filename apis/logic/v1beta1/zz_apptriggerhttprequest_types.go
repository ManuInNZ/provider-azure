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

type AppTriggerHTTPRequestInitParameters struct {

	// Specifies the HTTP Method which the request be using. Possible values include DELETE, GET, PATCH, POST or PUT.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Specifies the Relative Path used for this Request.
	RelativePath *string `json:"relativePath,omitempty" tf:"relative_path,omitempty"`

	// A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

type AppTriggerHTTPRequestObservation struct {

	// The URL for the workflow trigger
	CallbackURL *string `json:"callbackUrl,omitempty" tf:"callback_url,omitempty"`

	// The ID of the HTTP Request Trigger within the Logic App Workflow.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppID *string `json:"logicAppId,omitempty" tf:"logic_app_id,omitempty"`

	// Specifies the HTTP Method which the request be using. Possible values include DELETE, GET, PATCH, POST or PUT.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Specifies the Relative Path used for this Request.
	RelativePath *string `json:"relativePath,omitempty" tf:"relative_path,omitempty"`

	// A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

type AppTriggerHTTPRequestParameters struct {

	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/logic/v1beta1.AppWorkflow
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LogicAppID *string `json:"logicAppId,omitempty" tf:"logic_app_id,omitempty"`

	// Reference to a AppWorkflow in logic to populate logicAppId.
	// +kubebuilder:validation:Optional
	LogicAppIDRef *v1.Reference `json:"logicAppIdRef,omitempty" tf:"-"`

	// Selector for a AppWorkflow in logic to populate logicAppId.
	// +kubebuilder:validation:Optional
	LogicAppIDSelector *v1.Selector `json:"logicAppIdSelector,omitempty" tf:"-"`

	// Specifies the HTTP Method which the request be using. Possible values include DELETE, GET, PATCH, POST or PUT.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Specifies the Relative Path used for this Request.
	// +kubebuilder:validation:Optional
	RelativePath *string `json:"relativePath,omitempty" tf:"relative_path,omitempty"`

	// A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

// AppTriggerHTTPRequestSpec defines the desired state of AppTriggerHTTPRequest
type AppTriggerHTTPRequestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppTriggerHTTPRequestParameters `json:"forProvider"`
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
	InitProvider AppTriggerHTTPRequestInitParameters `json:"initProvider,omitempty"`
}

// AppTriggerHTTPRequestStatus defines the observed state of AppTriggerHTTPRequest.
type AppTriggerHTTPRequestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppTriggerHTTPRequestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppTriggerHTTPRequest is the Schema for the AppTriggerHTTPRequests API. Manages a HTTP Request Trigger within a Logic App Workflow
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppTriggerHTTPRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schema) || has(self.initProvider.schema)",message="schema is a required parameter"
	Spec   AppTriggerHTTPRequestSpec   `json:"spec"`
	Status AppTriggerHTTPRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppTriggerHTTPRequestList contains a list of AppTriggerHTTPRequests
type AppTriggerHTTPRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppTriggerHTTPRequest `json:"items"`
}

// Repository type metadata.
var (
	AppTriggerHTTPRequest_Kind             = "AppTriggerHTTPRequest"
	AppTriggerHTTPRequest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppTriggerHTTPRequest_Kind}.String()
	AppTriggerHTTPRequest_KindAPIVersion   = AppTriggerHTTPRequest_Kind + "." + CRDGroupVersion.String()
	AppTriggerHTTPRequest_GroupVersionKind = CRDGroupVersion.WithKind(AppTriggerHTTPRequest_Kind)
)

func init() {
	SchemeBuilder.Register(&AppTriggerHTTPRequest{}, &AppTriggerHTTPRequestList{})
}
