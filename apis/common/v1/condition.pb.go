// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: apis/common/v1/condition.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ConditionStatus int32

const (
	ConditionStatus_ConditionStatus_Unknown ConditionStatus = 0
	ConditionStatus_ConditionStatus_False   ConditionStatus = 1
	ConditionStatus_ConditionStatus_True    ConditionStatus = 2
)

var ConditionStatus_name = map[int32]string{
	0: "ConditionStatus_Unknown",
	1: "ConditionStatus_False",
	2: "ConditionStatus_True",
}

var ConditionStatus_value = map[string]int32{
	"ConditionStatus_Unknown": 0,
	"ConditionStatus_False":   1,
	"ConditionStatus_True":    2,
}

func (x ConditionStatus) String() string {
	return proto.EnumName(ConditionStatus_name, int32(x))
}

func (ConditionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c85e686f5022c5a9, []int{0}
}

type ConditionKind int32

const (
	ConditionKind_ConditionKind_Unspecified ConditionKind = 0
	ConditionKind_ConditionKind_RootPath    ConditionKind = 1
	ConditionKind_ConditionKind_Target      ConditionKind = 2
	ConditionKind_ConditionKind_Synced      ConditionKind = 3
	ConditionKind_ConditionKind_Ready       ConditionKind = 4
)

var ConditionKind_name = map[int32]string{
	0: "ConditionKind_Unspecified",
	1: "ConditionKind_RootPath",
	2: "ConditionKind_Target",
	3: "ConditionKind_Synced",
	4: "ConditionKind_Ready",
}

var ConditionKind_value = map[string]int32{
	"ConditionKind_Unspecified": 0,
	"ConditionKind_RootPath":    1,
	"ConditionKind_Target":      2,
	"ConditionKind_Synced":      3,
	"ConditionKind_Ready":       4,
}

func (x ConditionKind) String() string {
	return proto.EnumName(ConditionKind_name, int32(x))
}

func (ConditionKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c85e686f5022c5a9, []int{1}
}

type ConditionReason int32

const (
	ConditionReason_ConditionReason_Unspecified ConditionReason = 0
	// Reasons a resource validation failed
	ConditionReason_ConditionReason_Success ConditionReason = 1
	ConditionReason_ConditionReason_Failed  ConditionReason = 2
	// Reasons a resource is or is not ready
	ConditionReason_ConditionReason_Unknown     ConditionReason = 3
	ConditionReason_ConditionReason_Creating    ConditionReason = 4
	ConditionReason_ConditionReason_Deleting    ConditionReason = 5
	ConditionReason_ConditionReason_Updating    ConditionReason = 6
	ConditionReason_ConditionReason_Unavailable ConditionReason = 7
	ConditionReason_ConditionReason_Available   ConditionReason = 8
	ConditionReason_ConditionReason_Pending     ConditionReason = 9
	// Reasons a resource is or is not synced.
	ConditionReason_ConditionReason_ReconcileSuccess ConditionReason = 10
	ConditionReason_ConditionReason_ReconcileFailed  ConditionReason = 11
)

var ConditionReason_name = map[int32]string{
	0:  "ConditionReason_Unspecified",
	1:  "ConditionReason_Success",
	2:  "ConditionReason_Failed",
	3:  "ConditionReason_Unknown",
	4:  "ConditionReason_Creating",
	5:  "ConditionReason_Deleting",
	6:  "ConditionReason_Updating",
	7:  "ConditionReason_Unavailable",
	8:  "ConditionReason_Available",
	9:  "ConditionReason_Pending",
	10: "ConditionReason_ReconcileSuccess",
	11: "ConditionReason_ReconcileFailed",
}

var ConditionReason_value = map[string]int32{
	"ConditionReason_Unspecified":      0,
	"ConditionReason_Success":          1,
	"ConditionReason_Failed":           2,
	"ConditionReason_Unknown":          3,
	"ConditionReason_Creating":         4,
	"ConditionReason_Deleting":         5,
	"ConditionReason_Updating":         6,
	"ConditionReason_Unavailable":      7,
	"ConditionReason_Available":        8,
	"ConditionReason_Pending":          9,
	"ConditionReason_ReconcileSuccess": 10,
	"ConditionReason_ReconcileFailed":  11,
}

