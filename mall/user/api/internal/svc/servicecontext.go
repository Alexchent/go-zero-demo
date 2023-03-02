package svc

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/mall/user/api/internal/config"
	"go-zero-demo/mall/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	Auth      jwt.Claims
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
