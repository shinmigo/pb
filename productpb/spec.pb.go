// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spec.proto

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

type Spec struct {
	SpecId               uint64   `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	StoreId              uint64   `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	KindId               uint64   `protobuf:"varint,3,opt,name=kind_id,json=kindId,proto3" json:"kind_id,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Sort                 uint64   `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"`
	AdminId              uint64   `protobuf:"varint,6,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	Contents             []string `protobuf:"bytes,7,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{0}
}

func (m *Spec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spec.Unmarshal(m, b)
}
func (m *Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spec.Marshal(b, m, deterministic)
}
func (m *Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec.Merge(m, src)
}
func (m *Spec) XXX_Size() int {
	return xxx_messageInfo_Spec.Size(m)
}
func (m *Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_Spec proto.InternalMessageInfo

func (m *Spec) GetSpecId() uint64 {
	if m != nil {
		return m.SpecId
	}
	return 0
}

func (m *Spec) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Spec) GetKindId() uint64 {
	if m != nil {
		return m.KindId
	}
	return 0
}

func (m *Spec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Spec) GetSort() uint64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *Spec) GetAdminId() uint64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

func (m *Spec) GetContents() []string {
	if m != nil {
		return m.Contents
	}
	return nil
}

type SpecInfo struct {
	SpecId               uint64   `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sort                 uint64   `protobuf:"varint,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Contents             []string `protobuf:"bytes,4,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecInfo) Reset()         { *m = SpecInfo{} }
func (m *SpecInfo) String() string { return proto.CompactTextString(m) }
func (*SpecInfo) ProtoMessage()    {}
func (*SpecInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{1}
}

func (m *SpecInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecInfo.Unmarshal(m, b)
}
func (m *SpecInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecInfo.Marshal(b, m, deterministic)
}
func (m *SpecInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecInfo.Merge(m, src)
}
func (m *SpecInfo) XXX_Size() int {
	return xxx_messageInfo_SpecInfo.Size(m)
}
func (m *SpecInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SpecInfo proto.InternalMessageInfo

func (m *SpecInfo) GetSpecId() uint64 {
	if m != nil {
		return m.SpecId
	}
	return 0
}

func (m *SpecInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SpecInfo) GetSort() uint64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *SpecInfo) GetContents() []string {
	if m != nil {
		return m.Contents
	}
	return nil
}

type AddSpecReq struct {
	Spec                 *Spec    `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSpecReq) Reset()         { *m = AddSpecReq{} }
func (m *AddSpecReq) String() string { return proto.CompactTextString(m) }
func (*AddSpecReq) ProtoMessage()    {}
func (*AddSpecReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{2}
}

func (m *AddSpecReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSpecReq.Unmarshal(m, b)
}
func (m *AddSpecReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSpecReq.Marshal(b, m, deterministic)
}
func (m *AddSpecReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSpecReq.Merge(m, src)
}
func (m *AddSpecReq) XXX_Size() int {
	return xxx_messageInfo_AddSpecReq.Size(m)
}
func (m *AddSpecReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSpecReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddSpecReq proto.InternalMessageInfo

func (m *AddSpecReq) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type AddSpecRes struct {
	SpecId               uint64   `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSpecRes) Reset()         { *m = AddSpecRes{} }
func (m *AddSpecRes) String() string { return proto.CompactTextString(m) }
func (*AddSpecRes) ProtoMessage()    {}
func (*AddSpecRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{3}
}

func (m *AddSpecRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSpecRes.Unmarshal(m, b)
}
func (m *AddSpecRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSpecRes.Marshal(b, m, deterministic)
}
func (m *AddSpecRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSpecRes.Merge(m, src)
}
func (m *AddSpecRes) XXX_Size() int {
	return xxx_messageInfo_AddSpecRes.Size(m)
}
func (m *AddSpecRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSpecRes.DiscardUnknown(m)
}

var xxx_messageInfo_AddSpecRes proto.InternalMessageInfo

func (m *AddSpecRes) GetSpecId() uint64 {
	if m != nil {
		return m.SpecId
	}
	return 0
}