func (x ConditionReason) String() string {
	return proto.EnumName(ConditionReason_name, int32(x))
}

func (ConditionReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c85e686f5022c5a9, []int{2}
}

type Condition struct {
	Kind               ConditionKind   `protobuf:"varint,1,opt,name=kind,proto3,enum=common.v1.ConditionKind" json:"kind,omitempty"`
	Status             ConditionStatus `protobuf:"varint,2,opt,name=status,proto3,enum=common.v1.ConditionStatus" json:"status,omitempty"`
	LastTransitionTime *v1.Time        `protobuf:"bytes,3,opt,name=lastTransitionTime,proto3" json:"lastTransitionTime,omitempty"`
	//optional github.com.kubernetes.apimachinery.pkg.apis.meta.v1.Time lastTransitionTime = 3;
	Reason               ConditionReason `protobuf:"varint,4,opt,name=reason,proto3,enum=common.v1.ConditionReason" json:"reason,omitempty"`
	Message              string          `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85e686f5022c5a9, []int{0}
}
func (m *Condition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return m.Size()
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetKind() ConditionKind {
	if m != nil {
		return m.Kind
	}
	return ConditionKind_ConditionKind_Unspecified
}

func (m *Condition) GetStatus() ConditionStatus {
	if m != nil {
		return m.Status
	}
	return ConditionStatus_ConditionStatus_Unknown
}

func (m *Condition) GetLastTransitionTime() *v1.Time {
	if m != nil {
		return m.LastTransitionTime
	}
	return nil
}

func (m *Condition) GetReason() ConditionReason {
	if m != nil {
		return m.Reason
	}
	return ConditionReason_ConditionReason_Unspecified
}

func (m *Condition) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ConditionedStatus struct {
	Conditions           []*Condition `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConditionedStatus) Reset()         { *m = ConditionedStatus{} }
func (m *ConditionedStatus) String() string { return proto.CompactTextString(m) }
func (*ConditionedStatus) ProtoMessage()    {}
func (*ConditionedStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85e686f5022c5a9, []int{1}
}
func (m *ConditionedStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConditionedStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConditionedStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConditionedStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConditionedStatus.Merge(m, src)
}
func (m *ConditionedStatus) XXX_Size() int {
	return m.Size()
}
func (m *ConditionedStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ConditionedStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ConditionedStatus proto.InternalMessageInfo

func (m *ConditionedStatus) GetConditions() []*Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.v1.ConditionStatus", ConditionStatus_name, ConditionStatus_value)
	proto.RegisterEnum("common.v1.ConditionKind", ConditionKind_name, ConditionKind_value)
	proto.RegisterEnum("common.v1.ConditionReason", ConditionReason_name, ConditionReason_value)
	proto.RegisterType((*Condition)(nil), "common.v1.Condition")
	proto.RegisterType((*ConditionedStatus)(nil), "common.v1.ConditionedStatus")
}

func init() { proto.RegisterFile("apis/common/v1/condition.proto", fileDescriptor_c85e686f5022c5a9) }

