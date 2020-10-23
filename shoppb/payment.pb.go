// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shoppb/payment.proto

package shoppb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	basepb "github.com/shinmigo/pb/basepb"
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

type PaymentCode int32

const (
	PaymentCode_CodeNone PaymentCode = 0
	PaymentCode_Wechat   PaymentCode = 1
	PaymentCode_Alipay   PaymentCode = 2
)

var PaymentCode_name = map[int32]string{
	0: "CodeNone",
	1: "Wechat",
	2: "Alipay",
}

var PaymentCode_value = map[string]int32{
	"CodeNone": 0,
	"Wechat":   1,
	"Alipay":   2,
}

func (x PaymentCode) String() string {
	return proto.EnumName(PaymentCode_name, int32(x))
}

func (PaymentCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9ad8d900087ce41, []int{0}
}

type PaymentStatus int32

const (
	PaymentStatus_StatusNone PaymentStatus = 0
	PaymentStatus_Open       PaymentStatus = 1
	PaymentStatus_Close      PaymentStatus = 2
)

var PaymentStatus_name = map[int32]string{
	0: "StatusNone",
	1: "Open",
	2: "Close",
}

var PaymentStatus_value = map[string]int32{
	"StatusNone": 0,
	"Open":       1,
	"Close":      2,
}

func (x PaymentStatus) String() string {
	return proto.EnumName(PaymentStatus_name, int32(x))
}

func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9ad8d900087ce41, []int{1}
}

type ListPaymentReq struct {
	Status               PaymentStatus `protobuf:"varint,1,opt,name=status,proto3,enum=shoppb.PaymentStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListPaymentReq) Reset()         { *m = ListPaymentReq{} }
func (m *ListPaymentReq) String() string { return proto.CompactTextString(m) }
func (*ListPaymentReq) ProtoMessage()    {}
func (*ListPaymentReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ad8d900087ce41, []int{0}
}

func (m *ListPaymentReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPaymentReq.Unmarshal(m, b)
}
func (m *ListPaymentReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPaymentReq.Marshal(b, m, deterministic)
}
func (m *ListPaymentReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPaymentReq.Merge(m, src)
}
func (m *ListPaymentReq) XXX_Size() int {
	return xxx_messageInfo_ListPaymentReq.Size(m)
}
func (m *ListPaymentReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPaymentReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListPaymentReq proto.InternalMessageInfo

func (m *ListPaymentReq) GetStatus() PaymentStatus {
	if m != nil {
		return m.Status
	}
	return PaymentStatus_StatusNone
}

type Payment struct {
	Id                   uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 string        `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Params               string        `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
	Status               PaymentStatus `protobuf:"varint,5,opt,name=status,proto3,enum=shoppb.PaymentStatus" json:"status,omitempty"`
	Sort                 uint64        `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ad8d900087ce41, []int{1}
}

func (m *Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payment.Unmarshal(m, b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
}
func (m *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(m, src)
}
func (m *Payment) XXX_Size() int {
	return xxx_messageInfo_Payment.Size(m)
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Payment) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Payment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Payment) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func (m *Payment) GetStatus() PaymentStatus {
	if m != nil {
		return m.Status
	}
	return PaymentStatus_StatusNone
}

func (m *Payment) GetSort() uint64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

