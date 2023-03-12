//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "go.linka.cloud/k8s/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedCSIDriver) DeepCopyInto(out *AllowedCSIDriver) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedCSIDriver.
func (in *AllowedCSIDriver) DeepCopy() *AllowedCSIDriver {
	if in == nil {
		return nil
	}
	out := new(AllowedCSIDriver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedFlexVolume) DeepCopyInto(out *AllowedFlexVolume) {
	*out = *in
	if in.Driver != nil {
		in, out := &in.Driver, &out.Driver
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedFlexVolume.
func (in *AllowedFlexVolume) DeepCopy() *AllowedFlexVolume {
	if in == nil {
		return nil
	}
	out := new(AllowedFlexVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedHostPath) DeepCopyInto(out *AllowedHostPath) {
	*out = *in
	if in.PathPrefix != nil {
		in, out := &in.PathPrefix, &out.PathPrefix
		*out = new(string)
		**out = **in
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedHostPath.
func (in *AllowedHostPath) DeepCopy() *AllowedHostPath {
	if in == nil {
		return nil
	}
	out := new(AllowedHostPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Eviction) DeepCopyInto(out *Eviction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.DeleteOptions != nil {
		in, out := &in.DeleteOptions, &out.DeleteOptions
		*out = new(v1.DeleteOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Eviction.
func (in *Eviction) DeepCopy() *Eviction {
	if in == nil {
		return nil
	}
	out := new(Eviction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Eviction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FSGroupStrategyOptions) DeepCopyInto(out *FSGroupStrategyOptions) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(FSGroupStrategyType)
		**out = **in
	}
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]IDRange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FSGroupStrategyOptions.
func (in *FSGroupStrategyOptions) DeepCopy() *FSGroupStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(FSGroupStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPortRange) DeepCopyInto(out *HostPortRange) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(int32)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPortRange.
func (in *HostPortRange) DeepCopy() *HostPortRange {
	if in == nil {
		return nil
	}
	out := new(HostPortRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDRange) DeepCopyInto(out *IDRange) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(int64)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDRange.
func (in *IDRange) DeepCopy() *IDRange {
	if in == nil {
		return nil
	}
	out := new(IDRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudget) DeepCopyInto(out *PodDisruptionBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(PodDisruptionBudgetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(PodDisruptionBudgetStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudget.
func (in *PodDisruptionBudget) DeepCopy() *PodDisruptionBudget {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodDisruptionBudget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetList) DeepCopyInto(out *PodDisruptionBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodDisruptionBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetList.
func (in *PodDisruptionBudgetList) DeepCopy() *PodDisruptionBudgetList {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodDisruptionBudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetStatus) DeepCopyInto(out *PodDisruptionBudgetStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.DisruptedPods != nil {
		in, out := &in.DisruptedPods, &out.DisruptedPods
		*out = make(map[string]v1.Time, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.DisruptionsAllowed != nil {
		in, out := &in.DisruptionsAllowed, &out.DisruptionsAllowed
		*out = new(int32)
		**out = **in
	}
	if in.CurrentHealthy != nil {
		in, out := &in.CurrentHealthy, &out.CurrentHealthy
		*out = new(int32)
		**out = **in
	}
	if in.DesiredHealthy != nil {
		in, out := &in.DesiredHealthy, &out.DesiredHealthy
		*out = new(int32)
		**out = **in
	}
	if in.ExpectedPods != nil {
		in, out := &in.ExpectedPods, &out.ExpectedPods
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetStatus.
func (in *PodDisruptionBudgetStatus) DeepCopy() *PodDisruptionBudgetStatus {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicy) DeepCopyInto(out *PodSecurityPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(PodSecurityPolicySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicy.
func (in *PodSecurityPolicy) DeepCopy() *PodSecurityPolicy {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSecurityPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicyList) DeepCopyInto(out *PodSecurityPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodSecurityPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicyList.
func (in *PodSecurityPolicyList) DeepCopy() *PodSecurityPolicyList {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSecurityPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicySpec) DeepCopyInto(out *PodSecurityPolicySpec) {
	*out = *in
	if in.Privileged != nil {
		in, out := &in.Privileged, &out.Privileged
		*out = new(bool)
		**out = **in
	}
	if in.DefaultAddCapabilities != nil {
		in, out := &in.DefaultAddCapabilities, &out.DefaultAddCapabilities
		*out = make([]corev1.Capability, len(*in))
		copy(*out, *in)
	}
	if in.RequiredDropCapabilities != nil {
		in, out := &in.RequiredDropCapabilities, &out.RequiredDropCapabilities
		*out = make([]corev1.Capability, len(*in))
		copy(*out, *in)
	}
	if in.AllowedCapabilities != nil {
		in, out := &in.AllowedCapabilities, &out.AllowedCapabilities
		*out = make([]corev1.Capability, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]FSType, len(*in))
		copy(*out, *in)
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	if in.HostPorts != nil {
		in, out := &in.HostPorts, &out.HostPorts
		*out = make([]HostPortRange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostPID != nil {
		in, out := &in.HostPID, &out.HostPID
		*out = new(bool)
		**out = **in
	}
	if in.HostIPC != nil {
		in, out := &in.HostIPC, &out.HostIPC
		*out = new(bool)
		**out = **in
	}
	if in.SELinux != nil {
		in, out := &in.SELinux, &out.SELinux
		*out = new(SELinuxStrategyOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.RunAsUser != nil {
		in, out := &in.RunAsUser, &out.RunAsUser
		*out = new(RunAsUserStrategyOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.RunAsGroup != nil {
		in, out := &in.RunAsGroup, &out.RunAsGroup
		*out = new(RunAsGroupStrategyOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.SupplementalGroups != nil {
		in, out := &in.SupplementalGroups, &out.SupplementalGroups
		*out = new(SupplementalGroupsStrategyOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.FSGroup != nil {
		in, out := &in.FSGroup, &out.FSGroup
		*out = new(FSGroupStrategyOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadOnlyRootFilesystem != nil {
		in, out := &in.ReadOnlyRootFilesystem, &out.ReadOnlyRootFilesystem
		*out = new(bool)
		**out = **in
	}
	if in.DefaultAllowPrivilegeEscalation != nil {
		in, out := &in.DefaultAllowPrivilegeEscalation, &out.DefaultAllowPrivilegeEscalation
		*out = new(bool)
		**out = **in
	}
	if in.AllowPrivilegeEscalation != nil {
		in, out := &in.AllowPrivilegeEscalation, &out.AllowPrivilegeEscalation
		*out = new(bool)
		**out = **in
	}
	if in.AllowedHostPaths != nil {
		in, out := &in.AllowedHostPaths, &out.AllowedHostPaths
		*out = make([]AllowedHostPath, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedFlexVolumes != nil {
		in, out := &in.AllowedFlexVolumes, &out.AllowedFlexVolumes
		*out = make([]AllowedFlexVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedCSIDrivers != nil {
		in, out := &in.AllowedCSIDrivers, &out.AllowedCSIDrivers
		*out = make([]AllowedCSIDriver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedUnsafeSysctls != nil {
		in, out := &in.AllowedUnsafeSysctls, &out.AllowedUnsafeSysctls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ForbiddenSysctls != nil {
		in, out := &in.ForbiddenSysctls, &out.ForbiddenSysctls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedProcMountTypes != nil {
		in, out := &in.AllowedProcMountTypes, &out.AllowedProcMountTypes
		*out = make([]corev1.ProcMountType, len(*in))
		copy(*out, *in)
	}
	if in.RuntimeClass != nil {
		in, out := &in.RuntimeClass, &out.RuntimeClass
		*out = new(RuntimeClassStrategyOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicySpec.
func (in *PodSecurityPolicySpec) DeepCopy() *PodSecurityPolicySpec {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsGroupStrategyOptions) DeepCopyInto(out *RunAsGroupStrategyOptions) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(RunAsGroupStrategy)
		**out = **in
	}
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]IDRange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsGroupStrategyOptions.
func (in *RunAsGroupStrategyOptions) DeepCopy() *RunAsGroupStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(RunAsGroupStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsUserStrategyOptions) DeepCopyInto(out *RunAsUserStrategyOptions) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(RunAsUserStrategy)
		**out = **in
	}
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]IDRange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsUserStrategyOptions.
func (in *RunAsUserStrategyOptions) DeepCopy() *RunAsUserStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(RunAsUserStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeClassStrategyOptions) DeepCopyInto(out *RuntimeClassStrategyOptions) {
	*out = *in
	if in.AllowedRuntimeClassNames != nil {
		in, out := &in.AllowedRuntimeClassNames, &out.AllowedRuntimeClassNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultRuntimeClassName != nil {
		in, out := &in.DefaultRuntimeClassName, &out.DefaultRuntimeClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeClassStrategyOptions.
func (in *RuntimeClassStrategyOptions) DeepCopy() *RuntimeClassStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(RuntimeClassStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SELinuxStrategyOptions) DeepCopyInto(out *SELinuxStrategyOptions) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(SELinuxStrategy)
		**out = **in
	}
	if in.SELinuxOptions != nil {
		in, out := &in.SELinuxOptions, &out.SELinuxOptions
		*out = new(corev1.SELinuxOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SELinuxStrategyOptions.
func (in *SELinuxStrategyOptions) DeepCopy() *SELinuxStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(SELinuxStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupplementalGroupsStrategyOptions) DeepCopyInto(out *SupplementalGroupsStrategyOptions) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(SupplementalGroupsStrategyType)
		**out = **in
	}
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]IDRange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupplementalGroupsStrategyOptions.
func (in *SupplementalGroupsStrategyOptions) DeepCopy() *SupplementalGroupsStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(SupplementalGroupsStrategyOptions)
	in.DeepCopyInto(out)
	return out
}
