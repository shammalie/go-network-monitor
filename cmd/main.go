package main

import (
	"fmt"

	"github.com/shammalie/go-network-monitor/internal"
	"github.com/shammalie/go-network-monitor/pkg"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	db := internal.NewMongoClient()
	eventProcessor := internal.NewEventProcessor(db)

	server := pkg.NewGrpcServer()
	fmt.Printf("starting server %s:%d\n", server.Hostname, server.Port)
	go func() {
		for e := range server.NetworkCaptureServer.ClientEvents {
			eventProcessor.Events <- e
		}
	}()
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