var fileDescriptor_c85e686f5022c5a9 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0xeb, 0x24, 0x4d, 0xbf, 0xdc, 0xe8, 0x03, 0x33, 0x14, 0xea, 0xa6, 0x90, 0x46, 0x85,
	0x45, 0x14, 0x21, 0x5b, 0x0d, 0x5d, 0xb0, 0x60, 0x03, 0x45, 0x95, 0x10, 0x9b, 0xca, 0x49, 0x36,
	0xdd, 0x44, 0x13, 0xcf, 0xc5, 0x19, 0xc5, 0x9e, 0xb1, 0x3c, 0x93, 0xa0, 0xbc, 0x02, 0x2f, 0x50,
	0x1e, 0x89, 0x25, 0x8f, 0x80, 0xc2, 0x8b, 0x20, 0xff, 0x49, 0x14, 0x1b, 0x77, 0xe9, 0xfb, 0x3b,
	0xf7, 0xcc, 0xbd, 0x67, 0xac, 0x81, 0x2e, 0x8d, 0xb8, 0x72, 0x3c, 0x19, 0x86, 0x52, 0x38, 0xab,
	0x4b, 0xc7, 0x93, 0x82, 0x71, 0xcd, 0xa5, 0xb0, 0xa3, 0x58, 0x6a, 0x49, 0x5a, 0x19, 0xb2, 0x57,
	0x97, 0x9d, 0xab, 0xc5, 0x3b, 0x65, 0x73, 0xe9, 0xd0, 0x88, 0x87, 0xd4, 0x9b, 0x73, 0x81, 0xf1,
	0xda, 0x89, 0x16, 0xbe, 0x93, 0x5a, 0x84, 0xa8, 0x69, 0x62, 0xe0, 0xa3, 0xc0, 0x98, 0x6a, 0x64,
	0x99, 0xc1, 0xc5, 0x7d, 0x0d, 0x5a, 0xd7, 0x5b, 0x53, 0xf2, 0x06, 0x1a, 0x0b, 0x2e, 0x98, 0x65,
	0xf4, 0x8c, 0xfe, 0xa3, 0xa1, 0x65, 0xef, 0xdc, 0xed, 0x9d, 0xe6, 0x0b, 0x17, 0xcc, 0x4d, 0x55,
	0x64, 0x08, 0x4d, 0xa5, 0xa9, 0x5e, 0x2a, 0xab, 0x96, 0xea, 0x3b, 0x55, 0xfa, 0x51, 0xaa, 0x70,
	0x73, 0x25, 0xb9, 0x03, 0x12, 0x50, 0xa5, 0xc7, 0x31, 0x15, 0x2a, 0xe5, 0x63, 0x1e, 0xa2, 0x55,
	0xef, 0x19, 0xfd, 0xf6, 0x70, 0x60, 0x67, 0x2b, 0xd8, 0xfb, 0x2b, 0xd8, 0xd1, 0xc2, 0x4f, 0x0a,
	0xca, 0x4e, 0x56, 0x48, 0xac, 0x93, 0x0e, 0xb7, 0xc2, 0x25, 0x99, 0x27, 0x46, 0xaa, 0xa4, 0xb0,
	0x1a, 0x0f, 0xcf, 0xe3, 0xa6, 0x0a, 0x37, 0x57, 0x12, 0x0b, 0x8e, 0x42, 0x54, 0x8a, 0xfa, 0x68,
	0x1d, 0xf6, 0x8c, 0x7e, 0xcb, 0xdd, 0x7e, 0x5e, 0x7c, 0x86, 0x27, 0xbb, 0x26, 0x64, 0xd9, 0x1a,
	0xe4, 0x0a, 0x60, 0x77, 0x05, 0xca, 0x32, 0x7a, 0xf5, 0x7e, 0x7b, 0x78, 0x5c, 0x79, 0xcc, 0x9e,
	0x6e, 0xe0, 0xc1, 0xe3, 0x52, 0x1e, 0xe4, 0x0c, 0x4e, 0x4a, 0xa5, 0xe9, 0x44, 0x2c, 0x84, 0xfc,
	0x26, 0xcc, 0x03, 0x72, 0x0a, 0xcf, 0xca, 0xf0, 0x86, 0x06, 0x0a, 0x4d, 0x83, 0x58, 0x70, 0x5c,
	0x46, 0xe3, 0x78, 0x89, 0x66, 0x6d, 0x70, 0x6f, 0xc0, 0xff, 0x85, 0x5b, 0x22, 0x2f, 0xe1, 0xb4,
	0x50, 0x98, 0x4e, 0x84, 0x8a, 0xd0, 0xe3, 0x5f, 0x39, 0x32, 0xf3, 0x80, 0x74, 0xe0, 0x79, 0x11,
	0xbb, 0x52, 0xea, 0x5b, 0xaa, 0xe7, 0xa5, 0x63, 0x52, 0x36, 0xa6, 0xb1, 0x8f, 0xda, 0xac, 0xfd,
	0x4b, 0x46, 0x6b, 0xe1, 0x21, 0x33, 0xeb, 0xe4, 0x04, 0x9e, 0x96, 0xfc, 0x90, 0xb2, 0xb5, 0xd9,
	0x18, 0x7c, 0xaf, 0xef, 0xed, 0x9f, 0xe5, 0x4f, 0xce, 0xe1, 0xac, 0x54, 0x2a, 0x4d, 0xb7, 0x1f,
	0x50, 0x2e, 0x18, 0x2d, 0x3d, 0x0f, 0x95, 0x32, 0x8d, 0xc2, 0xe8, 0x39, 0xbc, 0xa1, 0x3c, 0x40,
	0x66, 0xd6, 0xaa, 0x1a, 0xb7, 0xc9, 0xd6, 0xc9, 0x0b, 0xb0, 0xca, 0xf0, 0x3a, 0x46, 0xaa, 0xb9,
	0xf0, 0xcd, 0x46, 0x15, 0xfd, 0x84, 0x01, 0xa6, 0xf4, 0xb0, 0x8a, 0x4e, 0x22, 0x96, 0xf5, 0x36,
	0xab, 0x17, 0xa2, 0x2b, 0xca, 0x03, 0x3a, 0x0b, 0xd0, 0x3c, 0x2a, 0xdc, 0x46, 0x2e, 0xf8, 0xb0,
	0xc3, 0xff, 0x55, 0x8d, 0x7d, 0x8b, 0x82, 0x25, 0xe6, 0x2d, 0xf2, 0x1a, 0x7a, 0x65, 0xe8, 0xa2,
	0x27, 0x85, 0xc7, 0x03, 0xdc, 0xa6, 0x02, 0xe4, 0x15, 0x9c, 0x3f, 0xa8, 0xca, 0xe3, 0x69, 0x7f,
	0x7c, 0xff, 0x73, 0xd3, 0x35, 0x7e, 0x6d, 0xba, 0xc6, 0xef, 0x4d, 0xd7, 0xf8, 0xf1, 0xa7, 0x7b,
	0x70, 0x37, 0xf0, 0xb9, 0x9e, 0x2f, 0x67, 0xc9, 0x5f, 0xec, 0xac, 0x05, 0x63, 0x8e, 0x60, 0x6c,
	0x1a, 0x2f, 0x85, 0xe6, 0x21, 0x3a, 0xc5, 0xf7, 0x67, 0xd6, 0x4c, 0x5f, 0x8d, 0xb7, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x67, 0xea, 0xb2, 0x98, 0x04, 0x00, 0x00,
}

