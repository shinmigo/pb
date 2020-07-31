// Code generated by protoc-gen-go. DO NOT EDIT.
// source: param.proto

package productpb

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

type ParamType int32

const (
	ParamType_Text     ParamType = 0
	ParamType_Radio    ParamType = 1
	ParamType_Checkbox ParamType = 2
)

var ParamType_name = map[int32]string{
	0: "Text",
	1: "Radio",
	2: "Checkbox",
}

var ParamType_value = map[string]int32{
	"Text":     0,
	"Radio":    1,
	"Checkbox": 2,
}

func (x ParamType) String() string {
	return proto.EnumName(ParamType_name, int32(x))
}

func (ParamType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{0}
}

type Param struct {
	ParamId              uint64    `protobuf:"varint,1,opt,name=param_id,json=paramId,proto3" json:"param_id,omitempty"`
	StoreId              uint64    `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	KindId               uint64    `protobuf:"varint,3,opt,name=kind_id,json=kindId,proto3" json:"kind_id,omitempty"`
	Name                 string    `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Type                 ParamType `protobuf:"varint,5,opt,name=type,proto3,enum=productpb.ParamType" json:"type,omitempty"`
	Sort                 uint64    `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`
	AdminId              uint64    `protobuf:"varint,7,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	Contents             []string  `protobuf:"bytes,8,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Param) Reset()         { *m = Param{} }
func (m *Param) String() string { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()    {}
func (*Param) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{0}
}

func (m *Param) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Param.Unmarshal(m, b)
}
func (m *Param) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Param.Marshal(b, m, deterministic)
}
func (m *Param) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Param.Merge(m, src)
}
func (m *Param) XXX_Size() int {
	return xxx_messageInfo_Param.Size(m)
}
func (m *Param) XXX_DiscardUnknown() {
	xxx_messageInfo_Param.DiscardUnknown(m)
}

var xxx_messageInfo_Param proto.InternalMessageInfo

func (m *Param) GetParamId() uint64 {
	if m != nil {
		return m.ParamId
	}
	return 0
}

func (m *Param) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Param) GetKindId() uint64 {
	if m != nil {
		return m.KindId
	}
	return 0
}

func (m *Param) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Param) GetType() ParamType {
	if m != nil {
		return m.Type
	}
	return ParamType_Text
}

func (m *Param) GetSort() uint64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *Param) GetAdminId() uint64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

func (m *Param) GetContents() []string {
	if m != nil {
		return m.Contents
	}
	return nil
}

type ParamInfo struct {
	ParamId              uint64    `protobuf:"varint,1,opt,name=param_id,json=paramId,proto3" json:"param_id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 ParamType `protobuf:"varint,3,opt,name=type,proto3,enum=productpb.ParamType" json:"type,omitempty"`
	Sort                 uint64    `protobuf:"varint,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Contents             []string  `protobuf:"bytes,5,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ParamInfo) Reset()         { *m = ParamInfo{} }
func (m *ParamInfo) String() string { return proto.CompactTextString(m) }
func (*ParamInfo) ProtoMessage()    {}
func (*ParamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{1}
}

func (m *ParamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParamInfo.Unmarshal(m, b)
}
func (m *ParamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParamInfo.Marshal(b, m, deterministic)
}
func (m *ParamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamInfo.Merge(m, src)
}
func (m *ParamInfo) XXX_Size() int {
	return xxx_messageInfo_ParamInfo.Size(m)
}
func (m *ParamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ParamInfo proto.InternalMessageInfo

func (m *ParamInfo) GetParamId() uint64 {
	if m != nil {
		return m.ParamId
	}
	return 0
}

func (m *ParamInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ParamInfo) GetType() ParamType {
	if m != nil {
		return m.Type
	}
	return ParamType_Text
}

func (m *ParamInfo) GetSort() uint64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *ParamInfo) GetContents() []string {
	if m != nil {
		return m.Contents
	}
	return nil
}

