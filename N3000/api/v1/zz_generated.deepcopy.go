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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FortvilleMAC) DeepCopyInto(out *FortvilleMAC) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FortvilleMAC.
func (in *FortvilleMAC) DeepCopy() *FortvilleMAC {
	if in == nil {
		return nil
	}
	out := new(FortvilleMAC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *N3000) DeepCopyInto(out *N3000) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new N3000.
func (in *N3000) DeepCopy() *N3000 {
	if in == nil {
		return nil
	}
	out := new(N3000)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *N3000) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *N3000Fortville) DeepCopyInto(out *N3000Fortville) {
	*out = *in
	if in.MACs != nil {
		in, out := &in.MACs, &out.MACs
		*out = make([]FortvilleMAC, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new N3000Fortville.
func (in *N3000Fortville) DeepCopy() *N3000Fortville {
	if in == nil {
		return nil
	}
	out := new(N3000Fortville)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *N3000Fpga) DeepCopyInto(out *N3000Fpga) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new N3000Fpga.
func (in *N3000Fpga) DeepCopy() *N3000Fpga {
	if in == nil {
		return nil
	}
	out := new(N3000Fpga)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *N3000List) DeepCopyInto(out *N3000List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]N3000, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new N3000List.
func (in *N3000List) DeepCopy() *N3000List {
	if in == nil {
		return nil
	}
	out := new(N3000List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *N3000List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *N3000Node) DeepCopyInto(out *N3000Node) {
	*out = *in
	if in.FPGA != nil {
		in, out := &in.FPGA, &out.FPGA
		*out = make([]N3000Fpga, len(*in))
		copy(*out, *in)
	}
	in.Fortville.DeepCopyInto(&out.Fortville)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new N3000Node.
func (in *N3000Node) DeepCopy() *N3000Node {
	if in == nil {
		return nil
	}
	out := new(N3000Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *N3000Spec) DeepCopyInto(out *N3000Spec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]N3000Node, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new N3000Spec.
func (in *N3000Spec) DeepCopy() *N3000Spec {
	if in == nil {
		return nil
	}
	out := new(N3000Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *N3000Status) DeepCopyInto(out *N3000Status) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new N3000Status.
func (in *N3000Status) DeepCopy() *N3000Status {
	if in == nil {
		return nil
	}
	out := new(N3000Status)
	in.DeepCopyInto(out)
	return out
}
