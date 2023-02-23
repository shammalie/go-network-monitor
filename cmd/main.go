package main

import (
	"fmt"
	"strings"

	"github.com/shammalie/go-network-monitor/internal/events"
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
		fmt.Println("no env file found, will use environment")
	}
	server := pkg.NewGrpcServer()
	go func() {
		ipIgnore := strings.Split(viper.GetString("IP_IGNORE"), ",")
		triageService := events.NewTriageService(ipIgnore)
		for event := range server.NetworkCaptureServer.ClientEvents {
			triageService.Triage(event)
		}
	}()
	fmt.Printf("starting server %s:%d\n", server.Hostname, server.Port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
