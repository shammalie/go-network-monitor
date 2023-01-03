package network_capture_v1

import (
	"io"
	"sync"
	"time"
)

type NetworkCaptureServer struct {
	UnimplementedNetworkCaptureServiceServer
	mu            sync.RWMutex
	clientEvents  []*NetworkCaptureRequest
	clientActions []*NetworkCaptureResponse
}

func (s *NetworkCaptureServer) NetworkCapture(stream NetworkCaptureService_NetworkCaptureServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		s.mu.Lock()
		s.clientEvents = append(s.clientEvents, in)
		s.mu.Unlock()
		for _, event := range s.clientActions {
			if err := stream.Send(event); err != nil {
				return err
			}
		}
	}
}

func (s *NetworkCaptureServer) SendClientAction(ip string, action string) {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.clientActions = append(s.clientActions, &NetworkCaptureResponse{
		Ip:        ip,
		Action:    action,
		Timestamp: time.Now().UnixMilli(),
	})
}
