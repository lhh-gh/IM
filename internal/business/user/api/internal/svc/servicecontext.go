package svc

import (
	"github/lhh-gh/IM/internal/business/user/api/internal/config"
	"github/lhh-gh/IM/internal/common/dao/myMysql"
)

type ServiceContext struct {
	Config config.Config
	MySQL  *myMysql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MySQL:  myMysql.MustNewMySQL(c.MySQL.DataSource),
	}
}
