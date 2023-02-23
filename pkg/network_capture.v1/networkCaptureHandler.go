package network_capture_v1

import (
	"io"
	"time"
)

type NetworkCaptureServer struct {
	UnimplementedNetworkCaptureServiceServer
	ClientEvents  chan *NetworkCaptureRequest
	ClientActions chan *NetworkCaptureResponse
}

func NewNetworkCaptureServer() *NetworkCaptureServer {
	return &NetworkCaptureServer{
		ClientEvents:  make(chan *NetworkCaptureRequest),
		ClientActions: make(chan *NetworkCaptureResponse),
	}
}

func (s *NetworkCaptureServer) NetworkCapture(stream NetworkCaptureService_NetworkCaptureServer) error {
	go func() error {
		for {
			select {
			case event := <-s.ClientActions:
				if err := stream.Send(event); err != nil {
					return err
				}
			case <-stream.Context().Done():
				return stream.Context().Err()
			}
		}
	}()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		s.ClientEvents <- in
	}
}

func (s *NetworkCaptureServer) SendClientAction(ip string, action string) {
	s.ClientActions <- &NetworkCaptureResponse{
		Ip:        ip,
		Action:    action,
		Timestamp: time.Now().UnixMilli(),
	}
}
