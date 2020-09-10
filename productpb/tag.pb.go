// Code generated by protoc-gen-go. DO NOT EDIT.
// source: productpb/tag.proto

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

type Tag struct {
	TagId                uint64   `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id"`
	StoreId              uint64   `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	AdminId              uint64   `protobuf:"varint,4,opt,name=admin_id,json=adminId,proto3" json:"admin_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_64608a916204a8a5, []int{0}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetTagId() uint64 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *Tag) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetAdminId() uint64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

type TagDetail struct {
	TagId                uint64   `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id"`
	StoreId              uint64   `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	CreatedBy            uint64   `protobuf:"varint,4,opt,name=created_by,json=createdBy,proto3" json:"created_by"`
	UpdatedBy            uint64   `protobuf:"varint,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagDetail) Reset()         { *m = TagDetail{} }
func (m *TagDetail) String() string { return proto.CompactTextString(m) }
func (*TagDetail) ProtoMessage()    {}
func (*TagDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_64608a916204a8a5, []int{1}
}

func (m *TagDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagDetail.Unmarshal(m, b)
}
func (m *TagDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagDetail.Marshal(b, m, deterministic)
}
func (m *TagDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagDetail.Merge(m, src)
}
func (m *TagDetail) XXX_Size() int {
	return xxx_messageInfo_TagDetail.Size(m)
}
func (m *TagDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TagDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TagDetail proto.InternalMessageInfo

func (m *TagDetail) GetTagId() uint64 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *TagDetail) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *TagDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TagDetail) GetCreatedBy() uint64 {
	if m != nil {
		return m.CreatedBy
	}
	return 0
}

func (m *TagDetail) GetUpdatedBy() uint64 {
	if m != nil {
		return m.UpdatedBy
	}
	return 0
}

func (m *TagDetail) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *TagDetail) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type DelTagReq struct {
	TagId                []uint64 `protobuf:"varint,1,rep,packed,name=tag_id,json=tagId,proto3" json:"tag_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelTagReq) Reset()         { *m = DelTagReq{} }
func (m *DelTagReq) String() string { return proto.CompactTextString(m) }
func (*DelTagReq) ProtoMessage()    {}
func (*DelTagReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_64608a916204a8a5, []int{2}
}

func (m *DelTagReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelTagReq.Unmarshal(m, b)
}
func (m *DelTagReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelTagReq.Marshal(b, m, deterministic)
}
func (m *DelTagReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelTagReq.Merge(m, src)
}
func (m *DelTagReq) XXX_Size() int {
	return xxx_messageInfo_DelTagReq.Size(m)
}
func (m *DelTagReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelTagReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelTagReq proto.InternalMessageInfo

func (m *DelTagReq) GetTagId() []uint64 {
	if m != nil {
		return m.TagId
	}
	return nil
}

type ListTagReq struct {
	Page                 uint64   `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	PageSize             uint64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id                   uint64   `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTagReq) Reset()         { *m = ListTagReq{} }
func (m *ListTagReq) String() string { return proto.CompactTextString(m) }
func (*ListTagReq) ProtoMessage()    {}
func (*ListTagReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_64608a916204a8a5, []int{3}
}

func (m *ListTagReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTagReq.Unmarshal(m, b)
}
func (m *ListTagReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTagReq.Marshal(b, m, deterministic)
}
func (m *ListTagReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTagReq.Merge(m, src)
}
func (m *ListTagReq) XXX_Size() int {
	return xxx_messageInfo_ListTagReq.Size(m)
}
func (m *ListTagReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTagReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListTagReq proto.InternalMessageInfo

func (m *ListTagReq) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListTagReq) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTagReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ListTagReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListTagRes struct {
	Total                uint64       `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	Tags                 []*TagDetail `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListTagRes) Reset()         { *m = ListTagRes{} }
func (m *ListTagRes) String() string { return proto.CompactTextString(m) }
func (*ListTagRes) ProtoMessage()    {}
func (*ListTagRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_64608a916204a8a5, []int{4}
}

func (m *ListTagRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTagRes.Unmarshal(m, b)
}
func (m *ListTagRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTagRes.Marshal(b, m, deterministic)
}
func (m *ListTagRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTagRes.Merge(m, src)
}
func (m *ListTagRes) XXX_Size() int {
	return xxx_messageInfo_ListTagRes.Size(m)
}
func (m *ListTagRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTagRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListTagRes proto.InternalMessageInfo

func (m *ListTagRes) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListTagRes) GetTags() []*TagDetail {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Tag)(nil), "productpb.Tag")
	proto.RegisterType((*TagDetail)(nil), "productpb.TagDetail")
	proto.RegisterType((*DelTagReq)(nil), "productpb.DelTagReq")
	proto.RegisterType((*ListTagReq)(nil), "productpb.ListTagReq")
	proto.RegisterType((*ListTagRes)(nil), "productpb.ListTagRes")
}

func init() { proto.RegisterFile("productpb/tag.proto", fileDescriptor_64608a916204a8a5) }

var fileDescriptor_64608a916204a8a5 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x45, 0xb6, 0x2c, 0x47, 0x53, 0xf0, 0x61, 0x93, 0x80, 0xea, 0x52, 0x30, 0x3a, 0xa9, 0x17,
	0x09, 0xd2, 0x5b, 0x6f, 0x0e, 0x29, 0xc5, 0x90, 0x93, 0xa2, 0x53, 0x2f, 0x66, 0xa4, 0xdd, 0x6e,
	0x16, 0xac, 0x8f, 0x6a, 0xc7, 0x05, 0xe7, 0x3f, 0xe6, 0x3f, 0x95, 0x5d, 0xad, 0x65, 0x85, 0xf6,
	0xd0, 0x43, 0x4f, 0xbb, 0xf3, 0xde, 0x9b, 0x1d, 0xbf, 0xe7, 0x11, 0x5c, 0x77, 0x7d, 0xcb, 0x8f,
	0x15, 0x75, 0x65, 0x46, 0x28, 0xd3, 0xae, 0x6f, 0xa9, 0x65, 0xe1, 0x08, 0xae, 0xaf, 0x4b, 0xd4,
	0xa2, 0x2b, 0xb3, 0xe1, 0x18, 0xf8, 0xf8, 0x07, 0xcc, 0x0b, 0x94, 0xec, 0x16, 0x02, 0x42, 0xb9,
	0x57, 0x3c, 0xf2, 0x36, 0x5e, 0xe2, 0xe7, 0x0b, 0x42, 0xb9, 0xe3, 0xec, 0x3d, 0x5c, 0x69, 0x6a,
	0x7b, 0x61, 0x88, 0x99, 0x25, 0x96, 0xb6, 0xde, 0x71, 0xc6, 0xc0, 0x6f, 0xb0, 0x16, 0xd1, 0x7c,
	0xe3, 0x25, 0x61, 0x6e, 0xef, 0x46, 0x8e, 0xbc, 0x56, 0x8d, 0x91, 0xfb, 0x83, 0xdc, 0xd6, 0x3b,
	0x1e, 0xbf, 0x7a, 0x10, 0x16, 0x28, 0x1f, 0x04, 0xa1, 0x3a, 0xfc, 0xa7, 0x71, 0x1f, 0x01, 0xaa,
	0x5e, 0x20, 0x09, 0xbe, 0x2f, 0x4f, 0x6e, 0x60, 0xe8, 0x90, 0xfb, 0x93, 0xa1, 0x8f, 0x1d, 0x3f,
	0xd3, 0x8b, 0x81, 0x76, 0xc8, 0x40, 0x9f, 0xbb, 0x91, 0xa2, 0xc0, 0xbe, 0x7b, 0xee, 0xde, 0xd2,
	0xb4, 0x1b, 0x29, 0x5a, 0x0e, 0xb4, 0x43, 0xb6, 0x14, 0xc7, 0x10, 0x3e, 0x88, 0x43, 0x81, 0x32,
	0x17, 0x3f, 0xdf, 0xd8, 0x99, 0x8f, 0x76, 0x62, 0x04, 0x78, 0x54, 0x9a, 0x9c, 0x88, 0x81, 0xdf,
	0xa1, 0x14, 0xce, 0xb1, 0xbd, 0xb3, 0x0f, 0x10, 0x9a, 0x73, 0xaf, 0xd5, 0x8b, 0x70, 0x8e, 0xaf,
	0x0c, 0xf0, 0xa4, 0x5e, 0x04, 0x5b, 0xc1, 0x4c, 0x71, 0x6b, 0xd8, 0xcf, 0x67, 0xea, 0x12, 0x81,
	0x7f, 0x89, 0x20, 0x7e, 0x9c, 0x8c, 0xd0, 0xec, 0x06, 0x16, 0xd4, 0x12, 0x1e, 0xc6, 0x54, 0x4d,
	0xc1, 0x12, 0xf0, 0x09, 0xa5, 0x8e, 0x66, 0x9b, 0x79, 0xf2, 0xee, 0xee, 0x26, 0x1d, 0x37, 0x22,
	0x1d, 0xff, 0x90, 0xdc, 0x2a, 0xee, 0x5e, 0x3d, 0x80, 0x02, 0xe5, 0x93, 0xe8, 0x7f, 0xa9, 0x4a,
	0xb0, 0x04, 0x82, 0x2d, 0xe7, 0x66, 0x3d, 0x56, 0x6f, 0x9b, 0xd6, 0xab, 0xd4, 0x2d, 0xd1, 0xb6,
	0x39, 0x99, 0xc1, 0x9f, 0x60, 0xf9, 0x95, 0x2b, 0xfa, 0x17, 0x69, 0x0a, 0xc1, 0x10, 0x1c, 0x9b,
	0xfe, 0x92, 0x31, 0xcb, 0x3f, 0xf4, 0x5f, 0x00, 0xbe, 0x09, 0xf3, 0xb2, 0xf1, 0xc9, 0x6e, 0x27,
	0x3d, 0x97, 0x6c, 0xd7, 0x7f, 0x85, 0xf5, 0x7d, 0xfc, 0x7d, 0x23, 0x15, 0x3d, 0x1f, 0xcb, 0xb4,
	0x6a, 0xeb, 0x4c, 0x3f, 0xab, 0xa6, 0x56, 0xb2, 0xcd, 0xba, 0x32, 0x1b, 0xe5, 0x65, 0x60, 0xbf,
	0x83, 0xcf, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xc6, 0x06, 0x6b, 0x3e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagServiceClient interface {
	AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	EditTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	DelTag(ctx context.Context, in *DelTagReq, opts ...grpc.CallOption) (*basepb.AnyRes, error)
	GetTagList(ctx context.Context, in *ListTagReq, opts ...grpc.CallOption) (*ListTagRes, error)
}

type tagServiceClient struct {
	cc *grpc.ClientConn
}

func NewTagServiceClient(cc *grpc.ClientConn) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.TagService/AddTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) EditTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.TagService/EditTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) DelTag(ctx context.Context, in *DelTagReq, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/productpb.TagService/DelTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) GetTagList(ctx context.Context, in *ListTagReq, opts ...grpc.CallOption) (*ListTagRes, error) {
	out := new(ListTagRes)
	err := c.cc.Invoke(ctx, "/productpb.TagService/GetTagList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
type TagServiceServer interface {
	AddTag(context.Context, *Tag) (*basepb.AnyRes, error)
	EditTag(context.Context, *Tag) (*basepb.AnyRes, error)
	DelTag(context.Context, *DelTagReq) (*basepb.AnyRes, error)
	GetTagList(context.Context, *ListTagReq) (*ListTagRes, error)
}

// UnimplementedTagServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (*UnimplementedTagServiceServer) AddTag(ctx context.Context, req *Tag) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTag not implemented")
}
func (*UnimplementedTagServiceServer) EditTag(ctx context.Context, req *Tag) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTag not implemented")
}
func (*UnimplementedTagServiceServer) DelTag(ctx context.Context, req *DelTagReq) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelTag not implemented")
}
func (*UnimplementedTagServiceServer) GetTagList(ctx context.Context, req *ListTagReq) (*ListTagRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagList not implemented")
}

func RegisterTagServiceServer(s *grpc.Server, srv TagServiceServer) {
	s.RegisterService(&_TagService_serviceDesc, srv)
}

func _TagService_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.TagService/AddTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).AddTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_EditTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).EditTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.TagService/EditTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).EditTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_DelTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).DelTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.TagService/DelTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).DelTag(ctx, req.(*DelTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_GetTagList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTagList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.TagService/GetTagList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTagList(ctx, req.(*ListTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "productpb.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTag",
			Handler:    _TagService_AddTag_Handler,
		},
		{
			MethodName: "EditTag",
			Handler:    _TagService_EditTag_Handler,
		},
		{
			MethodName: "DelTag",
			Handler:    _TagService_DelTag_Handler,
		},
		{
			MethodName: "GetTagList",
			Handler:    _TagService_GetTagList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "productpb/tag.proto",
}
