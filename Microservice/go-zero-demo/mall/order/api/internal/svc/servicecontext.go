package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/mall/order/api/internal/config"
	"go-zero-demo/mall/user/rpc/user"
)

//	服务依赖

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
