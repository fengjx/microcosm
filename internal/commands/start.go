package commands

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"io"
	"microcosm/internal/config"
	"microcosm/internal/db"
	"microcosm/internal/pkg/app"
	"microcosm/internal/routers"
	"os"
)

// Starts web server (user interface)
var StartCommand = cli.Command{
	Name:   "start",
	Usage:  "Starts web server",
	Flags:  startFlags,
	Action: startAction,
}

var startFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "http-listen-addr, a",
		Usage:  "HTTP server port",
		EnvVar: "MICROCOSM_HTTP_PORT",
	},
	cli.StringFlag{
		Name:   "config, c",
		Usage:  "app config files path",
		Value:  "conf/app.ini,conf/app-test.ini",
		EnvVar: "MICROCOSM_CONFIG_FILES",
	},
}

func startAction(ctx *cli.Context) error {
	conf := config.Init(ctx)
	f, _ := os.Create(conf.GetLogConfig().Accesslog)
	if conf.IsDebug() {
		// 如果需要同时将日志写入文件和控制台，请使用以下代码。
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	} else {
		gin.DefaultWriter = io.MultiWriter(f)
	}
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	server := app.New(engine)
	server.AddRouter(routers.NewAppRouter())
	server.AddStarter(db.New())
	server.Start(conf)
	return nil
}