type EditSpecReq struct {
	Spec                 *Spec    `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditSpecReq) Reset()         { *m = EditSpecReq{} }
func (m *EditSpecReq) String() string { return proto.CompactTextString(m) }
func (*EditSpecReq) ProtoMessage()    {}
func (*EditSpecReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{4}
}

func (m *EditSpecReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditSpecReq.Unmarshal(m, b)
}
func (m *EditSpecReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditSpecReq.Marshal(b, m, deterministic)
}
func (m *EditSpecReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditSpecReq.Merge(m, src)
}
func (m *EditSpecReq) XXX_Size() int {
	return xxx_messageInfo_EditSpecReq.Size(m)
}
func (m *EditSpecReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EditSpecReq.DiscardUnknown(m)
}

var xxx_messageInfo_EditSpecReq proto.InternalMessageInfo

func (m *EditSpecReq) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type EditSpecRes struct {
	// 在成功的情况下等于1
	Updated              int32    `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditSpecRes) Reset()         { *m = EditSpecRes{} }
func (m *EditSpecRes) String() string { return proto.CompactTextString(m) }
func (*EditSpecRes) ProtoMessage()    {}
func (*EditSpecRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{5}
}

func (m *EditSpecRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditSpecRes.Unmarshal(m, b)
}
func (m *EditSpecRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditSpecRes.Marshal(b, m, deterministic)
}
func (m *EditSpecRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditSpecRes.Merge(m, src)
}
func (m *EditSpecRes) XXX_Size() int {
	return xxx_messageInfo_EditSpecRes.Size(m)
}
func (m *EditSpecRes) XXX_DiscardUnknown() {
	xxx_messageInfo_EditSpecRes.DiscardUnknown(m)
}

var xxx_messageInfo_EditSpecRes proto.InternalMessageInfo

func (m *EditSpecRes) GetUpdated() int32 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type DelSpecReq struct {
	SpecId               uint64   `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelSpecReq) Reset()         { *m = DelSpecReq{} }
func (m *DelSpecReq) String() string { return proto.CompactTextString(m) }
func (*DelSpecReq) ProtoMessage()    {}
func (*DelSpecReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{6}
}

func (m *DelSpecReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelSpecReq.Unmarshal(m, b)
}
func (m *DelSpecReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelSpecReq.Marshal(b, m, deterministic)
}
func (m *DelSpecReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelSpecReq.Merge(m, src)
}
func (m *DelSpecReq) XXX_Size() int {
	return xxx_messageInfo_DelSpecReq.Size(m)
}
func (m *DelSpecReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelSpecReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelSpecReq proto.InternalMessageInfo

func (m *DelSpecReq) GetSpecId() uint64 {
	if m != nil {
		return m.SpecId
	}
	return 0
}

type DelSpecRes struct {
	// 在成功的情况下等于1
	Deleted              int32    `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelSpecRes) Reset()         { *m = DelSpecRes{} }
func (m *DelSpecRes) String() string { return proto.CompactTextString(m) }
func (*DelSpecRes) ProtoMessage()    {}
func (*DelSpecRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{7}
}

func (m *DelSpecRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelSpecRes.Unmarshal(m, b)
}
func (m *DelSpecRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelSpecRes.Marshal(b, m, deterministic)
}
func (m *DelSpecRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelSpecRes.Merge(m, src)
}
func (m *DelSpecRes) XXX_Size() int {
	return xxx_messageInfo_DelSpecRes.Size(m)
}
func (m *DelSpecRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DelSpecRes.DiscardUnknown(m)
}

var xxx_messageInfo_DelSpecRes proto.InternalMessageInfo

func (m *DelSpecRes) GetDeleted() int32 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

type ReadSpecReq struct {
	SpecId               uint64   `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadSpecReq) Reset()         { *m = ReadSpecReq{} }
func (m *ReadSpecReq) String() string { return proto.CompactTextString(m) }
func (*ReadSpecReq) ProtoMessage()    {}
func (*ReadSpecReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{8}
}

func (m *ReadSpecReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSpecReq.Unmarshal(m, b)
}
func (m *ReadSpecReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSpecReq.Marshal(b, m, deterministic)
}
func (m *ReadSpecReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSpecReq.Merge(m, src)
}
func (m *ReadSpecReq) XXX_Size() int {
	return xxx_messageInfo_ReadSpecReq.Size(m)
}
func (m *ReadSpecReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSpecReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSpecReq proto.InternalMessageInfo

func (m *ReadSpecReq) GetSpecId() uint64 {
	if m != nil {
		return m.SpecId
	}
	return 0
}

type ReadSpecRes struct {
	Spec                 *SpecInfo `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReadSpecRes) Reset()         { *m = ReadSpecRes{} }
func (m *ReadSpecRes) String() string { return proto.CompactTextString(m) }
func (*ReadSpecRes) ProtoMessage()    {}
func (*ReadSpecRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{9}
}

