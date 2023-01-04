package pkg

import (
	"fmt"
	"net"

	heartbeat_v1 "github.com/shammalie/go-network-monitor/pkg/heartbeat.v1"
	network_capture_v1 "github.com/shammalie/go-network-monitor/pkg/network_capture.v1"
	"google.golang.org/grpc"
)

const (
	logInfoPrefix  = "grpc_server - INFO: %v\n"
	logErrorPrefix = "grpc_server - ERROR: %v\n"
)

type grpcServer struct {
	Heartbeat            *heartbeat_v1.HeartbeatServer
	NetworkCaptureServer *network_capture_v1.NetworkCaptureServer
	listener             net.Listener
	server               *grpc.Server
}

// Create a new grpc server that implements the handler for network_capture.v1#
// Returns a grpc.Server pointer struct.
func NewGrpcServer(port int, hostname string) *grpcServer {
	if hostname == "" {
		hostname = "localhost"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		panic(fmt.Sprintf(logErrorPrefix, err))
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	heartbeatServer := heartbeat_v1.NewHeartbeatServer()
	networkCaptureServer := network_capture_v1.NewNetworkCaptureServer()

	network_capture_v1.RegisterNetworkCaptureServiceServer(s, networkCaptureServer)
	heartbeat_v1.RegisterHeartbeatServiceServer(s, heartbeatServer)
	return &grpcServer{
		Heartbeat:            heartbeatServer,
		NetworkCaptureServer: networkCaptureServer,
		listener:             lis,
		server:               s,
	}
}

func (opt *grpcServer) ListenAndServe() error {
	return opt.server.Serve(opt.listener)
}
