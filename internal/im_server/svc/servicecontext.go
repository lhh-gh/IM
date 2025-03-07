package svc

import (
	"context"
	"github/lhh-gh/IM/internal/msg_forward/gossip"
	"github/lhh-gh/IM/internal/msg_forward/gossip/gossippb"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"github/lhh-gh/IM/internal/common/dao/myRedis"
	"github/lhh-gh/IM/internal/im_server/internal/config"
	"github/lhh-gh/IM/pkg/servicehub"
)

type ServiceContext struct {
	Config       config.Config
	MsgForwarder *kafka.Writer
	Redis        *myRedis.Native
	RegisterHub  *servicehub.RegisterHub
	DiscovHub    *servicehub.DiscoveryHub
	Gossips      []*gossippb.GossipClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	discov := servicehub.NewDiscoveryHub(c.Etcd.Endpoints)
	return &ServiceContext{
		Config: c,
		MsgForwarder: &kafka.Writer{
			Addr:         kafka.TCP(c.MsgForwarder.Brokers...),
			Topic:        c.MsgForwarder.Topic,
			Balancer:     &kafka.LeastBytes{},
			BatchTimeout: 10 * time.Millisecond, // 低超时时间
			RequiredAcks: kafka.RequireOne,      // 仅等待 Leader 确认
			Compression:  compress.Zstd,         // Zstd压缩
			Async:        true,                  // 启用异步写入
			MaxAttempts:  1,                     // 限制重试次数
		},
		Redis:       myRedis.MustNewNativeRedis(c.RedisConf),
		RegisterHub: servicehub.NewRegisterHub(c.Etcd.Endpoints, 3),
		DiscovHub:   discov,
		Gossips:     gossip.NewGossipClients(discov.GetServiceEndpoints(context.Background(), "gossip")),
		// TODO: 另起goroutine进行异步etcd服务发现订阅
	}
}