func (m *ReadSpecRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSpecRes.Unmarshal(m, b)
}
func (m *ReadSpecRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSpecRes.Marshal(b, m, deterministic)
}
func (m *ReadSpecRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSpecRes.Merge(m, src)
}
func (m *ReadSpecRes) XXX_Size() int {
	return xxx_messageInfo_ReadSpecRes.Size(m)
}
func (m *ReadSpecRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSpecRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSpecRes proto.InternalMessageInfo

func (m *ReadSpecRes) GetSpec() *SpecInfo {
	if m != nil {
		return m.Spec
	}
	return nil
}

type ReadSpecsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadSpecsReq) Reset()         { *m = ReadSpecsReq{} }
func (m *ReadSpecsReq) String() string { return proto.CompactTextString(m) }
func (*ReadSpecsReq) ProtoMessage()    {}
func (*ReadSpecsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{10}
}

func (m *ReadSpecsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSpecsReq.Unmarshal(m, b)
}
func (m *ReadSpecsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSpecsReq.Marshal(b, m, deterministic)
}
func (m *ReadSpecsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSpecsReq.Merge(m, src)
}
func (m *ReadSpecsReq) XXX_Size() int {
	return xxx_messageInfo_ReadSpecsReq.Size(m)
}
func (m *ReadSpecsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSpecsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSpecsReq proto.InternalMessageInfo

type ReadSpecsRes struct {
	Specs                []*SpecInfo `protobuf:"bytes,1,rep,name=specs,proto3" json:"specs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReadSpecsRes) Reset()         { *m = ReadSpecsRes{} }
func (m *ReadSpecsRes) String() string { return proto.CompactTextString(m) }
func (*ReadSpecsRes) ProtoMessage()    {}
func (*ReadSpecsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{11}
}

func (m *ReadSpecsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSpecsRes.Unmarshal(m, b)
}
func (m *ReadSpecsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSpecsRes.Marshal(b, m, deterministic)
}
func (m *ReadSpecsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSpecsRes.Merge(m, src)
}
func (m *ReadSpecsRes) XXX_Size() int {
	return xxx_messageInfo_ReadSpecsRes.Size(m)
}
func (m *ReadSpecsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSpecsRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSpecsRes proto.InternalMessageInfo

func (m *ReadSpecsRes) GetSpecs() []*SpecInfo {
	if m != nil {
		return m.Specs
	}
	return nil
}

func init() {
	proto.RegisterType((*Spec)(nil), "productpb.Spec")
	proto.RegisterType((*SpecInfo)(nil), "productpb.SpecInfo")
	proto.RegisterType((*AddSpecReq)(nil), "productpb.AddSpecReq")
	proto.RegisterType((*AddSpecRes)(nil), "productpb.AddSpecRes")
	proto.RegisterType((*EditSpecReq)(nil), "productpb.EditSpecReq")
	proto.RegisterType((*EditSpecRes)(nil), "productpb.EditSpecRes")
	proto.RegisterType((*DelSpecReq)(nil), "productpb.DelSpecReq")
	proto.RegisterType((*DelSpecRes)(nil), "productpb.DelSpecRes")
	proto.RegisterType((*ReadSpecReq)(nil), "productpb.ReadSpecReq")
	proto.RegisterType((*ReadSpecRes)(nil), "productpb.ReadSpecRes")
	proto.RegisterType((*ReadSpecsReq)(nil), "productpb.ReadSpecsReq")
	proto.RegisterType((*ReadSpecsRes)(nil), "productpb.ReadSpecsRes")
}

func init() { proto.RegisterFile("spec.proto", fileDescriptor_423806180556987f) }

var fileDescriptor_423806180556987f = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x8a, 0x9b, 0x50,
	0x10, 0xc5, 0x68, 0xa2, 0x8e, 0xa5, 0x85, 0x29, 0x6d, 0xac, 0x4f, 0x62, 0x69, 0x92, 0xbe, 0x04,
	0x9a, 0x42, 0x4b, 0x0b, 0x7d, 0x28, 0xb4, 0x0f, 0xbe, 0x9a, 0x0f, 0x58, 0x8c, 0xf7, 0xee, 0x22,
	0x9b, 0xa8, 0xf1, 0xde, 0xec, 0x87, 0xed, 0xc3, 0x7e, 0xdf, 0x32, 0x57, 0x4d, 0xae, 0x41, 0x59,
	0xf6, 0xcd, 0x33, 0x73, 0xe6, 0x9c, 0x73, 0x67, 0x40, 0x00, 0x51, 0xf1, 0x6c, 0x5d, 0xd5, 0xa5,
	0x2c, 0xd1, 0xad, 0xea, 0x92, 0x9d, 0x32, 0x59, 0xed, 0xa2, 0x47, 0x03, 0xac, 0x6d, 0xc5, 0x33,
	0x9c, 0x83, 0x4d, 0x8c, 0x9b, 0x9c, 0xf9, 0x46, 0x68, 0xac, 0xac, 0x64, 0x46, 0x30, 0x66, 0xf8,
	0x09, 0x1c, 0x21, 0xcb, 0x9a, 0x53, 0x67, 0xa2, 0x3a, 0xb6, 0xc2, 0x31, 0xa3, 0x99, 0xfb, 0xbc,
	0x60, 0xd4, 0x31, 0x9b, 0x19, 0x82, 0x31, 0x43, 0x04, 0xab, 0x48, 0x0f, 0xdc, 0xb7, 0x42, 0x63,
	0xe5, 0x26, 0xea, 0x9b, 0x6a, 0xa2, 0xac, 0xa5, 0x3f, 0x55, 0x4c, 0xf5, 0x4d, 0xda, 0x29, 0x3b,
	0xe4, 0x05, 0x29, 0xcc, 0x1a, 0x6d, 0x85, 0x63, 0x86, 0x01, 0x38, 0x59, 0x59, 0x48, 0x5e, 0x48,
	0xe1, 0xdb, 0xa1, 0xb9, 0x72, 0x93, 0x33, 0x8e, 0xee, 0xc0, 0xa1, 0xcc, 0x71, 0x71, 0x5b, 0x8e,
	0xe7, 0xee, 0x32, 0x4c, 0x06, 0x32, 0x98, 0x5a, 0x06, 0xdd, 0xc8, 0xba, 0x32, 0xfa, 0x06, 0xf0,
	0x97, 0x31, 0xf2, 0x4a, 0xf8, 0x11, 0x3f, 0x83, 0x45, 0xda, 0xca, 0xc7, 0xdb, 0xbc, 0x5b, 0x9f,
	0xb7, 0xb8, 0x56, 0x0c, 0xd5, 0x8c, 0xbe, 0x68, 0x23, 0x62, 0x34, 0x5d, 0xb4, 0x01, 0xef, 0x3f,
	0xcb, 0xe5, 0xab, 0xa4, 0x97, 0xfa, 0x8c, 0x40, 0x1f, 0xec, 0x53, 0xc5, 0x52, 0xc9, 0x1b, 0xed,
	0x69, 0xd2, 0x41, 0xca, 0xf0, 0x8f, 0xef, 0x3b, 0xed, 0xd1, 0x0c, 0x0b, 0x8d, 0xa6, 0xe4, 0x18,
	0xdf, 0x73, 0x4d, 0xae, 0x85, 0xd1, 0x02, 0xbc, 0x84, 0xa7, 0xec, 0x45, 0xbd, 0x1f, 0x3a, 0x4f,
	0xe0, 0xb2, 0xf7, 0xa6, 0xf7, 0x57, 0x6f, 0xa2, 0xe3, 0xb5, 0xef, 0x7a, 0x0b, 0x6f, 0xba, 0x39,
	0x91, 0xf0, 0x63, 0xf4, 0xab, 0x87, 0x05, 0x7e, 0x85, 0x29, 0xf1, 0x84, 0x6f, 0x84, 0xe6, 0x98,
	0x52, 0xc3, 0xd8, 0x3c, 0x4d, 0xc0, 0xa3, 0xda, 0x96, 0xd7, 0x0f, 0x79, 0xc6, 0xf1, 0x27, 0xd8,
	0xed, 0x35, 0xf0, 0x83, 0x36, 0x76, 0x39, 0x6a, 0x30, 0x58, 0x16, 0xf8, 0x1b, 0x9c, 0x6e, 0xd7,
	0xf8, 0x51, 0xa3, 0x68, 0x47, 0x0b, 0x86, 0xeb, 0x82, 0x4c, 0xdb, 0xbd, 0xf6, 0x4c, 0x2f, 0x27,
	0x09, 0x06, 0xcb, 0xca, 0xb4, 0x7b, 0x78, 0xcf, 0x54, 0xdb, 0x7e, 0x30, 0x5c, 0x17, 0xf8, 0x07,
	0xdc, 0xf3, 0xd2, 0x70, 0x3e, 0x40, 0xa2, 0xd5, 0x06, 0x23, 0x0d, 0xb1, 0x9b, 0xa9, 0x3f, 0xc3,
	0xf7, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x38, 0x3b, 0xc6, 0x27, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SpecServiceClient is the client API for SpecService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpecServiceClient interface {
	AddSpec(ctx context.Context, in *AddSpecReq, opts ...grpc.CallOption) (*AddSpecRes, error)
	EditSpec(ctx context.Context, in *EditSpecReq, opts ...grpc.CallOption) (*EditSpecRes, error)
	DelSpec(ctx context.Context, in *DelSpecReq, opts ...grpc.CallOption) (*DelSpecRes, error)
	ReadSpec(ctx context.Context, in *ReadSpecReq, opts ...grpc.CallOption) (*ReadSpecRes, error)
	ReadSpecs(ctx context.Context, in *ReadSpecsReq, opts ...grpc.CallOption) (*ReadSpecsRes, error)
}

type specServiceClient struct {
	cc *grpc.ClientConn
}

func NewSpecServiceClient(cc *grpc.ClientConn) SpecServiceClient {
	return &specServiceClient{cc}
}

func (c *specServiceClient) AddSpec(ctx context.Context, in *AddSpecReq, opts ...grpc.CallOption) (*AddSpecRes, error) {
	out := new(AddSpecRes)
	err := c.cc.Invoke(ctx, "/productpb.SpecService/AddSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specServiceClient) EditSpec(ctx context.Context, in *EditSpecReq, opts ...grpc.CallOption) (*EditSpecRes, error) {
	out := new(EditSpecRes)
	err := c.cc.Invoke(ctx, "/productpb.SpecService/EditSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specServiceClient) DelSpec(ctx context.Context, in *DelSpecReq, opts ...grpc.CallOption) (*DelSpecRes, error) {
	out := new(DelSpecRes)
	err := c.cc.Invoke(ctx, "/productpb.SpecService/DelSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specServiceClient) ReadSpec(ctx context.Context, in *ReadSpecReq, opts ...grpc.CallOption) (*ReadSpecRes, error) {
	out := new(ReadSpecRes)
	err := c.cc.Invoke(ctx, "/productpb.SpecService/ReadSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specServiceClient) ReadSpecs(ctx context.Context, in *ReadSpecsReq, opts ...grpc.CallOption) (*ReadSpecsRes, error) {
	out := new(ReadSpecsRes)
	err := c.cc.Invoke(ctx, "/productpb.SpecService/ReadSpecs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpecServiceServer is the server API for SpecService service.
type SpecServiceServer interface {
	AddSpec(context.Context, *AddSpecReq) (*AddSpecRes, error)
	EditSpec(context.Context, *EditSpecReq) (*EditSpecRes, error)
	DelSpec(context.Context, *DelSpecReq) (*DelSpecRes, error)
	ReadSpec(context.Context, *ReadSpecReq) (*ReadSpecRes, error)
	ReadSpecs(context.Context, *ReadSpecsReq) (*ReadSpecsRes, error)
}

func RegisterSpecServiceServer(s *grpc.Server, srv SpecServiceServer) {
	s.RegisterService(&_SpecService_serviceDesc, srv)
}

func _SpecService_AddSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSpecReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecServiceServer).AddSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.SpecService/AddSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecServiceServer).AddSpec(ctx, req.(*AddSpecReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecService_EditSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditSpecReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecServiceServer).EditSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.SpecService/EditSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecServiceServer).EditSpec(ctx, req.(*EditSpecReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecService_DelSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelSpecReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecServiceServer).DelSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.SpecService/DelSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecServiceServer).DelSpec(ctx, req.(*DelSpecReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecService_ReadSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadSpecReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecServiceServer).ReadSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.SpecService/ReadSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecServiceServer).ReadSpec(ctx, req.(*ReadSpecReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecService_ReadSpecs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadSpecsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecServiceServer).ReadSpecs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.SpecService/ReadSpecs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecServiceServer).ReadSpecs(ctx, req.(*ReadSpecsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpecService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "productpb.SpecService",
	HandlerType: (*SpecServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSpec",
			Handler:    _SpecService_AddSpec_Handler,
		},
		{
			MethodName: "EditSpec",
			Handler:    _SpecService_EditSpec_Handler,
		},
		{
			MethodName: "DelSpec",
			Handler:    _SpecService_DelSpec_Handler,
		},
		{
			MethodName: "ReadSpec",
			Handler:    _SpecService_ReadSpec_Handler,
		},
		{
			MethodName: "ReadSpecs",
			Handler:    _SpecService_ReadSpecs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spec.proto",
}
