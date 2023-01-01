package network_capture_v1

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	networkCaptureClient NetworkCaptureServiceClient
	contextTime          time.Duration
}

func NewNetworkCaptureClient(serverAddr string, opts ...grpc.DialOption) *GrpcClient {
	contextTimeout := viper.GetInt("GPRC_CLIENT_DEADLINE_TIMEOUT_SECONDS")
	if contextTimeout == 0 {
		fmt.Println("no context timeout set, defaulting to 10 seconds")
		contextTimeout = 10
	}
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		panic(err)
	}
	return &GrpcClient{
		networkCaptureClient: NewNetworkCaptureServiceClient(conn),
		contextTime:          time.Duration(contextTimeout) * time.Second,
	}
}

func (c *GrpcClient) SendNetworkCapture(pcap *NetworkCaptureRequest) (*NetworkCaptureResponse, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(c.contextTime))
	defer cancel()
	resp, err := c.networkCaptureClient.NetworkCapture(ctx, pcap)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
