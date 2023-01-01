package network_capture_v1

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

const (
	contextTime = 5
)

type GrpcClient struct {
	networkCaptureClient NetworkCaptureServiceClient
}

func NewNetworkCaptureClient(serverAddr string, opts ...grpc.DialOption) *GrpcClient {
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		panic(err)
	}
	return &GrpcClient{
		networkCaptureClient: NewNetworkCaptureServiceClient(conn),
	}
}

func (c *GrpcClient) SendNetworkCapture(pcap *NetworkCaptureRequest) (*NetworkCaptureResponse, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(contextTime))
	defer cancel()
	resp, err := c.networkCaptureClient.NetworkCapture(ctx, pcap)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
