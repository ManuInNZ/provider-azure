//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Account) DeepCopyInto(out *Account) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Account.
func (in *Account) DeepCopy() *Account {
	if in == nil {
		return nil
	}
	out := new(Account)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Account) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountInitParameters) DeepCopyInto(out *AccountInitParameters) {
	*out = *in
	if in.CustomQuestionAnsweringSearchServiceID != nil {
		in, out := &in.CustomQuestionAnsweringSearchServiceID, &out.CustomQuestionAnsweringSearchServiceID
		*out = new(string)
		**out = **in
	}
	if in.CustomSubdomainName != nil {
		in, out := &in.CustomSubdomainName, &out.CustomSubdomainName
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKey != nil {
		in, out := &in.CustomerManagedKey, &out.CustomerManagedKey
		*out = make([]CustomerManagedKeyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DynamicThrottlingEnabled != nil {
		in, out := &in.DynamicThrottlingEnabled, &out.DynamicThrottlingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Fqdns != nil {
		in, out := &in.Fqdns, &out.Fqdns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.LocalAuthEnabled != nil {
		in, out := &in.LocalAuthEnabled, &out.LocalAuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorAADClientID != nil {
		in, out := &in.MetricsAdvisorAADClientID, &out.MetricsAdvisorAADClientID
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorAADTenantID != nil {
		in, out := &in.MetricsAdvisorAADTenantID, &out.MetricsAdvisorAADTenantID
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorSuperUserName != nil {
		in, out := &in.MetricsAdvisorSuperUserName, &out.MetricsAdvisorSuperUserName
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorWebsiteName != nil {
		in, out := &in.MetricsAdvisorWebsiteName, &out.MetricsAdvisorWebsiteName
		*out = new(string)
		**out = **in
	}
	if in.NetworkAcls != nil {
		in, out := &in.NetworkAcls, &out.NetworkAcls
		*out = make([]NetworkAclsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OutboundNetworkAccessRestricted != nil {
		in, out := &in.OutboundNetworkAccessRestricted, &out.OutboundNetworkAccessRestricted
		*out = new(bool)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.QnaRuntimeEndpoint != nil {
		in, out := &in.QnaRuntimeEndpoint, &out.QnaRuntimeEndpoint
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]StorageInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountInitParameters.
func (in *AccountInitParameters) DeepCopy() *AccountInitParameters {
	if in == nil {
		return nil
	}
	out := new(AccountInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountList) DeepCopyInto(out *AccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Account, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountList.
func (in *AccountList) DeepCopy() *AccountList {
	if in == nil {
		return nil
	}
	out := new(AccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountObservation) DeepCopyInto(out *AccountObservation) {
	*out = *in
	if in.CustomQuestionAnsweringSearchServiceID != nil {
		in, out := &in.CustomQuestionAnsweringSearchServiceID, &out.CustomQuestionAnsweringSearchServiceID
		*out = new(string)
		**out = **in
	}
	if in.CustomSubdomainName != nil {
		in, out := &in.CustomSubdomainName, &out.CustomSubdomainName
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKey != nil {
		in, out := &in.CustomerManagedKey, &out.CustomerManagedKey
		*out = make([]CustomerManagedKeyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DynamicThrottlingEnabled != nil {
		in, out := &in.DynamicThrottlingEnabled, &out.DynamicThrottlingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Fqdns != nil {
		in, out := &in.Fqdns, &out.Fqdns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.LocalAuthEnabled != nil {
		in, out := &in.LocalAuthEnabled, &out.LocalAuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorAADClientID != nil {
		in, out := &in.MetricsAdvisorAADClientID, &out.MetricsAdvisorAADClientID
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorAADTenantID != nil {
		in, out := &in.MetricsAdvisorAADTenantID, &out.MetricsAdvisorAADTenantID
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorSuperUserName != nil {
		in, out := &in.MetricsAdvisorSuperUserName, &out.MetricsAdvisorSuperUserName
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorWebsiteName != nil {
		in, out := &in.MetricsAdvisorWebsiteName, &out.MetricsAdvisorWebsiteName
		*out = new(string)
		**out = **in
	}
	if in.NetworkAcls != nil {
		in, out := &in.NetworkAcls, &out.NetworkAcls
		*out = make([]NetworkAclsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OutboundNetworkAccessRestricted != nil {
		in, out := &in.OutboundNetworkAccessRestricted, &out.OutboundNetworkAccessRestricted
		*out = new(bool)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.QnaRuntimeEndpoint != nil {
		in, out := &in.QnaRuntimeEndpoint, &out.QnaRuntimeEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]StorageObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountObservation.
func (in *AccountObservation) DeepCopy() *AccountObservation {
	if in == nil {
		return nil
	}
	out := new(AccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountParameters) DeepCopyInto(out *AccountParameters) {
	*out = *in
	if in.CustomQuestionAnsweringSearchServiceID != nil {
		in, out := &in.CustomQuestionAnsweringSearchServiceID, &out.CustomQuestionAnsweringSearchServiceID
		*out = new(string)
		**out = **in
	}
	if in.CustomQuestionAnsweringSearchServiceKeySecretRef != nil {
		in, out := &in.CustomQuestionAnsweringSearchServiceKeySecretRef, &out.CustomQuestionAnsweringSearchServiceKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.CustomSubdomainName != nil {
		in, out := &in.CustomSubdomainName, &out.CustomSubdomainName
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKey != nil {
		in, out := &in.CustomerManagedKey, &out.CustomerManagedKey
		*out = make([]CustomerManagedKeyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DynamicThrottlingEnabled != nil {
		in, out := &in.DynamicThrottlingEnabled, &out.DynamicThrottlingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Fqdns != nil {
		in, out := &in.Fqdns, &out.Fqdns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.LocalAuthEnabled != nil {
		in, out := &in.LocalAuthEnabled, &out.LocalAuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorAADClientID != nil {
		in, out := &in.MetricsAdvisorAADClientID, &out.MetricsAdvisorAADClientID
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorAADTenantID != nil {
		in, out := &in.MetricsAdvisorAADTenantID, &out.MetricsAdvisorAADTenantID
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorSuperUserName != nil {
		in, out := &in.MetricsAdvisorSuperUserName, &out.MetricsAdvisorSuperUserName
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorWebsiteName != nil {
		in, out := &in.MetricsAdvisorWebsiteName, &out.MetricsAdvisorWebsiteName
		*out = new(string)
		**out = **in
	}
	if in.NetworkAcls != nil {
		in, out := &in.NetworkAcls, &out.NetworkAcls
		*out = make([]NetworkAclsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OutboundNetworkAccessRestricted != nil {
		in, out := &in.OutboundNetworkAccessRestricted, &out.OutboundNetworkAccessRestricted
		*out = new(bool)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.QnaRuntimeEndpoint != nil {
		in, out := &in.QnaRuntimeEndpoint, &out.QnaRuntimeEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]StorageParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountParameters.
func (in *AccountParameters) DeepCopy() *AccountParameters {
	if in == nil {
		return nil
	}
	out := new(AccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpec) DeepCopyInto(out *AccountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpec.
func (in *AccountSpec) DeepCopy() *AccountSpec {
	if in == nil {
		return nil
	}
	out := new(AccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountStatus) DeepCopyInto(out *AccountStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountStatus.
func (in *AccountStatus) DeepCopy() *AccountStatus {
	if in == nil {
		return nil
	}
	out := new(AccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerManagedKeyInitParameters) DeepCopyInto(out *CustomerManagedKeyInitParameters) {
	*out = *in
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyID != nil {
		in, out := &in.KeyVaultKeyID, &out.KeyVaultKeyID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerManagedKeyInitParameters.
func (in *CustomerManagedKeyInitParameters) DeepCopy() *CustomerManagedKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(CustomerManagedKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerManagedKeyObservation) DeepCopyInto(out *CustomerManagedKeyObservation) {
	*out = *in
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyID != nil {
		in, out := &in.KeyVaultKeyID, &out.KeyVaultKeyID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerManagedKeyObservation.
func (in *CustomerManagedKeyObservation) DeepCopy() *CustomerManagedKeyObservation {
	if in == nil {
		return nil
	}
	out := new(CustomerManagedKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerManagedKeyParameters) DeepCopyInto(out *CustomerManagedKeyParameters) {
	*out = *in
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyID != nil {
		in, out := &in.KeyVaultKeyID, &out.KeyVaultKeyID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerManagedKeyParameters.
func (in *CustomerManagedKeyParameters) DeepCopy() *CustomerManagedKeyParameters {
	if in == nil {
		return nil
	}
	out := new(CustomerManagedKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityInitParameters) DeepCopyInto(out *IdentityInitParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityInitParameters.
func (in *IdentityInitParameters) DeepCopy() *IdentityInitParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclsInitParameters) DeepCopyInto(out *NetworkAclsInitParameters) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IPRules != nil {
		in, out := &in.IPRules, &out.IPRules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VirtualNetworkRules != nil {
		in, out := &in.VirtualNetworkRules, &out.VirtualNetworkRules
		*out = make([]VirtualNetworkRulesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclsInitParameters.
func (in *NetworkAclsInitParameters) DeepCopy() *NetworkAclsInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkAclsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclsObservation) DeepCopyInto(out *NetworkAclsObservation) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IPRules != nil {
		in, out := &in.IPRules, &out.IPRules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VirtualNetworkRules != nil {
		in, out := &in.VirtualNetworkRules, &out.VirtualNetworkRules
		*out = make([]VirtualNetworkRulesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclsObservation.
func (in *NetworkAclsObservation) DeepCopy() *NetworkAclsObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkAclsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclsParameters) DeepCopyInto(out *NetworkAclsParameters) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IPRules != nil {
		in, out := &in.IPRules, &out.IPRules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VirtualNetworkRules != nil {
		in, out := &in.VirtualNetworkRules, &out.VirtualNetworkRules
		*out = make([]VirtualNetworkRulesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclsParameters.
func (in *NetworkAclsParameters) DeepCopy() *NetworkAclsParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkAclsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageInitParameters) DeepCopyInto(out *StorageInitParameters) {
	*out = *in
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageInitParameters.
func (in *StorageInitParameters) DeepCopy() *StorageInitParameters {
	if in == nil {
		return nil
	}
	out := new(StorageInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObservation) DeepCopyInto(out *StorageObservation) {
	*out = *in
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObservation.
func (in *StorageObservation) DeepCopy() *StorageObservation {
	if in == nil {
		return nil
	}
	out := new(StorageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageParameters) DeepCopyInto(out *StorageParameters) {
	*out = *in
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageParameters.
func (in *StorageParameters) DeepCopy() *StorageParameters {
	if in == nil {
		return nil
	}
	out := new(StorageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRulesInitParameters) DeepCopyInto(out *VirtualNetworkRulesInitParameters) {
	*out = *in
	if in.IgnoreMissingVnetServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVnetServiceEndpoint, &out.IgnoreMissingVnetServiceEndpoint
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRulesInitParameters.
func (in *VirtualNetworkRulesInitParameters) DeepCopy() *VirtualNetworkRulesInitParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRulesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRulesObservation) DeepCopyInto(out *VirtualNetworkRulesObservation) {
	*out = *in
	if in.IgnoreMissingVnetServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVnetServiceEndpoint, &out.IgnoreMissingVnetServiceEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRulesObservation.
func (in *VirtualNetworkRulesObservation) DeepCopy() *VirtualNetworkRulesObservation {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRulesParameters) DeepCopyInto(out *VirtualNetworkRulesParameters) {
	*out = *in
	if in.IgnoreMissingVnetServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVnetServiceEndpoint, &out.IgnoreMissingVnetServiceEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRulesParameters.
func (in *VirtualNetworkRulesParameters) DeepCopy() *VirtualNetworkRulesParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRulesParameters)
	in.DeepCopyInto(out)
	return out
}
