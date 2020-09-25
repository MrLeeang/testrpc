package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"rpctest/cloud_api"
	"rpctest/cloud_rpc"
	"rpctest/cloud_service"
	"time"

	"google.golang.org/grpc"
)

const (
	port = "41005"
)

func rpcServer() {
	//起服务
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cloud_rpc.RegisterCloudVMServiceServer(s, &cloud_service.CloudVMService{})
	cloud_rpc.RegisterCloudDockerServiceServer(s, &cloud_service.CloudDockerService{})
	s.Serve(lis)
}

func registerClient() {
	// 注册信息
	for {
		time.Sleep(time.Duration(5) * time.Second)
		fmt.Println("register...")
	}
}

func main() {
	if len(os.Args) > 1 {
		operation := os.Args[1]
		switch operation {
		case "rpcserver":
			go registerClient()
			rpcServer()
		case "app":
			rpcServer()
		}

	}

	client := &cloud_api.CloudVMAPI{
		Server: "127.0.0.1",
		Port:   "41005",
	}
	client.Create("123123")
}
