// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/misc/proto/misc.proto

package miscpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VersionResponse struct {
	CommitHash           string           `protobuf:"bytes,1,opt,name=commit_hash,json=commitHash,proto3" json:"commit_hash,omitempty"`
	ShortHash            string           `protobuf:"bytes,2,opt,name=short_hash,json=shortHash,proto3" json:"short_hash,omitempty"`
	BuildDate            *types.Timestamp `protobuf:"bytes,3,opt,name=build_date,json=buildDate" json:"build_date,omitempty"`
	CommitDate           *types.Timestamp `protobuf:"bytes,4,opt,name=commit_date,json=commitDate" json:"commit_date,omitempty"`
	Count                int64            `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_misc_e882b0687e2fadfd, []int{0}
}
func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(dst, src)
}
func (m *VersionResponse) XXX_Size() int {
	return m.Size()
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetCommitHash() string {
	if m != nil {
		return m.CommitHash
	}
	return ""
}

func (m *VersionResponse) GetShortHash() string {
	if m != nil {
		return m.ShortHash
	}
	return ""
}

func (m *VersionResponse) GetBuildDate() *types.Timestamp {
	if m != nil {
		return m.BuildDate
	}
	return nil
}

func (m *VersionResponse) GetCommitDate() *types.Timestamp {
	if m != nil {
		return m.CommitDate
	}
	return nil
}

func (m *VersionResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_misc_e882b0687e2fadfd, []int{1}
}
func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(dst, src)
}
func (m *VersionRequest) XXX_Size() int {
	return m.Size()
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VersionResponse)(nil), "misc.VersionResponse")
	proto.RegisterType((*VersionRequest)(nil), "misc.VersionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MiscSystem service

type MiscSystemClient interface {
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type miscSystemClient struct {
	cc *grpc.ClientConn
}

func NewMiscSystemClient(cc *grpc.ClientConn) MiscSystemClient {
	return &miscSystemClient{cc}
}

func (c *miscSystemClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/misc.MiscSystem/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MiscSystem service

type MiscSystemServer interface {
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

func RegisterMiscSystemServer(s *grpc.Server, srv MiscSystemServer) {
	s.RegisterService(&_MiscSystem_serviceDesc, srv)
}

func _MiscSystem_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiscSystemServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/misc.MiscSystem/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiscSystemServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MiscSystem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "misc.MiscSystem",
	HandlerType: (*MiscSystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _MiscSystem_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/misc/proto/misc.proto",
}

func (m *VersionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CommitHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMisc(dAtA, i, uint64(len(m.CommitHash)))
		i += copy(dAtA[i:], m.CommitHash)
	}
	if len(m.ShortHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMisc(dAtA, i, uint64(len(m.ShortHash)))
		i += copy(dAtA[i:], m.ShortHash)
	}
	if m.BuildDate != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMisc(dAtA, i, uint64(m.BuildDate.Size()))
		n1, err := m.BuildDate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.CommitDate != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMisc(dAtA, i, uint64(m.CommitDate.Size()))
		n2, err := m.CommitDate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Count != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMisc(dAtA, i, uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *VersionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMisc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VersionResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.CommitHash)
	if l > 0 {
		n += 1 + l + sovMisc(uint64(l))
	}
	l = len(m.ShortHash)
	if l > 0 {
		n += 1 + l + sovMisc(uint64(l))
	}
	if m.BuildDate != nil {
		l = m.BuildDate.Size()
		n += 1 + l + sovMisc(uint64(l))
	}
	if m.CommitDate != nil {
		l = m.CommitDate.Size()
		n += 1 + l + sovMisc(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovMisc(uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VersionRequest) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMisc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMisc(x uint64) (n int) {
	return sovMisc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VersionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMisc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMisc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMisc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMisc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMisc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShortHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMisc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMisc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BuildDate == nil {
				m.BuildDate = &types.Timestamp{}
			}
			if err := m.BuildDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMisc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMisc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommitDate == nil {
				m.CommitDate = &types.Timestamp{}
			}
			if err := m.CommitDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMisc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMisc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMisc
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
func (m *VersionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMisc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VersionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMisc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMisc
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
func skipMisc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMisc
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
					return 0, ErrIntOverflowMisc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMisc
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMisc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMisc
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMisc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMisc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMisc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("modules/misc/proto/misc.proto", fileDescriptor_misc_e882b0687e2fadfd) }

var fileDescriptor_misc_e882b0687e2fadfd = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xfe, 0x65, 0xff, 0x7e, 0x36, 0x03, 0x1d, 0x61, 0x42, 0x29, 0xae, 0x1b, 0x3b, 0xed, 0x62,
	0x83, 0xf3, 0x24, 0xde, 0x86, 0x07, 0x2f, 0xc2, 0xa8, 0xe2, 0xc1, 0x83, 0x23, 0xed, 0x62, 0x1b,
	0x58, 0xfa, 0xd6, 0x25, 0x1d, 0x78, 0xf5, 0x2b, 0x78, 0xf1, 0x23, 0x79, 0x14, 0xfc, 0x00, 0xca,
	0xf4, 0x83, 0x48, 0x93, 0x6e, 0xa0, 0x17, 0x6f, 0x4f, 0x9e, 0x3f, 0xc9, 0xf3, 0xe6, 0xc5, 0x3d,
	0x09, 0xf3, 0x62, 0xc1, 0x15, 0x95, 0x42, 0xc5, 0x34, 0x5f, 0x82, 0x06, 0x03, 0x03, 0x03, 0x49,
	0xa3, 0xc4, 0xde, 0x61, 0x22, 0x74, 0x5a, 0x44, 0x41, 0x0c, 0x92, 0x26, 0x90, 0x80, 0xf5, 0x45,
	0xc5, 0x9d, 0x39, 0xd9, 0x50, 0x89, 0x6c, 0xc8, 0x3b, 0x48, 0x00, 0x92, 0x05, 0xa7, 0x2c, 0x17,
	0x94, 0x65, 0x19, 0x68, 0xa6, 0x05, 0x64, 0xaa, 0x52, 0xfb, 0x95, 0xba, 0xbd, 0x43, 0x0b, 0xc9,
	0x95, 0x66, 0x32, 0xb7, 0x86, 0xe1, 0x3b, 0xc2, 0x7b, 0xd7, 0x7c, 0xa9, 0x04, 0x64, 0x21, 0x57,
	0x39, 0x64, 0x8a, 0x93, 0x3e, 0x6e, 0xc7, 0x20, 0xa5, 0xd0, 0xb3, 0x94, 0xa9, 0xd4, 0x45, 0x03,
	0x34, 0x72, 0x42, 0x6c, 0xa9, 0x73, 0xa6, 0x52, 0xd2, 0xc3, 0x58, 0xa5, 0xb0, 0xac, 0xf4, 0x9a,
	0xd1, 0x1d, 0xc3, 0x18, 0xf9, 0x04, 0xe3, 0xa8, 0x10, 0x8b, 0xf9, 0x6c, 0xce, 0x34, 0x77, 0xeb,
	0x03, 0x34, 0x6a, 0x8f, 0xbd, 0xc0, 0x36, 0x09, 0x36, 0x4d, 0x82, 0xab, 0x4d, 0x93, 0xd0, 0x31,
	0xee, 0x33, 0xa6, 0x39, 0x39, 0xdd, 0x3e, 0x6d, 0xb2, 0x8d, 0x3f, 0xb3, 0x55, 0x2d, 0x13, 0xee,
	0xe2, 0x66, 0x0c, 0x45, 0xa6, 0xdd, 0xe6, 0x00, 0x8d, 0xea, 0xa1, 0x3d, 0x0c, 0x3b, 0x78, 0x77,
	0x3b, 0xe0, 0x7d, 0xc1, 0x95, 0x1e, 0xdf, 0x62, 0x7c, 0x21, 0x54, 0x7c, 0xf9, 0xa0, 0x34, 0x97,
	0x64, 0x8a, 0xff, 0x57, 0x3a, 0xe9, 0x06, 0x66, 0x1b, 0x3f, 0xed, 0xde, 0xfe, 0x2f, 0xd6, 0xfe,
	0xd2, 0xd0, 0x7d, 0x7c, 0xfb, 0x7a, 0xaa, 0x11, 0xd2, 0xa1, 0xab, 0x23, 0xbb, 0xd0, 0x95, 0x75,
	0x4c, 0xe8, 0xcb, 0xda, 0x47, 0xaf, 0x6b, 0x1f, 0x7d, 0xac, 0x7d, 0xf4, 0xfc, 0xe9, 0xff, 0xc3,
	0x3b, 0x31, 0x48, 0x73, 0xcb, 0xc4, 0x29, 0x5f, 0x9e, 0x96, 0x73, 0x4c, 0xd1, 0x4d, 0xab, 0xa4,
	0xf2, 0x28, 0x6a, 0x99, 0xc1, 0x8e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x82, 0xaa, 0x5f, 0x89,
	0x20, 0x02, 0x00, 0x00,
}
