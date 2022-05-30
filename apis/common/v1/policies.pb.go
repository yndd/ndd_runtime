// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: apis/common/v1/policies.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
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

type DeploymentPolicy int32

const (
	DeploymentPolicy_DeploymentPolicy_Active  DeploymentPolicy = 0
	DeploymentPolicy_DeploymentPolicy_Planned DeploymentPolicy = 1
)

var DeploymentPolicy_name = map[int32]string{
	0: "DeploymentPolicy_Active",
	1: "DeploymentPolicy_Planned",
}

var DeploymentPolicy_value = map[string]int32{
	"DeploymentPolicy_Active":  0,
	"DeploymentPolicy_Planned": 1,
}

func (x DeploymentPolicy) String() string {
	return proto.EnumName(DeploymentPolicy_name, int32(x))
}

func (DeploymentPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1dd1adc4aa8e1fa8, []int{0}
}

type DeletionPolicy int32

const (
	DeletionPolicy_DeletionPolicy_Orphan DeletionPolicy = 0
	DeletionPolicy_DeletionPolicy_Delete DeletionPolicy = 1
)

var DeletionPolicy_name = map[int32]string{
	0: "DeletionPolicy_Orphan",
	1: "DeletionPolicy_Delete",
}

var DeletionPolicy_value = map[string]int32{
	"DeletionPolicy_Orphan": 0,
	"DeletionPolicy_Delete": 1,
}

func (x DeletionPolicy) String() string {
	return proto.EnumName(DeletionPolicy_name, int32(x))
}

func (DeletionPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1dd1adc4aa8e1fa8, []int{1}
}

