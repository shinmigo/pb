// Code generated by protoc-gen-go. DO NOT EDIT.
// source: member.proto

//option go_package = ".;memberpb";

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

// 通用response
type Response struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 会员信息
type Member struct {
	MemberId             uint64   `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Status               uint32   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	StatusText           string   `protobuf:"bytes,4,opt,name=statusText,proto3" json:"statusText,omitempty"`
	Mobile               string   `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	MemberLevelId        string   `protobuf:"bytes,6,opt,name=memberLevelId,proto3" json:"memberLevelId,omitempty"`
	MemberLevelName      string   `protobuf:"bytes,7,opt,name=memberLevelName,proto3" json:"memberLevelName,omitempty"`
	LastLoginTime        string   `protobuf:"bytes,8,opt,name=lastLoginTime,proto3" json:"lastLoginTime,omitempty"`
	Gender               uint32   `protobuf:"varint,9,opt,name=gender,proto3" json:"gender,omitempty"`
	BirthDay             string   `protobuf:"bytes,10,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{1}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetMemberId() uint64 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

func (m *Member) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Member) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Member) GetStatusText() string {
	if m != nil {
		return m.StatusText
	}
	return ""
}

func (m *Member) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Member) GetMemberLevelId() string {
	if m != nil {
		return m.MemberLevelId
	}
	return ""
}

func (m *Member) GetMemberLevelName() string {
	if m != nil {
		return m.MemberLevelName
	}
	return ""
}

func (m *Member) GetLastLoginTime() string {
	if m != nil {
		return m.LastLoginTime
	}
	return ""
}

func (m *Member) GetGender() uint32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *Member) GetBirthDay() string {
	if m != nil {
		return m.BirthDay
	}
	return ""
}

// 会员列表的参数
type ListRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Status               uint32   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	MemberId             uint64   `protobuf:"varint,3,opt,name=memberId,proto3" json:"memberId,omitempty"`
	PageSize             uint32   `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Page                 uint32   `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *ListRequest) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ListRequest) GetMemberId() uint64 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

func (m *ListRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

// 列表返回数据
type ListResponse struct {
	Members              []*Member `protobuf:"bytes,1,rep,name=Members,proto3" json:"Members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{3}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type DetailRequest struct {
	MemberId             uint64   `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailRequest) Reset()         { *m = DetailRequest{} }
func (m *DetailRequest) String() string { return proto.CompactTextString(m) }
func (*DetailRequest) ProtoMessage()    {}
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{4}
}

func (m *DetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailRequest.Unmarshal(m, b)
}
func (m *DetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailRequest.Marshal(b, m, deterministic)
}
func (m *DetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailRequest.Merge(m, src)
}
func (m *DetailRequest) XXX_Size() int {
	return xxx_messageInfo_DetailRequest.Size(m)
}
func (m *DetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailRequest proto.InternalMessageInfo

func (m *DetailRequest) GetMemberId() uint64 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

// 新增会员的参数
type CreateRequest struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Status               uint32   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Gender               uint32   `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	BirthDay             string   `protobuf:"bytes,5,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	MemberLevelId        string   `protobuf:"bytes,6,opt,name=memberLevelId,proto3" json:"memberLevelId,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{5}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *CreateRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *CreateRequest) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateRequest) GetGender() uint32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *CreateRequest) GetBirthDay() string {
	if m != nil {
		return m.BirthDay
	}
	return ""
}

func (m *CreateRequest) GetMemberLevelId() string {
	if m != nil {
		return m.MemberLevelId
	}
	return ""
}

func (m *CreateRequest) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

// 更新会员的参数
type UpdateRequest struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Status               uint32   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Gender               uint32   `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	BirthDay             string   `protobuf:"bytes,5,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	MemberLevelId        string   `protobuf:"bytes,6,opt,name=memberLevelId,proto3" json:"memberLevelId,omitempty"`
	MemberId             uint64   `protobuf:"varint,7,opt,name=memberId,proto3" json:"memberId,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{6}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UpdateRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UpdateRequest) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateRequest) GetGender() uint32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UpdateRequest) GetBirthDay() string {
	if m != nil {
		return m.BirthDay
	}
	return ""
}

func (m *UpdateRequest) GetMemberLevelId() string {
	if m != nil {
		return m.MemberLevelId
	}
	return ""
}

func (m *UpdateRequest) GetMemberId() uint64 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

func (m *UpdateRequest) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

// 更新会员状态
type UpdateStatusRequest struct {
	MemberId             uint64   `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	Status               uint32   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStatusRequest) Reset()         { *m = UpdateStatusRequest{} }
func (m *UpdateStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateStatusRequest) ProtoMessage()    {}
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{7}
}

