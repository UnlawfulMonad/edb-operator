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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDatabase) DeepCopyInto(out *ExternalDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDatabase.
func (in *ExternalDatabase) DeepCopy() *ExternalDatabase {
	if in == nil {
		return nil
	}
	out := new(ExternalDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDatabaseList) DeepCopyInto(out *ExternalDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDatabaseList.
func (in *ExternalDatabaseList) DeepCopy() *ExternalDatabaseList {
	if in == nil {
		return nil
	}
	out := new(ExternalDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDatabaseReference) DeepCopyInto(out *ExternalDatabaseReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDatabaseReference.
func (in *ExternalDatabaseReference) DeepCopy() *ExternalDatabaseReference {
	if in == nil {
		return nil
	}
	out := new(ExternalDatabaseReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDatabaseSpec) DeepCopyInto(out *ExternalDatabaseSpec) {
	*out = *in
	out.AdminPassword = in.AdminPassword
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDatabaseSpec.
func (in *ExternalDatabaseSpec) DeepCopy() *ExternalDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDatabaseStatus) DeepCopyInto(out *ExternalDatabaseStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDatabaseStatus.
func (in *ExternalDatabaseStatus) DeepCopy() *ExternalDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLDatabase) DeepCopyInto(out *MySQLDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLDatabase.
func (in *MySQLDatabase) DeepCopy() *MySQLDatabase {
	if in == nil {
		return nil
	}
	out := new(MySQLDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLDatabaseList) DeepCopyInto(out *MySQLDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLDatabaseList.
func (in *MySQLDatabaseList) DeepCopy() *MySQLDatabaseList {
	if in == nil {
		return nil
	}
	out := new(MySQLDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLDatabaseSpec) DeepCopyInto(out *MySQLDatabaseSpec) {
	*out = *in
	out.ExternalDatabaseRef = in.ExternalDatabaseRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLDatabaseSpec.
func (in *MySQLDatabaseSpec) DeepCopy() *MySQLDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLDatabaseStatus) DeepCopyInto(out *MySQLDatabaseStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLDatabaseStatus.
func (in *MySQLDatabaseStatus) DeepCopy() *MySQLDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(MySQLDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLUser) DeepCopyInto(out *MySQLUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLUser.
func (in *MySQLUser) DeepCopy() *MySQLUser {
	if in == nil {
		return nil
	}
	out := new(MySQLUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLUserList) DeepCopyInto(out *MySQLUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLUserList.
func (in *MySQLUserList) DeepCopy() *MySQLUserList {
	if in == nil {
		return nil
	}
	out := new(MySQLUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLUserSpec) DeepCopyInto(out *MySQLUserSpec) {
	*out = *in
	if in.ExternalDatabaseRef != nil {
		in, out := &in.ExternalDatabaseRef, &out.ExternalDatabaseRef
		*out = new(ExternalDatabaseReference)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLUserSpec.
func (in *MySQLUserSpec) DeepCopy() *MySQLUserSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLUserStatus) DeepCopyInto(out *MySQLUserStatus) {
	*out = *in
	if in.PasswordSecretName != nil {
		in, out := &in.PasswordSecretName, &out.PasswordSecretName
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLUserStatus.
func (in *MySQLUserStatus) DeepCopy() *MySQLUserStatus {
	if in == nil {
		return nil
	}
	out := new(MySQLUserStatus)
	in.DeepCopyInto(out)
	return out
}
