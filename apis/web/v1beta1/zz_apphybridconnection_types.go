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

type AppHybridConnectionInitParameters struct {

	// The hostname of the endpoint.
	// The hostname of the endpoint.
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The port to use for the endpoint.
	// The port to use for the endpoint
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The name of the Relay key with Send permission to use. Defaults to RootManageSharedAccessKey
	// The name of the Relay key with `Send` permission to use. Defaults to `RootManageSharedAccessKey`
	SendKeyName *string `json:"sendKeyName,omitempty" tf:"send_key_name,omitempty"`
}

type AppHybridConnectionObservation struct {

	// The hostname of the endpoint.
	// The hostname of the endpoint.
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The ID of the Web App Hybrid Connection
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Relay Namespace.
	// The name of the Relay Namespace.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// The port to use for the endpoint.
	// The port to use for the endpoint
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the Relay Hybrid Connection to use. Changing this forces a new resource to be created.
	// The ID of the Relay Hybrid Connection to use.
	RelayID *string `json:"relayId,omitempty" tf:"relay_id,omitempty"`

	// The name of the Relay in use.
	// The name of the Relay in use.
	RelayName *string `json:"relayName,omitempty" tf:"relay_name,omitempty"`

	// The name of the Relay key with Send permission to use. Defaults to RootManageSharedAccessKey
	// The name of the Relay key with `Send` permission to use. Defaults to `RootManageSharedAccessKey`
	SendKeyName *string `json:"sendKeyName,omitempty" tf:"send_key_name,omitempty"`

	// The Service Bus Namespace.
	// The Service Bus Namespace.
	ServiceBusNamespace *string `json:"serviceBusNamespace,omitempty" tf:"service_bus_namespace,omitempty"`

	// The suffix for the endpoint.
	// The suffix for the endpoint.
	ServiceBusSuffix *string `json:"serviceBusSuffix,omitempty" tf:"service_bus_suffix,omitempty"`

	// The ID of the Web App for this Hybrid Connection. Changing this forces a new resource to be created.
	// The ID of the Web App for this Hybrid Connection.
	WebAppID *string `json:"webAppId,omitempty" tf:"web_app_id,omitempty"`
}

type AppHybridConnectionParameters struct {

	// The hostname of the endpoint.
	// The hostname of the endpoint.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The port to use for the endpoint.
	// The port to use for the endpoint
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the Relay Hybrid Connection to use. Changing this forces a new resource to be created.
	// The ID of the Relay Hybrid Connection to use.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/relay/v1beta1.HybridConnection
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RelayID *string `json:"relayId,omitempty" tf:"relay_id,omitempty"`

	// Reference to a HybridConnection in relay to populate relayId.
	// +kubebuilder:validation:Optional
	RelayIDRef *v1.Reference `json:"relayIdRef,omitempty" tf:"-"`

	// Selector for a HybridConnection in relay to populate relayId.
	// +kubebuilder:validation:Optional
	RelayIDSelector *v1.Selector `json:"relayIdSelector,omitempty" tf:"-"`

	// The name of the Relay key with Send permission to use. Defaults to RootManageSharedAccessKey
	// The name of the Relay key with `Send` permission to use. Defaults to `RootManageSharedAccessKey`
	// +kubebuilder:validation:Optional
	SendKeyName *string `json:"sendKeyName,omitempty" tf:"send_key_name,omitempty"`

	// The ID of the Web App for this Hybrid Connection. Changing this forces a new resource to be created.
	// The ID of the Web App for this Hybrid Connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/web/v1beta1.WindowsWebApp
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WebAppID *string `json:"webAppId,omitempty" tf:"web_app_id,omitempty"`

	// Reference to a WindowsWebApp in web to populate webAppId.
	// +kubebuilder:validation:Optional
	WebAppIDRef *v1.Reference `json:"webAppIdRef,omitempty" tf:"-"`

	// Selector for a WindowsWebApp in web to populate webAppId.
	// +kubebuilder:validation:Optional
	WebAppIDSelector *v1.Selector `json:"webAppIdSelector,omitempty" tf:"-"`
}

// AppHybridConnectionSpec defines the desired state of AppHybridConnection
type AppHybridConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppHybridConnectionParameters `json:"forProvider"`
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
	InitProvider AppHybridConnectionInitParameters `json:"initProvider,omitempty"`
}

// AppHybridConnectionStatus defines the observed state of AppHybridConnection.
type AppHybridConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppHybridConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppHybridConnection is the Schema for the AppHybridConnections API. Manages a Web App Hybrid Connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppHybridConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostname) || has(self.initProvider.hostname)",message="hostname is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.port) || has(self.initProvider.port)",message="port is a required parameter"
	Spec   AppHybridConnectionSpec   `json:"spec"`
	Status AppHybridConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppHybridConnectionList contains a list of AppHybridConnections
type AppHybridConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppHybridConnection `json:"items"`
}

// Repository type metadata.
var (
	AppHybridConnection_Kind             = "AppHybridConnection"
	AppHybridConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppHybridConnection_Kind}.String()
	AppHybridConnection_KindAPIVersion   = AppHybridConnection_Kind + "." + CRDGroupVersion.String()
	AppHybridConnection_GroupVersionKind = CRDGroupVersion.WithKind(AppHybridConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&AppHybridConnection{}, &AppHybridConnectionList{})
}