type ListPaymentRes struct {
	Payments             []*Payment `protobuf:"bytes,1,rep,name=payments,proto3" json:"payments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListPaymentRes) Reset()         { *m = ListPaymentRes{} }
func (m *ListPaymentRes) String() string { return proto.CompactTextString(m) }
func (*ListPaymentRes) ProtoMessage()    {}
func (*ListPaymentRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ad8d900087ce41, []int{2}
}

func (m *ListPaymentRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPaymentRes.Unmarshal(m, b)
}
func (m *ListPaymentRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPaymentRes.Marshal(b, m, deterministic)
}
func (m *ListPaymentRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPaymentRes.Merge(m, src)
}
func (m *ListPaymentRes) XXX_Size() int {
	return xxx_messageInfo_ListPaymentRes.Size(m)
}
func (m *ListPaymentRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPaymentRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListPaymentRes proto.InternalMessageInfo

func (m *ListPaymentRes) GetPayments() []*Payment {
	if m != nil {
		return m.Payments
	}
	return nil
}

type PaymentCodeReq struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentCodeReq) Reset()         { *m = PaymentCodeReq{} }
func (m *PaymentCodeReq) String() string { return proto.CompactTextString(m) }
func (*PaymentCodeReq) ProtoMessage()    {}
func (*PaymentCodeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ad8d900087ce41, []int{3}
}

func (m *PaymentCodeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentCodeReq.Unmarshal(m, b)
}
func (m *PaymentCodeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentCodeReq.Marshal(b, m, deterministic)
}
func (m *PaymentCodeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentCodeReq.Merge(m, src)
}
func (m *PaymentCodeReq) XXX_Size() int {
	return xxx_messageInfo_PaymentCodeReq.Size(m)
}
func (m *PaymentCodeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentCodeReq.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentCodeReq proto.InternalMessageInfo

func (m *PaymentCodeReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type DeletePaymentReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePaymentReq) Reset()         { *m = DeletePaymentReq{} }
func (m *DeletePaymentReq) String() string { return proto.CompactTextString(m) }
func (*DeletePaymentReq) ProtoMessage()    {}
func (*DeletePaymentReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ad8d900087ce41, []int{4}
}

func (m *DeletePaymentReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePaymentReq.Unmarshal(m, b)
}
func (m *DeletePaymentReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePaymentReq.Marshal(b, m, deterministic)
}
func (m *DeletePaymentReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePaymentReq.Merge(m, src)
}
func (m *DeletePaymentReq) XXX_Size() int {
	return xxx_messageInfo_DeletePaymentReq.Size(m)
}
func (m *DeletePaymentReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePaymentReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePaymentReq proto.InternalMessageInfo

func (m *DeletePaymentReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterEnum("shoppb.PaymentCode", PaymentCode_name, PaymentCode_value)
	proto.RegisterEnum("shoppb.PaymentStatus", PaymentStatus_name, PaymentStatus_value)
	proto.RegisterType((*ListPaymentReq)(nil), "shoppb.ListPaymentReq")
	proto.RegisterType((*Payment)(nil), "shoppb.Payment")
	proto.RegisterType((*ListPaymentRes)(nil), "shoppb.ListPaymentRes")
	proto.RegisterType((*PaymentCodeReq)(nil), "shoppb.PaymentCodeReq")
	proto.RegisterType((*DeletePaymentReq)(nil), "shoppb.DeletePaymentReq")
}

func init() { proto.RegisterFile("shoppb/payment.proto", fileDescriptor_a9ad8d900087ce41) }

var fileDescriptor_a9ad8d900087ce41 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0xc4, 0x6e, 0x6a, 0xd2, 0x17, 0x6a, 0xac, 0x07, 0x54, 0x56, 0x24, 0x44, 0x65, 0x71, 0x88,
	0x8a, 0x6a, 0x4b, 0x29, 0x17, 0x84, 0x10, 0x84, 0x16, 0x71, 0x41, 0x80, 0xdc, 0x03, 0x12, 0xb7,
	0x75, 0xfc, 0xd4, 0xac, 0x14, 0x7b, 0x17, 0xef, 0x16, 0x29, 0xff, 0xc2, 0x0f, 0xf0, 0x97, 0x68,
	0xbd, 0xeb, 0x26, 0x36, 0x3d, 0xe4, 0xb4, 0x2f, 0xb3, 0x33, 0x93, 0x37, 0x63, 0x1b, 0x9e, 0xaa,
	0x95, 0x90, 0xb2, 0xc8, 0x24, 0xdb, 0x54, 0x54, 0xeb, 0x54, 0x36, 0x42, 0x0b, 0x0c, 0x2c, 0x3a,
	0x7d, 0x52, 0x30, 0x45, 0xb2, 0xc8, 0xec, 0x61, 0x2f, 0x93, 0xf7, 0x10, 0x7e, 0xe1, 0x4a, 0x7f,
	0xb7, 0x8a, 0x9c, 0x7e, 0xe1, 0x39, 0x04, 0x4a, 0x33, 0x7d, 0xab, 0x62, 0xef, 0xd4, 0x9b, 0x85,
	0xf3, 0x67, 0xa9, 0xd5, 0xa7, 0x8e, 0x73, 0xdd, 0x5e, 0xe6, 0x8e, 0x94, 0xfc, 0xf1, 0xe0, 0xa1,
	0xbb, 0xc1, 0x10, 0x7c, 0x5e, 0xb6, 0xb2, 0x51, 0xee, 0xf3, 0x12, 0x11, 0x46, 0x4b, 0x51, 0x52,
	0xec, 0x9f, 0x7a, 0xb3, 0xa3, 0xbc, 0x9d, 0x0d, 0x56, 0xb3, 0x8a, 0xe2, 0x03, 0x8b, 0x99, 0x19,
	0x4f, 0x20, 0x90, 0xac, 0x61, 0x95, 0x8a, 0x47, 0x2d, 0xea, 0x7e, 0xed, 0xac, 0x72, 0xb8, 0xc7,
	0x2a, 0xc6, 0x5a, 0x89, 0x46, 0xc7, 0x41, 0xbb, 0x40, 0x3b, 0x27, 0xef, 0x06, 0xf9, 0x14, 0xbe,
	0x82, 0xb1, 0xeb, 0xc7, 0x24, 0x3c, 0x98, 0x4d, 0xe6, 0x8f, 0x07, 0xb6, 0xf9, 0x1d, 0x21, 0x79,
	0x09, 0xa1, 0x03, 0x2f, 0x45, 0x49, 0xa6, 0x9e, 0x2e, 0x93, 0xb7, 0xcd, 0x94, 0x24, 0x10, 0x5d,
	0xd1, 0x9a, 0x34, 0xed, 0xd4, 0x38, 0xe8, 0xe2, 0xec, 0x02, 0x26, 0x3b, 0x4e, 0xf8, 0x08, 0xc6,
	0xe6, 0xfc, 0x2a, 0x6a, 0x8a, 0x1e, 0x20, 0x40, 0xf0, 0x83, 0x96, 0x2b, 0xa6, 0x23, 0xcf, 0xcc,
	0x8b, 0x35, 0x97, 0x6c, 0x13, 0xf9, 0x67, 0xaf, 0xe1, 0xb8, 0x17, 0x15, 0x43, 0x00, 0x3b, 0x39,
	0xe1, 0x18, 0x46, 0xdf, 0x24, 0xd5, 0x91, 0x87, 0x47, 0x70, 0x78, 0xb9, 0x16, 0x8a, 0x22, 0x7f,
	0xfe, 0xd7, 0xbf, 0xdb, 0xfa, 0x9a, 0x9a, 0xdf, 0x7c, 0x49, 0xf8, 0x01, 0xc2, 0xcf, 0xd4, 0xb5,
	0x60, 0x0a, 0xc1, 0x93, 0x2e, 0x74, 0xff, 0xf1, 0x4f, 0xef, 0xc7, 0x15, 0xbe, 0x85, 0x68, 0xeb,
	0x70, 0x45, 0x9a, 0xf1, 0xf5, 0xd6, 0xa3, 0xdf, 0xd1, 0x74, 0x58, 0x28, 0xbe, 0x81, 0xe3, 0x5e,
	0x41, 0x18, 0x77, 0x8c, 0x61, 0x6f, 0xd3, 0x30, 0x75, 0xef, 0xe7, 0xa2, 0xde, 0x98, 0xff, 0x3d,
	0x07, 0x58, 0x94, 0x65, 0xa7, 0x1b, 0x3a, 0xff, 0x47, 0x4f, 0x61, 0xf2, 0xa9, 0xe4, 0x7a, 0x5f,
	0xfe, 0xc7, 0x17, 0x3f, 0x9f, 0xdf, 0x70, 0xbd, 0xba, 0x2d, 0xd2, 0xa5, 0xa8, 0x32, 0xb5, 0xe2,
	0x75, 0xc5, 0x6f, 0x44, 0x26, 0x8b, 0xcc, 0x0a, 0x8b, 0xa0, 0xfd, 0x4e, 0x2e, 0xfe, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x52, 0xbb, 0x04, 0xc5, 0x5c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentServiceClient interface {
	GetPaymentList(ctx context.Context, in *ListPaymentReq, opts ...grpc.CallOption) (*ListPaymentRes, error)
	GetPaymentDetail(ctx context.Context, in *PaymentCodeReq, opts ...grpc.CallOption) (*Payment, error)
	DeletePayment(ctx context.Context, in *DeletePaymentReq, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	AddPayment(ctx context.Context, in *Payment, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	EditPayment(ctx context.Context, in *Payment, opts ...grpc.CallOption) (*basepb.AnyRes, error)
}

type paymentServiceClient struct {
	cc *grpc.ClientConn
}

func NewPaymentServiceClient(cc *grpc.ClientConn) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) GetPaymentList(ctx context.Context, in *ListPaymentReq, opts ...grpc.CallOption) (*ListPaymentRes, error) {
	out := new(ListPaymentRes)
	err := c.cc.Invoke(ctx, "/shoppb.PaymentService/GetPaymentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPaymentDetail(ctx context.Context, in *PaymentCodeReq, opts ...grpc.CallOption) (*Payment, error) {
	out := new(Payment)
	err := c.cc.Invoke(ctx, "/shoppb.PaymentService/GetPaymentDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) DeletePayment(ctx context.Context, in *DeletePaymentReq, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/shoppb.PaymentService/DeletePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) AddPayment(ctx context.Context, in *Payment, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/shoppb.PaymentService/AddPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) EditPayment(ctx context.Context, in *Payment, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/shoppb.PaymentService/EditPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
type PaymentServiceServer interface {
	GetPaymentList(context.Context, *ListPaymentReq) (*ListPaymentRes, error)
	GetPaymentDetail(context.Context, *PaymentCodeReq) (*Payment, error)
	DeletePayment(context.Context, *DeletePaymentReq) (*basepb.AnyRes, error)
	AddPayment(context.Context, *Payment) (*basepb.AnyRes, error)
	EditPayment(context.Context, *Payment) (*basepb.AnyRes, error)
}

// UnimplementedPaymentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (*UnimplementedPaymentServiceServer) GetPaymentList(ctx context.Context, req *ListPaymentReq) (*ListPaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentList not implemented")
}
func (*UnimplementedPaymentServiceServer) GetPaymentDetail(ctx context.Context, req *PaymentCodeReq) (*Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentDetail not implemented")
}
func (*UnimplementedPaymentServiceServer) DeletePayment(ctx context.Context, req *DeletePaymentReq) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePayment not implemented")
}
func (*UnimplementedPaymentServiceServer) AddPayment(ctx context.Context, req *Payment) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPayment not implemented")
}
func (*UnimplementedPaymentServiceServer) EditPayment(ctx context.Context, req *Payment) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPayment not implemented")
}

func RegisterPaymentServiceServer(s *grpc.Server, srv PaymentServiceServer) {
	s.RegisterService(&_PaymentService_serviceDesc, srv)
}

func _PaymentService_GetPaymentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPaymentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.PaymentService/GetPaymentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPaymentList(ctx, req.(*ListPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPaymentDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPaymentDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.PaymentService/GetPaymentDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPaymentDetail(ctx, req.(*PaymentCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_DeletePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).DeletePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.PaymentService/DeletePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).DeletePayment(ctx, req.(*DeletePaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_AddPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).AddPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.PaymentService/AddPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).AddPayment(ctx, req.(*Payment))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_EditPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).EditPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.PaymentService/EditPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).EditPayment(ctx, req.(*Payment))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shoppb.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPaymentList",
			Handler:    _PaymentService_GetPaymentList_Handler,
		},
		{
			MethodName: "GetPaymentDetail",
			Handler:    _PaymentService_GetPaymentDetail_Handler,
		},
		{
			MethodName: "DeletePayment",
			Handler:    _PaymentService_DeletePayment_Handler,
		},
		{
			MethodName: "AddPayment",
			Handler:    _PaymentService_AddPayment_Handler,
		},
		{
			MethodName: "EditPayment",
			Handler:    _PaymentService_EditPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shoppb/payment.proto",
}
