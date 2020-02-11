// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type User struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Register struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{2}
}

func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Register) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Register) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterAck struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterAck) Reset()         { *m = RegisterAck{} }
func (m *RegisterAck) String() string { return proto.CompactTextString(m) }
func (*RegisterAck) ProtoMessage()    {}
func (*RegisterAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{3}
}

func (m *RegisterAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAck.Unmarshal(m, b)
}
func (m *RegisterAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAck.Marshal(b, m, deterministic)
}
func (m *RegisterAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAck.Merge(m, src)
}
func (m *RegisterAck) XXX_Size() int {
	return xxx_messageInfo_RegisterAck.Size(m)
}
func (m *RegisterAck) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAck.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAck proto.InternalMessageInfo

func (m *RegisterAck) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegisterAck) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Login struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}
func (*Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{4}
}

func (m *Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login.Unmarshal(m, b)
}
func (m *Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login.Marshal(b, m, deterministic)
}
func (m *Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login.Merge(m, src)
}
func (m *Login) XXX_Size() int {
	return xxx_messageInfo_Login.Size(m)
}
func (m *Login) XXX_DiscardUnknown() {
	xxx_messageInfo_Login.DiscardUnknown(m)
}

var xxx_messageInfo_Login proto.InternalMessageInfo

func (m *Login) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Login) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginAck struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginAck) Reset()         { *m = LoginAck{} }
func (m *LoginAck) String() string { return proto.CompactTextString(m) }
func (*LoginAck) ProtoMessage()    {}
func (*LoginAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{5}
}

func (m *LoginAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginAck.Unmarshal(m, b)
}
func (m *LoginAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginAck.Marshal(b, m, deterministic)
}
func (m *LoginAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginAck.Merge(m, src)
}
func (m *LoginAck) XXX_Size() int {
	return xxx_messageInfo_LoginAck.Size(m)
}
func (m *LoginAck) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginAck.DiscardUnknown(m)
}

var xxx_messageInfo_LoginAck proto.InternalMessageInfo

func (m *LoginAck) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *LoginAck) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "pb.Error")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*Register)(nil), "pb.Register")
	proto.RegisterType((*RegisterAck)(nil), "pb.RegisterAck")
	proto.RegisterType((*Login)(nil), "pb.Login")
	proto.RegisterType((*LoginAck)(nil), "pb.LoginAck")
}

func init() { proto.RegisterFile("pb/message.proto", fileDescriptor_8447775385e7eb85) }

var fileDescriptor_8447775385e7eb85 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x69, 0x6d, 0xa4, 0x9d, 0x5e, 0x24, 0x78, 0x28, 0x22, 0x58, 0x72, 0xda, 0x53, 0x85,
	0xf5, 0xe2, 0x45, 0x44, 0xd0, 0x83, 0xe0, 0x29, 0x20, 0x78, 0xed, 0x9f, 0xa1, 0x84, 0xb5, 0x4d,
	0x98, 0x54, 0xfc, 0xfa, 0x92, 0xc9, 0xae, 0x6c, 0x05, 0xf5, 0x96, 0xf7, 0x98, 0xf7, 0xf2, 0xcb,
	0x04, 0xce, 0x5c, 0x77, 0x3d, 0xa1, 0xf7, 0xed, 0x88, 0x8d, 0x23, 0xbb, 0x58, 0x99, 0xba, 0x4e,
	0xdd, 0x81, 0x78, 0x22, 0xb2, 0x24, 0x25, 0x64, 0xbd, 0x1d, 0xb0, 0x4a, 0xea, 0x64, 0x23, 0x34,
	0x9f, 0x65, 0x0d, 0xe5, 0x80, 0xbe, 0x27, 0xe3, 0x16, 0x63, 0xe7, 0x2a, 0xad, 0x93, 0x4d, 0xa1,
	0x8f, 0x2d, 0x75, 0x0b, 0xd9, 0xab, 0x47, 0x92, 0xe7, 0x20, 0x16, 0xbb, 0xc3, 0x99, 0xe3, 0x85,
	0x8e, 0x42, 0x5e, 0x40, 0xfe, 0xe1, 0x91, 0xe6, 0x76, 0xc2, 0x7d, 0xf8, 0x5b, 0xab, 0x37, 0xc8,
	0x35, 0x8e, 0xc6, 0x2f, 0x48, 0xab, 0xb9, 0x64, 0x3d, 0x17, 0x9a, 0x71, 0x6a, 0xcd, 0xfb, 0xbe,
	0x20, 0x8a, 0x90, 0x70, 0xad, 0xf7, 0x9f, 0x96, 0x86, 0xea, 0x24, 0x26, 0x0e, 0x5a, 0x3d, 0x42,
	0x79, 0x68, 0x7e, 0xe8, 0x77, 0xbf, 0xa0, 0x5d, 0x81, 0xc0, 0xf0, 0x6e, 0xae, 0x2d, 0xb7, 0x45,
	0xe3, 0xba, 0x86, 0x17, 0xa1, 0xa3, 0xaf, 0xee, 0x41, 0xbc, 0xd8, 0xd1, 0xcc, 0x7f, 0xc2, 0x1d,
	0x63, 0xa4, 0x3f, 0x30, 0x9e, 0x21, 0xe7, 0x82, 0xc0, 0x70, 0x09, 0x59, 0xc8, 0x70, 0xbe, 0xdc,
	0xe6, 0xe1, 0xb2, 0xb0, 0x36, 0xcd, 0xee, 0xbf, 0x2c, 0xdd, 0x29, 0xff, 0xd7, 0xcd, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x4f, 0x9f, 0xf8, 0x8a, 0xc3, 0x01, 0x00, 0x00,
}
