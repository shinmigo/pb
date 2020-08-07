// Code generated by protoc-gen-go. DO NOT EDIT.
// source: productpb/category.proto

package productpb

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

type CategoryStatus int32

const (
	CategoryStatus_InActive CategoryStatus = 0
	CategoryStatus_Active   CategoryStatus = 1
)

var CategoryStatus_name = map[int32]string{
	0: "InActive",
	1: "Active",
}

var CategoryStatus_value = map[string]int32{
	"InActive": 0,
	"Active":   1,
}

func (x CategoryStatus) String() string {
	return proto.EnumName(CategoryStatus_name, int32(x))
}

func (CategoryStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e95d069ce10e952, []int{0}
}

type Category struct {
	CategoryId           uint64         `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	StoreId              uint64         `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ParentId             uint64         `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name                 string         `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Icon                 string         `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	Status               CategoryStatus `protobuf:"varint,6,opt,name=status,proto3,enum=productpb.CategoryStatus" json:"status,omitempty"`
	Sort                 uint64         `protobuf:"varint,7,opt,name=sort,proto3" json:"sort,omitempty"`
	AdminId              uint64         `protobuf:"varint,8,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e95d069ce10e952, []int{0}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetCategoryId() uint64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *Category) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Category) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Category) GetStatus() CategoryStatus {
	if m != nil {
		return m.Status
	}
	return CategoryStatus_InActive
}

func (m *Category) GetSort() uint64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *Category) GetAdminId() uint64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

type CategoryDetail struct {
	CategoryId           uint64         `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	ParentId             uint64         `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name                 string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Icon                 string         `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	Status               CategoryStatus `protobuf:"varint,5,opt,name=status,proto3,enum=productpb.CategoryStatus" json:"status,omitempty"`
	Sort                 uint64         `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CategoryDetail) Reset()         { *m = CategoryDetail{} }
func (m *CategoryDetail) String() string { return proto.CompactTextString(m) }
func (*CategoryDetail) ProtoMessage()    {}
func (*CategoryDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e95d069ce10e952, []int{1}
}

func (m *CategoryDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryDetail.Unmarshal(m, b)
}
func (m *CategoryDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryDetail.Marshal(b, m, deterministic)
}
func (m *CategoryDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryDetail.Merge(m, src)
}
func (m *CategoryDetail) XXX_Size() int {
	return xxx_messageInfo_CategoryDetail.Size(m)
}
func (m *CategoryDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryDetail.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryDetail proto.InternalMessageInfo

func (m *CategoryDetail) GetCategoryId() uint64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *CategoryDetail) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *CategoryDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CategoryDetail) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *CategoryDetail) GetStatus() CategoryStatus {
	if m != nil {
		return m.Status
	}
	return CategoryStatus_InActive
}

func (m *CategoryDetail) GetSort() uint64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

