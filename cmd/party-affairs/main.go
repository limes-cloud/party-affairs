package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	_ "go.uber.org/automaxprocs"

	v1 "party-affairs/api/v1"
	systemconfig "party-affairs/config"
	"party-affairs/internal/service"
)

func main() {
	// 启动管理端
	app := kratosx.New(
		kratosx.RegistrarServer(RegisterServer),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

// RegisterServer 注册服务端
func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	conf := &systemconfig.Config{}
	if err := c.Value("business").Scan(conf); err != nil {
		panic("business config format error:" + err.Error())
	}
	c.Watch("business", func(value config.Value) {
		if err := value.Scan(conf); err != nil {
			log.Error("business 配置变更失败")
		}
	})

	srv := service.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)
}
