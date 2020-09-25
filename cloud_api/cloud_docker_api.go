package cloud_api

import (
	"context"
	"rpctest/cloud_rpc"

	"google.golang.org/grpc"
)

// CloudDockerAPI CloudDockerAPI
type CloudDockerAPI struct {
	Server string
	Port   string
}

// Create 创建docker
func (cloud *CloudDockerAPI) Create(jsonData string) *cloud_rpc.Msg {
	msg := &cloud_rpc.Msg{}
	msg.JsonData = jsonData
	conn, _ := grpc.Dial(cloud.Server+":"+cloud.Port, grpc.WithInsecure())
	defer conn.Close()
	client := cloud_rpc.NewCloudDockerServiceClient(conn)
	response, _ := client.Create(context.Background(), msg)
	return response
}
