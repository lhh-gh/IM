package svc

import (
	"github/lhh-gh/IM/internal/business/group/api/internal/config"
	"github/lhh-gh/IM/internal/common/dao/myMysql"
	"github/lhh-gh/IM/internal/common/dao/myRedis"
)

type ServiceContext struct {
	Config config.Config
	MySQL  *myMysql.DB
	Redis  *myRedis.Native
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MySQL:  myMysql.MustNewMySQL(c.MySQL.DataSource),
		Redis:  myRedis.MustNewNativeRedis(c.RedisConf),
	}
}
