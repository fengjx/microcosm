package app

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"microcosm/internal/config"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	config   *config.Config
	starters []Starter
	routers  []RouteStarter
	engine   *gin.Engine
}

func (app *App) AddStarter(s Starter) {
	app.starters = append(app.starters, s)
}

func (app *App) AddRouter(rs RouteStarter) {
	app.routers = append(app.routers, rs)
}

func (app *App) Start(config *config.Config) {
	app.config = config
	app.initLogger()
	app.initStarter()
	app.initEngine()
	app.initRouteStarter()

	serverConfig := app.config.GetServerConfig()
	server := &http.Server{
		Addr:    serverConfig.Addr,
		Handler: app.engine,
	}

	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Info("server started: ", serverConfig.Addr)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		app.Stop()
		cancel()
	}()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}

func (app *App) initLogger() {
	debug := app.config.IsDebug()
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func (app *App) initStarter() {
	for _, starter := range app.starters {
		starter.Start(app.config)
	}
}

func (app *App) initRouteStarter() {
	for _, router := range app.routers {
		router.Start(app.config, app.engine)
	}
}

func (app *App) initEngine() {
	serverConfig := app.config.GetServerConfig()
	if serverConfig.TemplatePath != "" {
		app.engine.LoadHTMLGlob(serverConfig.TemplatePath + "/*")
	}
	app.engine.Static("/front", "./static")
}

func (app *App) Stop() {
	for _, router := range app.routers {
		router.Stop(app.config, app.engine)
	}
	for _, starter := range app.starters {
		starter.Stop(app.config)
	}
}

func New(eng *gin.Engine) *App {
	app := new(App)
	app.starters = make([]Starter, 0)
	app.engine = eng
	return app
}
