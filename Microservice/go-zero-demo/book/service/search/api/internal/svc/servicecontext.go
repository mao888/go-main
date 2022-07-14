package svc

import (
	"book/service/search/api/internal/config"
	"book/service/search/api/internal/middleware"
	"book/service/user/rpc/user"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
