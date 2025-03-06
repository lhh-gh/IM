package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github/lhh-gh/IM/internal/business/info/api/internal/config"
	"github/lhh-gh/IM/internal/business/info/api/internal/handler"
	"github/lhh-gh/IM/internal/business/info/api/internal/svc"
)

var configFile = flag.String("f", "etc/infos.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
