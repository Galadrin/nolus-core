// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nolus/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Minter represents the minting state.
type Minter struct {
	NormTimePassed     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=norm_time_passed,json=normTimePassed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"norm_time_passed"`
	TotalMinted        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=total_minted,json=totalMinted,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_minted"`
	PrevBlockTimestamp int64                                  `protobuf:"varint,4,opt,name=prev_block_timestamp,json=prevBlockTimestamp,proto3" json:"prev_block_timestamp,omitempty"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9c8d0486b75e8ca, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetPrevBlockTimestamp() int64 {
	if m != nil {
		return m.PrevBlockTimestamp
	}
	return 0
}

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom              string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	MaxMintableNanoseconds int64  `protobuf:"varint,2,opt,name=max_mintable_nanoseconds,json=maxMintableNanoseconds,proto3" json:"max_mintable_nanoseconds,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9c8d0486b75e8ca, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetMaxMintableNanoseconds() int64 {
	if m != nil {
		return m.MaxMintableNanoseconds
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "nolus.mint.v1beta1.Minter")
	proto.RegisterType((*Params)(nil), "nolus.mint.v1beta1.Params")
}

func init() { proto.RegisterFile("nolus/mint/v1beta1/mint.proto", fileDescriptor_e9c8d0486b75e8ca) }

var fileDescriptor_e9c8d0486b75e8ca = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x93, 0xb7, 0x2f, 0x85, 0x9e, 0x22, 0x72, 0x14, 0x09, 0x42, 0xd3, 0xd2, 0x41, 0xba,
	0x34, 0xb1, 0x88, 0xe0, 0x5c, 0xba, 0x38, 0x54, 0x6b, 0xe8, 0x20, 0x2e, 0xe1, 0x92, 0x3c, 0xc4,
	0xa3, 0xb9, 0x7b, 0x42, 0xee, 0x5a, 0xaa, 0x9f, 0xc2, 0x8f, 0xd5, 0xb1, 0xa3, 0x38, 0x14, 0x69,
	0x07, 0xbf, 0x86, 0xdc, 0xb5, 0xe2, 0xec, 0x94, 0xdc, 0xff, 0x77, 0xfc, 0xf9, 0x3d, 0xf7, 0x90,
	0x96, 0xc4, 0x62, 0xae, 0x42, 0xc1, 0xa5, 0x0e, 0x17, 0x83, 0x04, 0x34, 0x1b, 0xd8, 0x43, 0x50,
	0x56, 0xa8, 0x91, 0x52, 0x8b, 0x03, 0x9b, 0x1c, 0xf0, 0x79, 0x33, 0xc7, 0x1c, 0x2d, 0x0e, 0xcd,
	0xdf, 0xfe, 0x66, 0xf7, 0xcb, 0x25, 0xf5, 0x31, 0x97, 0x1a, 0x2a, 0xfa, 0x48, 0x4e, 0x25, 0x56,
	0x22, 0xd6, 0x5c, 0x40, 0x5c, 0x32, 0xa5, 0x20, 0xf3, 0xfe, 0x75, 0xdc, 0x5e, 0x63, 0x18, 0xac,
	0x36, 0x6d, 0xe7, 0x63, 0xd3, 0xbe, 0xc8, 0xb9, 0x7e, 0x9e, 0x27, 0x41, 0x8a, 0x22, 0x4c, 0x51,
	0x09, 0x54, 0x87, 0x4f, 0x5f, 0x65, 0xb3, 0x50, 0xbf, 0x94, 0xa0, 0x82, 0x11, 0xa4, 0xd1, 0x89,
	0xe9, 0x99, 0x72, 0x01, 0x13, 0xdb, 0x42, 0x1f, 0xc8, 0xb1, 0x46, 0xcd, 0x8a, 0xd8, 0x08, 0x41,
	0xe6, 0xd5, 0xfe, 0xdc, 0x7a, 0x2b, 0x75, 0x74, 0x64, 0x3b, 0xac, 0x6c, 0x46, 0x2f, 0x49, 0xb3,
	0xac, 0x60, 0x11, 0x27, 0x05, 0xa6, 0x33, 0xab, 0xac, 0x34, 0x13, 0xa5, 0xf7, 0xbf, 0xe3, 0xf6,
	0x6a, 0x11, 0x35, 0x6c, 0x68, 0xd0, 0xf4, 0x87, 0x74, 0x19, 0xa9, 0x4f, 0x58, 0xc5, 0x84, 0xa2,
	0x2d, 0x42, 0x8c, 0x48, 0x9c, 0x81, 0x44, 0xe1, 0xb9, 0x46, 0x26, 0x6a, 0x98, 0x64, 0x64, 0x02,
	0x7a, 0x43, 0x3c, 0xc1, 0x96, 0xd6, 0x95, 0x25, 0x05, 0xc4, 0x92, 0x49, 0x54, 0x90, 0xa2, 0xcc,
	0x94, 0x7d, 0x8f, 0x5a, 0x74, 0x26, 0xd8, 0x72, 0x7c, 0xc0, 0x77, 0xbf, 0x74, 0x78, 0xbf, 0xda,
	0xfa, 0xee, 0x7a, 0xeb, 0xbb, 0x9f, 0x5b, 0xdf, 0x7d, 0xdb, 0xf9, 0xce, 0x7a, 0xe7, 0x3b, 0xef,
	0x3b, 0xdf, 0x79, 0xba, 0xce, 0xb9, 0x2e, 0x58, 0xd2, 0x97, 0x28, 0x30, 0x48, 0x2b, 0xc8, 0xb8,
	0x52, 0x5c, 0x60, 0x20, 0x41, 0x87, 0x26, 0xb3, 0x13, 0xbf, 0xa2, 0x84, 0x70, 0xb9, 0x5f, 0xaa,
	0x1d, 0x3b, 0xa9, 0xdb, 0x25, 0x5d, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x19, 0x8a, 0x33, 0x78,
	0xef, 0x01, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PrevBlockTimestamp != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.PrevBlockTimestamp))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.TotalMinted.Size()
		i -= size
		if _, err := m.TotalMinted.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.NormTimePassed.Size()
		i -= size
		if _, err := m.NormTimePassed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxMintableNanoseconds != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.MaxMintableNanoseconds))
		i--
		dAtA[i] = 0x10
	}
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NormTimePassed.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.TotalMinted.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.PrevBlockTimestamp != 0 {
		n += 1 + sovMint(uint64(m.PrevBlockTimestamp))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.MaxMintableNanoseconds != 0 {
		n += 1 + sovMint(uint64(m.MaxMintableNanoseconds))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NormTimePassed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NormTimePassed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalMinted", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalMinted.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevBlockTimestamp", wireType)
			}
			m.PrevBlockTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrevBlockTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMintableNanoseconds", wireType)
			}
			m.MaxMintableNanoseconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMintableNanoseconds |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
