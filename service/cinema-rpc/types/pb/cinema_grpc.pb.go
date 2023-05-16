// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: cinema.proto

package pb

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

// CinemaRpcClient is the client API for CinemaRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CinemaRpcClient interface {
	GetCinemaList(ctx context.Context, in *CinemaListRequest, opts ...grpc.CallOption) (*CinemaListResponse, error)
	GetConditionList(ctx context.Context, in *ConditionListRequest, opts ...grpc.CallOption) (*ConditionListResponse, error)
	GetCinema(ctx context.Context, in *GetCinemaRequest, opts ...grpc.CallOption) (*GetCinemaResponse, error)
	GetShowList(ctx context.Context, in *GetShowListRequest, opts ...grpc.CallOption) (*GetShowListResponse, error)
	GetHallSeats(ctx context.Context, in *GetHallSeatsRequest, opts ...grpc.CallOption) (*GetHallSeatsResponse, error)
}

type cinemaRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewCinemaRpcClient(cc grpc.ClientConnInterface) CinemaRpcClient {
	return &cinemaRpcClient{cc}
}

func (c *cinemaRpcClient) GetCinemaList(ctx context.Context, in *CinemaListRequest, opts ...grpc.CallOption) (*CinemaListResponse, error) {
	out := new(CinemaListResponse)
	err := c.cc.Invoke(ctx, "/cinema.cinemaRpc/GetCinemaList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaRpcClient) GetConditionList(ctx context.Context, in *ConditionListRequest, opts ...grpc.CallOption) (*ConditionListResponse, error) {
	out := new(ConditionListResponse)
	err := c.cc.Invoke(ctx, "/cinema.cinemaRpc/GetConditionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaRpcClient) GetCinema(ctx context.Context, in *GetCinemaRequest, opts ...grpc.CallOption) (*GetCinemaResponse, error) {
	out := new(GetCinemaResponse)
	err := c.cc.Invoke(ctx, "/cinema.cinemaRpc/GetCinema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaRpcClient) GetShowList(ctx context.Context, in *GetShowListRequest, opts ...grpc.CallOption) (*GetShowListResponse, error) {
	out := new(GetShowListResponse)
	err := c.cc.Invoke(ctx, "/cinema.cinemaRpc/GetShowList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaRpcClient) GetHallSeats(ctx context.Context, in *GetHallSeatsRequest, opts ...grpc.CallOption) (*GetHallSeatsResponse, error) {
	out := new(GetHallSeatsResponse)
	err := c.cc.Invoke(ctx, "/cinema.cinemaRpc/GetHallSeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CinemaRpcServer is the server API for CinemaRpc service.
// All implementations must embed UnimplementedCinemaRpcServer
// for forward compatibility
type CinemaRpcServer interface {
	GetCinemaList(context.Context, *CinemaListRequest) (*CinemaListResponse, error)
	GetConditionList(context.Context, *ConditionListRequest) (*ConditionListResponse, error)
	GetCinema(context.Context, *GetCinemaRequest) (*GetCinemaResponse, error)
	GetShowList(context.Context, *GetShowListRequest) (*GetShowListResponse, error)
	GetHallSeats(context.Context, *GetHallSeatsRequest) (*GetHallSeatsResponse, error)
	mustEmbedUnimplementedCinemaRpcServer()
}

// UnimplementedCinemaRpcServer must be embedded to have forward compatible implementations.
type UnimplementedCinemaRpcServer struct {
}

func (UnimplementedCinemaRpcServer) GetCinemaList(context.Context, *CinemaListRequest) (*CinemaListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCinemaList not implemented")
}
func (UnimplementedCinemaRpcServer) GetConditionList(context.Context, *ConditionListRequest) (*ConditionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConditionList not implemented")
}
func (UnimplementedCinemaRpcServer) GetCinema(context.Context, *GetCinemaRequest) (*GetCinemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCinema not implemented")
}
func (UnimplementedCinemaRpcServer) GetShowList(context.Context, *GetShowListRequest) (*GetShowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShowList not implemented")
}
func (UnimplementedCinemaRpcServer) GetHallSeats(context.Context, *GetHallSeatsRequest) (*GetHallSeatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHallSeats not implemented")
}
func (UnimplementedCinemaRpcServer) mustEmbedUnimplementedCinemaRpcServer() {}

// UnsafeCinemaRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CinemaRpcServer will
// result in compilation errors.
type UnsafeCinemaRpcServer interface {
	mustEmbedUnimplementedCinemaRpcServer()
}

func RegisterCinemaRpcServer(s grpc.ServiceRegistrar, srv CinemaRpcServer) {
	s.RegisterService(&CinemaRpc_ServiceDesc, srv)
}

func _CinemaRpc_GetCinemaList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CinemaListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaRpcServer).GetCinemaList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.cinemaRpc/GetCinemaList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaRpcServer).GetCinemaList(ctx, req.(*CinemaListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CinemaRpc_GetConditionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConditionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaRpcServer).GetConditionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.cinemaRpc/GetConditionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaRpcServer).GetConditionList(ctx, req.(*ConditionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CinemaRpc_GetCinema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCinemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaRpcServer).GetCinema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.cinemaRpc/GetCinema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaRpcServer).GetCinema(ctx, req.(*GetCinemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CinemaRpc_GetShowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaRpcServer).GetShowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.cinemaRpc/GetShowList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaRpcServer).GetShowList(ctx, req.(*GetShowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CinemaRpc_GetHallSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHallSeatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaRpcServer).GetHallSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.cinemaRpc/GetHallSeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaRpcServer).GetHallSeats(ctx, req.(*GetHallSeatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CinemaRpc_ServiceDesc is the grpc.ServiceDesc for CinemaRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CinemaRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cinema.cinemaRpc",
	HandlerType: (*CinemaRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCinemaList",
			Handler:    _CinemaRpc_GetCinemaList_Handler,
		},
		{
			MethodName: "GetConditionList",
			Handler:    _CinemaRpc_GetConditionList_Handler,
		},
		{
			MethodName: "GetCinema",
			Handler:    _CinemaRpc_GetCinema_Handler,
		},
		{
			MethodName: "GetShowList",
			Handler:    _CinemaRpc_GetShowList_Handler,
		},
		{
			MethodName: "GetHallSeats",
			Handler:    _CinemaRpc_GetHallSeats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cinema.proto",
}
