package client

import (
	"context"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	userpb "github.com/soraQaQ/blog/api/user/v1"
)

// NewUserClient 创建用户服务客户端
func NewUserClient() userpb.UserServiceClient {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	dis := consul.New(client)
	endpoint := "discovery:///user.service"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))
	if err != nil {
		panic(err)
	}
	return userpb.NewUserServiceClient(conn)
}
