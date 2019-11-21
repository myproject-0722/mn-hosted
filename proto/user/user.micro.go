// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package go_mnhosted_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...client.CallOption) (*SignUpResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...client.CallOption) (*SignInResponse, error)
	SignOut(ctx context.Context, in *SignOutRequest, opts ...client.CallOption) (*SignOutResponse, error)
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*GetInfoResponse, error)
	MailCode(ctx context.Context, in *MailCodeRequest, opts ...client.CallOption) (*MailCodeResponse, error)
	Reset(ctx context.Context, in *ResetRequest, opts ...client.CallOption) (*ResetResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.mnhosted.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) SignUp(ctx context.Context, in *SignUpRequest, opts ...client.CallOption) (*SignUpResponse, error) {
	req := c.c.NewRequest(c.name, "User.SignUp", in)
	out := new(SignUpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SignIn(ctx context.Context, in *SignInRequest, opts ...client.CallOption) (*SignInResponse, error) {
	req := c.c.NewRequest(c.name, "User.SignIn", in)
	out := new(SignInResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SignOut(ctx context.Context, in *SignOutRequest, opts ...client.CallOption) (*SignOutResponse, error) {
	req := c.c.NewRequest(c.name, "User.SignOut", in)
	out := new(SignOutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*GetInfoResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetInfo", in)
	out := new(GetInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) MailCode(ctx context.Context, in *MailCodeRequest, opts ...client.CallOption) (*MailCodeResponse, error) {
	req := c.c.NewRequest(c.name, "User.MailCode", in)
	out := new(MailCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Reset(ctx context.Context, in *ResetRequest, opts ...client.CallOption) (*ResetResponse, error) {
	req := c.c.NewRequest(c.name, "User.Reset", in)
	out := new(ResetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	SignUp(context.Context, *SignUpRequest, *SignUpResponse) error
	SignIn(context.Context, *SignInRequest, *SignInResponse) error
	SignOut(context.Context, *SignOutRequest, *SignOutResponse) error
	GetInfo(context.Context, *GetInfoRequest, *GetInfoResponse) error
	MailCode(context.Context, *MailCodeRequest, *MailCodeResponse) error
	Reset(context.Context, *ResetRequest, *ResetResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		SignUp(ctx context.Context, in *SignUpRequest, out *SignUpResponse) error
		SignIn(ctx context.Context, in *SignInRequest, out *SignInResponse) error
		SignOut(ctx context.Context, in *SignOutRequest, out *SignOutResponse) error
		GetInfo(ctx context.Context, in *GetInfoRequest, out *GetInfoResponse) error
		MailCode(ctx context.Context, in *MailCodeRequest, out *MailCodeResponse) error
		Reset(ctx context.Context, in *ResetRequest, out *ResetResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) SignUp(ctx context.Context, in *SignUpRequest, out *SignUpResponse) error {
	return h.UserHandler.SignUp(ctx, in, out)
}

func (h *userHandler) SignIn(ctx context.Context, in *SignInRequest, out *SignInResponse) error {
	return h.UserHandler.SignIn(ctx, in, out)
}

func (h *userHandler) SignOut(ctx context.Context, in *SignOutRequest, out *SignOutResponse) error {
	return h.UserHandler.SignOut(ctx, in, out)
}

func (h *userHandler) GetInfo(ctx context.Context, in *GetInfoRequest, out *GetInfoResponse) error {
	return h.UserHandler.GetInfo(ctx, in, out)
}

func (h *userHandler) MailCode(ctx context.Context, in *MailCodeRequest, out *MailCodeResponse) error {
	return h.UserHandler.MailCode(ctx, in, out)
}

func (h *userHandler) Reset(ctx context.Context, in *ResetRequest, out *ResetResponse) error {
	return h.UserHandler.Reset(ctx, in, out)
}
