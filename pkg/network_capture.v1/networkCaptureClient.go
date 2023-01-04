package network_capture_v1

import (
	context "context"
	"fmt"
	"io"
	"sync"
	"time"

	heartbeat_v1 "github.com/shammalie/go-network-monitor/pkg/heartbeat.v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	ReceivedActions      chan *NetworkCaptureResponse
	HeartbeatClient      *heartbeat_v1.GrpcHeartbeatClient
	networkCaptureClient NetworkCaptureServiceClient
	contextTime          time.Duration
	stream               NetworkCaptureService_NetworkCaptureClient
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
	client := &GrpcClient{
		ReceivedActions: make(chan *NetworkCaptureResponse),
		// HeartbeatClient:      heartbeat_v1.NewHeartbeatClient(conn),
		networkCaptureClient: NewNetworkCaptureServiceClient(conn),
		contextTime:          time.Duration(contextTimeout) * time.Second,
	}
	client.initaliseNetworkCaptureStream()
	return client
}

func (c *GrpcClient) initaliseNetworkCaptureStream() {
	var wg sync.WaitGroup
	stream, err := c.networkCaptureClient.NetworkCapture(context.Background())
	if err != nil {
		panic(err)
	}
	c.stream = stream
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			in, err := c.stream.Recv()
			if err == io.EOF {
				// Read done.
				wg.Done()
				return
			}
			if err != nil {
				panic(err)
			}
			c.ReceivedActions <- in
		}
	}()
}

func (c *GrpcClient) SendNetworkCapture(pcap *NetworkCaptureRequest) {
	if err := c.stream.Send(pcap); err != nil {
		panic(err)
	}
}

func (c *GrpcClient) CloseStream() error {
	return c.stream.CloseSend()
}
