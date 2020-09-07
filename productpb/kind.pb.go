// Code generated by protoc-gen-go. DO NOT EDIT.
// source: productpb/kind.proto

package productpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	basepb "github.com/shinmigo/pb/basepb"
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

type ParamRel struct {
	ParamId              uint64             `protobuf:"varint,1,opt,name=param_id,json=paramId,proto3" json:"param_id"`
	Name                 string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Type                 ParamType          `protobuf:"varint,3,opt,name=type,proto3,enum=productpb.ParamType" json:"type"`
	Contents             []*ParamRelContent `protobuf:"bytes,4,rep,name=contents,proto3" json:"contents"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ParamRel) Reset()         { *m = ParamRel{} }
func (m *ParamRel) String() string { return proto.CompactTextString(m) }
func (*ParamRel) ProtoMessage()    {}
func (*ParamRel) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{0}
}

func (m *ParamRel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParamRel.Unmarshal(m, b)
}
func (m *ParamRel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParamRel.Marshal(b, m, deterministic)
}
func (m *ParamRel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamRel.Merge(m, src)
}
func (m *ParamRel) XXX_Size() int {
	return xxx_messageInfo_ParamRel.Size(m)
}
func (m *ParamRel) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamRel.DiscardUnknown(m)
}

var xxx_messageInfo_ParamRel proto.InternalMessageInfo

func (m *ParamRel) GetParamId() uint64 {
	if m != nil {
		return m.ParamId
	}
	return 0
}

func (m *ParamRel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ParamRel) GetType() ParamType {
	if m != nil {
		return m.Type
	}
	return ParamType_Placeholder
}

func (m *ParamRel) GetContents() []*ParamRelContent {
	if m != nil {
		return m.Contents
	}
	return nil
}

type ParamRelContent struct {
	ParamValueId         uint64   `protobuf:"varint,1,opt,name=param_value_id,json=paramValueId,proto3" json:"param_value_id"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParamRelContent) Reset()         { *m = ParamRelContent{} }
func (m *ParamRelContent) String() string { return proto.CompactTextString(m) }
func (*ParamRelContent) ProtoMessage()    {}
func (*ParamRelContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{1}
}

func (m *ParamRelContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParamRelContent.Unmarshal(m, b)
}
func (m *ParamRelContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParamRelContent.Marshal(b, m, deterministic)
}
func (m *ParamRelContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamRelContent.Merge(m, src)
}
func (m *ParamRelContent) XXX_Size() int {
	return xxx_messageInfo_ParamRelContent.Size(m)
}
func (m *ParamRelContent) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamRelContent.DiscardUnknown(m)
}

var xxx_messageInfo_ParamRelContent proto.InternalMessageInfo

func (m *ParamRelContent) GetParamValueId() uint64 {
	if m != nil {
		return m.ParamValueId
	}
	return 0
}

func (m *ParamRelContent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type SpecRel struct {
	SpecId               uint64            `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Contents             []*SpecRelContent `protobuf:"bytes,3,rep,name=contents,proto3" json:"contents"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SpecRel) Reset()         { *m = SpecRel{} }
func (m *SpecRel) String() string { return proto.CompactTextString(m) }
func (*SpecRel) ProtoMessage()    {}
func (*SpecRel) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{2}
}

func (m *SpecRel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecRel.Unmarshal(m, b)
}
func (m *SpecRel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecRel.Marshal(b, m, deterministic)
}
func (m *SpecRel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecRel.Merge(m, src)
}
func (m *SpecRel) XXX_Size() int {
	return xxx_messageInfo_SpecRel.Size(m)
}
func (m *SpecRel) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecRel.DiscardUnknown(m)
}

var xxx_messageInfo_SpecRel proto.InternalMessageInfo

func (m *SpecRel) GetSpecId() uint64 {
	if m != nil {
		return m.SpecId
	}
	return 0
}

func (m *SpecRel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SpecRel) GetContents() []*SpecRelContent {
	if m != nil {
		return m.Contents
	}
	return nil
}

