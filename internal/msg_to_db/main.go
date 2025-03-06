package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github/lhh-gh/IM/internal/msg_to_db/internal/config"
	"github/lhh-gh/IM/internal/msg_to_db/internal/mqs"
	"github/lhh-gh/IM/internal/msg_to_db/internal/svc"
)

var configFile = flag.String("f", "etc/todb.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)

	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()

	mq := mqs.NewMsgToDB(ctx, svcCtx)

	// 处理退出信号，平滑关闭
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		mq.Close()
		os.Exit(0)
	}() // 处理退出信号，平滑关闭

	mq.Start()
}
