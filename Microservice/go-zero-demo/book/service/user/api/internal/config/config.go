package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf          // rest api配置
	Mysql         struct { // 数据库配置，除mysql外，可能还有mongo等其他数据库
		DataSource string
	}
	CacheRedis cache.CacheConf // redis缓存
	Auth       struct {        // jwt鉴权配置
		AccessSecret string // 生成jwt token的密钥，最简单的方式可以使用一个uuid值
		AccessExpire int64  // jwt token有效期，单位：秒
	}
}
