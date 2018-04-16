// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	proto/auth.proto

It has these top-level messages:
	Token
	Tokens
	LoginRequest
	UserInfo
	Empty
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/murillio4/qa-micro/user-service/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AuthService service

type AuthServiceClient interface {
	CreateTokens(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*Tokens, error)
	ValidateAuthToken(ctx context.Context, in *Token, opts ...client.CallOption) (*UserInfo, error)
	RefreshAuthToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Tokens, error)
	DeleteRefreshToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Empty, error)
}

type authServiceClient struct {
	c           client.Client
	serviceName string
}

func NewAuthServiceClient(serviceName string, c client.Client) AuthServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authServiceClient) CreateTokens(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*Tokens, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.CreateTokens", in)
	out := new(Tokens)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidateAuthToken(ctx context.Context, in *Token, opts ...client.CallOption) (*UserInfo, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.ValidateAuthToken", in)
	out := new(UserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshAuthToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Tokens, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.RefreshAuthToken", in)
	out := new(Tokens)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteRefreshToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.DeleteRefreshToken", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceHandler interface {
	CreateTokens(context.Context, *LoginRequest, *Tokens) error
	ValidateAuthToken(context.Context, *Token, *UserInfo) error
	RefreshAuthToken(context.Context, *Token, *Tokens) error
	DeleteRefreshToken(context.Context, *Token, *Empty) error
}

func RegisterAuthServiceHandler(s server.Server, hdlr AuthServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthService{hdlr}, opts...))
}

type AuthService struct {
	AuthServiceHandler
}

func (h *AuthService) CreateTokens(ctx context.Context, in *LoginRequest, out *Tokens) error {
	return h.AuthServiceHandler.CreateTokens(ctx, in, out)
}

func (h *AuthService) ValidateAuthToken(ctx context.Context, in *Token, out *UserInfo) error {
	return h.AuthServiceHandler.ValidateAuthToken(ctx, in, out)
}

func (h *AuthService) RefreshAuthToken(ctx context.Context, in *Token, out *Tokens) error {
	return h.AuthServiceHandler.RefreshAuthToken(ctx, in, out)
}

func (h *AuthService) DeleteRefreshToken(ctx context.Context, in *Token, out *Empty) error {
	return h.AuthServiceHandler.DeleteRefreshToken(ctx, in, out)
}
