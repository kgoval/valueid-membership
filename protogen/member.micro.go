// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: member.proto

/*
Package member is a generated protocol buffer package.

It is generated from these files:
	member.proto

It has these top-level messages:
	MemberDetailRequest
	MemberDetailResponse
*/
package member

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for MemberService service

type MemberService interface {
	Detail(ctx context.Context, in *MemberDetailRequest, opts ...client.CallOption) (*MemberDetailResponse, error)
}

type memberService struct {
	c    client.Client
	name string
}

func NewMemberService(name string, c client.Client) MemberService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "memberservice"
	}
	return &memberService{
		c:    c,
		name: name,
	}
}

func (c *memberService) Detail(ctx context.Context, in *MemberDetailRequest, opts ...client.CallOption) (*MemberDetailResponse, error) {
	req := c.c.NewRequest(c.name, "MemberService.Detail", in)
	out := new(MemberDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MemberService service

type MemberServiceHandler interface {
	Detail(context.Context, *MemberDetailRequest, *MemberDetailResponse) error
}

func RegisterMemberServiceHandler(s server.Server, hdlr MemberServiceHandler, opts ...server.HandlerOption) error {
	type memberService interface {
		Detail(ctx context.Context, in *MemberDetailRequest, out *MemberDetailResponse) error
	}
	type MemberService struct {
		memberService
	}
	h := &memberServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MemberService{h}, opts...))
}

type memberServiceHandler struct {
	MemberServiceHandler
}

func (h *memberServiceHandler) Detail(ctx context.Context, in *MemberDetailRequest, out *MemberDetailResponse) error {
	return h.MemberServiceHandler.Detail(ctx, in, out)
}
