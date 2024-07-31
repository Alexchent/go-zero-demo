package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/policy/internal/config"
	"go-zero-demo/policy/model"
	"time"
)

type ServiceContext struct {
	Config      config.Config
	PolicyModel model.PolicyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PolicyModel: model.NewPolicyModel(
			sqlx.NewMysql(c.DB.DataSource),
			c.CacheRedis,
			cache.WithExpiry(time.Second*10),
			cache.WithNotFoundExpiry(time.Second*1),
		),
	}
}
