// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using OdaInfo within kubernetes types, where deepcopy-gen is used.
func (in *OdaInfo) DeepCopyInto(out *OdaInfo) {
	p := proto.Clone(in).(*OdaInfo)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OdaInfo. Required by controller-gen.
func (in *OdaInfo) DeepCopy() *OdaInfo {
	if in == nil {
		return nil
	}
	out := new(OdaInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new OdaInfo. Required by controller-gen.
func (in *OdaInfo) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