func (m *UpdateStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStatusRequest.Unmarshal(m, b)
}
func (m *UpdateStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStatusRequest.Merge(m, src)
}
func (m *UpdateStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateStatusRequest.Size(m)
}
func (m *UpdateStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStatusRequest proto.InternalMessageInfo

func (m *UpdateStatusRequest) GetMemberId() uint64 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

func (m *UpdateStatusRequest) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateStatusRequest) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "memberpb.Response")
	proto.RegisterType((*Member)(nil), "memberpb.Member")
	proto.RegisterType((*ListRequest)(nil), "memberpb.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "memberpb.ListResponse")
	proto.RegisterType((*DetailRequest)(nil), "memberpb.DetailRequest")
	proto.RegisterType((*CreateRequest)(nil), "memberpb.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "memberpb.UpdateRequest")
	proto.RegisterType((*UpdateStatusRequest)(nil), "memberpb.UpdateStatusRequest")
}

func init() { proto.RegisterFile("member.proto", fileDescriptor_9b9836b7e13de206) }

var fileDescriptor_9b9836b7e13de206 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0xa4, 0x4b, 0xd3, 0xb3, 0x46, 0x4c, 0x46, 0x0c, 0x2b, 0x12, 0xa8, 0x8a, 0xb8,
	0xa8, 0x40, 0xaa, 0xd0, 0xb8, 0x18, 0xda, 0xed, 0x26, 0x4d, 0x93, 0x0a, 0x17, 0xe9, 0x78, 0x80,
	0xa4, 0x39, 0x2a, 0x16, 0xcd, 0x1f, 0x62, 0x6f, 0x02, 0xde, 0x80, 0x57, 0x42, 0xbc, 0x04, 0x0f,
	0xc2, 0x3b, 0x20, 0xdb, 0x69, 0x62, 0xb7, 0x74, 0xda, 0xed, 0xee, 0x7c, 0xce, 0xf1, 0xb1, 0xfd,
	0x7d, 0xe7, 0x97, 0xc0, 0xb8, 0xc0, 0x22, 0xc3, 0x66, 0x56, 0x37, 0x95, 0xa8, 0x48, 0xa0, 0xa3,
	0x3a, 0x8b, 0xdf, 0x42, 0x90, 0x20, 0xaf, 0xab, 0x92, 0x23, 0x21, 0x30, 0x58, 0x56, 0x39, 0x52,
	0x67, 0xe2, 0x4c, 0x47, 0x89, 0x5a, 0x93, 0x63, 0xf0, 0x0a, 0xbe, 0xa2, 0xae, 0x4a, 0xc9, 0x65,
	0xfc, 0xcb, 0x05, 0xff, 0x83, 0x6a, 0x27, 0x11, 0xb4, 0x07, 0x5d, 0xe7, 0xaa, 0x69, 0x90, 0x74,
	0xb1, 0xac, 0x95, 0x6c, 0xf9, 0xa5, 0x4c, 0x0b, 0x6c, 0xbb, 0xbb, 0x98, 0x9c, 0x80, 0xcf, 0x45,
	0x2a, 0x6e, 0x39, 0xf5, 0x26, 0xce, 0x34, 0x4c, 0xda, 0x88, 0xbc, 0x04, 0xd0, 0xab, 0x1b, 0xfc,
	0x26, 0xe8, 0x40, 0x75, 0x19, 0x19, 0xd9, 0x57, 0x54, 0x19, 0x5b, 0x23, 0x3d, 0x54, 0xb5, 0x36,
	0x22, 0xaf, 0x20, 0xd4, 0xf7, 0xce, 0xf1, 0x0e, 0xd7, 0xd7, 0x39, 0xf5, 0x55, 0xd9, 0x4e, 0x92,
	0x29, 0x3c, 0x31, 0x12, 0x1f, 0xe5, 0xc3, 0x86, 0x6a, 0xdf, 0x76, 0x5a, 0x9e, 0xb7, 0x4e, 0xb9,
	0x98, 0x57, 0x2b, 0x56, 0xde, 0xb0, 0x02, 0x69, 0xa0, 0xcf, 0xb3, 0x92, 0xf2, 0x35, 0x2b, 0x2c,
	0x73, 0x6c, 0xe8, 0x48, 0xab, 0xd0, 0x91, 0x54, 0x9e, 0xb1, 0x46, 0x7c, 0xbe, 0x4c, 0xbf, 0x53,
	0xd0, 0xca, 0x37, 0x71, 0xfc, 0xd3, 0x81, 0xa3, 0x39, 0xe3, 0x22, 0xc1, 0xaf, 0xb7, 0xc8, 0x4d,
	0x45, 0x8e, 0xa5, 0xa8, 0x77, 0xc8, 0xb5, 0x1c, 0x32, 0x1d, 0xf7, 0x76, 0x1d, 0xaf, 0xd3, 0x15,
	0x2e, 0xd8, 0x0f, 0x54, 0xde, 0x85, 0x49, 0x17, 0xcb, 0xd1, 0xca, 0xb5, 0xf2, 0x2d, 0x4c, 0xd4,
	0x3a, 0x3e, 0x87, 0xb1, 0x7e, 0x4a, 0x3b, 0xfe, 0xd7, 0x30, 0xd4, 0x73, 0xe5, 0xd4, 0x99, 0x78,
	0xd3, 0xa3, 0xd3, 0xe3, 0xd9, 0x06, 0x93, 0x99, 0x2e, 0x24, 0x9b, 0x0d, 0xf1, 0x1b, 0x08, 0x2f,
	0x51, 0xa4, 0x6c, 0xbd, 0x11, 0x72, 0x0f, 0x0a, 0xf1, 0x1f, 0x07, 0xc2, 0x8b, 0x06, 0x53, 0x81,
	0xc6, 0xee, 0x0e, 0x0e, 0x67, 0x17, 0x8e, 0xd6, 0x12, 0x77, 0x8f, 0x25, 0x36, 0x34, 0xfd, 0x18,
	0x06, 0x7b, 0xc7, 0x70, 0x68, 0x8f, 0xe1, 0x81, 0xc0, 0x44, 0x10, 0x54, 0x35, 0x36, 0xa9, 0xa8,
	0x9a, 0x96, 0x94, 0x2e, 0x8e, 0xff, 0x3a, 0x10, 0x7e, 0xaa, 0xf3, 0xc7, 0xa8, 0xa9, 0x9b, 0xd3,
	0x70, 0x17, 0xa0, 0x4e, 0x6f, 0xb0, 0xa5, 0x17, 0xe1, 0xa9, 0x96, 0xbb, 0x50, 0x2f, 0x7c, 0xc0,
	0xd8, 0xef, 0x63, 0xb8, 0xbb, 0xc6, 0xb3, 0xaf, 0x39, 0xfd, 0xed, 0x42, 0xa8, 0x19, 0x5b, 0x60,
	0x73, 0xc7, 0x96, 0x48, 0xce, 0x61, 0x78, 0x85, 0x42, 0x82, 0x4a, 0x9e, 0xf5, 0x3c, 0x1a, 0xdf,
	0x50, 0x74, 0xb2, 0x9d, 0xd6, 0x3c, 0xc7, 0x07, 0xe4, 0x3d, 0x8c, 0xae, 0x50, 0x68, 0x50, 0xc9,
	0xf3, 0x7e, 0x9b, 0x85, 0x6e, 0xb4, 0x83, 0x79, 0x7c, 0x40, 0xce, 0xc0, 0xd7, 0xc4, 0x9a, 0x6d,
	0x16, 0xc3, 0x11, 0xe9, 0x0b, 0xc6, 0x95, 0x67, 0xe0, 0x6b, 0x9f, 0xcc, 0x46, 0x0b, 0x94, 0x3d,
	0x8d, 0x17, 0x30, 0x36, 0x0d, 0x26, 0x2f, 0xb6, 0xdb, 0x2d, 0xe3, 0xff, 0x7f, 0x48, 0xe6, 0xab,
	0xdf, 0xfb, 0xbb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x72, 0xf6, 0x45, 0xee, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MemberServiceClient is the client API for MemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MemberServiceClient interface {
	GetList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	GetDetail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*Member, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*Response, error)
}

type memberServiceClient struct {
	cc *grpc.ClientConn
}

func NewMemberServiceClient(cc *grpc.ClientConn) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) GetList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/memberpb.MemberService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) GetDetail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/memberpb.MemberService/GetDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/memberpb.MemberService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/memberpb.MemberService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/memberpb.MemberService/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServiceServer is the server API for MemberService service.
type MemberServiceServer interface {
	GetList(context.Context, *ListRequest) (*ListResponse, error)
	GetDetail(context.Context, *DetailRequest) (*Member, error)
	Create(context.Context, *CreateRequest) (*Response, error)
	Update(context.Context, *UpdateRequest) (*Response, error)
	UpdateStatus(context.Context, *UpdateStatusRequest) (*Response, error)
}

func RegisterMemberServiceServer(s *grpc.Server, srv MemberServiceServer) {
	s.RegisterService(&_MemberService_serviceDesc, srv)
}

func _MemberService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memberpb.MemberService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetList(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memberpb.MemberService/GetDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetDetail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memberpb.MemberService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memberpb.MemberService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memberpb.MemberService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).UpdateStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MemberService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "memberpb.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _MemberService_GetList_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _MemberService_GetDetail_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _MemberService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MemberService_Update_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _MemberService_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "member.proto",
}
