package main

import (
	"fmt"

	"github.com/shammalie/go-network-monitor/pkg"
)

const (
	hostname = "localhost"
	port     = 4320
)

func main() {
	server := pkg.NewGrpcServer(port, hostname)
	fmt.Printf("starting server %s:%d\n", hostname, port)
	go func() {
		for e := range server.NetworkCaptureServer.ClientEvents {
			fmt.Println(e)
		}
	}()
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
