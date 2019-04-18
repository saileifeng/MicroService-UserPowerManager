// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/srvusers/proto/srvusers.proto

/*
Package SrvUsers is a generated protocol buffer package.

import public "google/protobuf/timestamp.proto";

It is generated from these files:
	srv/srvusers/proto/srvusers.proto

It has these top-level messages:
	User
	AddRequest
	AddResponse
	GetListRequest
	GetListResponse
	GetRequest
	GetResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
*/
package SrvUsers

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

// Client API for SrvUsers service

type SrvUsersService interface {
	// 添回用户
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	// 获取用户列表
	GetList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResponse, error)
	// 获取单个用户
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	// 修改用户信息
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	// 批量删除用户（假删除）
	BatchDelete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type srvUsersService struct {
	c    client.Client
	name string
}

func NewSrvUsersService(name string, c client.Client) SrvUsersService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "SrvUsers"
	}
	return &srvUsersService{
		c:    c,
		name: name,
	}
}

func (c *srvUsersService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "SrvUsers.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUsersService) GetList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResponse, error) {
	req := c.c.NewRequest(c.name, "SrvUsers.GetList", in)
	out := new(GetListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUsersService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "SrvUsers.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUsersService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "SrvUsers.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUsersService) BatchDelete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "SrvUsers.BatchDelete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SrvUsers service

type SrvUsersHandler interface {
	// 添回用户
	Add(context.Context, *AddRequest, *AddResponse) error
	// 获取用户列表
	GetList(context.Context, *GetListRequest, *GetListResponse) error
	// 获取单个用户
	Get(context.Context, *GetRequest, *GetResponse) error
	// 修改用户信息
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	// 批量删除用户（假删除）
	BatchDelete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterSrvUsersHandler(s server.Server, hdlr SrvUsersHandler, opts ...server.HandlerOption) error {
	type srvUsers interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		GetList(ctx context.Context, in *GetListRequest, out *GetListResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		BatchDelete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type SrvUsers struct {
		srvUsers
	}
	h := &srvUsersHandler{hdlr}
	return s.Handle(s.NewHandler(&SrvUsers{h}, opts...))
}

type srvUsersHandler struct {
	SrvUsersHandler
}

func (h *srvUsersHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.SrvUsersHandler.Add(ctx, in, out)
}

func (h *srvUsersHandler) GetList(ctx context.Context, in *GetListRequest, out *GetListResponse) error {
	return h.SrvUsersHandler.GetList(ctx, in, out)
}

func (h *srvUsersHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.SrvUsersHandler.Get(ctx, in, out)
}

func (h *srvUsersHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.SrvUsersHandler.Update(ctx, in, out)
}

func (h *srvUsersHandler) BatchDelete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.SrvUsersHandler.BatchDelete(ctx, in, out)
}
