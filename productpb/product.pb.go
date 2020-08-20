// Code generated by protoc-gen-go. DO NOT EDIT.
// source: productpb/product.proto

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

type ProductSpecType int32

const (
	ProductSpecType_SpecTypePlaceholder ProductSpecType = 0
	ProductSpecType_Single              ProductSpecType = 1
	ProductSpecType_Multiple            ProductSpecType = 2
)

var ProductSpecType_name = map[int32]string{
	0: "SpecTypePlaceholder",
	1: "Single",
	2: "Multiple",
}

var ProductSpecType_value = map[string]int32{
	"SpecTypePlaceholder": 0,
	"Single":              1,
	"Multiple":            2,
}

func (x ProductSpecType) String() string {
	return proto.EnumName(ProductSpecType_name, int32(x))
}

func (ProductSpecType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5303e89fc52b2032, []int{0}
}

type ProductStatus int32

const (
	ProductStatus_ProductStatusPlaceholder ProductStatus = 0
	ProductStatus_Disabled                 ProductStatus = 1
	ProductStatus_Enabled                  ProductStatus = 2
)

var ProductStatus_name = map[int32]string{
	0: "ProductStatusPlaceholder",
	1: "Disabled",
	2: "Enabled",
}

var ProductStatus_value = map[string]int32{
	"ProductStatusPlaceholder": 0,
	"Disabled":                 1,
	"Enabled":                  2,
}

func (x ProductStatus) String() string {
	return proto.EnumName(ProductStatus_name, int32(x))
}

func (ProductStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5303e89fc52b2032, []int{1}
}

type ProductSpec struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	OldPrice             float64  `protobuf:"fixed64,3,opt,name=old_price,json=oldPrice,proto3" json:"old_price,omitempty"`
	CostPrice            float64  `protobuf:"fixed64,4,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	Stock                uint64   `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Sku                  string   `protobuf:"bytes,6,opt,name=sku,proto3" json:"sku,omitempty"`
	Weight               float64  `protobuf:"fixed64,7,opt,name=weight,proto3" json:"weight,omitempty"`
	Volume               float64  `protobuf:"fixed64,8,opt,name=volume,proto3" json:"volume,omitempty"`
	SpecValueId          []uint64 `protobuf:"varint,9,rep,packed,name=spec_value_id,json=specValueId,proto3" json:"spec_value_id,omitempty"`
	ProductSpecId        uint64   `protobuf:"varint,10,opt,name=product_spec_id,json=productSpecId,proto3" json:"product_spec_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductSpec) Reset()         { *m = ProductSpec{} }
func (m *ProductSpec) String() string { return proto.CompactTextString(m) }
func (*ProductSpec) ProtoMessage()    {}
func (*ProductSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5303e89fc52b2032, []int{0}
}

func (m *ProductSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSpec.Unmarshal(m, b)
}
func (m *ProductSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSpec.Marshal(b, m, deterministic)
}
func (m *ProductSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSpec.Merge(m, src)
}
func (m *ProductSpec) XXX_Size() int {
	return xxx_messageInfo_ProductSpec.Size(m)
}
func (m *ProductSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSpec proto.InternalMessageInfo

func (m *ProductSpec) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ProductSpec) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductSpec) GetOldPrice() float64 {
	if m != nil {
		return m.OldPrice
	}
	return 0
}

func (m *ProductSpec) GetCostPrice() float64 {
	if m != nil {
		return m.CostPrice
	}
	return 0
}

func (m *ProductSpec) GetStock() uint64 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func (m *ProductSpec) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *ProductSpec) GetWeight() float64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *ProductSpec) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *ProductSpec) GetSpecValueId() []uint64 {
	if m != nil {
		return m.SpecValueId
	}
	return nil
}

func (m *ProductSpec) GetProductSpecId() uint64 {
	if m != nil {
		return m.ProductSpecId
	}
	return 0
}

type ProductParam struct {
	ParamId              uint64   `protobuf:"varint,1,opt,name=param_id,json=paramId,proto3" json:"param_id,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductParam) Reset()         { *m = ProductParam{} }
func (m *ProductParam) String() string { return proto.CompactTextString(m) }
func (*ProductParam) ProtoMessage()    {}
func (*ProductParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_5303e89fc52b2032, []int{1}
}