type EditCategoryStatusReq struct {
	CategoryId           []uint64       `protobuf:"varint,1,rep,packed,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Status               CategoryStatus `protobuf:"varint,2,opt,name=status,proto3,enum=productpb.CategoryStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EditCategoryStatusReq) Reset()         { *m = EditCategoryStatusReq{} }
func (m *EditCategoryStatusReq) String() string { return proto.CompactTextString(m) }
func (*EditCategoryStatusReq) ProtoMessage()    {}
func (*EditCategoryStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e95d069ce10e952, []int{2}
}

func (m *EditCategoryStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditCategoryStatusReq.Unmarshal(m, b)
}
func (m *EditCategoryStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditCategoryStatusReq.Marshal(b, m, deterministic)
}
func (m *EditCategoryStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditCategoryStatusReq.Merge(m, src)
}
func (m *EditCategoryStatusReq) XXX_Size() int {
	return xxx_messageInfo_EditCategoryStatusReq.Size(m)
}
func (m *EditCategoryStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EditCategoryStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_EditCategoryStatusReq proto.InternalMessageInfo

func (m *EditCategoryStatusReq) GetCategoryId() []uint64 {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

func (m *EditCategoryStatusReq) GetStatus() CategoryStatus {
	if m != nil {
		return m.Status
	}
	return CategoryStatus_InActive
}

type DelCategoryReq struct {
	CategoryId           []uint64 `protobuf:"varint,1,rep,packed,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelCategoryReq) Reset()         { *m = DelCategoryReq{} }
func (m *DelCategoryReq) String() string { return proto.CompactTextString(m) }
func (*DelCategoryReq) ProtoMessage()    {}
func (*DelCategoryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e95d069ce10e952, []int{3}
}

func (m *DelCategoryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelCategoryReq.Unmarshal(m, b)
}
func (m *DelCategoryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelCategoryReq.Marshal(b, m, deterministic)
}
func (m *DelCategoryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelCategoryReq.Merge(m, src)
}
func (m *DelCategoryReq) XXX_Size() int {
	return xxx_messageInfo_DelCategoryReq.Size(m)
}
func (m *DelCategoryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelCategoryReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelCategoryReq proto.InternalMessageInfo

func (m *DelCategoryReq) GetCategoryId() []uint64 {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

type ListCategoryReq struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int64    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Id                   uint64   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCategoryReq) Reset()         { *m = ListCategoryReq{} }
func (m *ListCategoryReq) String() string { return proto.CompactTextString(m) }
func (*ListCategoryReq) ProtoMessage()    {}
func (*ListCategoryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e95d069ce10e952, []int{4}
}

func (m *ListCategoryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCategoryReq.Unmarshal(m, b)
}
func (m *ListCategoryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCategoryReq.Marshal(b, m, deterministic)
}
func (m *ListCategoryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCategoryReq.Merge(m, src)
}
func (m *ListCategoryReq) XXX_Size() int {
	return xxx_messageInfo_ListCategoryReq.Size(m)
}
func (m *ListCategoryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCategoryReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListCategoryReq proto.InternalMessageInfo

func (m *ListCategoryReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListCategoryReq) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCategoryReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ListCategoryReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListCategoryRes struct {
	Total                uint64            `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Categories           []*CategoryDetail `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListCategoryRes) Reset()         { *m = ListCategoryRes{} }
func (m *ListCategoryRes) String() string { return proto.CompactTextString(m) }
func (*ListCategoryRes) ProtoMessage()    {}
func (*ListCategoryRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e95d069ce10e952, []int{5}
}

func (m *ListCategoryRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCategoryRes.Unmarshal(m, b)
}
func (m *ListCategoryRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCategoryRes.Marshal(b, m, deterministic)
}
func (m *ListCategoryRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCategoryRes.Merge(m, src)
}
func (m *ListCategoryRes) XXX_Size() int {
	return xxx_messageInfo_ListCategoryRes.Size(m)
}
func (m *ListCategoryRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCategoryRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListCategoryRes proto.InternalMessageInfo

func (m *ListCategoryRes) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListCategoryRes) GetCategories() []*CategoryDetail {
	if m != nil {
		return m.Categories
	}
	return nil
}

func init() {
	proto.RegisterEnum("productpb.CategoryStatus", CategoryStatus_name, CategoryStatus_value)
	proto.RegisterType((*Category)(nil), "productpb.Category")
	proto.RegisterType((*CategoryDetail)(nil), "productpb.CategoryDetail")
	proto.RegisterType((*EditCategoryStatusReq)(nil), "productpb.EditCategoryStatusReq")
	proto.RegisterType((*DelCategoryReq)(nil), "productpb.DelCategoryReq")
	proto.RegisterType((*ListCategoryReq)(nil), "productpb.ListCategoryReq")
	proto.RegisterType((*ListCategoryRes)(nil), "productpb.ListCategoryRes")
}

func init() { proto.RegisterFile("productpb/category.proto", fileDescriptor_5e95d069ce10e952) }

var fileDescriptor_5e95d069ce10e952 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6f, 0x9b, 0x30,
	0x18, 0x1d, 0x90, 0x50, 0xf2, 0xa5, 0x22, 0x93, 0xbb, 0x49, 0x34, 0x3b, 0x0c, 0x71, 0x8a, 0x7a,
	0x00, 0x35, 0xbd, 0x6c, 0xc7, 0x6c, 0xdd, 0x26, 0xa4, 0x9d, 0xe8, 0x6d, 0x97, 0x0a, 0xb0, 0x47,
	0xad, 0x25, 0x98, 0x61, 0xa7, 0xd2, 0xfa, 0xdf, 0x76, 0xde, 0xaf, 0xd9, 0x7f, 0x98, 0x6c, 0x03,
	0xa1, 0x09, 0xd5, 0xd2, 0x53, 0x3e, 0xfb, 0xf9, 0xf9, 0x7b, 0xef, 0xf3, 0x23, 0xe0, 0x55, 0x35,
	0xc3, 0xdb, 0x5c, 0x54, 0x59, 0x94, 0xa7, 0x82, 0x14, 0xac, 0xfe, 0x15, 0x56, 0x35, 0x13, 0x0c,
	0x4d, 0x3a, 0x64, 0x7e, 0x96, 0xa5, 0x9c, 0x54, 0x59, 0xa4, 0x7f, 0x34, 0x1e, 0xfc, 0x35, 0xc0,
	0xf9, 0xd8, 0x50, 0xd0, 0x5b, 0x98, 0xb6, 0xf4, 0x5b, 0x8a, 0x3d, 0xc3, 0x37, 0x16, 0xa3, 0x04,
	0xda, 0xad, 0x18, 0xa3, 0x73, 0x70, 0xb8, 0x60, 0x35, 0x91, 0xa8, 0xa9, 0xd0, 0x13, 0xb5, 0x8e,
	0x31, 0x7a, 0x03, 0x93, 0x2a, 0xad, 0x49, 0x29, 0x24, 0x66, 0x29, 0xcc, 0xd1, 0x1b, 0x31, 0x46,
	0x08, 0x46, 0x65, 0xba, 0x21, 0xde, 0xc8, 0x37, 0x16, 0x93, 0x44, 0xd5, 0x72, 0x8f, 0xe6, 0xac,
	0xf4, 0xc6, 0x7a, 0x4f, 0xd6, 0xe8, 0x12, 0x6c, 0x2e, 0x52, 0xb1, 0xe5, 0x9e, 0xed, 0x1b, 0x0b,
	0x77, 0x79, 0x1e, 0x76, 0xf2, 0xc3, 0x56, 0xe5, 0x8d, 0x3a, 0x90, 0x34, 0x07, 0xe5, 0x35, 0x9c,
	0xd5, 0xc2, 0x3b, 0x51, 0x2d, 0x55, 0x2d, 0x65, 0xa6, 0x78, 0x43, 0x4b, 0x29, 0xc5, 0xd1, 0x32,
	0xd5, 0x3a, 0xc6, 0xc1, 0x6f, 0x03, 0xdc, 0xf6, 0xa6, 0x6b, 0x22, 0x52, 0xba, 0xfe, 0xbf, 0xeb,
	0x47, 0xd6, 0xcc, 0x27, 0xac, 0x59, 0x03, 0xd6, 0x46, 0x83, 0xd6, 0xc6, 0xcf, 0xb5, 0x66, 0xef,
	0xac, 0x05, 0x3f, 0xe0, 0xf5, 0x27, 0x4c, 0xc5, 0x1e, 0x83, 0xfc, 0x3c, 0x74, 0x61, 0xed, 0xb9,
	0xd8, 0x09, 0x30, 0x8f, 0x14, 0x10, 0x5c, 0x82, 0x7b, 0x4d, 0xd6, 0x2d, 0x78, 0x4c, 0x97, 0xe0,
	0x3b, 0xcc, 0xbe, 0x52, 0x2e, 0xfa, 0x1c, 0x04, 0xa3, 0x2a, 0x2d, 0x88, 0x1a, 0xac, 0x95, 0xa8,
	0x5a, 0x8f, 0xb4, 0x20, 0xb7, 0x9c, 0x3e, 0x10, 0xa5, 0xc7, 0x92, 0x23, 0x2d, 0xc8, 0x0d, 0x7d,
	0x20, 0xc8, 0x05, 0xb3, 0xcb, 0x90, 0x49, 0x07, 0xd3, 0x13, 0x64, 0xfb, 0x7d, 0x38, 0x7a, 0x05,
	0x63, 0xc1, 0x44, 0xba, 0x6e, 0x5e, 0x50, 0x2f, 0xd0, 0x7b, 0x68, 0xe5, 0x51, 0x22, 0xad, 0x5b,
	0x8b, 0xe9, 0xa0, 0x75, 0x1d, 0x86, 0xa4, 0x77, 0xf8, 0xe2, 0x62, 0x17, 0x15, 0x3d, 0x18, 0x74,
	0x0a, 0x4e, 0x5c, 0xae, 0x72, 0x41, 0xef, 0xc9, 0xcb, 0x17, 0x08, 0xc0, 0x6e, 0x6a, 0x63, 0xf9,
	0xc7, 0x84, 0x59, 0x77, 0x98, 0xd4, 0xf7, 0x34, 0x27, 0x68, 0x09, 0xd3, 0x15, 0xc6, 0xdd, 0xd7,
	0x75, 0x36, 0xd0, 0x75, 0xee, 0x86, 0xcd, 0xe7, 0xb8, 0x2a, 0x95, 0x89, 0x2b, 0x38, 0xed, 0xbf,
	0xef, 0x71, 0xa4, 0xcf, 0x80, 0x0e, 0x43, 0x81, 0xfc, 0x1e, 0x75, 0x30, 0x33, 0x07, 0xf7, 0xbc,
	0x83, 0x69, 0xef, 0xbd, 0x51, 0x7f, 0x4c, 0x8f, 0x73, 0x70, 0xc0, 0x8c, 0x61, 0xf6, 0x85, 0x74,
	0x1d, 0xe4, 0xcb, 0xa0, 0x79, 0x8f, 0xbd, 0x17, 0x89, 0xf9, 0xd3, 0x18, 0xff, 0x10, 0x7c, 0xf3,
	0x0b, 0x2a, 0xee, 0xb6, 0x59, 0x98, 0xb3, 0x4d, 0xc4, 0xef, 0x68, 0xb9, 0xa1, 0x05, 0x8b, 0xaa,
	0x2c, 0xea, 0x38, 0x99, 0xad, 0xfe, 0xbc, 0xae, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xe1,
	0x5d, 0x49, 0xf8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CategoryServiceClient interface {
	AddCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	EditCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	EditCategoryStatus(ctx context.Context, in *EditCategoryStatusReq, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	DelCategory(ctx context.Context, in *DelCategoryReq, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	GetCategoryList(ctx context.Context, in *ListCategoryReq, opts ...grpc.CallOption) (*ListCategoryRes, error)
}

type categoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewCategoryServiceClient(cc *grpc.ClientConn) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) AddCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.CategoryService/AddCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) EditCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.CategoryService/EditCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) EditCategoryStatus(ctx context.Context, in *EditCategoryStatusReq, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.CategoryService/EditCategoryStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) DelCategory(ctx context.Context, in *DelCategoryReq, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.CategoryService/DelCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategoryList(ctx context.Context, in *ListCategoryReq, opts ...grpc.CallOption) (*ListCategoryRes, error) {
	out := new(ListCategoryRes)
	err := c.cc.Invoke(ctx, "/productpb.CategoryService/GetCategoryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
type CategoryServiceServer interface {
	AddCategory(context.Context, *Category) (*basepb.AnyRes, error)
	EditCategory(context.Context, *Category) (*basepb.AnyRes, error)
	EditCategoryStatus(context.Context, *EditCategoryStatusReq) (*basepb.AnyRes, error)
	DelCategory(context.Context, *DelCategoryReq) (*basepb.AnyRes, error)
	GetCategoryList(context.Context, *ListCategoryReq) (*ListCategoryRes, error)
}

// UnimplementedCategoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (*UnimplementedCategoryServiceServer) AddCategory(ctx context.Context, req *Category) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (*UnimplementedCategoryServiceServer) EditCategory(ctx context.Context, req *Category) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCategory not implemented")
}
func (*UnimplementedCategoryServiceServer) EditCategoryStatus(ctx context.Context, req *EditCategoryStatusReq) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCategoryStatus not implemented")
}
func (*UnimplementedCategoryServiceServer) DelCategory(ctx context.Context, req *DelCategoryReq) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelCategory not implemented")
}
func (*UnimplementedCategoryServiceServer) GetCategoryList(ctx context.Context, req *ListCategoryReq) (*ListCategoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryList not implemented")
}

func RegisterCategoryServiceServer(s *grpc.Server, srv CategoryServiceServer) {
	s.RegisterService(&_CategoryService_serviceDesc, srv)
}

func _CategoryService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.CategoryService/AddCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).AddCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_EditCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).EditCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.CategoryService/EditCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).EditCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_EditCategoryStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditCategoryStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).EditCategoryStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.CategoryService/EditCategoryStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).EditCategoryStatus(ctx, req.(*EditCategoryStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_DelCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).DelCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.CategoryService/DelCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).DelCategory(ctx, req.(*DelCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.CategoryService/GetCategoryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategoryList(ctx, req.(*ListCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CategoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "productpb.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCategory",
			Handler:    _CategoryService_AddCategory_Handler,
		},
		{
			MethodName: "EditCategory",
			Handler:    _CategoryService_EditCategory_Handler,
		},
		{
			MethodName: "EditCategoryStatus",
			Handler:    _CategoryService_EditCategoryStatus_Handler,
		},
		{
			MethodName: "DelCategory",
			Handler:    _CategoryService_DelCategory_Handler,
		},
		{
			MethodName: "GetCategoryList",
			Handler:    _CategoryService_GetCategoryList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "productpb/category.proto",
}
