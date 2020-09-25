package cloud_service

import (
	"context"
	"rpctest/cloud_rpc"
)

// CloudDockerService CloudDockerService
type CloudDockerService struct{}

// Create 创建
func (t *CloudDockerService) Create(ctx context.Context, request *cloud_rpc.Msg) (response *cloud_rpc.Msg, err error) {
	response = &cloud_rpc.Msg{
		JsonData: request.JsonData,
	}
	return response, err
}

// Delete 删除
func (t *CloudDockerService) Delete(ctx context.Context, request *cloud_rpc.Msg) (response *cloud_rpc.Msg, err error) {
	response = &cloud_rpc.Msg{
		JsonData: request.JsonData,
	}
	return response, err
}
