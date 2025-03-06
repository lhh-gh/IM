package svc

import (
	"fmt"
	"github/lhh-gh/IM/internal/common/dao/myMongo"
	"github/lhh-gh/IM/internal/common/dao/myMysql"
	"github/lhh-gh/IM/internal/common/dao/myRedis"
	"github/lhh-gh/IM/internal/msg_forward/internal/config"
	"github/lhh-gh/IM/pkg/servicehub"
)

type ServiceContext struct {
	Config config.Config
	Redis  *myRedis.Native
	Mongo  *myMongo.DB

	MySQL  *myMysql.DB
	Regist *servicehub.RegisterHub
	Discov *servicehub.DiscoveryHub
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  myRedis.MustNewNativeRedis(c.RedisConf),
		Mongo:  myMongo.MustNewMongo(fmt.Sprintf("mongodb://%s", c.MongoConf.Host)),
		MySQL:  myMysql.MustNewMySQL(c.MySQL.DataSource),
		Regist: servicehub.NewRegisterHub(c.Etcd.Endpoints, 3),
		Discov: servicehub.NewDiscoveryHub(c.Etcd.Endpoints),
	}
}