type Lifecycle struct {
	DeploymentPolicy     DeploymentPolicy `protobuf:"varint,1,opt,name=deploymentPolicy,proto3,enum=common.v1.DeploymentPolicy" json:"deploymentPolicy,omitempty"`
	DeletionPolicy       DeletionPolicy   `protobuf:"varint,2,opt,name=deletionPolicy,proto3,enum=common.v1.DeletionPolicy" json:"deletionPolicy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Lifecycle) Reset()         { *m = Lifecycle{} }
func (m *Lifecycle) String() string { return proto.CompactTextString(m) }
func (*Lifecycle) ProtoMessage()    {}
func (*Lifecycle) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1adc4aa8e1fa8, []int{0}
}
func (m *Lifecycle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lifecycle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lifecycle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lifecycle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lifecycle.Merge(m, src)
}
func (m *Lifecycle) XXX_Size() int {
	return m.Size()
}
func (m *Lifecycle) XXX_DiscardUnknown() {
	xxx_messageInfo_Lifecycle.DiscardUnknown(m)
}

var xxx_messageInfo_Lifecycle proto.InternalMessageInfo

func (m *Lifecycle) GetDeploymentPolicy() DeploymentPolicy {
	if m != nil {
		return m.DeploymentPolicy
	}
	return DeploymentPolicy_DeploymentPolicy_Active
}

func (m *Lifecycle) GetDeletionPolicy() DeletionPolicy {
	if m != nil {
		return m.DeletionPolicy
	}
	return DeletionPolicy_DeletionPolicy_Orphan
}

func init() {
	proto.RegisterEnum("common.v1.DeploymentPolicy", DeploymentPolicy_name, DeploymentPolicy_value)
	proto.RegisterEnum("common.v1.DeletionPolicy", DeletionPolicy_name, DeletionPolicy_value)
	proto.RegisterType((*Lifecycle)(nil), "common.v1.Lifecycle")
}

func init() { proto.RegisterFile("apis/common/v1/policies.proto", fileDescriptor_1dd1adc4aa8e1fa8) }

var fileDescriptor_1dd1adc4aa8e1fa8 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0xc8, 0x2c,
	0xd6, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2f, 0x33, 0xd4, 0x2f, 0xc8, 0xcf, 0xc9, 0x4c,
	0xce, 0x4c, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xc8, 0xe8, 0x95, 0x19,
	0x2a, 0x4d, 0x67, 0xe4, 0xe2, 0xf4, 0xc9, 0x4c, 0x4b, 0x4d, 0xae, 0x4c, 0xce, 0x49, 0x15, 0x72,
	0xe7, 0x12, 0x48, 0x49, 0x2d, 0xc8, 0xc9, 0xaf, 0xcc, 0x4d, 0xcd, 0x2b, 0x09, 0x00, 0x69, 0xaa,
	0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x92, 0xd6, 0x83, 0xeb, 0xd1, 0x73, 0x41, 0x53, 0x12,
	0x84, 0xa1, 0x49, 0xc8, 0x91, 0x8b, 0x2f, 0x25, 0x35, 0x27, 0xb5, 0x24, 0x33, 0x3f, 0x0f, 0x6a,
	0x0c, 0x13, 0xd8, 0x18, 0x49, 0x14, 0x63, 0x90, 0x15, 0x04, 0xa1, 0x69, 0xd0, 0xf2, 0xe5, 0x12,
	0x40, 0xb7, 0x48, 0x48, 0x9a, 0x4b, 0x1c, 0x5d, 0x2c, 0xde, 0x31, 0xb9, 0x24, 0xb3, 0x2c, 0x55,
	0x80, 0x41, 0x48, 0x86, 0x4b, 0x02, 0x43, 0x32, 0x20, 0x27, 0x31, 0x2f, 0x2f, 0x35, 0x45, 0x80,
	0x51, 0xcb, 0x8d, 0x8b, 0x0f, 0xd5, 0x42, 0x21, 0x49, 0x2e, 0x51, 0x54, 0x91, 0x78, 0xff, 0xa2,
	0x82, 0x8c, 0xc4, 0x3c, 0x01, 0x06, 0x2c, 0x52, 0x60, 0x6e, 0xaa, 0x00, 0xa3, 0x93, 0xcd, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43,
	0x94, 0x56, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x12, 0xc8, 0x67, 0xfa, 0x95, 0x79, 0x29, 0x29, 0xfa,
	0x79, 0x29, 0x29, 0xf1, 0x45, 0xa5, 0x79, 0x25, 0x99, 0xb9, 0xa9, 0xfa, 0xa8, 0x11, 0x91, 0xc4,
	0x06, 0x8e, 0x00, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0xe9, 0xc0, 0x6d, 0xa1, 0x01,
	0x00, 0x00,
}

func (m *Lifecycle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lifecycle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lifecycle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.DeletionPolicy != 0 {
		i = encodeVarintPolicies(dAtA, i, uint64(m.DeletionPolicy))
		i--
		dAtA[i] = 0x10
	}
	if m.DeploymentPolicy != 0 {
		i = encodeVarintPolicies(dAtA, i, uint64(m.DeploymentPolicy))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPolicies(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolicies(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lifecycle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeploymentPolicy != 0 {
		n += 1 + sovPolicies(uint64(m.DeploymentPolicy))
	}
	if m.DeletionPolicy != 0 {
		n += 1 + sovPolicies(uint64(m.DeletionPolicy))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPolicies(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolicies(x uint64) (n int) {
	return sovPolicies(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lifecycle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicies
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
			return fmt.Errorf("proto: Lifecycle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lifecycle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeploymentPolicy", wireType)
			}
			m.DeploymentPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicies
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeploymentPolicy |= DeploymentPolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletionPolicy", wireType)
			}
			m.DeletionPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicies
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeletionPolicy |= DeletionPolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPolicies(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicies
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
func skipPolicies(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolicies
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
					return 0, ErrIntOverflowPolicies
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
					return 0, ErrIntOverflowPolicies
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
				return 0, ErrInvalidLengthPolicies
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolicies
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolicies
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolicies        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolicies          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolicies = fmt.Errorf("proto: unexpected end of group")
)
