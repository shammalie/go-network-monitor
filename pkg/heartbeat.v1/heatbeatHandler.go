package heartbeat_v1

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type HeartbeatServer struct {
	UnimplementedHeartbeatServiceServer
	up     bool
	status ServiceStatus
}

func NewHeartbeatServer() *HeartbeatServer {
	return &HeartbeatServer{
		up:     true,
		status: ServiceStatus_SERVICE_UP,
	}
}

func (s *HeartbeatServer) HeartbeatRequest(ctx context.Context, _ *emptypb.Empty) (*HeartbeatServiceResponse, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return &HeartbeatServiceResponse{
		Up:     s.up,
		Status: s.status,
	}, nil
}

func (s *HeartbeatServer) UpdateServerStatus(isUp bool, statusCode ServiceStatus) {
	s.up = isUp
	s.status = statusCode
}
