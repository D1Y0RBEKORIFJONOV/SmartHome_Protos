// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: device/TV/tv.proto

package tv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	TVService_AddTV_FullMethodName            = "/TVService/AddTV"
	TVService_AddChannel_FullMethodName       = "/TVService/AddChannel"
	TVService_GetUserChannel_FullMethodName   = "/TVService/GetUserChannel"
	TVService_DeleteChannel_FullMethodName    = "/TVService/DeleteChannel"
	TVService_DownOrUpVVoiceTv_FullMethodName = "/TVService/DownOrUpVVoiceTv"
	TVService_PreviousAndNext_FullMethodName  = "/TVService/PreviousAndNext"
	TVService_OnAndOffUserTv_FullMethodName   = "/TVService/OnAndOffUserTv"
)

// TVServiceClient is the client API for TVService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TVServiceClient interface {
	AddTV(ctx context.Context, in *AddTVReq, opts ...grpc.CallOption) (*TvStatusMessage, error)
	AddChannel(ctx context.Context, in *AddChannelReq, opts ...grpc.CallOption) (*TvStatusMessage, error)
	GetUserChannel(ctx context.Context, in *GetUserChannelReq, opts ...grpc.CallOption) (*GetUserChannelRes, error)
	DeleteChannel(ctx context.Context, in *DeleteChannelReq, opts ...grpc.CallOption) (*TvStatusMessage, error)
	DownOrUpVVoiceTv(ctx context.Context, in *DownOrUpVVoiceTvReq, opts ...grpc.CallOption) (*DownOrUpVVoiceTvRes, error)
	PreviousAndNext(ctx context.Context, in *PreviousAndNextReq, opts ...grpc.CallOption) (*PreviousAndNextRes, error)
	OnAndOffUserTv(ctx context.Context, in *OnAndOffUserTvReq, opts ...grpc.CallOption) (*OnAndOffUserTvRes, error)
}

type tVServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTVServiceClient(cc grpc.ClientConnInterface) TVServiceClient {
	return &tVServiceClient{cc}
}

func (c *tVServiceClient) AddTV(ctx context.Context, in *AddTVReq, opts ...grpc.CallOption) (*TvStatusMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TvStatusMessage)
	err := c.cc.Invoke(ctx, TVService_AddTV_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tVServiceClient) AddChannel(ctx context.Context, in *AddChannelReq, opts ...grpc.CallOption) (*TvStatusMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TvStatusMessage)
	err := c.cc.Invoke(ctx, TVService_AddChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tVServiceClient) GetUserChannel(ctx context.Context, in *GetUserChannelReq, opts ...grpc.CallOption) (*GetUserChannelRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserChannelRes)
	err := c.cc.Invoke(ctx, TVService_GetUserChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tVServiceClient) DeleteChannel(ctx context.Context, in *DeleteChannelReq, opts ...grpc.CallOption) (*TvStatusMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TvStatusMessage)
	err := c.cc.Invoke(ctx, TVService_DeleteChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tVServiceClient) DownOrUpVVoiceTv(ctx context.Context, in *DownOrUpVVoiceTvReq, opts ...grpc.CallOption) (*DownOrUpVVoiceTvRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownOrUpVVoiceTvRes)
	err := c.cc.Invoke(ctx, TVService_DownOrUpVVoiceTv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tVServiceClient) PreviousAndNext(ctx context.Context, in *PreviousAndNextReq, opts ...grpc.CallOption) (*PreviousAndNextRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PreviousAndNextRes)
	err := c.cc.Invoke(ctx, TVService_PreviousAndNext_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tVServiceClient) OnAndOffUserTv(ctx context.Context, in *OnAndOffUserTvReq, opts ...grpc.CallOption) (*OnAndOffUserTvRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnAndOffUserTvRes)
	err := c.cc.Invoke(ctx, TVService_OnAndOffUserTv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TVServiceServer is the server API for TVService service.
// All implementations must embed UnimplementedTVServiceServer
// for forward compatibility
type TVServiceServer interface {
	AddTV(context.Context, *AddTVReq) (*TvStatusMessage, error)
	AddChannel(context.Context, *AddChannelReq) (*TvStatusMessage, error)
	GetUserChannel(context.Context, *GetUserChannelReq) (*GetUserChannelRes, error)
	DeleteChannel(context.Context, *DeleteChannelReq) (*TvStatusMessage, error)
	DownOrUpVVoiceTv(context.Context, *DownOrUpVVoiceTvReq) (*DownOrUpVVoiceTvRes, error)
	PreviousAndNext(context.Context, *PreviousAndNextReq) (*PreviousAndNextRes, error)
	OnAndOffUserTv(context.Context, *OnAndOffUserTvReq) (*OnAndOffUserTvRes, error)
	mustEmbedUnimplementedTVServiceServer()
}

// UnimplementedTVServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTVServiceServer struct {
}

