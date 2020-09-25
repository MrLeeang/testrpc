package cloud_api

import (
	"context"
	"rpctest/cloud_rpc"

	"google.golang.org/grpc"
)

// CloudVMAPI CloudVMAPI
type CloudVMAPI struct {
	Server string
	Port   string
}

// Create Create
func (cloud *CloudVMAPI) Create(jsonData string) *cloud_rpc.Msg {
	msg := &cloud_rpc.Msg{}
	msg.JsonData = jsonData
	conn, _ := grpc.Dial(cloud.Server+":"+cloud.Port, grpc.WithInsecure())
	defer conn.Close()
	client := cloud_rpc.NewCloudVMServiceClient(conn)
	response, _ := client.Create(context.Background(), msg)
	return response
}
