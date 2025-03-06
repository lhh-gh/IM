package svc

import (
	"github/lhh-gh/IM/internal/business/wsget/api/internal/config"
	"github/lhh-gh/IM/pkg/servicehub"
)

type ServiceContext struct {
	Config       config.Config
	DiscoveryHub *servicehub.DiscoveryHub
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		DiscoveryHub: servicehub.NewDiscoveryHub(c.Etcd.Host),
	}
}