func (m *ProductParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductParam.Unmarshal(m, b)
}
func (m *ProductParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductParam.Marshal(b, m, deterministic)
}
func (m *ProductParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductParam.Merge(m, src)
}
func (m *ProductParam) XXX_Size() int {
	return xxx_messageInfo_ProductParam.Size(m)
}
func (m *ProductParam) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductParam.DiscardUnknown(m)
}

var xxx_messageInfo_ProductParam proto.InternalMessageInfo

func (m *ProductParam) GetParamId() uint64 {
	if m != nil {
		return m.ParamId
	}
	return 0
}

func (m *ProductParam) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Product struct {
	ProductId            uint64          `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	CategoryId           uint64          `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	StoreId              uint64          `protobuf:"varint,3,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	KindId               uint64          `protobuf:"varint,4,opt,name=kind_id,json=kindId,proto3" json:"kind_id,omitempty"`
	Name                 string          `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	ShortDescription     string          `protobuf:"bytes,6,opt,name=short_description,json=shortDescription,proto3" json:"short_description,omitempty"`
	Unit                 string          `protobuf:"bytes,7,opt,name=unit,proto3" json:"unit,omitempty"`
	Images               []string        `protobuf:"bytes,8,rep,name=images,proto3" json:"images,omitempty"`
	SpecType             ProductSpecType `protobuf:"varint,9,opt,name=spec_type,json=specType,proto3,enum=productpb.ProductSpecType" json:"spec_type,omitempty"`
	Status               ProductStatus   `protobuf:"varint,10,opt,name=status,proto3,enum=productpb.ProductStatus" json:"status,omitempty"`
	Tags                 []uint64        `protobuf:"varint,11,rep,packed,name=tags,proto3" json:"tags,omitempty"`
	AdminId              uint64          `protobuf:"varint,12,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	Spec                 []*ProductSpec  `protobuf:"bytes,13,rep,name=spec,proto3" json:"spec,omitempty"`
	Param                []*ProductParam `protobuf:"bytes,14,rep,name=param,proto3" json:"param,omitempty"`
	Description          string          `protobuf:"bytes,15,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_5303e89fc52b2032, []int{2}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetProductId() uint64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *Product) GetCategoryId() uint64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *Product) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Product) GetKindId() uint64 {
	if m != nil {
		return m.KindId
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetShortDescription() string {
	if m != nil {
		return m.ShortDescription
	}
	return ""
}

func (m *Product) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Product) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Product) GetSpecType() ProductSpecType {
	if m != nil {
		return m.SpecType
	}
	return ProductSpecType_SpecTypePlaceholder
}

func (m *Product) GetStatus() ProductStatus {
	if m != nil {
		return m.Status
	}
	return ProductStatus_ProductStatusPlaceholder
}

func (m *Product) GetTags() []uint64 {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Product) GetAdminId() uint64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

func (m *Product) GetSpec() []*ProductSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Product) GetParam() []*ProductParam {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("productpb.ProductSpecType", ProductSpecType_name, ProductSpecType_value)
	proto.RegisterEnum("productpb.ProductStatus", ProductStatus_name, ProductStatus_value)
	proto.RegisterType((*ProductSpec)(nil), "productpb.ProductSpec")
	proto.RegisterType((*ProductParam)(nil), "productpb.ProductParam")
	proto.RegisterType((*Product)(nil), "productpb.Product")
}

func init() { proto.RegisterFile("productpb/product.proto", fileDescriptor_5303e89fc52b2032) }

