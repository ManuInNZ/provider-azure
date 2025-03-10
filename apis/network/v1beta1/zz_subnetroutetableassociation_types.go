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

type SubnetRouteTableAssociationInitParameters struct {
}

type SubnetRouteTableAssociationObservation struct {

	// The ID of the Subnet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Route Table which should be associated with the Subnet. Changing this forces a new resource to be created.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SubnetRouteTableAssociationParameters struct {

	// The ID of the Route Table which should be associated with the Subnet. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=RouteTable
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// The ID of the Subnet. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// SubnetRouteTableAssociationSpec defines the desired state of SubnetRouteTableAssociation
type SubnetRouteTableAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetRouteTableAssociationParameters `json:"forProvider"`
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
	InitProvider SubnetRouteTableAssociationInitParameters `json:"initProvider,omitempty"`
}

// SubnetRouteTableAssociationStatus defines the observed state of SubnetRouteTableAssociation.
type SubnetRouteTableAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetRouteTableAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetRouteTableAssociation is the Schema for the SubnetRouteTableAssociations API. Associates a
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SubnetRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetRouteTableAssociationSpec   `json:"spec"`
	Status            SubnetRouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetRouteTableAssociationList contains a list of SubnetRouteTableAssociations
type SubnetRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetRouteTableAssociation `json:"items"`
}

// Repository type metadata.
var (
	SubnetRouteTableAssociation_Kind             = "SubnetRouteTableAssociation"
	SubnetRouteTableAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetRouteTableAssociation_Kind}.String()
	SubnetRouteTableAssociation_KindAPIVersion   = SubnetRouteTableAssociation_Kind + "." + CRDGroupVersion.String()
	SubnetRouteTableAssociation_GroupVersionKind = CRDGroupVersion.WithKind(SubnetRouteTableAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetRouteTableAssociation{}, &SubnetRouteTableAssociationList{})
}
