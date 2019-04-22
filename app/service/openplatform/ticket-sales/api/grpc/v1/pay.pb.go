// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/openplatform/ticket-sales/api/grpc/v1/pay.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PayNotifyRequest struct {
	MsgID      string `protobuf:"bytes,1,opt,name=msgID,proto3" json:"msgID,omitempty"`
	MsgContent string `protobuf:"bytes,2,opt,name=msgContent,proto3" json:"msgContent,omitempty"`
	TestMode   bool   `protobuf:"varint,3,opt,name=testMode,proto3" json:"testMode,omitempty"`
}

func (m *PayNotifyRequest) Reset()                    { *m = PayNotifyRequest{} }
func (*PayNotifyRequest) ProtoMessage()               {}
func (*PayNotifyRequest) Descriptor() ([]byte, []int) { return fileDescriptorPay, []int{0} }

type PayNotifyResponse struct {
}

func (m *PayNotifyResponse) Reset()                    { *m = PayNotifyResponse{} }
func (*PayNotifyResponse) ProtoMessage()               {}
func (*PayNotifyResponse) Descriptor() ([]byte, []int) { return fileDescriptorPay, []int{1} }

func init() {
	proto.RegisterType((*PayNotifyRequest)(nil), "ticket.service.sales.v1.PayNotifyRequest")
	proto.RegisterType((*PayNotifyResponse)(nil), "ticket.service.sales.v1.PayNotifyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pay service

type PayClient interface {
	PayNotify(ctx context.Context, in *PayNotifyRequest, opts ...grpc.CallOption) (*PayNotifyResponse, error)
}

type payClient struct {
	cc *grpc.ClientConn
}

func NewPayClient(cc *grpc.ClientConn) PayClient {
	return &payClient{cc}
}

func (c *payClient) PayNotify(ctx context.Context, in *PayNotifyRequest, opts ...grpc.CallOption) (*PayNotifyResponse, error) {
	out := new(PayNotifyResponse)
	err := grpc.Invoke(ctx, "/ticket.service.sales.v1.Pay/PayNotify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pay service

type PayServer interface {
	PayNotify(context.Context, *PayNotifyRequest) (*PayNotifyResponse, error)
}

func RegisterPayServer(s *grpc.Server, srv PayServer) {
	s.RegisterService(&_Pay_serviceDesc, srv)
}

func _Pay_PayNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).PayNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket.service.sales.v1.Pay/PayNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).PayNotify(ctx, req.(*PayNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ticket.service.sales.v1.Pay",
	HandlerType: (*PayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayNotify",
			Handler:    _Pay_PayNotify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/openplatform/ticket-sales/api/grpc/v1/pay.proto",
}

func (m *PayNotifyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PayNotifyRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MsgID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPay(dAtA, i, uint64(len(m.MsgID)))
		i += copy(dAtA[i:], m.MsgID)
	}
	if len(m.MsgContent) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPay(dAtA, i, uint64(len(m.MsgContent)))
		i += copy(dAtA[i:], m.MsgContent)
	}
	if m.TestMode {
		dAtA[i] = 0x18
		i++
		if m.TestMode {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *PayNotifyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PayNotifyResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintPay(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PayNotifyRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.MsgID)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	l = len(m.MsgContent)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	if m.TestMode {
		n += 2
	}
	return n
}

func (m *PayNotifyResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovPay(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPay(x uint64) (n int) {
	return sovPay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PayNotifyRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PayNotifyRequest{`,
		`MsgID:` + fmt.Sprintf("%v", this.MsgID) + `,`,
		`MsgContent:` + fmt.Sprintf("%v", this.MsgContent) + `,`,
		`TestMode:` + fmt.Sprintf("%v", this.TestMode) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PayNotifyResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PayNotifyResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringPay(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PayNotifyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPay
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
			return fmt.Errorf("proto: PayNotifyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PayNotifyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
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
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgContent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
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
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgContent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestMode", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TestMode = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPay
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
func (m *PayNotifyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPay
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
			return fmt.Errorf("proto: PayNotifyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PayNotifyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPay
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
func skipPay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPay
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
					return 0, ErrIntOverflowPay
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
					return 0, ErrIntOverflowPay
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
				return 0, ErrInvalidLengthPay
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPay
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
				next, err := skipPay(dAtA[start:])
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
	ErrInvalidLengthPay = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPay   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("app/service/openplatform/ticket-sales/api/grpc/v1/pay.proto", fileDescriptorPay)
}

var fileDescriptorPay = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xeb, 0x56, 0x45, 0xad, 0x27, 0x30, 0x03, 0x55, 0x84, 0xdc, 0xaa, 0x53, 0x41, 0xaa,
	0xad, 0xc0, 0xc8, 0x56, 0x58, 0x18, 0x8a, 0xaa, 0x8c, 0x6c, 0x4e, 0x70, 0x8c, 0x45, 0x93, 0x33,
	0xb1, 0x13, 0x29, 0x1b, 0x3f, 0xaf, 0x23, 0x23, 0x13, 0xa2, 0xf9, 0x25, 0xa8, 0x0e, 0xaa, 0x2a,
	0x24, 0x24, 0x36, 0xbf, 0xe7, 0xef, 0xee, 0xdd, 0x1d, 0xbe, 0x11, 0xc6, 0x70, 0x2b, 0x8b, 0x4a,
	0x27, 0x92, 0x83, 0x91, 0xb9, 0x59, 0x0b, 0x97, 0x42, 0x91, 0x71, 0xa7, 0x93, 0x17, 0xe9, 0xe6,
	0x56, 0xac, 0xa5, 0xe5, 0xc2, 0x68, 0xae, 0x0a, 0x93, 0xf0, 0x2a, 0xe4, 0x46, 0xd4, 0xcc, 0x14,
	0xe0, 0x80, 0x9c, 0xb5, 0x0c, 0xfb, 0xa9, 0x67, 0x9e, 0x65, 0x55, 0x18, 0xcc, 0x95, 0x76, 0xcf,
	0x65, 0xcc, 0x12, 0xc8, 0xb8, 0x02, 0x05, 0xdc, 0xf3, 0x71, 0x99, 0x7a, 0xe5, 0x85, 0x7f, 0xb5,
	0x7d, 0xa6, 0x80, 0x8f, 0x57, 0xa2, 0x7e, 0x00, 0xa7, 0xd3, 0x3a, 0x92, 0xaf, 0xa5, 0xb4, 0x8e,
	0x8c, 0x71, 0x3f, 0xb3, 0xea, 0xfe, 0x6e, 0x84, 0x26, 0x68, 0x36, 0x5c, 0x0c, 0x9b, 0xcf, 0x71,
	0x7f, 0xb9, 0x33, 0xa2, 0xd6, 0x27, 0x14, 0xe3, 0xcc, 0xaa, 0x5b, 0xc8, 0x9d, 0xcc, 0xdd, 0xa8,
	0xbb, 0xa3, 0xa2, 0x03, 0x87, 0x04, 0x78, 0xe0, 0xa4, 0x75, 0x4b, 0x78, 0x92, 0xa3, 0xde, 0x04,
	0xcd, 0x06, 0xd1, 0x5e, 0x4f, 0x4f, 0xf1, 0xc9, 0x41, 0xa0, 0x35, 0x90, 0x5b, 0x79, 0xa5, 0x71,
	0x6f, 0x25, 0x6a, 0x12, 0xe3, 0xe1, 0xfe, 0x8f, 0x5c, 0xb0, 0x3f, 0x56, 0x64, 0xbf, 0x07, 0x0e,
	0x2e, 0xff, 0x83, 0xb6, 0x51, 0x8b, 0xf3, 0xcd, 0x96, 0x76, 0x3e, 0xb6, 0xb4, 0xf3, 0xd6, 0x50,
	0xb4, 0x69, 0x28, 0x7a, 0x6f, 0x28, 0xfa, 0x6a, 0x28, 0x7a, 0xec, 0x56, 0x61, 0x7c, 0xe4, 0xaf,
	0x72, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x78, 0x04, 0x37, 0x9c, 0x01, 0x00, 0x00,
}