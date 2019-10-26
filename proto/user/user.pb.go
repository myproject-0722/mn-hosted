// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package go_mnhosted_srv_user

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

type SignUpRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *SignUpRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type SignUpResponse struct {
	Rescode              int32    `protobuf:"varint,1,opt,name=rescode,proto3" json:"rescode,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpResponse) Reset()         { *m = SignUpResponse{} }
func (m *SignUpResponse) String() string { return proto.CompactTextString(m) }
func (*SignUpResponse) ProtoMessage()    {}
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *SignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResponse.Unmarshal(m, b)
}
func (m *SignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResponse.Marshal(b, m, deterministic)
}
func (m *SignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResponse.Merge(m, src)
}
func (m *SignUpResponse) XXX_Size() int {
	return xxx_messageInfo_SignUpResponse.Size(m)
}
func (m *SignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResponse proto.InternalMessageInfo

func (m *SignUpResponse) GetRescode() int32 {
	if m != nil {
		return m.Rescode
	}
	return 0
}

func (m *SignUpResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SignUpResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SignInRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (m *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(m, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *SignInRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type SignInResponse struct {
	Rescode              int32    `protobuf:"varint,1,opt,name=rescode,proto3" json:"rescode,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResponse) Reset()         { *m = SignInResponse{} }
func (m *SignInResponse) String() string { return proto.CompactTextString(m) }
func (*SignInResponse) ProtoMessage()    {}
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *SignInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResponse.Unmarshal(m, b)
}
func (m *SignInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResponse.Marshal(b, m, deterministic)
}
func (m *SignInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResponse.Merge(m, src)
}
func (m *SignInResponse) XXX_Size() int {
	return xxx_messageInfo_SignInResponse.Size(m)
}
func (m *SignInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResponse proto.InternalMessageInfo

func (m *SignInResponse) GetRescode() int32 {
	if m != nil {
		return m.Rescode
	}
	return 0
}

func (m *SignInResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SignInResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SignInResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*SignUpRequest)(nil), "go.mnhosted.srv.user.SignUpRequest")
	proto.RegisterType((*SignUpResponse)(nil), "go.mnhosted.srv.user.SignUpResponse")
	proto.RegisterType((*SignInRequest)(nil), "go.mnhosted.srv.user.SignInRequest")
	proto.RegisterType((*SignInResponse)(nil), "go.mnhosted.srv.user.SignInResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x9b, 0xa4, 0x8d, 0x38, 0x60, 0x91, 0xa1, 0xc8, 0xe2, 0xa9, 0x44, 0x0f, 0x3d, 0xed,
	0x41, 0x9f, 0xc0, 0x63, 0xc1, 0x53, 0x24, 0x0f, 0xd0, 0x66, 0x87, 0x18, 0xa4, 0x3b, 0x71, 0x67,
	0xa3, 0x2f, 0xe6, 0x03, 0x4a, 0xb2, 0x59, 0x45, 0x90, 0x1c, 0xf4, 0xb6, 0x3f, 0x33, 0x7c, 0x7c,
	0xfb, 0x33, 0x00, 0xbd, 0x90, 0xd3, 0x9d, 0x63, 0xcf, 0xb8, 0x69, 0x58, 0x9f, 0xec, 0x33, 0x8b,
	0x27, 0xa3, 0xc5, 0xbd, 0xe9, 0x61, 0x56, 0x3c, 0xc0, 0xc5, 0x53, 0xdb, 0xd8, 0xaa, 0x2b, 0xe9,
	0xb5, 0x27, 0xf1, 0xa8, 0xe0, 0xec, 0x50, 0xd7, 0xdc, 0x5b, 0xaf, 0x92, 0x6d, 0xb2, 0x3b, 0x2f,
	0x63, 0xc4, 0x2b, 0xc8, 0xbb, 0x83, 0xc8, 0xbb, 0x51, 0xe9, 0x38, 0x98, 0x52, 0xf1, 0x08, 0xeb,
	0x88, 0x90, 0x8e, 0xad, 0xd0, 0xc0, 0x70, 0x24, 0x35, 0x1b, 0x1a, 0x19, 0xab, 0x32, 0x46, 0xbc,
	0x84, 0xec, 0x24, 0xcd, 0x04, 0x18, 0x9e, 0xb8, 0x86, 0xb4, 0x35, 0x2a, 0xdb, 0x26, 0xbb, 0xac,
	0x4c, 0x5b, 0x13, 0x85, 0xf6, 0xf6, 0xef, 0x42, 0xc7, 0x20, 0x34, 0x20, 0xfe, 0x2f, 0x84, 0x1b,
	0x58, 0x79, 0x7e, 0x21, 0xab, 0x96, 0xe3, 0x4e, 0x08, 0x77, 0x1f, 0x09, 0x2c, 0x2b, 0x21, 0x87,
	0x15, 0xe4, 0xe1, 0xf7, 0x78, 0xa3, 0x7f, 0x6b, 0x58, 0xff, 0xa8, 0xf7, 0xfa, 0x76, 0x7e, 0x29,
	0xf8, 0x16, 0x8b, 0x88, 0xdd, 0xdb, 0x39, 0xec, 0x57, 0x49, 0x73, 0xd8, 0xef, 0x1a, 0x8a, 0xc5,
	0x31, 0x1f, 0x6f, 0xe1, 0xfe, 0x33, 0x00, 0x00, 0xff, 0xff, 0xac, 0x5a, 0xc9, 0xba, 0x19, 0x02,
	0x00, 0x00,
}
