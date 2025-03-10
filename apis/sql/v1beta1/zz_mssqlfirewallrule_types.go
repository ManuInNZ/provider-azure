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

type MSSQLFirewallRuleInitParameters struct {

	// The ending IP address to allow through the firewall for this rule.
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The starting IP address to allow through the firewall for this rule.
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`
}

type MSSQLFirewallRuleObservation struct {

	// The ending IP address to allow through the firewall for this rule.
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The SQL Firewall Rule ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource ID of the SQL Server on which to create the Firewall Rule. Changing this forces a new resource to be created.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// The starting IP address to allow through the firewall for this rule.
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`
}

type MSSQLFirewallRuleParameters struct {

	// The ending IP address to allow through the firewall for this rule.
	// +kubebuilder:validation:Optional
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The resource ID of the SQL Server on which to create the Firewall Rule. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/sql/v1beta1.MSSQLServer
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a MSSQLServer in sql to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer in sql to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// The starting IP address to allow through the firewall for this rule.
	// +kubebuilder:validation:Optional
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`
}

// MSSQLFirewallRuleSpec defines the desired state of MSSQLFirewallRule
type MSSQLFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLFirewallRuleParameters `json:"forProvider"`
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
	InitProvider MSSQLFirewallRuleInitParameters `json:"initProvider,omitempty"`
}

// MSSQLFirewallRuleStatus defines the observed state of MSSQLFirewallRule.
type MSSQLFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLFirewallRule is the Schema for the MSSQLFirewallRules API. Manages an Azure SQL Firewall Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endIpAddress) || has(self.initProvider.endIpAddress)",message="endIpAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.startIpAddress) || has(self.initProvider.startIpAddress)",message="startIpAddress is a required parameter"
	Spec   MSSQLFirewallRuleSpec   `json:"spec"`
	Status MSSQLFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLFirewallRuleList contains a list of MSSQLFirewallRules
type MSSQLFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	MSSQLFirewallRule_Kind             = "MSSQLFirewallRule"
	MSSQLFirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLFirewallRule_Kind}.String()
	MSSQLFirewallRule_KindAPIVersion   = MSSQLFirewallRule_Kind + "." + CRDGroupVersion.String()
	MSSQLFirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLFirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLFirewallRule{}, &MSSQLFirewallRuleList{})
}
