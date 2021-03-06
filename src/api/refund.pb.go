// Code generated by protoc-gen-go. DO NOT EDIT.
// source: refund.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type RefundRequest struct {
	Contract             string   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Transaction          string   `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefundRequest) Reset()         { *m = RefundRequest{} }
func (m *RefundRequest) String() string { return proto.CompactTextString(m) }
func (*RefundRequest) ProtoMessage()    {}
func (*RefundRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27efa219a51643e4, []int{0}
}

func (m *RefundRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefundRequest.Unmarshal(m, b)
}
func (m *RefundRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefundRequest.Marshal(b, m, deterministic)
}
func (m *RefundRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefundRequest.Merge(m, src)
}
func (m *RefundRequest) XXX_Size() int {
	return xxx_messageInfo_RefundRequest.Size(m)
}
func (m *RefundRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefundRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefundRequest proto.InternalMessageInfo

func (m *RefundRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *RefundRequest) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

func (m *RefundRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type RefundResponse struct {
	RefundFee            string   `protobuf:"bytes,1,opt,name=refundFee,proto3" json:"refundFee,omitempty"`
	RefundTransaction    string   `protobuf:"bytes,2,opt,name=refundTransaction,proto3" json:"refundTransaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefundResponse) Reset()         { *m = RefundResponse{} }
func (m *RefundResponse) String() string { return proto.CompactTextString(m) }
func (*RefundResponse) ProtoMessage()    {}
func (*RefundResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27efa219a51643e4, []int{1}
}

func (m *RefundResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefundResponse.Unmarshal(m, b)
}
func (m *RefundResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefundResponse.Marshal(b, m, deterministic)
}
func (m *RefundResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefundResponse.Merge(m, src)
}
func (m *RefundResponse) XXX_Size() int {
	return xxx_messageInfo_RefundResponse.Size(m)
}
func (m *RefundResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefundResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefundResponse proto.InternalMessageInfo

func (m *RefundResponse) GetRefundFee() string {
	if m != nil {
		return m.RefundFee
	}
	return ""
}

func (m *RefundResponse) GetRefundTransaction() string {
	if m != nil {
		return m.RefundTransaction
	}
	return ""
}

func init() {
	proto.RegisterType((*RefundRequest)(nil), "api.RefundRequest")
	proto.RegisterType((*RefundResponse)(nil), "api.RefundResponse")
}

func init() { proto.RegisterFile("refund.proto", fileDescriptor_27efa219a51643e4) }

var fileDescriptor_27efa219a51643e4 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x4d, 0x2b,
	0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x49,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49,
	0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x51, 0xca, 0xe4, 0xe2, 0x0d, 0x02, 0x6b, 0x09, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x48, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c,
	0x2e, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x14, 0xb8, 0xb8, 0x4b, 0x8a,
	0x12, 0xf3, 0x8a, 0x13, 0x93, 0x41, 0x46, 0x48, 0x30, 0x81, 0xa5, 0x91, 0x85, 0xc0, 0xba, 0x4b,
	0x8b, 0x8a, 0x52, 0xf3, 0x92, 0x2b, 0x25, 0x98, 0xa1, 0xba, 0xa1, 0x7c, 0xa5, 0x18, 0x2e, 0x3e,
	0x98, 0x55, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x32, 0x5c, 0x9c, 0x10, 0xf7, 0xba, 0xa5,
	0xa6, 0x42, 0x2d, 0x43, 0x08, 0x08, 0xe9, 0x70, 0x09, 0x42, 0x38, 0x21, 0x18, 0x76, 0x62, 0x4a,
	0x18, 0x05, 0x71, 0xb1, 0x41, 0x4c, 0x17, 0xf2, 0xe0, 0x62, 0x83, 0x48, 0x0b, 0x09, 0xe9, 0x25,
	0x16, 0x64, 0xea, 0xa1, 0xf8, 0x4f, 0x4a, 0x18, 0x45, 0x0c, 0xe2, 0x10, 0x25, 0xd1, 0xa6, 0xcb,
	0x4f, 0x26, 0x33, 0xf1, 0x2b, 0x71, 0xe9, 0x97, 0x19, 0xea, 0x43, 0x34, 0x5b, 0x31, 0x6a, 0x25,
	0xb1, 0x81, 0xc3, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x5c, 0x25, 0xf5, 0x56, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RefundClient is the client API for Refund service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RefundClient interface {
	Refund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundResponse, error)
}

type refundClient struct {
	cc *grpc.ClientConn
}

func NewRefundClient(cc *grpc.ClientConn) RefundClient {
	return &refundClient{cc}
}

func (c *refundClient) Refund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundResponse, error) {
	out := new(RefundResponse)
	err := c.cc.Invoke(ctx, "/api.Refund/refund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RefundServer is the server API for Refund service.
type RefundServer interface {
	Refund(context.Context, *RefundRequest) (*RefundResponse, error)
}

// UnimplementedRefundServer can be embedded to have forward compatible implementations.
type UnimplementedRefundServer struct {
}

func (*UnimplementedRefundServer) Refund(ctx context.Context, req *RefundRequest) (*RefundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refund not implemented")
}

func RegisterRefundServer(s *grpc.Server, srv RefundServer) {
	s.RegisterService(&_Refund_serviceDesc, srv)
}

func _Refund_Refund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefundServer).Refund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Refund/Refund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefundServer).Refund(ctx, req.(*RefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Refund_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Refund",
	HandlerType: (*RefundServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "refund",
			Handler:    _Refund_Refund_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "refund.proto",
}
