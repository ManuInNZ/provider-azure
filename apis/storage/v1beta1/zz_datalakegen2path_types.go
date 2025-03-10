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

type DataLakeGen2PathAceInitParameters struct {

	// Specifies the Object ID of the Azure Active Directory User or Group that the entry relates to. Only valid for user or group entries.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the permissions for the entry in rwx form. For example, rwx gives full permissions but r-- only gives read permissions.
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Specifies whether the ACE represents an access entry or a default entry. Default value is access.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Specifies the type of entry. Can be user, group, mask or other.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataLakeGen2PathAceObservation struct {

	// Specifies the Object ID of the Azure Active Directory User or Group that the entry relates to. Only valid for user or group entries.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the permissions for the entry in rwx form. For example, rwx gives full permissions but r-- only gives read permissions.
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Specifies whether the ACE represents an access entry or a default entry. Default value is access.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Specifies the type of entry. Can be user, group, mask or other.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataLakeGen2PathAceParameters struct {

	// Specifies the Object ID of the Azure Active Directory User or Group that the entry relates to. Only valid for user or group entries.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the permissions for the entry in rwx form. For example, rwx gives full permissions but r-- only gives read permissions.
	// +kubebuilder:validation:Optional
	Permissions *string `json:"permissions" tf:"permissions,omitempty"`

	// Specifies whether the ACE represents an access entry or a default entry. Default value is access.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Specifies the type of entry. Can be user, group, mask or other.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type DataLakeGen2PathInitParameters struct {

	// One or more ace blocks as defined below to specify the entries for the ACL for the path.
	Ace []DataLakeGen2PathAceInitParameters `json:"ace,omitempty" tf:"ace,omitempty"`

	// Specifies the Object ID of the Azure Active Directory Group to make the owning group. Possible values also include $superuser.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Specifies the Object ID of the Azure Active Directory User to make the owning user. Possible values also include $superuser.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The path which should be created within the Data Lake Gen2 File System in the Storage Account. Changing this forces a new resource to be created.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the type for path to create. Currently only directory is supported. Changing this forces a new resource to be created.
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`
}

type DataLakeGen2PathObservation struct {

	// One or more ace blocks as defined below to specify the entries for the ACL for the path.
	Ace []DataLakeGen2PathAceObservation `json:"ace,omitempty" tf:"ace,omitempty"`

	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	FileSystemName *string `json:"filesystemName,omitempty" tf:"filesystem_name,omitempty"`

	// Specifies the Object ID of the Azure Active Directory Group to make the owning group. Possible values also include $superuser.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// The ID of the Data Lake Gen2 File System.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Object ID of the Azure Active Directory User to make the owning user. Possible values also include $superuser.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The path which should be created within the Data Lake Gen2 File System in the Storage Account. Changing this forces a new resource to be created.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the type for path to create. Currently only directory is supported. Changing this forces a new resource to be created.
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type DataLakeGen2PathParameters struct {

	// One or more ace blocks as defined below to specify the entries for the ACL for the path.
	// +kubebuilder:validation:Optional
	Ace []DataLakeGen2PathAceParameters `json:"ace,omitempty" tf:"ace,omitempty"`

	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.DataLakeGen2FileSystem
	// +kubebuilder:validation:Optional
	FileSystemName *string `json:"filesystemName,omitempty" tf:"filesystem_name,omitempty"`

	// Reference to a DataLakeGen2FileSystem in storage to populate filesystemName.
	// +kubebuilder:validation:Optional
	FileSystemNameRef *v1.Reference `json:"filesystemNameRef,omitempty" tf:"-"`

	// Selector for a DataLakeGen2FileSystem in storage to populate filesystemName.
	// +kubebuilder:validation:Optional
	FileSystemNameSelector *v1.Selector `json:"filesystemNameSelector,omitempty" tf:"-"`

	// Specifies the Object ID of the Azure Active Directory Group to make the owning group. Possible values also include $superuser.
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Specifies the Object ID of the Azure Active Directory User to make the owning user. Possible values also include $superuser.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The path which should be created within the Data Lake Gen2 File System in the Storage Account. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the type for path to create. Currently only directory is supported. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

// DataLakeGen2PathSpec defines the desired state of DataLakeGen2Path
type DataLakeGen2PathSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataLakeGen2PathParameters `json:"forProvider"`
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
	InitProvider DataLakeGen2PathInitParameters `json:"initProvider,omitempty"`
}

// DataLakeGen2PathStatus defines the observed state of DataLakeGen2Path.
type DataLakeGen2PathStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataLakeGen2PathObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataLakeGen2Path is the Schema for the DataLakeGen2Paths API. Manages a Data Lake Gen2 Path in a File System within an Azure Storage Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataLakeGen2Path struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || has(self.initProvider.path)",message="path is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resource) || has(self.initProvider.resource)",message="resource is a required parameter"
	Spec   DataLakeGen2PathSpec   `json:"spec"`
	Status DataLakeGen2PathStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataLakeGen2PathList contains a list of DataLakeGen2Paths
type DataLakeGen2PathList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataLakeGen2Path `json:"items"`
}

// Repository type metadata.
var (
	DataLakeGen2Path_Kind             = "DataLakeGen2Path"
	DataLakeGen2Path_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataLakeGen2Path_Kind}.String()
	DataLakeGen2Path_KindAPIVersion   = DataLakeGen2Path_Kind + "." + CRDGroupVersion.String()
	DataLakeGen2Path_GroupVersionKind = CRDGroupVersion.WithKind(DataLakeGen2Path_Kind)
)

func init() {
	SchemeBuilder.Register(&DataLakeGen2Path{}, &DataLakeGen2PathList{})
}
