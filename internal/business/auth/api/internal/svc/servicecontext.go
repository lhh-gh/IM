package svc

import (
	"github/lhh-gh/IM/internal/business/auth/api/internal/config"
	"github/lhh-gh/IM/internal/common/dao/myMysql"
)

type ServiceContext struct {
	Config config.Config
	MySQL  *myMysql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqldb := myMysql.MustNewMySQL(c.MySQL.DataSource)
	return &ServiceContext{
		Config: c,
		MySQL:  mysqldb,
	}
}