func (m *Condition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Condition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Condition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintCondition(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Reason != 0 {
		i = encodeVarintCondition(dAtA, i, uint64(m.Reason))
		i--
		dAtA[i] = 0x20
	}
	if m.LastTransitionTime != nil {
		{
			size, err := m.LastTransitionTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCondition(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Status != 0 {
		i = encodeVarintCondition(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.Kind != 0 {
		i = encodeVarintCondition(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ConditionedStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConditionedStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConditionedStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Conditions) > 0 {
		for iNdEx := len(m.Conditions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Conditions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCondition(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCondition(dAtA []byte, offset int, v uint64) int {
	offset -= sovCondition(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Condition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovCondition(uint64(m.Kind))
	}
	if m.Status != 0 {
		n += 1 + sovCondition(uint64(m.Status))
	}
	if m.LastTransitionTime != nil {
		l = m.LastTransitionTime.Size()
		n += 1 + l + sovCondition(uint64(l))
	}
	if m.Reason != 0 {
		n += 1 + sovCondition(uint64(m.Reason))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovCondition(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConditionedStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Conditions) > 0 {
		for _, e := range m.Conditions {
			l = e.Size()
			n += 1 + l + sovCondition(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCondition(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCondition(x uint64) (n int) {
	return sovCondition(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Condition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCondition
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Condition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Condition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCondition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= ConditionKind(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCondition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ConditionStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTransitionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCondition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCondition
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCondition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastTransitionTime == nil {
				m.LastTransitionTime = &v1.Time{}
			}
			if err := m.LastTransitionTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			m.Reason = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCondition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reason |= ConditionReason(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCondition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCondition
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCondition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCondition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCondition
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConditionedStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCondition
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConditionedStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConditionedStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCondition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCondition
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCondition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conditions = append(m.Conditions, &Condition{})
			if err := m.Conditions[len(m.Conditions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCondition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCondition
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCondition(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCondition
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCondition
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCondition
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCondition
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCondition
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCondition
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCondition        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCondition          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCondition = fmt.Errorf("proto: unexpected end of group")
)
