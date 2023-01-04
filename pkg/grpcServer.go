package pkg

import (
	"fmt"
	"net"
	"os"

	heartbeat_v1 "github.com/shammalie/go-network-monitor/pkg/heartbeat.v1"
	network_capture_v1 "github.com/shammalie/go-network-monitor/pkg/network_capture.v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	logInfoPrefix  = "grpc_server - INFO: %v\n"
	logErrorPrefix = "grpc_server - ERROR: %v\n"
)

type grpcServer struct {
	Heartbeat            *heartbeat_v1.HeartbeatServer
	NetworkCaptureServer *network_capture_v1.NetworkCaptureServer
	Hostname             string
	Port                 int
	listener             net.Listener
	server               *grpc.Server
}

// Create a new grpc server that implements the handler for network_capture.v1#
// Returns a grpc.Server pointer struct.
func NewGrpcServer() *grpcServer {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	port := viper.GetInt("SERVER_PORT")
	if port == 0 {
		port = 3000
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
		Hostname:             hostname,
		Port:                 port,
		listener:             lis,
		server:               s,
	}
}

func (opt *grpcServer) ListenAndServe() error {
	return opt.server.Serve(opt.listener)
}
