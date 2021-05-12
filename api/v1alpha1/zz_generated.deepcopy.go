// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACL) DeepCopyInto(out *ACL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACL.
func (in *ACL) DeepCopy() *ACL {
	if in == nil {
		return nil
	}
	out := new(ACL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OVSSwitch) DeepCopyInto(out *OVSSwitch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OVSSwitch.
func (in *OVSSwitch) DeepCopy() *OVSSwitch {
	if in == nil {
		return nil
	}
	out := new(OVSSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OVSSwitch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OVSSwitchList) DeepCopyInto(out *OVSSwitchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OVSSwitch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OVSSwitchList.
func (in *OVSSwitchList) DeepCopy() *OVSSwitchList {
	if in == nil {
		return nil
	}
	out := new(OVSSwitchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OVSSwitchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OVSSwitchSpec) DeepCopyInto(out *OVSSwitchSpec) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OVSSwitchSpec.
func (in *OVSSwitchSpec) DeepCopy() *OVSSwitchSpec {
	if in == nil {
		return nil
	}
	out := new(OVSSwitchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OVSSwitchStatus) DeepCopyInto(out *OVSSwitchStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OVSSwitchStatus.
func (in *OVSSwitchStatus) DeepCopy() *OVSSwitchStatus {
	if in == nil {
		return nil
	}
	out := new(OVSSwitchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSwitchRef) DeepCopyInto(out *ProviderSwitchRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSwitchRef.
func (in *ProviderSwitchRef) DeepCopy() *ProviderSwitchRef {
	if in == nil {
		return nil
	}
	out := new(ProviderSwitchRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Switch) DeepCopyInto(out *Switch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Switch.
func (in *Switch) DeepCopy() *Switch {
	if in == nil {
		return nil
	}
	out := new(Switch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Switch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchList) DeepCopyInto(out *SwitchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Switch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchList.
func (in *SwitchList) DeepCopy() *SwitchList {
	if in == nil {
		return nil
	}
	out := new(SwitchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPort) DeepCopyInto(out *SwitchPort) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPort.
func (in *SwitchPort) DeepCopy() *SwitchPort {
	if in == nil {
		return nil
	}
	out := new(SwitchPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchPort) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortConfiguration) DeepCopyInto(out *SwitchPortConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortConfiguration.
func (in *SwitchPortConfiguration) DeepCopy() *SwitchPortConfiguration {
	if in == nil {
		return nil
	}
	out := new(SwitchPortConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchPortConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortConfigurationList) DeepCopyInto(out *SwitchPortConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SwitchPortConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortConfigurationList.
func (in *SwitchPortConfigurationList) DeepCopy() *SwitchPortConfigurationList {
	if in == nil {
		return nil
	}
	out := new(SwitchPortConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchPortConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortConfigurationReference) DeepCopyInto(out *SwitchPortConfigurationReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortConfigurationReference.
func (in *SwitchPortConfigurationReference) DeepCopy() *SwitchPortConfigurationReference {
	if in == nil {
		return nil
	}
	out := new(SwitchPortConfigurationReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortConfigurationSpec) DeepCopyInto(out *SwitchPortConfigurationSpec) {
	*out = *in
	if in.ACLs != nil {
		in, out := &in.ACLs, &out.ACLs
		*out = make([]ACL, len(*in))
		copy(*out, *in)
	}
	if in.Vlans != nil {
		in, out := &in.Vlans, &out.Vlans
		*out = make([]VLAN, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortConfigurationSpec.
func (in *SwitchPortConfigurationSpec) DeepCopy() *SwitchPortConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(SwitchPortConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortConfigurationStatus) DeepCopyInto(out *SwitchPortConfigurationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortConfigurationStatus.
func (in *SwitchPortConfigurationStatus) DeepCopy() *SwitchPortConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(SwitchPortConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortList) DeepCopyInto(out *SwitchPortList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SwitchPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortList.
func (in *SwitchPortList) DeepCopy() *SwitchPortList {
	if in == nil {
		return nil
	}
	out := new(SwitchPortList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchPortList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortReference) DeepCopyInto(out *SwitchPortReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortReference.
func (in *SwitchPortReference) DeepCopy() *SwitchPortReference {
	if in == nil {
		return nil
	}
	out := new(SwitchPortReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortSpec) DeepCopyInto(out *SwitchPortSpec) {
	*out = *in
	if in.ConfigurationRef != nil {
		in, out := &in.ConfigurationRef, &out.ConfigurationRef
		*out = new(SwitchPortConfigurationReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortSpec.
func (in *SwitchPortSpec) DeepCopy() *SwitchPortSpec {
	if in == nil {
		return nil
	}
	out := new(SwitchPortSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortStatus) DeepCopyInto(out *SwitchPortStatus) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(SwitchPortConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortStatus.
func (in *SwitchPortStatus) DeepCopy() *SwitchPortStatus {
	if in == nil {
		return nil
	}
	out := new(SwitchPortStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSpec) DeepCopyInto(out *SwitchSpec) {
	*out = *in
	if in.ProviderSwitch != nil {
		in, out := &in.ProviderSwitch, &out.ProviderSwitch
		*out = new(ProviderSwitchRef)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make(map[string]Port, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchSpec.
func (in *SwitchSpec) DeepCopy() *SwitchSpec {
	if in == nil {
		return nil
	}
	out := new(SwitchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchStatus) DeepCopyInto(out *SwitchStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchStatus.
func (in *SwitchStatus) DeepCopy() *SwitchStatus {
	if in == nil {
		return nil
	}
	out := new(SwitchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VLAN) DeepCopyInto(out *VLAN) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VLAN.
func (in *VLAN) DeepCopy() *VLAN {
	if in == nil {
		return nil
	}
	out := new(VLAN)
	in.DeepCopyInto(out)
	return out
}
