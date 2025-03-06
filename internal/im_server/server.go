package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"github/lhh-gh/IM/internal/im_server/internal/config"
	"github/lhh-gh/IM/internal/im_server/internal/mqs"
	"github/lhh-gh/IM/internal/im_server/internal/server"
	"github/lhh-gh/IM/internal/im_server/svc"
)

var (
	configFile = flag.String("f", "etc/server.yaml", "the config file")

	c  config.Config
	s  *server.Server
	mq *mqs.MsgSender
)

func main() {
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)

	idOptions := idgen.NewIdGeneratorOptions(c.WorkID)
	idgen.SetIdGenerator(idOptions)

	ctx := context.Background()
	svcCtx := svc.NewServiceContext(c)

	s = server.MustNewServer(ctx, svcCtx)
	threading.GoSafe(func() {
		s.Start()
	})

	// 处理退出信号，平滑关闭
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		mq.Close()
		os.Exit(0)
	}() // 处理退出信号，平滑关闭

	mq = mqs.NewMsgSender(ctx, svcCtx).WithManager(s.Manager)
	mq.Start()
}