func (UnimplementedTVServiceServer) AddTV(context.Context, *AddTVReq) (*TvStatusMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTV not implemented")
}
func (UnimplementedTVServiceServer) AddChannel(context.Context, *AddChannelReq) (*TvStatusMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChannel not implemented")
}
func (UnimplementedTVServiceServer) GetUserChannel(context.Context, *GetUserChannelReq) (*GetUserChannelRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserChannel not implemented")
}
func (UnimplementedTVServiceServer) DeleteChannel(context.Context, *DeleteChannelReq) (*TvStatusMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedTVServiceServer) DownOrUpVVoiceTv(context.Context, *DownOrUpVVoiceTvReq) (*DownOrUpVVoiceTvRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownOrUpVVoiceTv not implemented")
}
func (UnimplementedTVServiceServer) PreviousAndNext(context.Context, *PreviousAndNextReq) (*PreviousAndNextRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviousAndNext not implemented")
}
func (UnimplementedTVServiceServer) OnAndOffUserTv(context.Context, *OnAndOffUserTvReq) (*OnAndOffUserTvRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnAndOffUserTv not implemented")
}
func (UnimplementedTVServiceServer) mustEmbedUnimplementedTVServiceServer() {}

// UnsafeTVServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TVServiceServer will
// result in compilation errors.
type UnsafeTVServiceServer interface {
	mustEmbedUnimplementedTVServiceServer()
}

func RegisterTVServiceServer(s grpc.ServiceRegistrar, srv TVServiceServer) {
	s.RegisterService(&TVService_ServiceDesc, srv)
}

func _TVService_AddTV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTVReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServiceServer).AddTV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TVService_AddTV_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServiceServer).AddTV(ctx, req.(*AddTVReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TVService_AddChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChannelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServiceServer).AddChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TVService_AddChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServiceServer).AddChannel(ctx, req.(*AddChannelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TVService_GetUserChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserChannelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServiceServer).GetUserChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TVService_GetUserChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServiceServer).GetUserChannel(ctx, req.(*GetUserChannelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TVService_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServiceServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TVService_DeleteChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServiceServer).DeleteChannel(ctx, req.(*DeleteChannelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TVService_DownOrUpVVoiceTv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownOrUpVVoiceTvReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServiceServer).DownOrUpVVoiceTv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TVService_DownOrUpVVoiceTv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServiceServer).DownOrUpVVoiceTv(ctx, req.(*DownOrUpVVoiceTvReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TVService_PreviousAndNext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreviousAndNextReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServiceServer).PreviousAndNext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TVService_PreviousAndNext_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServiceServer).PreviousAndNext(ctx, req.(*PreviousAndNextReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TVService_OnAndOffUserTv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnAndOffUserTvReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServiceServer).OnAndOffUserTv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TVService_OnAndOffUserTv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServiceServer).OnAndOffUserTv(ctx, req.(*OnAndOffUserTvReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TVService_ServiceDesc is the grpc.ServiceDesc for TVService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TVService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TVService",
	HandlerType: (*TVServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTV",
			Handler:    _TVService_AddTV_Handler,
		},
		{
			MethodName: "AddChannel",
			Handler:    _TVService_AddChannel_Handler,
		},
		{
			MethodName: "GetUserChannel",
			Handler:    _TVService_GetUserChannel_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _TVService_DeleteChannel_Handler,
		},
		{
			MethodName: "DownOrUpVVoiceTv",
			Handler:    _TVService_DownOrUpVVoiceTv_Handler,
		},
		{
			MethodName: "PreviousAndNext",
			Handler:    _TVService_PreviousAndNext_Handler,
		},
		{
			MethodName: "OnAndOffUserTv",
			Handler:    _TVService_OnAndOffUserTv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device/TV/tv.proto",
}
