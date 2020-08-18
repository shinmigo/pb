// Code generated by protoc-gen-go. DO NOT EDIT.
// source: productpb/category.proto

package productpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/shinmigo/pb/basepb"
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
	Page     int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Id       uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to StatusPresent:
	//	*ListCategoryReq_Status
	StatusPresent        isListCategoryReq_StatusPresent `protobuf_oneof:"status_present"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
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

type isListCategoryReq_StatusPresent interface {
	isListCategoryReq_StatusPresent()
}

type ListCategoryReq_Status struct {
	Status CategoryStatus `protobuf:"varint,5,opt,name=status,proto3,enum=productpb.CategoryStatus,oneof"`
}

func (*ListCategoryReq_Status) isListCategoryReq_StatusPresent() {}

func (m *ListCategoryReq) GetStatusPresent() isListCategoryReq_StatusPresent {
	if m != nil {
		return m.StatusPresent
	}
	return nil
}

func (m *ListCategoryReq) GetStatus() CategoryStatus {
	if x, ok := m.GetStatusPresent().(*ListCategoryReq_Status); ok {
		return x.Status
	}
	return CategoryStatus_InActive
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListCategoryReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListCategoryReq_Status)(nil),
	}
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
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x8d, 0x6d, 0x20, 0x30, 0x44, 0x26, 0xda, 0xb4, 0x92, 0x43, 0x0f, 0x45, 0x3e, 0xa1, 0x1c,
	0x40, 0x81, 0x4b, 0x7b, 0x24, 0x4d, 0x3f, 0x2c, 0xf5, 0x44, 0x6e, 0xbd, 0x20, 0xdb, 0x3b, 0x72,
	0x56, 0x05, 0xaf, 0xeb, 0x5d, 0x22, 0x35, 0xff, 0xa9, 0x3f, 0xa1, 0xe7, 0xfe, 0x9a, 0xfe, 0x87,
	0x6a, 0x77, 0x6d, 0xe3, 0x80, 0xa3, 0x92, 0x13, 0xb3, 0xfb, 0xf6, 0xed, 0xbc, 0x37, 0xbc, 0x35,
	0x78, 0x59, 0xce, 0xe9, 0x36, 0x96, 0x59, 0x34, 0x8d, 0x43, 0x89, 0x09, 0xcf, 0x7f, 0x4e, 0xb2,
	0x9c, 0x4b, 0x4e, 0x7a, 0x15, 0x32, 0xbc, 0x88, 0x42, 0x81, 0x59, 0x34, 0x35, 0x3f, 0x06, 0xf7,
	0xff, 0x5a, 0xd0, 0xfd, 0x50, 0x50, 0xc8, 0x5b, 0xe8, 0x97, 0xf4, 0x15, 0xa3, 0x9e, 0x35, 0xb2,
	0xc6, 0xad, 0x25, 0x94, 0x5b, 0x01, 0x25, 0x97, 0xd0, 0x15, 0x92, 0xe7, 0xa8, 0x50, 0x5b, 0xa3,
	0xa7, 0x7a, 0x1d, 0x50, 0xf2, 0x06, 0x7a, 0x59, 0x98, 0x63, 0x2a, 0x15, 0xe6, 0x68, 0xac, 0x6b,
	0x36, 0x02, 0x4a, 0x08, 0xb4, 0xd2, 0x70, 0x83, 0x5e, 0x6b, 0x64, 0x8d, 0x7b, 0x4b, 0x5d, 0xab,
	0x3d, 0x16, 0xf3, 0xd4, 0x6b, 0x9b, 0x3d, 0x55, 0x93, 0x6b, 0xe8, 0x08, 0x19, 0xca, 0xad, 0xf0,
	0x3a, 0x23, 0x6b, 0xec, 0xce, 0x2e, 0x27, 0x95, 0xfc, 0x49, 0xa9, 0xf2, 0x4e, 0x1f, 0x58, 0x16,
	0x07, 0xd5, 0x35, 0x82, 0xe7, 0xd2, 0x3b, 0xd5, 0x2d, 0x75, 0xad, 0x64, 0x86, 0x74, 0xc3, 0x52,
	0x25, 0xa5, 0x6b, 0x64, 0xea, 0x75, 0x40, 0xfd, 0xdf, 0x16, 0xb8, 0xe5, 0x4d, 0xb7, 0x28, 0x43,
	0xb6, 0xfe, 0xbf, 0xeb, 0x27, 0xd6, 0xec, 0x67, 0xac, 0x39, 0x0d, 0xd6, 0x5a, 0x8d, 0xd6, 0xda,
	0x2f, 0xb5, 0xd6, 0xd9, 0x59, 0xf3, 0xbf, 0xc3, 0xeb, 0x8f, 0x94, 0xc9, 0x3d, 0x06, 0xfe, 0x38,
	0x74, 0xe1, 0xec, 0xb9, 0xd8, 0x09, 0xb0, 0x8f, 0x14, 0xe0, 0x5f, 0x83, 0x7b, 0x8b, 0xeb, 0x12,
	0x3c, 0xa6, 0x8b, 0xff, 0xcb, 0x82, 0xc1, 0x57, 0x26, 0x64, 0x9d, 0x44, 0xa0, 0x95, 0x85, 0x09,
	0xea, 0xc9, 0x3a, 0x4b, 0x5d, 0x9b, 0x99, 0x26, 0xb8, 0x12, 0xec, 0x11, 0xb5, 0x20, 0x47, 0xcd,
	0x34, 0xc1, 0x3b, 0xf6, 0x88, 0xc4, 0x05, 0xbb, 0x0a, 0x91, 0xcd, 0x9a, 0xe3, 0x33, 0x3f, 0x7a,
	0x9e, 0x5f, 0x4e, 0x4a, 0x43, 0x37, 0xe7, 0xe0, 0x9a, 0x6a, 0x95, 0xe5, 0x28, 0x30, 0x95, 0x7e,
	0xb4, 0x2f, 0x57, 0x90, 0x57, 0xd0, 0x96, 0x5c, 0x86, 0xeb, 0x22, 0x09, 0x66, 0x41, 0xde, 0x43,
	0x69, 0x93, 0xa1, 0x1a, 0xa1, 0x33, 0xee, 0x37, 0xf6, 0x34, 0xa1, 0x5a, 0xd6, 0x0e, 0x5f, 0x5d,
	0xed, 0x22, 0x67, 0x14, 0x91, 0x33, 0xe8, 0x06, 0xe9, 0x22, 0x96, 0xec, 0x01, 0xcf, 0x4f, 0x08,
	0x40, 0xa7, 0xa8, 0xad, 0xd9, 0x1f, 0x1b, 0x06, 0xd5, 0x61, 0xcc, 0x1f, 0x58, 0x8c, 0x64, 0x06,
	0xfd, 0x05, 0xa5, 0xd5, 0x2b, 0xbd, 0x68, 0xe8, 0x3a, 0x74, 0x27, 0xc5, 0xb3, 0x5e, 0xa4, 0xda,
	0xc4, 0x1c, 0xce, 0xea, 0x39, 0x39, 0x8e, 0xf4, 0x09, 0xc8, 0x61, 0xb8, 0xc8, 0xa8, 0x46, 0x6d,
	0xcc, 0xde, 0xc1, 0x3d, 0xef, 0xa0, 0x5f, 0xcb, 0x0d, 0xa9, 0x8f, 0xe9, 0x69, 0x9e, 0x0e, 0x98,
	0x01, 0x0c, 0x3e, 0x63, 0xd5, 0x41, 0xfd, 0x33, 0x64, 0x58, 0x63, 0xef, 0x25, 0x6b, 0xf8, 0x3c,
	0x26, 0x6e, 0xfc, 0x6f, 0xa3, 0x84, 0xc9, 0xfb, 0x6d, 0x34, 0x89, 0xf9, 0x66, 0x2a, 0xee, 0x59,
	0xba, 0x61, 0x09, 0x9f, 0x66, 0xd1, 0xb4, 0xe2, 0x44, 0x1d, 0xfd, 0x11, 0x9c, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x90, 0x12, 0x20, 0xcf, 0x40, 0x05, 0x00, 0x00,
}