type AddParamReq struct {
	Param                *Param   `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddParamReq) Reset()         { *m = AddParamReq{} }
func (m *AddParamReq) String() string { return proto.CompactTextString(m) }
func (*AddParamReq) ProtoMessage()    {}
func (*AddParamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{2}
}

func (m *AddParamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddParamReq.Unmarshal(m, b)
}
func (m *AddParamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddParamReq.Marshal(b, m, deterministic)
}
func (m *AddParamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddParamReq.Merge(m, src)
}
func (m *AddParamReq) XXX_Size() int {
	return xxx_messageInfo_AddParamReq.Size(m)
}
func (m *AddParamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddParamReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddParamReq proto.InternalMessageInfo

func (m *AddParamReq) GetParam() *Param {
	if m != nil {
		return m.Param
	}
	return nil
}

type AddParamRes struct {
	ParamId              uint64   `protobuf:"varint,1,opt,name=param_id,json=paramId,proto3" json:"param_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddParamRes) Reset()         { *m = AddParamRes{} }
func (m *AddParamRes) String() string { return proto.CompactTextString(m) }
func (*AddParamRes) ProtoMessage()    {}
func (*AddParamRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{3}
}

func (m *AddParamRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddParamRes.Unmarshal(m, b)
}
func (m *AddParamRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddParamRes.Marshal(b, m, deterministic)
}
func (m *AddParamRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddParamRes.Merge(m, src)
}
func (m *AddParamRes) XXX_Size() int {
	return xxx_messageInfo_AddParamRes.Size(m)
}
func (m *AddParamRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AddParamRes.DiscardUnknown(m)
}

var xxx_messageInfo_AddParamRes proto.InternalMessageInfo

func (m *AddParamRes) GetParamId() uint64 {
	if m != nil {
		return m.ParamId
	}
	return 0
}

type EditParamReq struct {
	Param                *Param   `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditParamReq) Reset()         { *m = EditParamReq{} }
func (m *EditParamReq) String() string { return proto.CompactTextString(m) }
func (*EditParamReq) ProtoMessage()    {}
func (*EditParamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{4}
}

func (m *EditParamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditParamReq.Unmarshal(m, b)
}
func (m *EditParamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditParamReq.Marshal(b, m, deterministic)
}
func (m *EditParamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditParamReq.Merge(m, src)
}
func (m *EditParamReq) XXX_Size() int {
	return xxx_messageInfo_EditParamReq.Size(m)
}
func (m *EditParamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EditParamReq.DiscardUnknown(m)
}

var xxx_messageInfo_EditParamReq proto.InternalMessageInfo

func (m *EditParamReq) GetParam() *Param {
	if m != nil {
		return m.Param
	}
	return nil
}

type EditParamRes struct {
	// 在成功的情况下等于1
	Updated              int32    `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditParamRes) Reset()         { *m = EditParamRes{} }
func (m *EditParamRes) String() string { return proto.CompactTextString(m) }
func (*EditParamRes) ProtoMessage()    {}
func (*EditParamRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{5}
}

func (m *EditParamRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditParamRes.Unmarshal(m, b)
}
func (m *EditParamRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditParamRes.Marshal(b, m, deterministic)
}
func (m *EditParamRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditParamRes.Merge(m, src)
}
func (m *EditParamRes) XXX_Size() int {
	return xxx_messageInfo_EditParamRes.Size(m)
}
func (m *EditParamRes) XXX_DiscardUnknown() {
	xxx_messageInfo_EditParamRes.DiscardUnknown(m)
}

var xxx_messageInfo_EditParamRes proto.InternalMessageInfo

func (m *EditParamRes) GetUpdated() int32 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type DelParamReq struct {
	ParamId              uint64   `protobuf:"varint,1,opt,name=param_id,json=paramId,proto3" json:"param_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelParamReq) Reset()         { *m = DelParamReq{} }
func (m *DelParamReq) String() string { return proto.CompactTextString(m) }
func (*DelParamReq) ProtoMessage()    {}
func (*DelParamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{6}
}

func (m *DelParamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelParamReq.Unmarshal(m, b)
}
func (m *DelParamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelParamReq.Marshal(b, m, deterministic)
}
func (m *DelParamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelParamReq.Merge(m, src)
}
func (m *DelParamReq) XXX_Size() int {
	return xxx_messageInfo_DelParamReq.Size(m)
}
func (m *DelParamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelParamReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelParamReq proto.InternalMessageInfo

func (m *DelParamReq) GetParamId() uint64 {
	if m != nil {
		return m.ParamId
	}
	return 0
}

type DelParamRes struct {
	// 在成功的情况下等于1
	Deleted              int32    `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelParamRes) Reset()         { *m = DelParamRes{} }
func (m *DelParamRes) String() string { return proto.CompactTextString(m) }
func (*DelParamRes) ProtoMessage()    {}
func (*DelParamRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{7}
}

func (m *DelParamRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelParamRes.Unmarshal(m, b)
}
func (m *DelParamRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelParamRes.Marshal(b, m, deterministic)
}
func (m *DelParamRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelParamRes.Merge(m, src)
}
func (m *DelParamRes) XXX_Size() int {
	return xxx_messageInfo_DelParamRes.Size(m)
}
func (m *DelParamRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DelParamRes.DiscardUnknown(m)
}

var xxx_messageInfo_DelParamRes proto.InternalMessageInfo

func (m *DelParamRes) GetDeleted() int32 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

type ReadParamReq struct {
	ParamId              uint64   `protobuf:"varint,1,opt,name=param_id,json=paramId,proto3" json:"param_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadParamReq) Reset()         { *m = ReadParamReq{} }
func (m *ReadParamReq) String() string { return proto.CompactTextString(m) }
func (*ReadParamReq) ProtoMessage()    {}
func (*ReadParamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{8}
}

func (m *ReadParamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadParamReq.Unmarshal(m, b)
}
func (m *ReadParamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadParamReq.Marshal(b, m, deterministic)
}
func (m *ReadParamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadParamReq.Merge(m, src)
}
func (m *ReadParamReq) XXX_Size() int {
	return xxx_messageInfo_ReadParamReq.Size(m)
}
func (m *ReadParamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadParamReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadParamReq proto.InternalMessageInfo

func (m *ReadParamReq) GetParamId() uint64 {
	if m != nil {
		return m.ParamId
	}
	return 0
}

type ReadParamRes struct {
	Param                *ParamInfo `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadParamRes) Reset()         { *m = ReadParamRes{} }
func (m *ReadParamRes) String() string { return proto.CompactTextString(m) }
func (*ReadParamRes) ProtoMessage()    {}
func (*ReadParamRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{9}
}

func (m *ReadParamRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadParamRes.Unmarshal(m, b)
}
func (m *ReadParamRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadParamRes.Marshal(b, m, deterministic)
}
func (m *ReadParamRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadParamRes.Merge(m, src)
}
func (m *ReadParamRes) XXX_Size() int {
	return xxx_messageInfo_ReadParamRes.Size(m)
}
func (m *ReadParamRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadParamRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadParamRes proto.InternalMessageInfo

func (m *ReadParamRes) GetParam() *ParamInfo {
	if m != nil {
		return m.Param
	}
	return nil
}

type ReadParamsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadParamsReq) Reset()         { *m = ReadParamsReq{} }
func (m *ReadParamsReq) String() string { return proto.CompactTextString(m) }
func (*ReadParamsReq) ProtoMessage()    {}
func (*ReadParamsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{10}
}

func (m *ReadParamsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadParamsReq.Unmarshal(m, b)
}
func (m *ReadParamsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadParamsReq.Marshal(b, m, deterministic)
}
func (m *ReadParamsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadParamsReq.Merge(m, src)
}
func (m *ReadParamsReq) XXX_Size() int {
	return xxx_messageInfo_ReadParamsReq.Size(m)
}
func (m *ReadParamsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadParamsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadParamsReq proto.InternalMessageInfo

type ReadParamsRes struct {
	Params               []*ParamInfo `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReadParamsRes) Reset()         { *m = ReadParamsRes{} }
func (m *ReadParamsRes) String() string { return proto.CompactTextString(m) }
func (*ReadParamsRes) ProtoMessage()    {}
func (*ReadParamsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1747c98d2fc1d5, []int{11}
}

func (m *ReadParamsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadParamsRes.Unmarshal(m, b)
}
func (m *ReadParamsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadParamsRes.Marshal(b, m, deterministic)
}
func (m *ReadParamsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadParamsRes.Merge(m, src)
}
func (m *ReadParamsRes) XXX_Size() int {
	return xxx_messageInfo_ReadParamsRes.Size(m)
}
func (m *ReadParamsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadParamsRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadParamsRes proto.InternalMessageInfo

func (m *ReadParamsRes) GetParams() []*ParamInfo {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterEnum("productpb.ParamType", ParamType_name, ParamType_value)
	proto.RegisterType((*Param)(nil), "productpb.Param")
	proto.RegisterType((*ParamInfo)(nil), "productpb.ParamInfo")
	proto.RegisterType((*AddParamReq)(nil), "productpb.AddParamReq")
	proto.RegisterType((*AddParamRes)(nil), "productpb.AddParamRes")
	proto.RegisterType((*EditParamReq)(nil), "productpb.EditParamReq")
	proto.RegisterType((*EditParamRes)(nil), "productpb.EditParamRes")
	proto.RegisterType((*DelParamReq)(nil), "productpb.DelParamReq")
	proto.RegisterType((*DelParamRes)(nil), "productpb.DelParamRes")
	proto.RegisterType((*ReadParamReq)(nil), "productpb.ReadParamReq")
	proto.RegisterType((*ReadParamRes)(nil), "productpb.ReadParamRes")
	proto.RegisterType((*ReadParamsReq)(nil), "productpb.ReadParamsReq")
	proto.RegisterType((*ReadParamsRes)(nil), "productpb.ReadParamsRes")
}

func init() { proto.RegisterFile("param.proto", fileDescriptor_1e1747c98d2fc1d5) }

var fileDescriptor_1e1747c98d2fc1d5 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x89, 0x9d, 0xd8, 0x93, 0x00, 0xd1, 0x08, 0xd1, 0xc5, 0x27, 0xcb, 0x07, 0x30, 0x15,
	0xca, 0x21, 0x08, 0x0e, 0x95, 0x2a, 0x81, 0x80, 0x83, 0x6f, 0x68, 0xe9, 0x1d, 0x39, 0xd9, 0x45,
	0x58, 0x6d, 0xbc, 0xae, 0x77, 0x8b, 0xda, 0xff, 0xe0, 0x9b, 0xf8, 0x0d, 0x7e, 0x05, 0xed, 0x24,
	0x36, 0xeb, 0x92, 0x44, 0x11, 0x37, 0xcf, 0xbe, 0x37, 0xf3, 0xde, 0xec, 0x5b, 0x19, 0x26, 0x75,
	0xd1, 0x14, 0xeb, 0x79, 0xdd, 0x28, 0xa3, 0x30, 0xaa, 0x1b, 0x25, 0x6e, 0x56, 0xa6, 0x5e, 0xa6,
	0xbf, 0x3d, 0x08, 0x3e, 0x5b, 0x08, 0x9f, 0x41, 0x48, 0x9c, 0xaf, 0xa5, 0x60, 0x5e, 0xe2, 0x65,
	0x3e, 0x1f, 0x53, 0x9d, 0x0b, 0x0b, 0x69, 0xa3, 0x1a, 0x69, 0xa1, 0xc1, 0x06, 0xa2, 0x3a, 0x17,
	0x78, 0x02, 0xe3, 0xcb, 0xb2, 0x12, 0x16, 0x19, 0x12, 0x32, 0xb2, 0x65, 0x2e, 0x10, 0xc1, 0xaf,
	0x8a, 0xb5, 0x64, 0x7e, 0xe2, 0x65, 0x11, 0xa7, 0x6f, 0xcc, 0xc0, 0x37, 0x77, 0xb5, 0x64, 0x41,
	0xe2, 0x65, 0x8f, 0x16, 0x4f, 0xe6, 0x9d, 0x8d, 0x39, 0x59, 0xb8, 0xb8, 0xab, 0x25, 0x27, 0x86,
	0xed, 0xd6, 0xaa, 0x31, 0x6c, 0x44, 0x33, 0xe9, 0xdb, 0xba, 0x28, 0xc4, 0xba, 0xac, 0xac, 0xd6,
	0x78, 0xe3, 0x82, 0xea, 0x5c, 0x60, 0x0c, 0xe1, 0x4a, 0x55, 0x46, 0x56, 0x46, 0xb3, 0x30, 0x19,
	0x66, 0x11, 0xef, 0xea, 0xf4, 0xa7, 0x07, 0x11, 0x8d, 0xcf, 0xab, 0x6f, 0xea, 0xd0, 0x96, 0xad,
	0xe3, 0xc1, 0x0e, 0xc7, 0xc3, 0xa3, 0x1d, 0xfb, 0x8e, 0x63, 0xd7, 0x56, 0x70, 0xcf, 0xd6, 0x1b,
	0x98, 0xbc, 0x17, 0x82, 0xa6, 0x70, 0x79, 0x8d, 0xcf, 0x21, 0x20, 0x1f, 0x64, 0x6a, 0xb2, 0x98,
	0xdd, 0x57, 0xe2, 0x1b, 0x38, 0xcd, 0xdc, 0x36, 0x7d, 0x60, 0x9d, 0xf4, 0x2d, 0x4c, 0x3f, 0x89,
	0xd2, 0xfc, 0x87, 0x82, 0xdb, 0xa7, 0x91, 0xc1, 0xf8, 0xa6, 0x16, 0x85, 0x91, 0x1b, 0x85, 0x80,
	0xb7, 0xa5, 0xf5, 0xf2, 0x51, 0x5e, 0x75, 0x02, 0x07, 0xbc, 0xbc, 0x70, 0x99, 0x34, 0x52, 0xc8,
	0x2b, 0xe9, 0x8c, 0xdc, 0x96, 0xe9, 0x4b, 0x98, 0x72, 0x59, 0x88, 0x63, 0x66, 0x9e, 0xf5, 0xa8,
	0x1a, 0x4f, 0xfb, 0xfb, 0xfd, 0x93, 0x95, 0x8d, 0xbf, 0xdd, 0xf1, 0x31, 0x3c, 0xec, 0x7a, 0x35,
	0x97, 0xd7, 0xe9, 0x79, 0xff, 0x40, 0xe3, 0x2b, 0x18, 0x11, 0x55, 0x33, 0x2f, 0x19, 0xee, 0x1d,
	0xb7, 0xe5, 0x9c, 0xce, 0xb7, 0x4f, 0xcc, 0xbe, 0x07, 0x0c, 0xc1, 0xbf, 0x90, 0xb7, 0x66, 0xf6,
	0x00, 0x23, 0x08, 0x78, 0x21, 0x4a, 0x35, 0xf3, 0x70, 0x0a, 0xe1, 0x87, 0xef, 0x72, 0x75, 0xb9,
	0x54, 0xb7, 0xb3, 0xc1, 0xe2, 0xd7, 0x00, 0xa6, 0xd4, 0xf0, 0x45, 0x36, 0x3f, 0xca, 0x95, 0xc4,
	0x33, 0x08, 0xdb, 0x58, 0xf1, 0xa9, 0x23, 0xe5, 0x3c, 0x91, 0x78, 0xf7, 0xb9, 0xc6, 0x73, 0x88,
	0xba, 0xc0, 0xf0, 0xc4, 0x21, 0xb9, 0xf1, 0xc7, 0x7b, 0x00, 0x6d, 0xa5, 0xdb, 0x6c, 0x7a, 0xd2,
	0x4e, 0xb4, 0xf1, 0xee, 0x73, 0x92, 0xee, 0xae, 0xad, 0x27, 0xed, 0x86, 0x18, 0xef, 0x01, 0x34,
	0xbe, 0x03, 0xf8, 0x7b, 0xeb, 0xc8, 0x76, 0xd1, 0x6c, 0x3a, 0xf1, 0x3e, 0x44, 0x2f, 0x47, 0xf4,
	0x43, 0x7b, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xb9, 0x83, 0xc4, 0xdf, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ParamServiceClient is the client API for ParamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParamServiceClient interface {
	AddParam(ctx context.Context, in *AddParamReq, opts ...grpc.CallOption) (*AddParamRes, error)
	EditParam(ctx context.Context, in *EditParamReq, opts ...grpc.CallOption) (*EditParamRes, error)
	DelParam(ctx context.Context, in *DelParamReq, opts ...grpc.CallOption) (*DelParamRes, error)
	ReadParam(ctx context.Context, in *ReadParamReq, opts ...grpc.CallOption) (*ReadParamRes, error)
	ReadParams(ctx context.Context, in *ReadParamsReq, opts ...grpc.CallOption) (*ReadParamsRes, error)
}

type paramServiceClient struct {
	cc *grpc.ClientConn
}

func NewParamServiceClient(cc *grpc.ClientConn) ParamServiceClient {
	return &paramServiceClient{cc}
}

func (c *paramServiceClient) AddParam(ctx context.Context, in *AddParamReq, opts ...grpc.CallOption) (*AddParamRes, error) {
	out := new(AddParamRes)
	err := c.cc.Invoke(ctx, "/productpb.ParamService/AddParam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramServiceClient) EditParam(ctx context.Context, in *EditParamReq, opts ...grpc.CallOption) (*EditParamRes, error) {
	out := new(EditParamRes)
	err := c.cc.Invoke(ctx, "/productpb.ParamService/EditParam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramServiceClient) DelParam(ctx context.Context, in *DelParamReq, opts ...grpc.CallOption) (*DelParamRes, error) {
	out := new(DelParamRes)
	err := c.cc.Invoke(ctx, "/productpb.ParamService/DelParam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramServiceClient) ReadParam(ctx context.Context, in *ReadParamReq, opts ...grpc.CallOption) (*ReadParamRes, error) {
	out := new(ReadParamRes)
	err := c.cc.Invoke(ctx, "/productpb.ParamService/ReadParam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramServiceClient) ReadParams(ctx context.Context, in *ReadParamsReq, opts ...grpc.CallOption) (*ReadParamsRes, error) {
	out := new(ReadParamsRes)
	err := c.cc.Invoke(ctx, "/productpb.ParamService/ReadParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParamServiceServer is the server API for ParamService service.
type ParamServiceServer interface {
	AddParam(context.Context, *AddParamReq) (*AddParamRes, error)
	EditParam(context.Context, *EditParamReq) (*EditParamRes, error)
	DelParam(context.Context, *DelParamReq) (*DelParamRes, error)
	ReadParam(context.Context, *ReadParamReq) (*ReadParamRes, error)
	ReadParams(context.Context, *ReadParamsReq) (*ReadParamsRes, error)
}

func RegisterParamServiceServer(s *grpc.Server, srv ParamServiceServer) {
	s.RegisterService(&_ParamService_serviceDesc, srv)
}

func _ParamService_AddParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddParamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamServiceServer).AddParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ParamService/AddParam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamServiceServer).AddParam(ctx, req.(*AddParamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamService_EditParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditParamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamServiceServer).EditParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ParamService/EditParam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamServiceServer).EditParam(ctx, req.(*EditParamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamService_DelParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelParamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamServiceServer).DelParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ParamService/DelParam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamServiceServer).DelParam(ctx, req.(*DelParamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamService_ReadParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadParamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamServiceServer).ReadParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ParamService/ReadParam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamServiceServer).ReadParam(ctx, req.(*ReadParamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamService_ReadParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadParamsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamServiceServer).ReadParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ParamService/ReadParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamServiceServer).ReadParams(ctx, req.(*ReadParamsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "productpb.ParamService",
	HandlerType: (*ParamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddParam",
			Handler:    _ParamService_AddParam_Handler,
		},
		{
			MethodName: "EditParam",
			Handler:    _ParamService_EditParam_Handler,
		},
		{
			MethodName: "DelParam",
			Handler:    _ParamService_DelParam_Handler,
		},
		{
			MethodName: "ReadParam",
			Handler:    _ParamService_ReadParam_Handler,
		},
		{
			MethodName: "ReadParams",
			Handler:    _ParamService_ReadParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "param.proto",
}