// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using Condition within kubernetes types, where deepcopy-gen is used.
func (in *Condition) DeepCopyInto(out *Condition) {
	p := proto.Clone(in).(*Condition)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition. Required by controller-gen.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Condition. Required by controller-gen.
func (in *Condition) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConditionedStatus within kubernetes types, where deepcopy-gen is used.
func (in *ConditionedStatus) DeepCopyInto(out *ConditionedStatus) {
	p := proto.Clone(in).(*ConditionedStatus)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionedStatus. Required by controller-gen.
func (in *ConditionedStatus) DeepCopy() *ConditionedStatus {
	if in == nil {
		return nil
	}
	out := new(ConditionedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConditionedStatus. Required by controller-gen.
func (in *ConditionedStatus) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
