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

type ConfigurationInitParameters struct {

	// Specifies the name of the MariaDB Configuration, which needs to be a valid MariaDB configuration name. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the value of the MariaDB Configuration. See the MariaDB documentation for valid values. Changing this forces a new resource to be created.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigurationObservation struct {

	// The ID of the MariaDB Configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the MariaDB Configuration, which needs to be a valid MariaDB configuration name. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Specifies the value of the MariaDB Configuration. See the MariaDB documentation for valid values. Changing this forces a new resource to be created.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigurationParameters struct {

	// Specifies the name of the MariaDB Configuration, which needs to be a valid MariaDB configuration name. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dbformariadb/v1beta1.Server
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Reference to a Server in dbformariadb to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// Selector for a Server in dbformariadb to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// Specifies the value of the MariaDB Configuration. See the MariaDB documentation for valid values. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// ConfigurationSpec defines the desired state of Configuration
type ConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationParameters `json:"forProvider"`
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
	InitProvider ConfigurationInitParameters `json:"initProvider,omitempty"`
}

// ConfigurationStatus defines the observed state of Configuration.
type ConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Configuration is the Schema for the Configurations API. Sets a MariaDB Configuration value on a MariaDB Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || has(self.initProvider.value)",message="value is a required parameter"
	Spec   ConfigurationSpec   `json:"spec"`
	Status ConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationList contains a list of Configurations
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configuration `json:"items"`
}

// Repository type metadata.
var (
	Configuration_Kind             = "Configuration"
	Configuration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Configuration_Kind}.String()
	Configuration_KindAPIVersion   = Configuration_Kind + "." + CRDGroupVersion.String()
	Configuration_GroupVersionKind = CRDGroupVersion.WithKind(Configuration_Kind)
)

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
}
