// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/network_capture/v1/pcap_event.proto

package network_capture_v1

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

// NetworkCaptureServiceClient is the client API for NetworkCaptureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkCaptureServiceClient interface {
	// Submit a pcap capture and relay an action for the client.
	NetworkCapture(ctx context.Context, opts ...grpc.CallOption) (NetworkCaptureService_NetworkCaptureClient, error)
}

type networkCaptureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkCaptureServiceClient(cc grpc.ClientConnInterface) NetworkCaptureServiceClient {
	return &networkCaptureServiceClient{cc}
}

func (c *networkCaptureServiceClient) NetworkCapture(ctx context.Context, opts ...grpc.CallOption) (NetworkCaptureService_NetworkCaptureClient, error) {
	stream, err := c.cc.NewStream(ctx, &NetworkCaptureService_ServiceDesc.Streams[0], "/network_capture.v1.NetworkCaptureService/NetworkCapture", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkCaptureServiceNetworkCaptureClient{stream}
	return x, nil
}

type NetworkCaptureService_NetworkCaptureClient interface {
	Send(*NetworkCaptureRequest) error
	Recv() (*NetworkCaptureResponse, error)
	grpc.ClientStream
}

type networkCaptureServiceNetworkCaptureClient struct {
	grpc.ClientStream
}

func (x *networkCaptureServiceNetworkCaptureClient) Send(m *NetworkCaptureRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *networkCaptureServiceNetworkCaptureClient) Recv() (*NetworkCaptureResponse, error) {
	m := new(NetworkCaptureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetworkCaptureServiceServer is the server API for NetworkCaptureService service.
// All implementations must embed UnimplementedNetworkCaptureServiceServer
// for forward compatibility
type NetworkCaptureServiceServer interface {
	// Submit a pcap capture and relay an action for the client.
	NetworkCapture(NetworkCaptureService_NetworkCaptureServer) error
	mustEmbedUnimplementedNetworkCaptureServiceServer()
}

// UnimplementedNetworkCaptureServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkCaptureServiceServer struct {
}

func (UnimplementedNetworkCaptureServiceServer) NetworkCapture(NetworkCaptureService_NetworkCaptureServer) error {
	return status.Errorf(codes.Unimplemented, "method NetworkCapture not implemented")
}
func (UnimplementedNetworkCaptureServiceServer) mustEmbedUnimplementedNetworkCaptureServiceServer() {}

// UnsafeNetworkCaptureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkCaptureServiceServer will
// result in compilation errors.
type UnsafeNetworkCaptureServiceServer interface {
	mustEmbedUnimplementedNetworkCaptureServiceServer()
}

func RegisterNetworkCaptureServiceServer(s grpc.ServiceRegistrar, srv NetworkCaptureServiceServer) {
	s.RegisterService(&NetworkCaptureService_ServiceDesc, srv)
}

func _NetworkCaptureService_NetworkCapture_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkCaptureServiceServer).NetworkCapture(&networkCaptureServiceNetworkCaptureServer{stream})
}

type NetworkCaptureService_NetworkCaptureServer interface {
	Send(*NetworkCaptureResponse) error
	Recv() (*NetworkCaptureRequest, error)
	grpc.ServerStream
}

type networkCaptureServiceNetworkCaptureServer struct {
	grpc.ServerStream
}

func (x *networkCaptureServiceNetworkCaptureServer) Send(m *NetworkCaptureResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *networkCaptureServiceNetworkCaptureServer) Recv() (*NetworkCaptureRequest, error) {
	m := new(NetworkCaptureRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetworkCaptureService_ServiceDesc is the grpc.ServiceDesc for NetworkCaptureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkCaptureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "network_capture.v1.NetworkCaptureService",
	HandlerType: (*NetworkCaptureServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NetworkCapture",
			Handler:       _NetworkCaptureService_NetworkCapture_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/network_capture/v1/pcap_event.proto",
}