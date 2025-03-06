package svc

import (
	"github/lhh-gh/IM/internal/business/info/api/internal/config"
	"github/lhh-gh/IM/internal/common/dao/myMongo"
)

type ServiceContext struct {
	Config config.Config
	Mongo  *myMongo.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mongo:  myMongo.MustNewMongo(c.Mongo.Host),
	}
}
