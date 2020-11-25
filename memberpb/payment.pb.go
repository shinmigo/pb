// Code generated by protoc-gen-go. DO NOT EDIT.
// source: memberpb/payment.proto

package memberpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
	return fileDescriptor_b614f87eb7e5cee8, []int{0}
}

type PaymentType int32

const (
	PaymentType_TypeNone PaymentType = 0
	PaymentType_Order    PaymentType = 1
	PaymentType_Recharge PaymentType = 2
)

var PaymentType_name = map[int32]string{
	0: "TypeNone",
	1: "Order",
	2: "Recharge",
}

var PaymentType_value = map[string]int32{
	"TypeNone": 0,
	"Order":    1,
	"Recharge": 2,
}

func (x PaymentType) String() string {
	return proto.EnumName(PaymentType_name, int32(x))
}

func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b614f87eb7e5cee8, []int{1}
}

type PaymentStatus int32

const (
	PaymentStatus_StatusNone  PaymentStatus = 0
	PaymentStatus_Unpaid      PaymentStatus = 1
	PaymentStatus_PaySuccess  PaymentStatus = 2
	PaymentStatus_StatusOther PaymentStatus = 3
)

var PaymentStatus_name = map[int32]string{
	0: "StatusNone",
	1: "Unpaid",
	2: "PaySuccess",
	3: "StatusOther",
}

var PaymentStatus_value = map[string]int32{
	"StatusNone":  0,
	"Unpaid":      1,
	"PaySuccess":  2,
	"StatusOther": 3,
}

func (x PaymentStatus) String() string {
	return proto.EnumName(PaymentStatus_name, int32(x))
}

func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b614f87eb7e5cee8, []int{2}
}

type PaymentIdReq struct {
	PaymentId            string   `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentIdReq) Reset()         { *m = PaymentIdReq{} }
func (m *PaymentIdReq) String() string { return proto.CompactTextString(m) }
func (*PaymentIdReq) ProtoMessage()    {}
func (*PaymentIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b614f87eb7e5cee8, []int{0}
}

func (m *PaymentIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentIdReq.Unmarshal(m, b)
}
func (m *PaymentIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentIdReq.Marshal(b, m, deterministic)
}
func (m *PaymentIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentIdReq.Merge(m, src)
}
func (m *PaymentIdReq) XXX_Size() int {
	return xxx_messageInfo_PaymentIdReq.Size(m)
}
func (m *PaymentIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentIdReq proto.InternalMessageInfo

func (m *PaymentIdReq) GetPaymentId() string {
	if m != nil {
		return m.PaymentId
	}
	return ""
}

type PaymentParams struct {
	SourceId             string   `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id"`
	Money                float64  `protobuf:"fixed64,2,opt,name=money,proto3" json:"money"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentParams) Reset()         { *m = PaymentParams{} }
func (m *PaymentParams) String() string { return proto.CompactTextString(m) }
func (*PaymentParams) ProtoMessage()    {}
func (*PaymentParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_b614f87eb7e5cee8, []int{1}
}

func (m *PaymentParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentParams.Unmarshal(m, b)
}
func (m *PaymentParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentParams.Marshal(b, m, deterministic)
}
func (m *PaymentParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentParams.Merge(m, src)
}
func (m *PaymentParams) XXX_Size() int {
	return xxx_messageInfo_PaymentParams.Size(m)
}
func (m *PaymentParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentParams.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentParams proto.InternalMessageInfo

func (m *PaymentParams) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *PaymentParams) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

type PaymentRelList struct {
	List                 []*PaymentParams `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PaymentRelList) Reset()         { *m = PaymentRelList{} }
func (m *PaymentRelList) String() string { return proto.CompactTextString(m) }
func (*PaymentRelList) ProtoMessage()    {}
func (*PaymentRelList) Descriptor() ([]byte, []int) {
	return fileDescriptor_b614f87eb7e5cee8, []int{2}
}

