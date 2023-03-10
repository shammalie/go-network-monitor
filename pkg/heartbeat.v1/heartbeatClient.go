package heartbeat_v1

import (
	"context"
	"net"
	"os"
	"sync"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcHeartbeatClient struct {
	heartbeatClient HeartbeatServiceClient
	ctxTimeSeconds  int
	clientInfo      clientInfo
	Heartbeat       Heartbeat
	stop            chan struct{}
	mu              sync.RWMutex
}

type clientInfo struct {
	ip       string
	hostname string
}

type Heartbeat struct {
	serviceUp         bool
	serviceStatusCode ServiceStatus
	latency           int64
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}

func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func NewHeartbeatClient(conn *grpc.ClientConn) *GrpcHeartbeatClient {
	var wg sync.WaitGroup
	server := &GrpcHeartbeatClient{
		heartbeatClient: NewHeartbeatServiceClient(conn),
		ctxTimeSeconds:  5,
		Heartbeat:       Heartbeat{},
		clientInfo: clientInfo{
			ip:       getOutboundIP(),
			hostname: getHostname(),
		},
	}
	wg.Add(1)
	go func() {
		for {
			select {
			case <-time.NewTicker(5 * time.Second).C:
				startTime := time.Now()
				ctx, close := context.WithDeadline(context.Background(), startTime.Add(time.Duration(server.ctxTimeSeconds)*time.Second))
				resp, err := server.heartbeatClient.HeartbeatRequest(ctx, &emptypb.Empty{})
				if err != nil {
					panic(err)
				}
				close()
				server.updateHeartbeat(resp, startTime)
			case <-server.stop:
				return
			}
		}
	}()
	return server
}

func (s *GrpcHeartbeatClient) updateHeartbeat(resp *HeartbeatServiceResponse, startTime time.Time) {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.Heartbeat.serviceUp = resp.Up
	s.Heartbeat.serviceStatusCode = resp.Status
	s.Heartbeat.latency = time.Now().UnixMilli() - startTime.UnixMilli()
}
