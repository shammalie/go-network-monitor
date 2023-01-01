package pkg

import (
	"fmt"
	"net"

	network_capture_v1 "github.com/shammalie/go-network-monitor/pkg/network_capture.v1"
	"google.golang.org/grpc"
)

const (
	logInfoPrefix  = "grpc_server - INFO: %v\n"
	logErrorPrefix = "grpc_server - ERROR: %v\n"
)

// Create a new grpc server that implements the handler for network_capture.v1#
// Returns a grpc.Server pointer struct.
func NewGrpcServer(port int, hostname string) *grpc.Server {
	if hostname == "" {
		hostname = "localhost"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		panic(fmt.Sprintf(logErrorPrefix, err))
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	network_capture_v1.RegisterNetworkCaptureServiceServer(s, &network_capture_v1.Server{})
	go func() {
		s.Serve(lis)
	}()
	return s
}
