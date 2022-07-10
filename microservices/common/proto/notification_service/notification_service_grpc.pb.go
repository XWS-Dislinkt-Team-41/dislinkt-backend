// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: notification_service/notification_service.proto

package notification

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

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	GetAllNotifications(ctx context.Context, in *GetAllNotificationsRequest, opts ...grpc.CallOption) (*GetAllNotificationsResponse, error)
	InsertNotification(ctx context.Context, in *InsertNotificationRequest, opts ...grpc.CallOption) (*InsertNotificationResponse, error)
	MarkAllAsSeen(ctx context.Context, in *MarkAllAsSeenRequest, opts ...grpc.CallOption) (*EmptyRespones, error)
	GetUserSettings(ctx context.Context, in *GetUserSettingsRequest, opts ...grpc.CallOption) (*GetUserSettingsResponse, error)
	UpdateUserSettings(ctx context.Context, in *UpdateUserSettingsRequest, opts ...grpc.CallOption) (*GetUserSettingsResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) GetAllNotifications(ctx context.Context, in *GetAllNotificationsRequest, opts ...grpc.CallOption) (*GetAllNotificationsResponse, error) {
	out := new(GetAllNotificationsResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/GetAllNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) InsertNotification(ctx context.Context, in *InsertNotificationRequest, opts ...grpc.CallOption) (*InsertNotificationResponse, error) {
	out := new(InsertNotificationResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/InsertNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) MarkAllAsSeen(ctx context.Context, in *MarkAllAsSeenRequest, opts ...grpc.CallOption) (*EmptyRespones, error) {
	out := new(EmptyRespones)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/MarkAllAsSeen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetUserSettings(ctx context.Context, in *GetUserSettingsRequest, opts ...grpc.CallOption) (*GetUserSettingsResponse, error) {
	out := new(GetUserSettingsResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/GetUserSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) UpdateUserSettings(ctx context.Context, in *UpdateUserSettingsRequest, opts ...grpc.CallOption) (*GetUserSettingsResponse, error) {
	out := new(GetUserSettingsResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/UpdateUserSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	GetAllNotifications(context.Context, *GetAllNotificationsRequest) (*GetAllNotificationsResponse, error)
	InsertNotification(context.Context, *InsertNotificationRequest) (*InsertNotificationResponse, error)
	MarkAllAsSeen(context.Context, *MarkAllAsSeenRequest) (*EmptyRespones, error)
	GetUserSettings(context.Context, *GetUserSettingsRequest) (*GetUserSettingsResponse, error)
	UpdateUserSettings(context.Context, *UpdateUserSettingsRequest) (*GetUserSettingsResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) GetAllNotifications(context.Context, *GetAllNotificationsRequest) (*GetAllNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotifications not implemented")
}
func (UnimplementedNotificationServiceServer) InsertNotification(context.Context, *InsertNotificationRequest) (*InsertNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertNotification not implemented")
}
func (UnimplementedNotificationServiceServer) MarkAllAsSeen(context.Context, *MarkAllAsSeenRequest) (*EmptyRespones, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAllAsSeen not implemented")
}
func (UnimplementedNotificationServiceServer) GetUserSettings(context.Context, *GetUserSettingsRequest) (*GetUserSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSettings not implemented")
}
func (UnimplementedNotificationServiceServer) UpdateUserSettings(context.Context, *UpdateUserSettingsRequest) (*GetUserSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserSettings not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_GetAllNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetAllNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/GetAllNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetAllNotifications(ctx, req.(*GetAllNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_InsertNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).InsertNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/InsertNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).InsertNotification(ctx, req.(*InsertNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_MarkAllAsSeen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAllAsSeenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).MarkAllAsSeen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/MarkAllAsSeen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).MarkAllAsSeen(ctx, req.(*MarkAllAsSeenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetUserSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetUserSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/GetUserSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetUserSettings(ctx, req.(*GetUserSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_UpdateUserSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).UpdateUserSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/UpdateUserSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).UpdateUserSettings(ctx, req.(*UpdateUserSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllNotifications",
			Handler:    _NotificationService_GetAllNotifications_Handler,
		},
		{
			MethodName: "InsertNotification",
			Handler:    _NotificationService_InsertNotification_Handler,
		},
		{
			MethodName: "MarkAllAsSeen",
			Handler:    _NotificationService_MarkAllAsSeen_Handler,
		},
		{
			MethodName: "GetUserSettings",
			Handler:    _NotificationService_GetUserSettings_Handler,
		},
		{
			MethodName: "UpdateUserSettings",
			Handler:    _NotificationService_UpdateUserSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification_service/notification_service.proto",
}
