package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

//	添加user rpc配置
type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
}
