// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: proto/users.proto

package session

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Users_CreateUser_FullMethodName               = "/session.Users/CreateUser"
	Users_RemoveUser_FullMethodName               = "/session.Users/RemoveUser"
	Users_HasUser_FullMethodName                  = "/session.Users/HasUser"
	Users_GetUser_FullMethodName                  = "/session.Users/GetUser"
	Users_ChangeUserPassword_FullMethodName       = "/session.Users/ChangeUserPassword"
	Users_ChangeUserName_FullMethodName           = "/session.Users/ChangeUserName"
	Users_GetUserDataByUuid_FullMethodName        = "/session.Users/GetUserDataByUuid"
	Users_GetUserPreview_FullMethodName           = "/session.Users/GetUserPreview"
	Users_ChangeUserPasswordByUuid_FullMethodName = "/session.Users/ChangeUserPasswordByUuid"
	Users_ChangeUserNameByUuid_FullMethodName     = "/session.Users/ChangeUserNameByUuid"
	Users_ChangeUserAvatarByUuid_FullMethodName   = "/session.Users/ChangeUserAvatarByUuid"
)

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	HasUser(ctx context.Context, in *HasUserRequest, opts ...grpc.CallOption) (*HasUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*ChangeUserPasswordResponse, error)
	ChangeUserName(ctx context.Context, in *ChangeUserNameRequest, opts ...grpc.CallOption) (*ChangeUserNameResponse, error)
	GetUserDataByUuid(ctx context.Context, in *GetUserDataByUuidRequest, opts ...grpc.CallOption) (*GetUserDataByUuidResponse, error)
	GetUserPreview(ctx context.Context, in *GetUserPreviewRequest, opts ...grpc.CallOption) (*GetUserPreviewResponse, error)
	ChangeUserPasswordByUuid(ctx context.Context, in *ChangeUserPasswordByUuidRequest, opts ...grpc.CallOption) (*ChangeUserPasswordByUuidResponse, error)
	ChangeUserNameByUuid(ctx context.Context, in *ChangeUserNameByUuidRequest, opts ...grpc.CallOption) (*ChangeUserNameByUuidResponse, error)
	ChangeUserAvatarByUuid(ctx context.Context, in *ChangeUserAvatarByUuidRequest, opts ...grpc.CallOption) (*ChangeUserAvatarByUuidResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Users_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, Users_RemoveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) HasUser(ctx context.Context, in *HasUserRequest, opts ...grpc.CallOption) (*HasUserResponse, error) {
	out := new(HasUserResponse)
	err := c.cc.Invoke(ctx, Users_HasUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, Users_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*ChangeUserPasswordResponse, error) {
	out := new(ChangeUserPasswordResponse)
	err := c.cc.Invoke(ctx, Users_ChangeUserPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ChangeUserName(ctx context.Context, in *ChangeUserNameRequest, opts ...grpc.CallOption) (*ChangeUserNameResponse, error) {
	out := new(ChangeUserNameResponse)
	err := c.cc.Invoke(ctx, Users_ChangeUserName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserDataByUuid(ctx context.Context, in *GetUserDataByUuidRequest, opts ...grpc.CallOption) (*GetUserDataByUuidResponse, error) {
	out := new(GetUserDataByUuidResponse)
	err := c.cc.Invoke(ctx, Users_GetUserDataByUuid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserPreview(ctx context.Context, in *GetUserPreviewRequest, opts ...grpc.CallOption) (*GetUserPreviewResponse, error) {
	out := new(GetUserPreviewResponse)
	err := c.cc.Invoke(ctx, Users_GetUserPreview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ChangeUserPasswordByUuid(ctx context.Context, in *ChangeUserPasswordByUuidRequest, opts ...grpc.CallOption) (*ChangeUserPasswordByUuidResponse, error) {
	out := new(ChangeUserPasswordByUuidResponse)
	err := c.cc.Invoke(ctx, Users_ChangeUserPasswordByUuid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ChangeUserNameByUuid(ctx context.Context, in *ChangeUserNameByUuidRequest, opts ...grpc.CallOption) (*ChangeUserNameByUuidResponse, error) {
	out := new(ChangeUserNameByUuidResponse)
	err := c.cc.Invoke(ctx, Users_ChangeUserNameByUuid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ChangeUserAvatarByUuid(ctx context.Context, in *ChangeUserAvatarByUuidRequest, opts ...grpc.CallOption) (*ChangeUserAvatarByUuidResponse, error) {
	out := new(ChangeUserAvatarByUuidResponse)
	err := c.cc.Invoke(ctx, Users_ChangeUserAvatarByUuid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	HasUser(context.Context, *HasUserRequest) (*HasUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error)
	ChangeUserName(context.Context, *ChangeUserNameRequest) (*ChangeUserNameResponse, error)
	GetUserDataByUuid(context.Context, *GetUserDataByUuidRequest) (*GetUserDataByUuidResponse, error)
	GetUserPreview(context.Context, *GetUserPreviewRequest) (*GetUserPreviewResponse, error)
	ChangeUserPasswordByUuid(context.Context, *ChangeUserPasswordByUuidRequest) (*ChangeUserPasswordByUuidResponse, error)
	ChangeUserNameByUuid(context.Context, *ChangeUserNameByUuidRequest) (*ChangeUserNameByUuidResponse, error)
	ChangeUserAvatarByUuid(context.Context, *ChangeUserAvatarByUuidRequest) (*ChangeUserAvatarByUuidResponse, error)
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedUsersServer) HasUser(context.Context, *HasUserRequest) (*HasUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasUser not implemented")
}
func (UnimplementedUsersServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUsersServer) ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPassword not implemented")
}
func (UnimplementedUsersServer) ChangeUserName(context.Context, *ChangeUserNameRequest) (*ChangeUserNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserName not implemented")
}
func (UnimplementedUsersServer) GetUserDataByUuid(context.Context, *GetUserDataByUuidRequest) (*GetUserDataByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDataByUuid not implemented")
}
func (UnimplementedUsersServer) GetUserPreview(context.Context, *GetUserPreviewRequest) (*GetUserPreviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPreview not implemented")
}
func (UnimplementedUsersServer) ChangeUserPasswordByUuid(context.Context, *ChangeUserPasswordByUuidRequest) (*ChangeUserPasswordByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPasswordByUuid not implemented")
}
func (UnimplementedUsersServer) ChangeUserNameByUuid(context.Context, *ChangeUserNameByUuidRequest) (*ChangeUserNameByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserNameByUuid not implemented")
}
func (UnimplementedUsersServer) ChangeUserAvatarByUuid(context.Context, *ChangeUserAvatarByUuidRequest) (*ChangeUserAvatarByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserAvatarByUuid not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_HasUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).HasUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_HasUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).HasUser(ctx, req.(*HasUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ChangeUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ChangeUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ChangeUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ChangeUserPassword(ctx, req.(*ChangeUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ChangeUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ChangeUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ChangeUserName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ChangeUserName(ctx, req.(*ChangeUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserDataByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDataByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserDataByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetUserDataByUuid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserDataByUuid(ctx, req.(*GetUserDataByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserPreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPreviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserPreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetUserPreview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserPreview(ctx, req.(*GetUserPreviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ChangeUserPasswordByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserPasswordByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ChangeUserPasswordByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ChangeUserPasswordByUuid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ChangeUserPasswordByUuid(ctx, req.(*ChangeUserPasswordByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ChangeUserNameByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserNameByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ChangeUserNameByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ChangeUserNameByUuid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ChangeUserNameByUuid(ctx, req.(*ChangeUserNameByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ChangeUserAvatarByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserAvatarByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ChangeUserAvatarByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ChangeUserAvatarByUuid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ChangeUserAvatarByUuid(ctx, req.(*ChangeUserAvatarByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "session.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _Users_RemoveUser_Handler,
		},
		{
			MethodName: "HasUser",
			Handler:    _Users_HasUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Users_GetUser_Handler,
		},
		{
			MethodName: "ChangeUserPassword",
			Handler:    _Users_ChangeUserPassword_Handler,
		},
		{
			MethodName: "ChangeUserName",
			Handler:    _Users_ChangeUserName_Handler,
		},
		{
			MethodName: "GetUserDataByUuid",
			Handler:    _Users_GetUserDataByUuid_Handler,
		},
		{
			MethodName: "GetUserPreview",
			Handler:    _Users_GetUserPreview_Handler,
		},
		{
			MethodName: "ChangeUserPasswordByUuid",
			Handler:    _Users_ChangeUserPasswordByUuid_Handler,
		},
		{
			MethodName: "ChangeUserNameByUuid",
			Handler:    _Users_ChangeUserNameByUuid_Handler,
		},
		{
			MethodName: "ChangeUserAvatarByUuid",
			Handler:    _Users_ChangeUserAvatarByUuid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/users.proto",
}
