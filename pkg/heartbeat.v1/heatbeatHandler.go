package heartbeat_v1

import "context"

type HeartbeatServer struct {
	UnimplementedHeartbeatServiceServer
	up     bool
	status ServiceStatus
}

func (s *HeartbeatServer) GetHeartbeat(ctx context.Context, req *HeartbeatServiceRequest) (*HeartbeatServiceResponse, error) {
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