var fileDescriptor_5303e89fc52b2032 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x26, 0xfd, 0x9d, 0x97, 0xb6, 0x2b, 0xde, 0xb4, 0x9a, 0x01, 0x22, 0xea, 0x01, 0x55, 0x43,
	0xb4, 0xa3, 0x1c, 0x38, 0xa2, 0xa1, 0x4d, 0x22, 0x07, 0xa4, 0x2a, 0x43, 0x1c, 0xb8, 0x54, 0x69,
	0x6c, 0xb5, 0xd6, 0x92, 0xd8, 0x8a, 0x9d, 0xa1, 0x1e, 0xf9, 0x33, 0xf8, 0x6f, 0x91, 0x9f, 0xd3,
	0xfd, 0x60, 0xbb, 0x70, 0xca, 0xfb, 0xbe, 0xef, 0x3d, 0x3f, 0xfb, 0xf3, 0x8b, 0x61, 0xac, 0x4a,
	0xc9, 0xaa, 0xd4, 0xa8, 0xf5, 0xbc, 0x8e, 0x66, 0xaa, 0x94, 0x46, 0x12, 0xff, 0x56, 0x38, 0x39,
	0x5c, 0x27, 0x9a, 0xab, 0xf5, 0xdc, 0x7d, 0x9c, 0x3e, 0xf9, 0xd3, 0x80, 0x60, 0xe9, 0x52, 0xae,
	0x14, 0x4f, 0xc9, 0x11, 0xb4, 0x45, 0x9e, 0x6c, 0x38, 0xf5, 0x42, 0x6f, 0xea, 0xc7, 0x0e, 0x58,
	0x56, 0x95, 0x22, 0xe5, 0xb4, 0x11, 0x7a, 0x53, 0x2f, 0x76, 0x80, 0xbc, 0x04, 0x5f, 0x66, 0x6c,
	0xe5, 0x94, 0x26, 0x2a, 0x3d, 0x99, 0xb1, 0x25, 0x8a, 0xaf, 0x01, 0x52, 0xa9, 0x4d, 0xad, 0xb6,
	0x50, 0xf5, 0x2d, 0xe3, 0xe4, 0x23, 0x68, 0x6b, 0x23, 0xd3, 0x6b, 0xda, 0x0e, 0xbd, 0x69, 0x2b,
	0x76, 0x80, 0x8c, 0xa0, 0xa9, 0xaf, 0x2b, 0xda, 0xc1, 0xde, 0x36, 0x24, 0xc7, 0xd0, 0xf9, 0xc5,
	0xc5, 0x66, 0x6b, 0x68, 0x17, 0x97, 0xa8, 0x91, 0xe5, 0x6f, 0x64, 0x56, 0xe5, 0x9c, 0xf6, 0x1c,
	0xef, 0x10, 0x99, 0xc0, 0x40, 0x2b, 0x9e, 0xae, 0x6e, 0x92, 0xac, 0xe2, 0x2b, 0xc1, 0xa8, 0x1f,
	0x36, 0xa7, 0xad, 0x38, 0xb0, 0xe4, 0x0f, 0xcb, 0x45, 0x8c, 0xbc, 0x85, 0x83, 0xda, 0x95, 0x15,
	0xe6, 0x0a, 0x46, 0x01, 0x77, 0x31, 0x50, 0x77, 0x4e, 0x44, 0x6c, 0xf2, 0x19, 0xfa, 0xb5, 0x35,
	0xcb, 0xa4, 0x4c, 0x72, 0xf2, 0x02, 0x7a, 0xca, 0x06, 0xb6, 0xc0, 0xc3, 0x82, 0x2e, 0xe2, 0x88,
	0xd9, 0xe3, 0x60, 0x47, 0x34, 0xc8, 0x8f, 0x1d, 0x98, 0xfc, 0x6e, 0x41, 0xb7, 0x5e, 0xc1, 0xfa,
	0xb1, 0x6f, 0x7a, 0x5b, 0xbe, 0xbf, 0x9c, 0x88, 0x91, 0x37, 0x10, 0xa4, 0x89, 0xe1, 0x1b, 0x59,
	0xee, 0xac, 0xde, 0x40, 0x1d, 0xf6, 0x54, 0xc4, 0x6c, 0x73, 0x6d, 0x64, 0x89, 0x67, 0x6a, 0xba,
	0xe6, 0x88, 0x23, 0x46, 0xc6, 0xd0, 0xbd, 0x16, 0x05, 0xb3, 0x4a, 0x0b, 0x95, 0x8e, 0x85, 0x11,
	0x23, 0x04, 0x5a, 0x45, 0x92, 0x73, 0xf4, 0xd8, 0x8f, 0x31, 0x26, 0xef, 0xe0, 0xb9, 0xde, 0xca,
	0xd2, 0xac, 0x18, 0xd7, 0x69, 0x29, 0x94, 0x11, 0xb2, 0xa8, 0x0d, 0x1f, 0xa1, 0x70, 0x71, 0xc7,
	0xdb, 0x05, 0xaa, 0x42, 0x38, 0xef, 0xfd, 0x18, 0x63, 0xeb, 0x3c, 0x0e, 0x85, 0xa6, 0xbd, 0xb0,
	0x39, 0xf5, 0xe3, 0x1a, 0x91, 0x4f, 0xe0, 0xa3, 0x9b, 0x66, 0xa7, 0x38, 0xf5, 0x43, 0x6f, 0x3a,
	0x5c, 0x9c, 0xcc, 0x6e, 0xa7, 0x6f, 0x76, 0x6f, 0xc8, 0xbe, 0xef, 0x14, 0x8f, 0x7b, 0xba, 0x8e,
	0xc8, 0x19, 0x74, 0xb4, 0x49, 0x4c, 0xa5, 0xf1, 0x16, 0x86, 0x0b, 0xfa, 0x44, 0x15, 0xea, 0x71,
	0x9d, 0x67, 0xb7, 0x65, 0x92, 0x8d, 0xa6, 0x01, 0xde, 0x2d, 0xc6, 0xd6, 0x9f, 0x84, 0xe5, 0xa2,
	0xb0, 0x2e, 0xf4, 0x9d, 0x3f, 0x88, 0x23, 0x46, 0x4e, 0xa1, 0x65, 0x9b, 0xd1, 0x41, 0xd8, 0x9c,
	0x06, 0x8b, 0xe3, 0xa7, 0x37, 0x15, 0x63, 0x0e, 0x79, 0x0f, 0x6d, 0xbc, 0x53, 0x3a, 0xc4, 0xe4,
	0xf1, 0xe3, 0x64, 0x9c, 0x85, 0xd8, 0x65, 0x91, 0x10, 0x82, 0xfb, 0x3e, 0x1e, 0xa0, 0x4f, 0xf7,
	0xa9, 0xd3, 0x0b, 0x38, 0xf8, 0xe7, 0xe8, 0x64, 0x0c, 0x87, 0xfb, 0x78, 0x99, 0x25, 0x29, 0xdf,
	0xca, 0x8c, 0xf1, 0x72, 0xf4, 0x8c, 0x00, 0x74, 0xae, 0x44, 0xb1, 0xc9, 0xf8, 0xc8, 0x23, 0x7d,
	0xe8, 0x7d, 0xab, 0x32, 0x23, 0x54, 0xc6, 0x47, 0x8d, 0xd3, 0xaf, 0x30, 0x78, 0x60, 0x05, 0x79,
	0x05, 0xf4, 0x01, 0xf1, 0x70, 0xa1, 0x3e, 0xf4, 0x2e, 0x84, 0x4e, 0xd6, 0x19, 0x67, 0x23, 0x8f,
	0x04, 0xd0, 0xbd, 0x2c, 0x1c, 0x68, 0x2c, 0x2a, 0x18, 0xee, 0x0b, 0x79, 0x79, 0x63, 0x7f, 0xc5,
	0x33, 0x80, 0x73, 0xc6, 0xf6, 0x73, 0x4a, 0x1e, 0x9f, 0xf8, 0x64, 0x38, 0xab, 0xdf, 0x8c, 0xf3,
	0x62, 0x17, 0x73, 0x4d, 0x3e, 0x40, 0x70, 0xc9, 0x84, 0xf9, 0x8f, 0x92, 0x2f, 0x93, 0x9f, 0xe1,
	0x46, 0x98, 0x6d, 0xb5, 0x9e, 0xa5, 0x32, 0x9f, 0xeb, 0xad, 0x28, 0x72, 0xb1, 0x91, 0xf3, 0xbb,
	0xf7, 0x4a, 0xad, 0xd7, 0x1d, 0x7c, 0x92, 0x3e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x41, 0x50,
	0x6c, 0x1b, 0xcd, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	EditProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*basepb.AnyRes, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.ProductService/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) EditProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.ProductService/EditProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	AddProduct(context.Context, *Product) (*basepb.AnyRes, error)
	EditProduct(context.Context, *Product) (*basepb.AnyRes, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) AddProduct(ctx context.Context, req *Product) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (*UnimplementedProductServiceServer) EditProduct(ctx context.Context, req *Product) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProduct not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ProductService/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_EditProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).EditProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ProductService/EditProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).EditProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "productpb.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProduct",
			Handler:    _ProductService_AddProduct_Handler,
		},
		{
			MethodName: "EditProduct",
			Handler:    _ProductService_EditProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "productpb/product.proto",
}