func (m *PaymentRelList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentRelList.Unmarshal(m, b)
}
func (m *PaymentRelList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentRelList.Marshal(b, m, deterministic)
}
func (m *PaymentRelList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentRelList.Merge(m, src)
}
func (m *PaymentRelList) XXX_Size() int {
	return xxx_messageInfo_PaymentRelList.Size(m)
}
func (m *PaymentRelList) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentRelList.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentRelList proto.InternalMessageInfo

func (m *PaymentRelList) GetList() []*PaymentParams {
	if m != nil {
		return m.List
	}
	return nil
}

type ToAdd struct {
	MemberId             uint64           `protobuf:"varint,1,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	Type                 PaymentType      `protobuf:"varint,2,opt,name=type,proto3,enum=memberpb.PaymentType" json:"type"`
	PaymentCode          PaymentCode      `protobuf:"varint,3,opt,name=payment_code,json=paymentCode,proto3,enum=memberpb.PaymentCode" json:"payment_code"`
	Ip                   string           `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip"`
	Params               []*PaymentParams `protobuf:"bytes,5,rep,name=params,proto3" json:"params"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ToAdd) Reset()         { *m = ToAdd{} }
func (m *ToAdd) String() string { return proto.CompactTextString(m) }
func (*ToAdd) ProtoMessage()    {}
func (*ToAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_b614f87eb7e5cee8, []int{3}
}

func (m *ToAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToAdd.Unmarshal(m, b)
}
func (m *ToAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToAdd.Marshal(b, m, deterministic)
}
func (m *ToAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToAdd.Merge(m, src)
}
func (m *ToAdd) XXX_Size() int {
	return xxx_messageInfo_ToAdd.Size(m)
}
func (m *ToAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ToAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ToAdd proto.InternalMessageInfo

func (m *ToAdd) GetMemberId() uint64 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

func (m *ToAdd) GetType() PaymentType {
	if m != nil {
		return m.Type
	}
	return PaymentType_TypeNone
}

func (m *ToAdd) GetPaymentCode() PaymentCode {
	if m != nil {
		return m.PaymentCode
	}
	return PaymentCode_CodeNone
}

func (m *ToAdd) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ToAdd) GetParams() []*PaymentParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type ToEdit struct {
	PaymentId            string        `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id"`
	Status               PaymentStatus `protobuf:"varint,2,opt,name=status,proto3,enum=memberpb.PaymentStatus" json:"status"`
	PaymentCode          PaymentCode   `protobuf:"varint,3,opt,name=payment_code,json=paymentCode,proto3,enum=memberpb.PaymentCode" json:"payment_code"`
	Money                float64       `protobuf:"fixed64,4,opt,name=money,proto3" json:"money"`
	PayedMsg             string        `protobuf:"bytes,5,opt,name=payed_msg,json=payedMsg,proto3" json:"payed_msg"`
	TradeNo              string        `protobuf:"bytes,6,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ToEdit) Reset()         { *m = ToEdit{} }
func (m *ToEdit) String() string { return proto.CompactTextString(m) }
func (*ToEdit) ProtoMessage()    {}
func (*ToEdit) Descriptor() ([]byte, []int) {
	return fileDescriptor_b614f87eb7e5cee8, []int{4}
}

func (m *ToEdit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToEdit.Unmarshal(m, b)
}
func (m *ToEdit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToEdit.Marshal(b, m, deterministic)
}
func (m *ToEdit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToEdit.Merge(m, src)
}
func (m *ToEdit) XXX_Size() int {
	return xxx_messageInfo_ToEdit.Size(m)
}
func (m *ToEdit) XXX_DiscardUnknown() {
	xxx_messageInfo_ToEdit.DiscardUnknown(m)
}

var xxx_messageInfo_ToEdit proto.InternalMessageInfo

func (m *ToEdit) GetPaymentId() string {
	if m != nil {
		return m.PaymentId
	}
	return ""
}

func (m *ToEdit) GetStatus() PaymentStatus {
	if m != nil {
		return m.Status
	}
	return PaymentStatus_StatusNone
}

func (m *ToEdit) GetPaymentCode() PaymentCode {
	if m != nil {
		return m.PaymentCode
	}
	return PaymentCode_CodeNone
}

func (m *ToEdit) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *ToEdit) GetPayedMsg() string {
	if m != nil {
		return m.PayedMsg
	}
	return ""
}

func (m *ToEdit) GetTradeNo() string {
	if m != nil {
		return m.TradeNo
	}
	return ""
}

type PaymentRes struct {
	PaymentId            string   `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id"`
	Money                float64  `protobuf:"fixed64,3,opt,name=money,proto3" json:"money"`
	State                uint32   `protobuf:"varint,2,opt,name=state,proto3" json:"state"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentRes) Reset()         { *m = PaymentRes{} }
func (m *PaymentRes) String() string { return proto.CompactTextString(m) }
func (*PaymentRes) ProtoMessage()    {}
func (*PaymentRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b614f87eb7e5cee8, []int{5}
}

func (m *PaymentRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentRes.Unmarshal(m, b)
}
func (m *PaymentRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentRes.Marshal(b, m, deterministic)
}
func (m *PaymentRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentRes.Merge(m, src)
}
func (m *PaymentRes) XXX_Size() int {
	return xxx_messageInfo_PaymentRes.Size(m)
}
func (m *PaymentRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentRes.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentRes proto.InternalMessageInfo

func (m *PaymentRes) GetPaymentId() string {
	if m != nil {
		return m.PaymentId
	}
	return ""
}

func (m *PaymentRes) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *PaymentRes) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func init() {
	proto.RegisterEnum("memberpb.PaymentCode", PaymentCode_name, PaymentCode_value)
	proto.RegisterEnum("memberpb.PaymentType", PaymentType_name, PaymentType_value)
	proto.RegisterEnum("memberpb.PaymentStatus", PaymentStatus_name, PaymentStatus_value)
	proto.RegisterType((*PaymentIdReq)(nil), "memberpb.PaymentIdReq")
	proto.RegisterType((*PaymentParams)(nil), "memberpb.PaymentParams")
	proto.RegisterType((*PaymentRelList)(nil), "memberpb.PaymentRelList")
	proto.RegisterType((*ToAdd)(nil), "memberpb.ToAdd")
	proto.RegisterType((*ToEdit)(nil), "memberpb.ToEdit")
	proto.RegisterType((*PaymentRes)(nil), "memberpb.PaymentRes")
}

func init() { proto.RegisterFile("memberpb/payment.proto", fileDescriptor_b614f87eb7e5cee8) }

var fileDescriptor_b614f87eb7e5cee8 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0x69, 0x93, 0xb5, 0xa7, 0x5b, 0x17, 0x59, 0x63, 0x84, 0x21, 0xc4, 0xc8, 0x55, 0x29,
	0xa2, 0x11, 0x1d, 0x17, 0x08, 0x89, 0x8b, 0x0e, 0x21, 0x54, 0x69, 0x6c, 0x55, 0x5a, 0x34, 0x89,
	0x9b, 0xca, 0x8d, 0xad, 0xd6, 0x52, 0x13, 0x9b, 0xd8, 0x45, 0xca, 0x5b, 0xf1, 0x1c, 0x3c, 0x05,
	0x8f, 0x82, 0xec, 0xfc, 0xb4, 0xa8, 0x83, 0x5d, 0x70, 0x95, 0x9c, 0x73, 0xbe, 0xef, 0x7c, 0xf9,
	0x7c, 0x7c, 0x02, 0xa7, 0x09, 0x4d, 0x16, 0x34, 0x13, 0x8b, 0x50, 0xe0, 0x3c, 0xa1, 0xa9, 0x1a,
	0x88, 0x8c, 0x2b, 0x8e, 0x5a, 0x55, 0x3e, 0x78, 0x05, 0x87, 0x93, 0xa2, 0x34, 0x26, 0x11, 0xfd,
	0x86, 0x9e, 0x02, 0x94, 0xd0, 0x39, 0x23, 0xbe, 0x75, 0x6e, 0xf5, 0xda, 0x51, 0x5b, 0x54, 0x88,
	0xe0, 0x12, 0x8e, 0x4a, 0xf8, 0x04, 0x67, 0x38, 0x91, 0xe8, 0x09, 0xb4, 0x25, 0xdf, 0x64, 0x31,
	0xdd, 0xc2, 0x5b, 0x45, 0x62, 0x4c, 0xd0, 0x09, 0x38, 0x09, 0x4f, 0x69, 0xee, 0xdb, 0xe7, 0x56,
	0xcf, 0x8a, 0x8a, 0x20, 0x78, 0x0f, 0xdd, 0xb2, 0x47, 0x44, 0xd7, 0x57, 0x4c, 0x2a, 0xf4, 0x12,
	0x9a, 0x6b, 0x26, 0x95, 0x6f, 0x9d, 0x37, 0x7a, 0x9d, 0xe1, 0xa3, 0x41, 0xf5, 0x75, 0x83, 0x3f,
	0xb4, 0x22, 0x03, 0x0a, 0x7e, 0x5a, 0xe0, 0xcc, 0xf8, 0x88, 0x10, 0xad, 0x5d, 0x20, 0x2b, 0xed,
	0x66, 0x54, 0x1a, 0x1b, 0x13, 0xf4, 0x02, 0x9a, 0x2a, 0x17, 0xd4, 0x48, 0x77, 0x87, 0x0f, 0xf7,
	0x7a, 0xce, 0x72, 0x41, 0x23, 0x03, 0x41, 0x6f, 0xe1, 0xb0, 0xf2, 0x1c, 0x73, 0x42, 0xfd, 0xc6,
	0x5f, 0x28, 0x1f, 0x38, 0xa1, 0x51, 0x47, 0x6c, 0x03, 0xd4, 0x05, 0x9b, 0x09, 0xbf, 0x69, 0x6c,
	0xdb, 0x4c, 0xa0, 0x10, 0x5c, 0x61, 0xbe, 0xd5, 0x77, 0xfe, 0x6d, 0xa5, 0x84, 0x05, 0xbf, 0x2c,
	0x70, 0x67, 0xfc, 0x23, 0x61, 0xea, 0x9e, 0x93, 0xd7, 0xad, 0xa5, 0xc2, 0x6a, 0x23, 0x4b, 0x47,
	0xfb, 0xad, 0xa7, 0xa6, 0x1c, 0x95, 0xb0, 0xff, 0x70, 0x55, 0x8f, 0xad, 0xb9, 0x33, 0x36, 0x7d,
	0xda, 0x02, 0xe7, 0x94, 0xcc, 0x13, 0xb9, 0xf4, 0x9d, 0x62, 0xd2, 0x26, 0xf1, 0x59, 0x2e, 0xd1,
	0x63, 0x68, 0xa9, 0x0c, 0x13, 0x3a, 0x4f, 0xb9, 0xef, 0x9a, 0xda, 0x81, 0x89, 0xaf, 0x79, 0x70,
	0x0b, 0x50, 0x8f, 0x5b, 0xde, 0xe7, 0xb2, 0x96, 0x6e, 0xec, 0x4a, 0x9f, 0x80, 0xa3, 0x4d, 0x15,
	0xc3, 0x3c, 0x8a, 0x8a, 0xa0, 0x7f, 0x01, 0x9d, 0x1d, 0x0b, 0xe8, 0x10, 0x5a, 0xfa, 0x79, 0xcd,
	0x53, 0xea, 0x3d, 0x40, 0x00, 0xee, 0x2d, 0x8d, 0x57, 0x58, 0x79, 0x96, 0x7e, 0x1f, 0xad, 0x99,
	0xc0, 0xb9, 0x67, 0xf7, 0xdf, 0xd4, 0x24, 0x7d, 0x01, 0x34, 0x49, 0x3f, 0x4b, 0x52, 0x1b, 0x9c,
	0x9b, 0x8c, 0xd0, 0xcc, 0xb3, 0x74, 0x21, 0xd2, 0xfc, 0x6c, 0x49, 0x3d, 0xbb, 0x7f, 0x55, 0x5f,
	0xfb, 0xe2, 0x90, 0x51, 0x17, 0xa0, 0x78, 0xdb, 0xca, 0x7d, 0x49, 0x05, 0x66, 0xc4, 0xb3, 0x74,
	0x6d, 0x82, 0xf3, 0xe9, 0x26, 0x8e, 0xa9, 0x94, 0x9e, 0x8d, 0x8e, 0xa1, 0x53, 0x60, 0x6f, 0xd4,
	0x8a, 0x66, 0x5e, 0x63, 0xf8, 0xc3, 0xaa, 0x37, 0x60, 0x4a, 0xb3, 0xef, 0x2c, 0xa6, 0xe8, 0x1d,
	0xb8, 0x9f, 0xa8, 0x9a, 0xe0, 0x1c, 0x9d, 0xee, 0x0d, 0xc8, 0x2c, 0xe6, 0x99, 0xbf, 0x97, 0xaf,
	0xb6, 0x27, 0x04, 0x77, 0x44, 0x88, 0xe6, 0x1e, 0x6f, 0x31, 0x66, 0x43, 0xce, 0x4e, 0xee, 0x20,
	0x49, 0xf4, 0x1a, 0x0e, 0xf4, 0x8d, 0xd3, 0x0c, 0x6f, 0x97, 0xa1, 0x93, 0x77, 0x53, 0x2e, 0x9f,
	0x7f, 0x7d, 0xb6, 0x64, 0x6a, 0xb5, 0x59, 0x0c, 0x62, 0x9e, 0x84, 0x72, 0xc5, 0xd2, 0x84, 0x2d,
	0x79, 0x28, 0x16, 0x61, 0x85, 0x5e, 0xb8, 0xe6, 0xd7, 0x72, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x12, 0x96, 0xab, 0x79, 0x74, 0x04, 0x00, 0x00,
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
	GetPay(ctx context.Context, in *PaymentIdReq, opts ...grpc.CallOption) (*PaymentRelList, error)
	AddPay(ctx context.Context, in *ToAdd, opts ...grpc.CallOption) (*PaymentRes, error)
	EditPay(ctx context.Context, in *ToEdit, opts ...grpc.CallOption) (*PaymentRes, error)
}

type paymentServiceClient struct {
	cc *grpc.ClientConn
}

func NewPaymentServiceClient(cc *grpc.ClientConn) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) GetPay(ctx context.Context, in *PaymentIdReq, opts ...grpc.CallOption) (*PaymentRelList, error) {
	out := new(PaymentRelList)
	err := c.cc.Invoke(ctx, "/memberpb.PaymentService/GetPay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) AddPay(ctx context.Context, in *ToAdd, opts ...grpc.CallOption) (*PaymentRes, error) {
	out := new(PaymentRes)
	err := c.cc.Invoke(ctx, "/memberpb.PaymentService/AddPay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) EditPay(ctx context.Context, in *ToEdit, opts ...grpc.CallOption) (*PaymentRes, error) {
	out := new(PaymentRes)
	err := c.cc.Invoke(ctx, "/memberpb.PaymentService/EditPay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
type PaymentServiceServer interface {
	GetPay(context.Context, *PaymentIdReq) (*PaymentRelList, error)
	AddPay(context.Context, *ToAdd) (*PaymentRes, error)
	EditPay(context.Context, *ToEdit) (*PaymentRes, error)
}

func RegisterPaymentServiceServer(s *grpc.Server, srv PaymentServiceServer) {
	s.RegisterService(&_PaymentService_serviceDesc, srv)
}

func _PaymentService_GetPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memberpb.PaymentService/GetPay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPay(ctx, req.(*PaymentIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_AddPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToAdd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).AddPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memberpb.PaymentService/AddPay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).AddPay(ctx, req.(*ToAdd))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_EditPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToEdit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).EditPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memberpb.PaymentService/EditPay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).EditPay(ctx, req.(*ToEdit))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "memberpb.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPay",
			Handler:    _PaymentService_GetPay_Handler,
		},
		{
			MethodName: "AddPay",
			Handler:    _PaymentService_AddPay_Handler,
		},
		{
			MethodName: "EditPay",
			Handler:    _PaymentService_EditPay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "memberpb/payment.proto",
}