type SpecRelContent struct {
	SpecValueId          uint64   `protobuf:"varint,1,opt,name=spec_value_id,json=specValueId,proto3" json:"spec_value_id"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecRelContent) Reset()         { *m = SpecRelContent{} }
func (m *SpecRelContent) String() string { return proto.CompactTextString(m) }
func (*SpecRelContent) ProtoMessage()    {}
func (*SpecRelContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{3}
}

func (m *SpecRelContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecRelContent.Unmarshal(m, b)
}
func (m *SpecRelContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecRelContent.Marshal(b, m, deterministic)
}
func (m *SpecRelContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecRelContent.Merge(m, src)
}
func (m *SpecRelContent) XXX_Size() int {
	return xxx_messageInfo_SpecRelContent.Size(m)
}
func (m *SpecRelContent) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecRelContent.DiscardUnknown(m)
}

var xxx_messageInfo_SpecRelContent proto.InternalMessageInfo

func (m *SpecRelContent) GetSpecValueId() uint64 {
	if m != nil {
		return m.SpecValueId
	}
	return 0
}

func (m *SpecRelContent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Kind struct {
	KindId               uint64   `protobuf:"varint,1,opt,name=kind_id,json=kindId,proto3" json:"kind_id"`
	StoreId              uint64   `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	AdminId              uint64   `protobuf:"varint,4,opt,name=admin_id,json=adminId,proto3" json:"admin_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Kind) Reset()         { *m = Kind{} }
func (m *Kind) String() string { return proto.CompactTextString(m) }
func (*Kind) ProtoMessage()    {}
func (*Kind) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{4}
}

func (m *Kind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Kind.Unmarshal(m, b)
}
func (m *Kind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Kind.Marshal(b, m, deterministic)
}
func (m *Kind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Kind.Merge(m, src)
}
func (m *Kind) XXX_Size() int {
	return xxx_messageInfo_Kind.Size(m)
}
func (m *Kind) XXX_DiscardUnknown() {
	xxx_messageInfo_Kind.DiscardUnknown(m)
}

var xxx_messageInfo_Kind proto.InternalMessageInfo

func (m *Kind) GetKindId() uint64 {
	if m != nil {
		return m.KindId
	}
	return 0
}

func (m *Kind) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Kind) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Kind) GetAdminId() uint64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

type KindDetail struct {
	KindId               uint64      `protobuf:"varint,1,opt,name=kind_id,json=kindId,proto3" json:"kind_id"`
	StoreId              uint64      `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id"`
	Name                 string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	ParamQty             uint64      `protobuf:"varint,4,opt,name=param_qty,json=paramQty,proto3" json:"param_qty"`
	SpecQty              uint64      `protobuf:"varint,5,opt,name=spec_qty,json=specQty,proto3" json:"spec_qty"`
	CreatedBy            uint64      `protobuf:"varint,6,opt,name=created_by,json=createdBy,proto3" json:"created_by"`
	UpdatedBy            uint64      `protobuf:"varint,7,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by"`
	CreatedAt            string      `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt            string      `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Params               []*ParamRel `protobuf:"bytes,10,rep,name=params,proto3" json:"params"`
	Specs                []*SpecRel  `protobuf:"bytes,11,rep,name=specs,proto3" json:"specs"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *KindDetail) Reset()         { *m = KindDetail{} }
func (m *KindDetail) String() string { return proto.CompactTextString(m) }
func (*KindDetail) ProtoMessage()    {}
func (*KindDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{5}
}

func (m *KindDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KindDetail.Unmarshal(m, b)
}
func (m *KindDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KindDetail.Marshal(b, m, deterministic)
}
func (m *KindDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KindDetail.Merge(m, src)
}
func (m *KindDetail) XXX_Size() int {
	return xxx_messageInfo_KindDetail.Size(m)
}
func (m *KindDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_KindDetail.DiscardUnknown(m)
}

var xxx_messageInfo_KindDetail proto.InternalMessageInfo

func (m *KindDetail) GetKindId() uint64 {
	if m != nil {
		return m.KindId
	}
	return 0
}

func (m *KindDetail) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *KindDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KindDetail) GetParamQty() uint64 {
	if m != nil {
		return m.ParamQty
	}
	return 0
}

func (m *KindDetail) GetSpecQty() uint64 {
	if m != nil {
		return m.SpecQty
	}
	return 0
}

func (m *KindDetail) GetCreatedBy() uint64 {
	if m != nil {
		return m.CreatedBy
	}
	return 0
}

func (m *KindDetail) GetUpdatedBy() uint64 {
	if m != nil {
		return m.UpdatedBy
	}
	return 0
}

func (m *KindDetail) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *KindDetail) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *KindDetail) GetParams() []*ParamRel {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *KindDetail) GetSpecs() []*SpecRel {
	if m != nil {
		return m.Specs
	}
	return nil
}

type DelKindReq struct {
	KindId               uint64   `protobuf:"varint,1,opt,name=kind_id,json=kindId,proto3" json:"kind_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelKindReq) Reset()         { *m = DelKindReq{} }
func (m *DelKindReq) String() string { return proto.CompactTextString(m) }
func (*DelKindReq) ProtoMessage()    {}
func (*DelKindReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{6}
}

func (m *DelKindReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelKindReq.Unmarshal(m, b)
}
func (m *DelKindReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelKindReq.Marshal(b, m, deterministic)
}
func (m *DelKindReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelKindReq.Merge(m, src)
}
func (m *DelKindReq) XXX_Size() int {
	return xxx_messageInfo_DelKindReq.Size(m)
}
func (m *DelKindReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelKindReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelKindReq proto.InternalMessageInfo

func (m *DelKindReq) GetKindId() uint64 {
	if m != nil {
		return m.KindId
	}
	return 0
}

type ListKindReq struct {
	Page                 uint64   `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	PageSize             uint64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id                   uint64   `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListKindReq) Reset()         { *m = ListKindReq{} }
func (m *ListKindReq) String() string { return proto.CompactTextString(m) }
func (*ListKindReq) ProtoMessage()    {}
func (*ListKindReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{7}
}

func (m *ListKindReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKindReq.Unmarshal(m, b)
}
func (m *ListKindReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKindReq.Marshal(b, m, deterministic)
}
func (m *ListKindReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKindReq.Merge(m, src)
}
func (m *ListKindReq) XXX_Size() int {
	return xxx_messageInfo_ListKindReq.Size(m)
}
func (m *ListKindReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKindReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListKindReq proto.InternalMessageInfo

func (m *ListKindReq) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListKindReq) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListKindReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ListKindReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListKindRes struct {
	Total                uint64        `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	Kinds                []*KindDetail `protobuf:"bytes,2,rep,name=kinds,proto3" json:"kinds"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListKindRes) Reset()         { *m = ListKindRes{} }
func (m *ListKindRes) String() string { return proto.CompactTextString(m) }
func (*ListKindRes) ProtoMessage()    {}
func (*ListKindRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{8}
}

func (m *ListKindRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKindRes.Unmarshal(m, b)
}
func (m *ListKindRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKindRes.Marshal(b, m, deterministic)
}
func (m *ListKindRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKindRes.Merge(m, src)
}
func (m *ListKindRes) XXX_Size() int {
	return xxx_messageInfo_ListKindRes.Size(m)
}
func (m *ListKindRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKindRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListKindRes proto.InternalMessageInfo

func (m *ListKindRes) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListKindRes) GetKinds() []*KindDetail {
	if m != nil {
		return m.Kinds
	}
	return nil
}

type BindParamReq struct {
	KindId               uint64   `protobuf:"varint,1,opt,name=kind_id,json=kindId,proto3" json:"kind_id"`
	ParamIds             []uint64 `protobuf:"varint,2,rep,packed,name=param_ids,json=paramIds,proto3" json:"param_ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindParamReq) Reset()         { *m = BindParamReq{} }
func (m *BindParamReq) String() string { return proto.CompactTextString(m) }
func (*BindParamReq) ProtoMessage()    {}
func (*BindParamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{9}
}

func (m *BindParamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindParamReq.Unmarshal(m, b)
}
func (m *BindParamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindParamReq.Marshal(b, m, deterministic)
}
func (m *BindParamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindParamReq.Merge(m, src)
}
func (m *BindParamReq) XXX_Size() int {
	return xxx_messageInfo_BindParamReq.Size(m)
}
func (m *BindParamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BindParamReq.DiscardUnknown(m)
}

var xxx_messageInfo_BindParamReq proto.InternalMessageInfo

func (m *BindParamReq) GetKindId() uint64 {
	if m != nil {
		return m.KindId
	}
	return 0
}

func (m *BindParamReq) GetParamIds() []uint64 {
	if m != nil {
		return m.ParamIds
	}
	return nil
}

type BindSpecReq struct {
	KindId               uint64   `protobuf:"varint,1,opt,name=kind_id,json=kindId,proto3" json:"kind_id"`
	SpecIds              []uint64 `protobuf:"varint,2,rep,packed,name=spec_ids,json=specIds,proto3" json:"spec_ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindSpecReq) Reset()         { *m = BindSpecReq{} }
func (m *BindSpecReq) String() string { return proto.CompactTextString(m) }
func (*BindSpecReq) ProtoMessage()    {}
func (*BindSpecReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{10}
}

func (m *BindSpecReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindSpecReq.Unmarshal(m, b)
}
func (m *BindSpecReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindSpecReq.Marshal(b, m, deterministic)
}
func (m *BindSpecReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindSpecReq.Merge(m, src)
}
func (m *BindSpecReq) XXX_Size() int {
	return xxx_messageInfo_BindSpecReq.Size(m)
}
func (m *BindSpecReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BindSpecReq.DiscardUnknown(m)
}

var xxx_messageInfo_BindSpecReq proto.InternalMessageInfo

func (m *BindSpecReq) GetKindId() uint64 {
	if m != nil {
		return m.KindId
	}
	return 0
}

func (m *BindSpecReq) GetSpecIds() []uint64 {
	if m != nil {
		return m.SpecIds
	}
	return nil
}

type ListParamAndSpecReq struct {
	KindId               uint64   `protobuf:"varint,1,opt,name=kind_id,json=kindId,proto3" json:"kind_id"`
	StoreId              uint64   `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListParamAndSpecReq) Reset()         { *m = ListParamAndSpecReq{} }
func (m *ListParamAndSpecReq) String() string { return proto.CompactTextString(m) }
func (*ListParamAndSpecReq) ProtoMessage()    {}
func (*ListParamAndSpecReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9eb9475b40e833a, []int{11}
}

func (m *ListParamAndSpecReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListParamAndSpecReq.Unmarshal(m, b)
}
func (m *ListParamAndSpecReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListParamAndSpecReq.Marshal(b, m, deterministic)
}
func (m *ListParamAndSpecReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListParamAndSpecReq.Merge(m, src)
}
func (m *ListParamAndSpecReq) XXX_Size() int {
	return xxx_messageInfo_ListParamAndSpecReq.Size(m)
}
func (m *ListParamAndSpecReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListParamAndSpecReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListParamAndSpecReq proto.InternalMessageInfo

func (m *ListParamAndSpecReq) GetKindId() uint64 {
	if m != nil {
		return m.KindId
	}
	return 0
}

func (m *ListParamAndSpecReq) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func init() {
	proto.RegisterType((*ParamRel)(nil), "productpb.ParamRel")
	proto.RegisterType((*ParamRelContent)(nil), "productpb.ParamRelContent")
	proto.RegisterType((*SpecRel)(nil), "productpb.SpecRel")
	proto.RegisterType((*SpecRelContent)(nil), "productpb.SpecRelContent")
	proto.RegisterType((*Kind)(nil), "productpb.Kind")
	proto.RegisterType((*KindDetail)(nil), "productpb.KindDetail")
	proto.RegisterType((*DelKindReq)(nil), "productpb.DelKindReq")
	proto.RegisterType((*ListKindReq)(nil), "productpb.ListKindReq")
	proto.RegisterType((*ListKindRes)(nil), "productpb.ListKindRes")
	proto.RegisterType((*BindParamReq)(nil), "productpb.BindParamReq")
	proto.RegisterType((*BindSpecReq)(nil), "productpb.BindSpecReq")
	proto.RegisterType((*ListParamAndSpecReq)(nil), "productpb.ListParamAndSpecReq")
}

func init() { proto.RegisterFile("productpb/kind.proto", fileDescriptor_b9eb9475b40e833a) }

var fileDescriptor_b9eb9475b40e833a = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x1c, 0x27, 0x76, 0xc6, 0x25, 0x95, 0xb6, 0x2d, 0x75, 0x83, 0x90, 0x22, 0x0b, 0xa4,
	0x88, 0x4a, 0x09, 0x0a, 0x3f, 0x37, 0x0e, 0x09, 0x45, 0x28, 0x02, 0xa1, 0xd6, 0x45, 0x1c, 0xb8,
	0x54, 0xb6, 0x77, 0x95, 0xae, 0x48, 0xec, 0x6d, 0x76, 0x53, 0x29, 0x7d, 0x15, 0x9e, 0x86, 0x27,
	0xe1, 0x55, 0xd0, 0xfe, 0x38, 0xde, 0x26, 0x2d, 0xf4, 0xc0, 0x29, 0x9e, 0x99, 0x6f, 0x7e, 0xbe,
	0xf9, 0xc6, 0x31, 0xec, 0xb3, 0x45, 0x81, 0x97, 0x99, 0x60, 0xe9, 0xe0, 0x07, 0xcd, 0x71, 0x9f,
	0x2d, 0x0a, 0x51, 0xa0, 0xd6, 0xda, 0xdb, 0xd9, 0x4b, 0x13, 0x4e, 0x58, 0x3a, 0xd0, 0x3f, 0x3a,
	0xde, 0x39, 0xa8, 0xb2, 0x58, 0xb2, 0x48, 0xe6, 0xda, 0x1d, 0xfd, 0xac, 0x81, 0x7f, 0x2a, 0xed,
	0x98, 0xcc, 0xd0, 0x11, 0xf8, 0x2a, 0x76, 0x41, 0x71, 0x58, 0xeb, 0xd6, 0x7a, 0x6e, 0xec, 0x29,
	0x7b, 0x82, 0x11, 0x02, 0x37, 0x4f, 0xe6, 0x24, 0x74, 0xba, 0xb5, 0x5e, 0x2b, 0x56, 0xcf, 0xa8,
	0x07, 0xae, 0x58, 0x31, 0x12, 0xd6, 0xbb, 0xb5, 0x5e, 0x7b, 0xb8, 0xdf, 0x5f, 0x77, 0xe8, 0xab,
	0x8a, 0x5f, 0x57, 0x8c, 0xc4, 0x0a, 0x81, 0xde, 0x82, 0x9f, 0x15, 0xb9, 0x20, 0xb9, 0xe0, 0xa1,
	0xdb, 0xad, 0xf7, 0x82, 0x61, 0x67, 0x13, 0x1d, 0x93, 0xd9, 0x7b, 0x0d, 0x89, 0xd7, 0xd8, 0xe8,
	0x0c, 0x76, 0x37, 0x82, 0xe8, 0x19, 0xb4, 0xf5, 0x8c, 0xd7, 0xc9, 0x6c, 0x49, 0xaa, 0x49, 0x77,
	0x94, 0xf7, 0x9b, 0x74, 0x4e, 0x30, 0x0a, 0xc1, 0x33, 0x45, 0xcc, 0xc4, 0xa5, 0x19, 0xcd, 0xc1,
	0x3b, 0x67, 0x24, 0x93, 0x74, 0x0f, 0xc1, 0xe3, 0x8c, 0x64, 0x55, 0x8d, 0xa6, 0x34, 0xef, 0x21,
	0xfb, 0xc6, 0xa2, 0x50, 0x57, 0x14, 0x8e, 0x2c, 0x0a, 0xa6, 0xe4, 0x36, 0x83, 0x2f, 0xd0, 0xbe,
	0x1d, 0x43, 0x11, 0x3c, 0x52, 0x5d, 0x37, 0xe6, 0x0f, 0xa4, 0xf3, 0xdf, 0xe3, 0x53, 0x70, 0x3f,
	0xd1, 0x1c, 0xcb, 0xd9, 0xa5, 0xf8, 0xd6, 0xec, 0xd2, 0x9c, 0x60, 0xa9, 0x21, 0x17, 0xc5, 0x42,
	0x55, 0x76, 0xb4, 0x86, 0xca, 0xb6, 0x68, 0xd5, 0x2d, 0x5a, 0x47, 0xe0, 0x27, 0x78, 0x4e, 0x73,
	0x09, 0x77, 0x35, 0x5c, 0xd9, 0x13, 0x1c, 0xfd, 0x76, 0x00, 0x64, 0xaf, 0x13, 0x22, 0x12, 0x3a,
	0xfb, 0x6f, 0x1d, 0x9f, 0x40, 0x4b, 0x0b, 0x78, 0x25, 0x56, 0xa6, 0xa5, 0xbe, 0xba, 0x33, 0xb1,
	0x52, 0xb5, 0xe4, 0x72, 0x64, 0xac, 0x61, 0x6a, 0x31, 0x92, 0xc9, 0xd0, 0x53, 0x80, 0x6c, 0x41,
	0x12, 0x41, 0xf0, 0x45, 0xba, 0x0a, 0x9b, 0x2a, 0xd8, 0x32, 0x9e, 0xb1, 0x0a, 0x2f, 0x19, 0x2e,
	0xc3, 0x9e, 0x0e, 0x1b, 0xcf, 0xf8, 0x56, 0x76, 0x22, 0x42, 0x5f, 0xcd, 0x53, 0x66, 0x8f, 0x84,
	0x9d, 0x9d, 0x88, 0xb0, 0xa5, 0xc3, 0xc6, 0x33, 0x12, 0xe8, 0x18, 0x9a, 0x6a, 0x44, 0x1e, 0x82,
	0x92, 0x7e, 0xef, 0x8e, 0xeb, 0x8d, 0x0d, 0x04, 0xf5, 0xa0, 0x21, 0x67, 0xe6, 0x61, 0xa0, 0xb0,
	0x68, 0xfb, 0x4c, 0x62, 0x0d, 0x88, 0x9e, 0x03, 0x9c, 0x90, 0x99, 0xdc, 0x71, 0x4c, 0xae, 0xee,
	0x5d, 0x70, 0x94, 0x42, 0xf0, 0x99, 0x72, 0x51, 0xe2, 0x10, 0xb8, 0x2c, 0x99, 0x12, 0x03, 0x52,
	0xcf, 0x7a, 0xa9, 0x53, 0x72, 0xc1, 0xe9, 0x0d, 0x31, 0x22, 0xf8, 0xd2, 0x71, 0x4e, 0x6f, 0x08,
	0x6a, 0x83, 0x43, 0xb1, 0xd2, 0xc0, 0x8d, 0x1d, 0x5a, 0xa9, 0xe2, 0x56, 0xaa, 0x44, 0xa7, 0x76,
	0x0f, 0x8e, 0xf6, 0xa1, 0x21, 0x0a, 0x91, 0xcc, 0x4c, 0x13, 0x6d, 0xa0, 0x63, 0x68, 0xc8, 0x91,
	0x78, 0xe8, 0x28, 0x66, 0x07, 0x16, 0xb3, 0xea, 0x50, 0x62, 0x8d, 0x89, 0x4e, 0x60, 0x67, 0x4c,
	0x73, 0x6c, 0xd6, 0x73, 0x3f, 0xbd, 0xea, 0x20, 0xa8, 0xa9, 0x5c, 0x1e, 0xc4, 0x04, 0xf3, 0x68,
	0x04, 0x81, 0xac, 0xa2, 0x17, 0x77, 0xf5, 0xf7, 0x23, 0xd4, 0xef, 0x72, 0x59, 0xc3, 0xd3, 0x2f,
	0x33, 0x8f, 0x26, 0xb0, 0x27, 0xa9, 0xa9, 0x41, 0x46, 0x0f, 0x2c, 0x75, 0xf7, 0x3d, 0x0f, 0x7f,
	0x39, 0x10, 0x48, 0xa6, 0xe7, 0x64, 0x71, 0x4d, 0x33, 0x82, 0x5e, 0x80, 0x37, 0xc2, 0x58, 0xbd,
	0x90, 0xbb, 0x1b, 0xcb, 0xe8, 0xb4, 0xfb, 0xe6, 0xff, 0x77, 0x94, 0xaf, 0xe4, 0x4a, 0x8f, 0xc1,
	0xff, 0x80, 0xa9, 0x78, 0x18, 0xf8, 0x25, 0x78, 0xe6, 0x32, 0x90, 0xbd, 0xe5, 0xea, 0x5a, 0xb6,
	0x32, 0xde, 0x41, 0xf0, 0x91, 0xa8, 0xea, 0x92, 0x2c, 0x7a, 0x6c, 0x65, 0x59, 0xc7, 0xd3, 0xb9,
	0xdb, 0xcf, 0xd1, 0x6b, 0x68, 0xad, 0xd5, 0x42, 0x87, 0x16, 0xc8, 0xd6, 0x70, 0xab, 0xe9, 0x10,
	0xfc, 0x52, 0x9d, 0x5b, 0x1d, 0x2d, 0xc9, 0x36, 0x73, 0xc6, 0xd1, 0xf7, 0xee, 0x94, 0x8a, 0xcb,
	0x65, 0xda, 0xcf, 0x8a, 0xf9, 0x80, 0x5f, 0xd2, 0x7c, 0x4e, 0xa7, 0xc5, 0x40, 0x7e, 0x97, 0xca,
	0xfc, 0xb4, 0xa9, 0x3e, 0x4e, 0xaf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x44, 0x43, 0x72, 0x48,
	0xeb, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KindServiceClient is the client API for KindService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KindServiceClient interface {
	AddKind(ctx context.Context, in *Kind, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	EditKind(ctx context.Context, in *Kind, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	DelKind(ctx context.Context, in *DelKindReq, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	GetKindList(ctx context.Context, in *ListKindReq, opts ...grpc.CallOption) (*ListKindRes, error)
	BindParam(ctx context.Context, in *BindParamReq, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	BindSpec(ctx context.Context, in *BindSpecReq, opts ...grpc.CallOption) (*basepb.AnyRes, error)
}

type kindServiceClient struct {
	cc *grpc.ClientConn
}

func NewKindServiceClient(cc *grpc.ClientConn) KindServiceClient {
	return &kindServiceClient{cc}
}

func (c *kindServiceClient) AddKind(ctx context.Context, in *Kind, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.KindService/AddKind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kindServiceClient) EditKind(ctx context.Context, in *Kind, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.KindService/EditKind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kindServiceClient) DelKind(ctx context.Context, in *DelKindReq, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.KindService/DelKind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kindServiceClient) GetKindList(ctx context.Context, in *ListKindReq, opts ...grpc.CallOption) (*ListKindRes, error) {
	out := new(ListKindRes)
	err := c.cc.Invoke(ctx, "/productpb.KindService/GetKindList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kindServiceClient) BindParam(ctx context.Context, in *BindParamReq, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.KindService/BindParam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kindServiceClient) BindSpec(ctx context.Context, in *BindSpecReq, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.KindService/BindSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KindServiceServer is the server API for KindService service.
type KindServiceServer interface {
	AddKind(context.Context, *Kind) (*basepb.AnyRes, error)
	EditKind(context.Context, *Kind) (*basepb.AnyRes, error)
	DelKind(context.Context, *DelKindReq) (*basepb.AnyRes, error)
	GetKindList(context.Context, *ListKindReq) (*ListKindRes, error)
	BindParam(context.Context, *BindParamReq) (*basepb.AnyRes, error)
	BindSpec(context.Context, *BindSpecReq) (*basepb.AnyRes, error)
}

func RegisterKindServiceServer(s *grpc.Server, srv KindServiceServer) {
	s.RegisterService(&_KindService_serviceDesc, srv)
}

func _KindService_AddKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Kind)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KindServiceServer).AddKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.KindService/AddKind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KindServiceServer).AddKind(ctx, req.(*Kind))
	}
	return interceptor(ctx, in, info, handler)
}

func _KindService_EditKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Kind)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KindServiceServer).EditKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.KindService/EditKind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KindServiceServer).EditKind(ctx, req.(*Kind))
	}
	return interceptor(ctx, in, info, handler)
}

func _KindService_DelKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelKindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KindServiceServer).DelKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.KindService/DelKind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KindServiceServer).DelKind(ctx, req.(*DelKindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KindService_GetKindList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KindServiceServer).GetKindList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.KindService/GetKindList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KindServiceServer).GetKindList(ctx, req.(*ListKindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KindService_BindParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindParamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KindServiceServer).BindParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.KindService/BindParam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KindServiceServer).BindParam(ctx, req.(*BindParamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KindService_BindSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindSpecReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KindServiceServer).BindSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.KindService/BindSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KindServiceServer).BindSpec(ctx, req.(*BindSpecReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _KindService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "productpb.KindService",
	HandlerType: (*KindServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddKind",
			Handler:    _KindService_AddKind_Handler,
		},
		{
			MethodName: "EditKind",
			Handler:    _KindService_EditKind_Handler,
		},
		{
			MethodName: "DelKind",
			Handler:    _KindService_DelKind_Handler,
		},
		{
			MethodName: "GetKindList",
			Handler:    _KindService_GetKindList_Handler,
		},
		{
			MethodName: "BindParam",
			Handler:    _KindService_BindParam_Handler,
		},
		{
			MethodName: "BindSpec",
			Handler:    _KindService_BindSpec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "productpb/kind.proto",
}
