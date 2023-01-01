package network_capture_v1

import (
	"context"
	"fmt"
)

type Server struct {
	UnimplementedNetworkCaptureServiceServer
}

func (s *Server) NetworkCapture(ctx context.Context, req *NetworkCaptureRequest) (*NetworkCaptureResponse, error) {
	fmt.Println(req)
	return &NetworkCaptureResponse{
		Outcome: "RECORD",
	}, nil
}
