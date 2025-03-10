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

type IOTHubEnrichmentInitParameters struct {

	// The list of endpoints which will be enriched.
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// The key of the enrichment. Changing this forces a new resource to be created.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use $iothubname) or information from the device twin (ex: $twin.tags.latitude)
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type IOTHubEnrichmentObservation struct {

	// The list of endpoints which will be enriched.
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// The ID of the IoTHub Enrichment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IoTHub name of the enrichment. Changing this forces a new resource to be created.
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// The key of the enrichment. Changing this forces a new resource to be created.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The name of the resource group under which the IoTHub resource is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use $iothubname) or information from the device twin (ex: $twin.tags.latitude)
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type IOTHubEnrichmentParameters struct {

	// The list of endpoints which will be enriched.
	// +kubebuilder:validation:Optional
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// The IoTHub name of the enrichment. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Reference to a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iothubNameRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iothubNameSelector,omitempty" tf:"-"`

	// The key of the enrichment. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The name of the resource group under which the IoTHub resource is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use $iothubname) or information from the device twin (ex: $twin.tags.latitude)
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// IOTHubEnrichmentSpec defines the desired state of IOTHubEnrichment
type IOTHubEnrichmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubEnrichmentParameters `json:"forProvider"`
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
	InitProvider IOTHubEnrichmentInitParameters `json:"initProvider,omitempty"`
}

// IOTHubEnrichmentStatus defines the observed state of IOTHubEnrichment.
type IOTHubEnrichmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubEnrichmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEnrichment is the Schema for the IOTHubEnrichments API. Manages an IotHub Enrichment
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubEnrichment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endpointNames) || has(self.initProvider.endpointNames)",message="endpointNames is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || has(self.initProvider.key)",message="key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || has(self.initProvider.value)",message="value is a required parameter"
	Spec   IOTHubEnrichmentSpec   `json:"spec"`
	Status IOTHubEnrichmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEnrichmentList contains a list of IOTHubEnrichments
type IOTHubEnrichmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubEnrichment `json:"items"`
}

// Repository type metadata.
var (
	IOTHubEnrichment_Kind             = "IOTHubEnrichment"
	IOTHubEnrichment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubEnrichment_Kind}.String()
	IOTHubEnrichment_KindAPIVersion   = IOTHubEnrichment_Kind + "." + CRDGroupVersion.String()
	IOTHubEnrichment_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubEnrichment_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubEnrichment{}, &IOTHubEnrichmentList{})
}
