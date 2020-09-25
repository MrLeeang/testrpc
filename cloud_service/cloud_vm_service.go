package cloud_service

import (
	"context"
	"rpctest/cloud_rpc"
)

// CloudVMService CloudVMService
type CloudVMService struct{}

// Create 创建
func (t *CloudVMService) Create(ctx context.Context, request *cloud_rpc.Msg) (response *cloud_rpc.Msg, err error) {
	response = &cloud_rpc.Msg{
		JsonData: request.JsonData,
	}
	return response, err
}

// Delete 删除
func (t *CloudVMService) Delete(ctx context.Context, request *cloud_rpc.Msg) (response *cloud_rpc.Msg, err error) {
	response = &cloud_rpc.Msg{
		JsonData: request.JsonData,
	}
	return response, err
}
