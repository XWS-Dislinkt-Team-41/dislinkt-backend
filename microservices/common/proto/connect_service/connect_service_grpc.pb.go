// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: connect_service.proto

package connections

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

// ConnectServiceClient is the client API for ConnectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectServiceClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectionResponse, error)
	UnConnect(ctx context.Context, in *UnConnectRequest, opts ...grpc.CallOption) (*EmptyRespones, error)
	GetUserConnections(ctx context.Context, in *GetUserConnectionsRequest, opts ...grpc.CallOption) (*GetUserConnectionsResponse, error)
	AcceptInvitation(ctx context.Context, in *AcceptInvitationRequest, opts ...grpc.CallOption) (*ConnectionResponse, error)
	DeclineInvitation(ctx context.Context, in *DeclineInvitationRequest, opts ...grpc.CallOption) (*EmptyRespones, error)
	GetAllInvitations(ctx context.Context, in *GetAllUserInvitationsRequest, opts ...grpc.CallOption) (*GetAllInvitationsResponse, error)
	GetAllSentInvitations(ctx context.Context, in *GetAllSentInvitationsRequest, opts ...grpc.CallOption) (*GetAllInvitationsResponse, error)
}

type connectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectServiceClient(cc grpc.ClientConnInterface) ConnectServiceClient {
	return &connectServiceClient{cc}
}

func (c *connectServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connections.ConnectService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) UnConnect(ctx context.Context, in *UnConnectRequest, opts ...grpc.CallOption) (*EmptyRespones, error) {
	out := new(EmptyRespones)
	err := c.cc.Invoke(ctx, "/connections.ConnectService/UnConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) GetUserConnections(ctx context.Context, in *GetUserConnectionsRequest, opts ...grpc.CallOption) (*GetUserConnectionsResponse, error) {
	out := new(GetUserConnectionsResponse)
	err := c.cc.Invoke(ctx, "/connections.ConnectService/GetUserConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) AcceptInvitation(ctx context.Context, in *AcceptInvitationRequest, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connections.ConnectService/AcceptInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) DeclineInvitation(ctx context.Context, in *DeclineInvitationRequest, opts ...grpc.CallOption) (*EmptyRespones, error) {
	out := new(EmptyRespones)
	err := c.cc.Invoke(ctx, "/connections.ConnectService/DeclineInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) GetAllInvitations(ctx context.Context, in *GetAllUserInvitationsRequest, opts ...grpc.CallOption) (*GetAllInvitationsResponse, error) {
	out := new(GetAllInvitationsResponse)
	err := c.cc.Invoke(ctx, "/connections.ConnectService/GetAllInvitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) GetAllSentInvitations(ctx context.Context, in *GetAllSentInvitationsRequest, opts ...grpc.CallOption) (*GetAllInvitationsResponse, error) {
	out := new(GetAllInvitationsResponse)
	err := c.cc.Invoke(ctx, "/connections.ConnectService/GetAllSentInvitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectServiceServer is the server API for ConnectService service.
// All implementations must embed UnimplementedConnectServiceServer
// for forward compatibility
type ConnectServiceServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectionResponse, error)
	UnConnect(context.Context, *UnConnectRequest) (*EmptyRespones, error)
	GetUserConnections(context.Context, *GetUserConnectionsRequest) (*GetUserConnectionsResponse, error)
	AcceptInvitation(context.Context, *AcceptInvitationRequest) (*ConnectionResponse, error)
	DeclineInvitation(context.Context, *DeclineInvitationRequest) (*EmptyRespones, error)
	GetAllInvitations(context.Context, *GetAllUserInvitationsRequest) (*GetAllInvitationsResponse, error)
	GetAllSentInvitations(context.Context, *GetAllSentInvitationsRequest) (*GetAllInvitationsResponse, error)
	mustEmbedUnimplementedConnectServiceServer()
}

// UnimplementedConnectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectServiceServer struct {
}

func (UnimplementedConnectServiceServer) Connect(context.Context, *ConnectRequest) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedConnectServiceServer) UnConnect(context.Context, *UnConnectRequest) (*EmptyRespones, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnConnect not implemented")
}
func (UnimplementedConnectServiceServer) GetUserConnections(context.Context, *GetUserConnectionsRequest) (*GetUserConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserConnections not implemented")
}
func (UnimplementedConnectServiceServer) AcceptInvitation(context.Context, *AcceptInvitationRequest) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptInvitation not implemented")
}
func (UnimplementedConnectServiceServer) DeclineInvitation(context.Context, *DeclineInvitationRequest) (*EmptyRespones, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeclineInvitation not implemented")
}
func (UnimplementedConnectServiceServer) GetAllInvitations(context.Context, *GetAllUserInvitationsRequest) (*GetAllInvitationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllInvitations not implemented")
}
func (UnimplementedConnectServiceServer) GetAllSentInvitations(context.Context, *GetAllSentInvitationsRequest) (*GetAllInvitationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSentInvitations not implemented")
}
func (UnimplementedConnectServiceServer) mustEmbedUnimplementedConnectServiceServer() {}

// UnsafeConnectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectServiceServer will
// result in compilation errors.
type UnsafeConnectServiceServer interface {
	mustEmbedUnimplementedConnectServiceServer()
}

func RegisterConnectServiceServer(s grpc.ServiceRegistrar, srv ConnectServiceServer) {
	s.RegisterService(&ConnectService_ServiceDesc, srv)
}

func _ConnectService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connections.ConnectService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_UnConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).UnConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connections.ConnectService/UnConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).UnConnect(ctx, req.(*UnConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_GetUserConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserConnectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).GetUserConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connections.ConnectService/GetUserConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).GetUserConnections(ctx, req.(*GetUserConnectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_AcceptInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).AcceptInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connections.ConnectService/AcceptInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).AcceptInvitation(ctx, req.(*AcceptInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_DeclineInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclineInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).DeclineInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connections.ConnectService/DeclineInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).DeclineInvitation(ctx, req.(*DeclineInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_GetAllInvitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserInvitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).GetAllInvitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connections.ConnectService/GetAllInvitations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).GetAllInvitations(ctx, req.(*GetAllUserInvitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_GetAllSentInvitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSentInvitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).GetAllSentInvitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connections.ConnectService/GetAllSentInvitations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).GetAllSentInvitations(ctx, req.(*GetAllSentInvitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectService_ServiceDesc is the grpc.ServiceDesc for ConnectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connections.ConnectService",
	HandlerType: (*ConnectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _ConnectService_Connect_Handler,
		},
		{
			MethodName: "UnConnect",
			Handler:    _ConnectService_UnConnect_Handler,
		},
		{
			MethodName: "GetUserConnections",
			Handler:    _ConnectService_GetUserConnections_Handler,
		},
		{
			MethodName: "AcceptInvitation",
			Handler:    _ConnectService_AcceptInvitation_Handler,
		},
		{
			MethodName: "DeclineInvitation",
			Handler:    _ConnectService_DeclineInvitation_Handler,
		},
		{
			MethodName: "GetAllInvitations",
			Handler:    _ConnectService_GetAllInvitations_Handler,
		},
		{
			MethodName: "GetAllSentInvitations",
			Handler:    _ConnectService_GetAllSentInvitations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connect_service.proto",
}
