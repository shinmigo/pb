// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shoppb/banner.proto

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

type BannerAdStatus int32

const (
	BannerAdStatus_BannerAdStatusPlaceholder BannerAdStatus = 0
	BannerAdStatus_Enabled                   BannerAdStatus = 1
	BannerAdStatus_Disabled                  BannerAdStatus = 2
)

var BannerAdStatus_name = map[int32]string{
	0: "BannerAdStatusPlaceholder",
	1: "Enabled",
	2: "Disabled",
}

var BannerAdStatus_value = map[string]int32{
	"BannerAdStatusPlaceholder": 0,
	"Enabled":                   1,
	"Disabled":                  2,
}

func (x BannerAdStatus) String() string {
	return proto.EnumName(BannerAdStatus_name, int32(x))
}

func (BannerAdStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fb5d95530256fb08, []int{0}
}

type BannerAd struct {
	Id                   uint64         `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	EleType              uint32         `protobuf:"varint,2,opt,name=ele_type,json=eleType,proto3" json:"ele_type"`
	ImageUrl             string         `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
	RedirectUrl          string         `protobuf:"bytes,4,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url"`
	Sort                 uint32         `protobuf:"varint,5,opt,name=sort,proto3" json:"sort"`
	Status               BannerAdStatus `protobuf:"varint,6,opt,name=status,proto3,enum=shoppb.BannerAdStatus" json:"status"`
	AdminId              uint64         `protobuf:"varint,7,opt,name=admin_id,json=adminId,proto3" json:"admin_id"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BannerAd) Reset()         { *m = BannerAd{} }
func (m *BannerAd) String() string { return proto.CompactTextString(m) }
func (*BannerAd) ProtoMessage()    {}
func (*BannerAd) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb5d95530256fb08, []int{0}
}

func (m *BannerAd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BannerAd.Unmarshal(m, b)
}
func (m *BannerAd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BannerAd.Marshal(b, m, deterministic)
}
func (m *BannerAd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannerAd.Merge(m, src)
}
func (m *BannerAd) XXX_Size() int {
	return xxx_messageInfo_BannerAd.Size(m)
}
func (m *BannerAd) XXX_DiscardUnknown() {
	xxx_messageInfo_BannerAd.DiscardUnknown(m)
}

var xxx_messageInfo_BannerAd proto.InternalMessageInfo

func (m *BannerAd) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BannerAd) GetEleType() uint32 {
	if m != nil {
		return m.EleType
	}
	return 0
}

func (m *BannerAd) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *BannerAd) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

func (m *BannerAd) GetSort() uint32 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *BannerAd) GetStatus() BannerAdStatus {
	if m != nil {
		return m.Status
	}
	return BannerAdStatus_BannerAdStatusPlaceholder
}

func (m *BannerAd) GetAdminId() uint64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

type EditBannerAdStatusReq struct {
	Id                   []uint64       `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id"`
	Status               BannerAdStatus `protobuf:"varint,2,opt,name=status,proto3,enum=shoppb.BannerAdStatus" json:"status"`
	AdminId              uint64         `protobuf:"varint,3,opt,name=admin_id,json=adminId,proto3" json:"admin_id"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EditBannerAdStatusReq) Reset()         { *m = EditBannerAdStatusReq{} }
func (m *EditBannerAdStatusReq) String() string { return proto.CompactTextString(m) }
func (*EditBannerAdStatusReq) ProtoMessage()    {}
func (*EditBannerAdStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb5d95530256fb08, []int{1}
}

func (m *EditBannerAdStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditBannerAdStatusReq.Unmarshal(m, b)
}
func (m *EditBannerAdStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditBannerAdStatusReq.Marshal(b, m, deterministic)
}
func (m *EditBannerAdStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditBannerAdStatusReq.Merge(m, src)
}
func (m *EditBannerAdStatusReq) XXX_Size() int {
	return xxx_messageInfo_EditBannerAdStatusReq.Size(m)
}
func (m *EditBannerAdStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EditBannerAdStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_EditBannerAdStatusReq proto.InternalMessageInfo

func (m *EditBannerAdStatusReq) GetId() []uint64 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *EditBannerAdStatusReq) GetStatus() BannerAdStatus {
	if m != nil {
		return m.Status
	}
	return BannerAdStatus_BannerAdStatusPlaceholder
}

func (m *EditBannerAdStatusReq) GetAdminId() uint64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

type DelBannerAdReq struct {
	Id                   []uint64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelBannerAdReq) Reset()         { *m = DelBannerAdReq{} }
func (m *DelBannerAdReq) String() string { return proto.CompactTextString(m) }
func (*DelBannerAdReq) ProtoMessage()    {}
func (*DelBannerAdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb5d95530256fb08, []int{2}
}

func (m *DelBannerAdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelBannerAdReq.Unmarshal(m, b)
}
func (m *DelBannerAdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelBannerAdReq.Marshal(b, m, deterministic)
}
func (m *DelBannerAdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelBannerAdReq.Merge(m, src)
}
func (m *DelBannerAdReq) XXX_Size() int {
	return xxx_messageInfo_DelBannerAdReq.Size(m)
}
func (m *DelBannerAdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelBannerAdReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelBannerAdReq proto.InternalMessageInfo

func (m *DelBannerAdReq) GetId() []uint64 {
	if m != nil {
		return m.Id
	}
	return nil
}

type ListBannerAdReq struct {
	Page                 uint64         `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	PageSize             uint64         `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id                   uint64         `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	EleType              uint32         `protobuf:"varint,4,opt,name=ele_type,json=eleType,proto3" json:"ele_type"`
	Status               BannerAdStatus `protobuf:"varint,5,opt,name=status,proto3,enum=shoppb.BannerAdStatus" json:"status"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListBannerAdReq) Reset()         { *m = ListBannerAdReq{} }
func (m *ListBannerAdReq) String() string { return proto.CompactTextString(m) }
func (*ListBannerAdReq) ProtoMessage()    {}
func (*ListBannerAdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb5d95530256fb08, []int{3}
}

func (m *ListBannerAdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBannerAdReq.Unmarshal(m, b)
}
func (m *ListBannerAdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBannerAdReq.Marshal(b, m, deterministic)
}
func (m *ListBannerAdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBannerAdReq.Merge(m, src)
}
func (m *ListBannerAdReq) XXX_Size() int {
	return xxx_messageInfo_ListBannerAdReq.Size(m)
}
func (m *ListBannerAdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBannerAdReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListBannerAdReq proto.InternalMessageInfo

func (m *ListBannerAdReq) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListBannerAdReq) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListBannerAdReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ListBannerAdReq) GetEleType() uint32 {
	if m != nil {
		return m.EleType
	}
	return 0
}

func (m *ListBannerAdReq) GetStatus() BannerAdStatus {
	if m != nil {
		return m.Status
	}
	return BannerAdStatus_BannerAdStatusPlaceholder
}

type BannerAdDetail struct {
	Id                   uint64         `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	EleType              uint32         `protobuf:"varint,2,opt,name=ele_type,json=eleType,proto3" json:"ele_type"`
	ImageUrl             string         `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
	RedirectUrl          string         `protobuf:"bytes,4,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url"`
	Sort                 uint32         `protobuf:"varint,5,opt,name=sort,proto3" json:"sort"`
	Status               BannerAdStatus `protobuf:"varint,6,opt,name=status,proto3,enum=shoppb.BannerAdStatus" json:"status"`
	CreatedBy            uint64         `protobuf:"varint,7,opt,name=created_by,json=createdBy,proto3" json:"created_by"`
	UpdatedBy            uint64         `protobuf:"varint,8,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by"`
	CreatedAt            string         `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt            string         `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BannerAdDetail) Reset()         { *m = BannerAdDetail{} }
func (m *BannerAdDetail) String() string { return proto.CompactTextString(m) }
func (*BannerAdDetail) ProtoMessage()    {}
func (*BannerAdDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb5d95530256fb08, []int{4}
}

func (m *BannerAdDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BannerAdDetail.Unmarshal(m, b)
}
func (m *BannerAdDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BannerAdDetail.Marshal(b, m, deterministic)
}
func (m *BannerAdDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannerAdDetail.Merge(m, src)
}
func (m *BannerAdDetail) XXX_Size() int {
	return xxx_messageInfo_BannerAdDetail.Size(m)
}
func (m *BannerAdDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_BannerAdDetail.DiscardUnknown(m)
}

var xxx_messageInfo_BannerAdDetail proto.InternalMessageInfo

func (m *BannerAdDetail) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BannerAdDetail) GetEleType() uint32 {
	if m != nil {
		return m.EleType
	}
	return 0
}

func (m *BannerAdDetail) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *BannerAdDetail) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

func (m *BannerAdDetail) GetSort() uint32 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *BannerAdDetail) GetStatus() BannerAdStatus {
	if m != nil {
		return m.Status
	}
	return BannerAdStatus_BannerAdStatusPlaceholder
}

func (m *BannerAdDetail) GetCreatedBy() uint64 {
	if m != nil {
		return m.CreatedBy
	}
	return 0
}

func (m *BannerAdDetail) GetUpdatedBy() uint64 {
	if m != nil {
		return m.UpdatedBy
	}
	return 0
}

func (m *BannerAdDetail) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *BannerAdDetail) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListBannerAdRes struct {
	Total                uint64            `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	BannerAds            []*BannerAdDetail `protobuf:"bytes,2,rep,name=banner_ads,json=bannerAds,proto3" json:"banner_ads"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListBannerAdRes) Reset()         { *m = ListBannerAdRes{} }
func (m *ListBannerAdRes) String() string { return proto.CompactTextString(m) }
func (*ListBannerAdRes) ProtoMessage()    {}
func (*ListBannerAdRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb5d95530256fb08, []int{5}
}

func (m *ListBannerAdRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBannerAdRes.Unmarshal(m, b)
}
func (m *ListBannerAdRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBannerAdRes.Marshal(b, m, deterministic)
}
func (m *ListBannerAdRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBannerAdRes.Merge(m, src)
}
func (m *ListBannerAdRes) XXX_Size() int {
	return xxx_messageInfo_ListBannerAdRes.Size(m)
}
func (m *ListBannerAdRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBannerAdRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListBannerAdRes proto.InternalMessageInfo

func (m *ListBannerAdRes) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListBannerAdRes) GetBannerAds() []*BannerAdDetail {
	if m != nil {
		return m.BannerAds
	}
	return nil
}

func init() {
	proto.RegisterEnum("shoppb.BannerAdStatus", BannerAdStatus_name, BannerAdStatus_value)
	proto.RegisterType((*BannerAd)(nil), "shoppb.BannerAd")
	proto.RegisterType((*EditBannerAdStatusReq)(nil), "shoppb.EditBannerAdStatusReq")
	proto.RegisterType((*DelBannerAdReq)(nil), "shoppb.DelBannerAdReq")
	proto.RegisterType((*ListBannerAdReq)(nil), "shoppb.ListBannerAdReq")
	proto.RegisterType((*BannerAdDetail)(nil), "shoppb.BannerAdDetail")
	proto.RegisterType((*ListBannerAdRes)(nil), "shoppb.ListBannerAdRes")
}

func init() { proto.RegisterFile("shoppb/banner.proto", fileDescriptor_fb5d95530256fb08) }

var fileDescriptor_fb5d95530256fb08 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xcd, 0x8e, 0xd3, 0x3c,
	0x14, 0xfd, 0x92, 0x66, 0xda, 0xf4, 0xb6, 0x5f, 0x5b, 0x79, 0xf8, 0xc9, 0x14, 0x55, 0x94, 0xae,
	0x2a, 0x16, 0x29, 0x1a, 0x34, 0x0f, 0xd0, 0x4e, 0x47, 0x08, 0xc4, 0x02, 0xa5, 0xb0, 0x61, 0x41,
	0xe4, 0xd4, 0x57, 0xad, 0xa5, 0x34, 0x09, 0xb1, 0x8b, 0xd4, 0x79, 0x17, 0x1e, 0x83, 0x15, 0xaf,
	0xc1, 0x03, 0xa1, 0xd8, 0x4e, 0xff, 0x91, 0x86, 0x25, 0xab, 0xc4, 0xe7, 0xf8, 0xd8, 0xf7, 0xdc,
	0x1f, 0xc3, 0xa5, 0x58, 0xa6, 0x59, 0x16, 0x8d, 0x22, 0x9a, 0x24, 0x98, 0xfb, 0x59, 0x9e, 0xca,
	0x94, 0x54, 0x35, 0xd8, 0xbd, 0x8c, 0xa8, 0x40, 0x45, 0x16, 0x1f, 0x4d, 0x0e, 0x7e, 0x59, 0xe0,
	0x4e, 0xd4, 0xee, 0x31, 0x23, 0x2d, 0xb0, 0x39, 0xf3, 0xac, 0xbe, 0x35, 0x74, 0x02, 0x9b, 0x33,
	0x72, 0x05, 0x2e, 0xc6, 0x18, 0xca, 0x4d, 0x86, 0x9e, 0xdd, 0xb7, 0x86, 0xff, 0x07, 0x35, 0x8c,
	0xf1, 0xe3, 0x26, 0x43, 0xf2, 0x0c, 0xea, 0x7c, 0x45, 0x17, 0x18, 0xae, 0xf3, 0xd8, 0xab, 0xf4,
	0xad, 0x61, 0x3d, 0x70, 0x15, 0xf0, 0x29, 0x8f, 0xc9, 0x0b, 0x68, 0xe6, 0xc8, 0x78, 0x8e, 0x73,
	0xa9, 0x78, 0x47, 0xf1, 0x8d, 0x12, 0x2b, 0xb6, 0x10, 0x70, 0x44, 0x9a, 0x4b, 0xef, 0x42, 0x1d,
	0xab, 0xfe, 0x89, 0x0f, 0x55, 0x21, 0xa9, 0x5c, 0x0b, 0xaf, 0xda, 0xb7, 0x86, 0xad, 0xeb, 0x27,
	0xbe, 0x8e, 0xdc, 0x2f, 0x03, 0x9c, 0x29, 0x36, 0x30, 0xbb, 0x8a, 0xf0, 0x28, 0x5b, 0xf1, 0x24,
	0xe4, 0xcc, 0xab, 0xa9, 0xa0, 0x6b, 0x6a, 0xfd, 0x96, 0x0d, 0x72, 0x78, 0x7c, 0xc7, 0xb8, 0x3c,
	0x12, 0xe2, 0xd7, 0xad, 0xc5, 0x8a, 0xb1, 0xb8, 0xbb, 0xd3, 0xfe, 0xeb, 0x3b, 0x2b, 0x87, 0x77,
	0xf6, 0xa1, 0x35, 0xc5, 0xb8, 0xd4, 0x9d, 0xb9, 0x6c, 0xf0, 0xdd, 0x82, 0xf6, 0x7b, 0x2e, 0xe4,
	0xfe, 0x1e, 0x02, 0x4e, 0x46, 0x17, 0x68, 0xb2, 0xae, 0xfe, 0x8b, 0xe4, 0x16, 0xdf, 0x50, 0xf0,
	0x7b, 0x9d, 0x78, 0x27, 0x70, 0x0b, 0x60, 0xc6, 0xef, 0xd1, 0x1c, 0x5a, 0x39, 0x5b, 0x24, 0xe7,
	0xb0, 0x48, 0x3b, 0x73, 0x17, 0x0f, 0x31, 0x37, 0xf8, 0x69, 0x43, 0xab, 0xa4, 0xa6, 0x28, 0x29,
	0x8f, 0xff, 0xb5, 0x96, 0xe8, 0x01, 0xcc, 0x73, 0xa4, 0x12, 0x59, 0x18, 0x6d, 0x4c, 0x53, 0xd4,
	0x0d, 0x32, 0xd9, 0x14, 0xf4, 0x3a, 0x63, 0x25, 0xed, 0x6a, 0xda, 0x20, 0x9a, 0x2e, 0xd5, 0x54,
	0x7a, 0x75, 0x15, 0x62, 0xa9, 0x1e, 0xcb, 0x7d, 0x35, 0x95, 0x1e, 0x68, 0xda, 0x20, 0x63, 0x39,
	0xf8, 0x72, 0x5c, 0x5c, 0x41, 0x1e, 0xc1, 0x85, 0x4c, 0x25, 0x8d, 0x4d, 0x02, 0xf5, 0x82, 0xdc,
	0x00, 0xe8, 0x01, 0x0d, 0x29, 0x2b, 0xfa, 0xae, 0x32, 0x6c, 0x9c, 0x1a, 0xd3, 0xf9, 0x0f, 0xea,
	0x91, 0x59, 0x8b, 0x97, 0xef, 0x76, 0xc5, 0x99, 0x95, 0x6e, 0xaf, 0x0e, 0x91, 0x0f, 0x31, 0x9d,
	0xe3, 0x32, 0x8d, 0x19, 0xe6, 0x9d, 0xff, 0x48, 0x03, 0x6a, 0x77, 0x09, 0x8d, 0x62, 0x64, 0x1d,
	0x8b, 0x34, 0xc1, 0x9d, 0x72, 0xa1, 0x57, 0xf6, 0xf5, 0x0f, 0x1b, 0xda, 0x5b, 0x29, 0xe6, 0xdf,
	0xf8, 0x1c, 0xc9, 0x2d, 0xb4, 0xdf, 0xe0, 0x36, 0xfc, 0xc2, 0x0a, 0x79, 0x5a, 0x46, 0x75, 0xd4,
	0xb5, 0xdd, 0x3f, 0x10, 0x82, 0x8c, 0xa0, 0x31, 0x66, 0x6c, 0xfb, 0xa2, 0x74, 0x8e, 0x6d, 0x75,
	0x5b, 0xbe, 0x79, 0x7f, 0xc6, 0xc9, 0xa6, 0x10, 0xbc, 0x82, 0xe6, 0xfe, 0xa4, 0x3e, 0x40, 0x71,
	0x0b, 0xe4, 0x74, 0xb6, 0x49, 0xaf, 0xd4, 0x9d, 0x9d, 0xfb, 0x93, 0x43, 0x6e, 0xa0, 0xb1, 0x37,
	0xac, 0x64, 0x9b, 0xfe, 0xc3, 0x09, 0x3e, 0x96, 0x4d, 0x9e, 0x7f, 0xee, 0x2d, 0xb8, 0x5c, 0xae,
	0x23, 0x7f, 0x9e, 0xae, 0x46, 0x62, 0xc9, 0x93, 0x15, 0x5f, 0xa4, 0xa3, 0x2c, 0x1a, 0x69, 0x7d,
	0x54, 0x55, 0xcf, 0xea, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0xde, 0x92, 0x93, 0x8a,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BannerAdServiceClient is the client API for BannerAdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BannerAdServiceClient interface {
	GetBannerAdList(ctx context.Context, in *ListBannerAdReq, opts ...grpc.CallOption) (*ListBannerAdRes, error)
	AddBannerAd(ctx context.Context, in *BannerAd, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	EditBannerAd(ctx context.Context, in *BannerAd, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	EditBannerAdStatus(ctx context.Context, in *EditBannerAdStatusReq, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	DelBannerAd(ctx context.Context, in *DelBannerAdReq, opts ...grpc.CallOption) (*basepb.AnyRes, error)
}

type bannerAdServiceClient struct {
	cc *grpc.ClientConn
}

func NewBannerAdServiceClient(cc *grpc.ClientConn) BannerAdServiceClient {
	return &bannerAdServiceClient{cc}
}

func (c *bannerAdServiceClient) GetBannerAdList(ctx context.Context, in *ListBannerAdReq, opts ...grpc.CallOption) (*ListBannerAdRes, error) {
	out := new(ListBannerAdRes)
	err := c.cc.Invoke(ctx, "/shoppb.BannerAdService/GetBannerAdList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerAdServiceClient) AddBannerAd(ctx context.Context, in *BannerAd, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/shoppb.BannerAdService/AddBannerAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerAdServiceClient) EditBannerAd(ctx context.Context, in *BannerAd, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/shoppb.BannerAdService/EditBannerAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerAdServiceClient) EditBannerAdStatus(ctx context.Context, in *EditBannerAdStatusReq, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/shoppb.BannerAdService/EditBannerAdStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerAdServiceClient) DelBannerAd(ctx context.Context, in *DelBannerAdReq, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/shoppb.BannerAdService/DelBannerAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BannerAdServiceServer is the server API for BannerAdService service.
type BannerAdServiceServer interface {
	GetBannerAdList(context.Context, *ListBannerAdReq) (*ListBannerAdRes, error)
	AddBannerAd(context.Context, *BannerAd) (*basepb.AnyRes, error)
	EditBannerAd(context.Context, *BannerAd) (*basepb.AnyRes, error)
	EditBannerAdStatus(context.Context, *EditBannerAdStatusReq) (*basepb.AnyRes, error)
	DelBannerAd(context.Context, *DelBannerAdReq) (*basepb.AnyRes, error)
}

// UnimplementedBannerAdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBannerAdServiceServer struct {
}

func (*UnimplementedBannerAdServiceServer) GetBannerAdList(ctx context.Context, req *ListBannerAdReq) (*ListBannerAdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBannerAdList not implemented")
}
func (*UnimplementedBannerAdServiceServer) AddBannerAd(ctx context.Context, req *BannerAd) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBannerAd not implemented")
}
func (*UnimplementedBannerAdServiceServer) EditBannerAd(ctx context.Context, req *BannerAd) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditBannerAd not implemented")
}
func (*UnimplementedBannerAdServiceServer) EditBannerAdStatus(ctx context.Context, req *EditBannerAdStatusReq) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditBannerAdStatus not implemented")
}
func (*UnimplementedBannerAdServiceServer) DelBannerAd(ctx context.Context, req *DelBannerAdReq) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelBannerAd not implemented")
}

func RegisterBannerAdServiceServer(s *grpc.Server, srv BannerAdServiceServer) {
	s.RegisterService(&_BannerAdService_serviceDesc, srv)
}

func _BannerAdService_GetBannerAdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBannerAdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerAdServiceServer).GetBannerAdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.BannerAdService/GetBannerAdList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerAdServiceServer).GetBannerAdList(ctx, req.(*ListBannerAdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerAdService_AddBannerAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerAd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerAdServiceServer).AddBannerAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.BannerAdService/AddBannerAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerAdServiceServer).AddBannerAd(ctx, req.(*BannerAd))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerAdService_EditBannerAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerAd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerAdServiceServer).EditBannerAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.BannerAdService/EditBannerAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerAdServiceServer).EditBannerAd(ctx, req.(*BannerAd))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerAdService_EditBannerAdStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditBannerAdStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerAdServiceServer).EditBannerAdStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.BannerAdService/EditBannerAdStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerAdServiceServer).EditBannerAdStatus(ctx, req.(*EditBannerAdStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerAdService_DelBannerAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelBannerAdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerAdServiceServer).DelBannerAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.BannerAdService/DelBannerAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerAdServiceServer).DelBannerAd(ctx, req.(*DelBannerAdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BannerAdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shoppb.BannerAdService",
	HandlerType: (*BannerAdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBannerAdList",
			Handler:    _BannerAdService_GetBannerAdList_Handler,
		},
		{
			MethodName: "AddBannerAd",
			Handler:    _BannerAdService_AddBannerAd_Handler,
		},
		{
			MethodName: "EditBannerAd",
			Handler:    _BannerAdService_EditBannerAd_Handler,
		},
		{
			MethodName: "EditBannerAdStatus",
			Handler:    _BannerAdService_EditBannerAdStatus_Handler,
		},
		{
			MethodName: "DelBannerAd",
			Handler:    _BannerAdService_DelBannerAd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shoppb/banner.proto",
}
