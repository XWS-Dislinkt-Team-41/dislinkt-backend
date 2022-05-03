// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user_service/user_service.proto

package user

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetAllPublicUserId(ctx context.Context, in *GetAllPublicUserIdRequest, opts ...grpc.CallOption) (*GetAllPublicUserIdResponse, error)
	IsPrivate(ctx context.Context, in *IsPrivateRequest, opts ...grpc.CallOption) (*IsPrivateResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	SearchPublic(ctx context.Context, in *SearchPublicRequest, opts ...grpc.CallOption) (*SearchPublicResponse, error)
	UpdatePersonalInfo(ctx context.Context, in *UpdatePersonalInfoRequest, opts ...grpc.CallOption) (*UpdatePersonalInfoResponse, error)
	UpdateCareerInfo(ctx context.Context, in *UpdateCareerInfoRequest, opts ...grpc.CallOption) (*UpdateCareerInfoResponse, error)
	UpdateInterestsInfo(ctx context.Context, in *UpdateInterestsInfoRequest, opts ...grpc.CallOption) (*UpdateInterestsInfoResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllPublicUserId(ctx context.Context, in *GetAllPublicUserIdRequest, opts ...grpc.CallOption) (*GetAllPublicUserIdResponse, error) {
	out := new(GetAllPublicUserIdResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetAllPublicUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsPrivate(ctx context.Context, in *IsPrivateRequest, opts ...grpc.CallOption) (*IsPrivateResponse, error) {
	out := new(IsPrivateResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/IsPrivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SearchPublic(ctx context.Context, in *SearchPublicRequest, opts ...grpc.CallOption) (*SearchPublicResponse, error) {
	out := new(SearchPublicResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/SearchPublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePersonalInfo(ctx context.Context, in *UpdatePersonalInfoRequest, opts ...grpc.CallOption) (*UpdatePersonalInfoResponse, error) {
	out := new(UpdatePersonalInfoResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdatePersonalInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateCareerInfo(ctx context.Context, in *UpdateCareerInfoRequest, opts ...grpc.CallOption) (*UpdateCareerInfoResponse, error) {
	out := new(UpdateCareerInfoResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateCareerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateInterestsInfo(ctx context.Context, in *UpdateInterestsInfoRequest, opts ...grpc.CallOption) (*UpdateInterestsInfoResponse, error) {
	out := new(UpdateInterestsInfoResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateInterestsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetAllPublicUserId(context.Context, *GetAllPublicUserIdRequest) (*GetAllPublicUserIdResponse, error)
	IsPrivate(context.Context, *IsPrivateRequest) (*IsPrivateResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	SearchPublic(context.Context, *SearchPublicRequest) (*SearchPublicResponse, error)
	UpdatePersonalInfo(context.Context, *UpdatePersonalInfoRequest) (*UpdatePersonalInfoResponse, error)
	UpdateCareerInfo(context.Context, *UpdateCareerInfoRequest) (*UpdateCareerInfoResponse, error)
	UpdateInterestsInfo(context.Context, *UpdateInterestsInfoRequest) (*UpdateInterestsInfoResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserServiceServer) GetAllPublicUserId(context.Context, *GetAllPublicUserIdRequest) (*GetAllPublicUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPublicUserId not implemented")
}
func (UnimplementedUserServiceServer) IsPrivate(context.Context, *IsPrivateRequest) (*IsPrivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPrivate not implemented")
}
func (UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) SearchPublic(context.Context, *SearchPublicRequest) (*SearchPublicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPublic not implemented")
}
func (UnimplementedUserServiceServer) UpdatePersonalInfo(context.Context, *UpdatePersonalInfoRequest) (*UpdatePersonalInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePersonalInfo not implemented")
}
func (UnimplementedUserServiceServer) UpdateCareerInfo(context.Context, *UpdateCareerInfoRequest) (*UpdateCareerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCareerInfo not implemented")
}
func (UnimplementedUserServiceServer) UpdateInterestsInfo(context.Context, *UpdateInterestsInfoRequest) (*UpdateInterestsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInterestsInfo not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllPublicUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPublicUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllPublicUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetAllPublicUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllPublicUserId(ctx, req.(*GetAllPublicUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsPrivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsPrivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsPrivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/IsPrivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsPrivate(ctx, req.(*IsPrivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SearchPublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPublicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SearchPublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/SearchPublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SearchPublic(ctx, req.(*SearchPublicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePersonalInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonalInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePersonalInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdatePersonalInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePersonalInfo(ctx, req.(*UpdatePersonalInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateCareerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCareerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateCareerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateCareerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateCareerInfo(ctx, req.(*UpdateCareerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateInterestsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInterestsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateInterestsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateInterestsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateInterestsInfo(ctx, req.(*UpdateInterestsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "GetAllPublicUserId",
			Handler:    _UserService_GetAllPublicUserId_Handler,
		},
		{
			MethodName: "IsPrivate",
			Handler:    _UserService_IsPrivate_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "SearchPublic",
			Handler:    _UserService_SearchPublic_Handler,
		},
		{
			MethodName: "UpdatePersonalInfo",
			Handler:    _UserService_UpdatePersonalInfo_Handler,
		},
		{
			MethodName: "UpdateCareerInfo",
			Handler:    _UserService_UpdateCareerInfo_Handler,
		},
		{
			MethodName: "UpdateInterestsInfo",
			Handler:    _UserService_UpdateInterestsInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service/user_service.proto",
}
