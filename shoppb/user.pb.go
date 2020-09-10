// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shoppb/user.proto

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

type LoginReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_038558042afd3979, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserRes struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRes) Reset()         { *m = UserRes{} }
func (m *UserRes) String() string { return proto.CompactTextString(m) }
func (*UserRes) ProtoMessage()    {}
func (*UserRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_038558042afd3979, []int{1}
}

func (m *UserRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRes.Unmarshal(m, b)
}
func (m *UserRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRes.Marshal(b, m, deterministic)
}
func (m *UserRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRes.Merge(m, src)
}
func (m *UserRes) XXX_Size() int {
	return xxx_messageInfo_UserRes.Size(m)
}
func (m *UserRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRes.DiscardUnknown(m)
}

var xxx_messageInfo_UserRes proto.InternalMessageInfo

func (m *UserRes) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserRes) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EditUserReq struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	AdminId              uint64   `protobuf:"varint,4,opt,name=admin_id,json=adminId,proto3" json:"admin_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditUserReq) Reset()         { *m = EditUserReq{} }
func (m *EditUserReq) String() string { return proto.CompactTextString(m) }
func (*EditUserReq) ProtoMessage()    {}
func (*EditUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_038558042afd3979, []int{2}
}

func (m *EditUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditUserReq.Unmarshal(m, b)
}
func (m *EditUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditUserReq.Marshal(b, m, deterministic)
}
func (m *EditUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditUserReq.Merge(m, src)
}
func (m *EditUserReq) XXX_Size() int {
	return xxx_messageInfo_EditUserReq.Size(m)
}
func (m *EditUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EditUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_EditUserReq proto.InternalMessageInfo

func (m *EditUserReq) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *EditUserReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *EditUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EditUserReq) GetAdminId() uint64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "shoppb.LoginReq")
	proto.RegisterType((*UserRes)(nil), "shoppb.UserRes")
	proto.RegisterType((*EditUserReq)(nil), "shoppb.EditUserReq")
}

func init() { proto.RegisterFile("shoppb/user.proto", fileDescriptor_038558042afd3979) }

var fileDescriptor_038558042afd3979 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x14, 0xc4, 0xd9, 0xb5, 0xb6, 0xf5, 0x2d, 0xf8, 0x27, 0x7b, 0x70, 0x2d, 0x88, 0xd2, 0x93, 0x78,
	0x48, 0x40, 0x3f, 0x81, 0x0b, 0x1e, 0x16, 0x3c, 0x45, 0xbc, 0x78, 0x91, 0x66, 0x13, 0xda, 0x08,
	0x4d, 0xd2, 0xa4, 0x55, 0xfc, 0xf6, 0x92, 0xc4, 0x88, 0x8a, 0x7a, 0x6a, 0x67, 0x86, 0xf7, 0xcb,
	0xcb, 0x04, 0x8e, 0x5c, 0xa7, 0x8d, 0x61, 0x64, 0x72, 0xc2, 0x62, 0x63, 0xf5, 0xa8, 0x51, 0x1e,
	0xad, 0x6a, 0xc9, 0x1a, 0x27, 0x0c, 0x23, 0xf1, 0x13, 0xc3, 0x7a, 0x0d, 0xe5, 0x9d, 0x6e, 0xa5,
	0xa2, 0x62, 0x40, 0x15, 0x94, 0x7e, 0x4c, 0x35, 0xbd, 0x58, 0xcd, 0xce, 0x67, 0x17, 0x7b, 0xf4,
	0x53, 0xfb, 0xcc, 0x34, 0xce, 0xbd, 0x6a, 0xcb, 0x57, 0xf3, 0x98, 0x25, 0x5d, 0x53, 0x28, 0x1e,
	0x9c, 0xb0, 0x54, 0x38, 0x74, 0x0c, 0x85, 0x1f, 0x79, 0x92, 0x3c, 0x10, 0x32, 0x9a, 0x7b, 0xb9,
	0xe1, 0xdf, 0xd8, 0xf3, 0x1f, 0x6c, 0x04, 0x59, 0xf0, 0x77, 0x82, 0x1f, 0xfe, 0xeb, 0x01, 0x16,
	0xb7, 0x5c, 0x8e, 0x91, 0x3b, 0xfc, 0xcb, 0xfd, 0x6b, 0xaf, 0xdf, 0xb8, 0xe8, 0x04, 0xca, 0x86,
	0xf7, 0x52, 0x79, 0x52, 0x16, 0x48, 0x45, 0xd0, 0x1b, 0x7e, 0xf5, 0x0c, 0x0b, 0x7f, 0xdc, 0xbd,
	0xb0, 0x2f, 0x72, 0x2b, 0xd0, 0x25, 0xec, 0x86, 0x66, 0xd0, 0x21, 0x8e, 0x05, 0xe2, 0x54, 0x54,
	0x75, 0x90, 0x9c, 0x74, 0x6d, 0x02, 0x65, 0xda, 0x16, 0x2d, 0x53, 0xf8, 0x65, 0xff, 0x6a, 0x1f,
	0x7f, 0xb4, 0x7e, 0xa3, 0xde, 0xa8, 0x70, 0xeb, 0xb3, 0xc7, 0xd3, 0x56, 0x8e, 0xdd, 0xc4, 0xf0,
	0x56, 0xf7, 0xc4, 0x75, 0x52, 0xf5, 0xb2, 0xd5, 0xc4, 0x30, 0x12, 0x87, 0x59, 0x1e, 0x9e, 0xe7,
	0xfa, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xa0, 0xf8, 0x56, 0xd0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*UserRes, error)
	EditUser(ctx context.Context, in *EditUserReq, opts ...grpc.CallOption) (*basepb.AnyRes, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/shoppb.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditUser(ctx context.Context, in *EditUserReq, opts ...grpc.CallOption) (*basepb.AnyRes, error) {
	out := new(basepb.AnyRes)
	err := c.cc.Invoke(ctx, "/shoppb.UserService/EditUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Login(context.Context, *LoginReq) (*UserRes, error)
	EditUser(context.Context, *EditUserReq) (*basepb.AnyRes, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Login(ctx context.Context, req *LoginReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServiceServer) EditUser(ctx context.Context, req *EditUserReq) (*basepb.AnyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppb.UserService/EditUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditUser(ctx, req.(*EditUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shoppb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "EditUser",
			Handler:    _UserService_EditUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shoppb/user.proto",
}
