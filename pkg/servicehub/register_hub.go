package servicehub

import (
	"context"
	"errors"
	"fmt"
	"github/lhh-gh/IM/pkg/utils"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// RegisterHub 服务注册中心
type RegisterHub struct {
	client             *clientv3.Client
	heartbeatFrequency int64 // 续租间隔
}

var (
	registerHub  *RegisterHub
	registerOnce sync.Once
)

// NewRegisterHub 单例模式创建一个RegisterHub
func NewRegisterHub(etcdServers []string, heartbeatFrequency int64) *RegisterHub {
	if registerHub == nil {
		registerOnce.Do(func() {
			if client, err := clientv3.New(
				clientv3.Config{
					Endpoints:   etcdServers,
					DialTimeout: 3 * time.Second,
				}); err != nil {
				logx.Error("[GetRegisterHub] Connect to etcd failed, error: ", err)
			} else {
				registerHub = &RegisterHub{
					client:             client,
					heartbeatFrequency: heartbeatFrequency,
				}
			}
		})
	}
	return registerHub
}

// Register 注册加续租
func (hub *RegisterHub) Register(ctx context.Context, service string, port int, workerID uint16) {
	ip, err := utils.GetLocalIP()
	if err != nil {
		panic(err)
	}
	addr := fmt.Sprintf("%s:%d", ip, port)
	// 先注册一遍
	leaseID, err := hub.register(ctx, service, int(workerID), addr, 0)
	if err != nil {
		panic(err)
	}
	go func() {
		// 异步持续续租
		for {
			hub.register(ctx, service, int(workerID), addr, leaseID)
			time.Sleep(time.Duration(3)*time.Second - 200*time.Millisecond)
		}
	}()
}

func (hub *RegisterHub) register(ctx context.Context, service string, workerID int, endpoint string, leaseID clientv3.LeaseID) (clientv3.LeaseID, error) {
	// 还未进行注册
	if leaseID <= 0 {
		// 先获取租赁leaseID
		if lease, err := hub.client.Grant(ctx, hub.heartbeatFrequency); err != nil {
			logx.Error("[Register] Regist failed, error: ", err)
			return 0, err
		} else {
			// 在etcd进行服务注册，key是当前服务器id,value为当前服务器地址
			key := fmt.Sprintf("%s/%d", service, workerID)
			if _, err = hub.client.Put(ctx, key, endpoint, clientv3.WithLease(lease.ID)); err != nil {
				logx.Error("[Register] Put service to etcd failed, error: ", err)
				return lease.ID, err
			} else {
				logx.Debugf("[Register] Service %s lease success", service)
				return lease.ID, nil
			}
		}
	} else {
		// 之前注册过了，这里是续租
		if _, err := hub.client.KeepAliveOnce(ctx, leaseID); errors.Is(err, rpctypes.ErrLeaseNotFound) {
			return hub.register(ctx, service, workerID, endpoint, 0) // 走注册流程
		} else if err != nil {
			logx.Error("[Register] Keep alive failed, error: ", err)
			return 0, err
		} else {
			return leaseID, nil
		}
	}
}

func (hub *RegisterHub) Unregister(ctx context.Context, service string, workerID int, endpoint string) error {
	key := fmt.Sprintf("%s/%d", service, workerID)
	if _, err := hub.client.Delete(ctx, key); err != nil {
		logx.Errorf("[Unregister] Unregist service %s at %s failed, error: %v", service, endpoint, err)
		return err
	} else {
		logx.Infof("[Unregister] Unregist service %s at %s success", service, endpoint)
		return nil
	}
}

func (hub *RegisterHub) Close() {
	hub.client.Close()
}
