package config

import (
	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"strings"
	"sync"
	"time"
)

type Config struct {
	env          string
	cfg          *ini.File
	appConfig    *AppConfig
	serverConfig *ServerConfig
	logConfig    *LogConfig
}

type AppConfig struct {
	Name string
}

type ServerConfig struct {
	Addr         string
	TemplatePath string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type LogConfig struct {
	Accesslog string
	AppLog    string
}

func (config *Config) init(ctx *cli.Context) {
	var err error
	configFiles := strings.Split(ctx.String("config"), ",")
	cfg, err := ini.Load(configFiles[0])
	if err != nil {
		log.Fatalf("config init [0] error: %v \n", err)
	}

	if len(configFiles) > 1 {
		for _, item := range configFiles[1:] {
			err := cfg.Append(item)
			if err != nil {
				log.Fatalf("Append config error", err, item)
			}
		}
	}

	config.cfg = cfg
	if err != nil {
		log.Fatalf("config init error: %v \n", err)
	}
	config.initAppConfig()
	config.initServerConfig()
	config.initLogConfig()
	config.coverCliConfig(ctx)
}

func (config *Config) initAppConfig() {
	appConfig := new(AppConfig)
	appConfig.Name = "my-app"
	config.appConfig = appConfig
	err := config.cfg.Section("app").MapTo(appConfig)
	if err != nil {
		log.Fatalln("initAppConfig error", err)
	}
}

func (config *Config) initLogConfig() {
	logConfig := new(LogConfig)
	logConfig.Accesslog = "gin.log"
	logConfig.AppLog = "app.log"
	config.logConfig = logConfig
	err := config.cfg.Section("log").MapTo(logConfig)
	if err != nil {
		log.Fatalln("initLogConfig error", err)
	}
}

func (config *Config) initServerConfig() {
	serverConfig := new(ServerConfig)
	serverConfig.Addr = ":8000"
	config.serverConfig = serverConfig
	err := config.cfg.Section("server").MapTo(serverConfig)

	if err != nil {
		log.Fatalln("initServerConfig error", err)
	}
}

func (config *Config) coverCliConfig(ctx *cli.Context) {
	addr := ctx.String("http-listen-addr")
	if addr != "" {
		config.serverConfig.Addr = addr
	}
}

func (config *Config) GetCfg() *ini.File {
	return config.cfg
}

var once sync.Once
var appConfig *Config

func Init(ctx *cli.Context) *Config {
	once.Do(func() {
		appConfig = new(Config)
		appConfig.init(ctx)
	})
	return appConfig
}

func GetConfig() *Config {
	return appConfig
}

func (config *Config) GetString(section string, key string) string {
	val, err := config.cfg.Section(section).GetKey(key)
	if err != nil {
		return ""
	}
	return val.String()
}

func (config *Config) GetAppConfig() AppConfig {
	return *config.appConfig
}

func (config *Config) GetServerConfig() ServerConfig {
	return *config.serverConfig
}

func (config *Config) GetLogConfig() LogConfig {
	return *config.logConfig
}

func (config *Config) IsDebug() bool {
	return "prod" != config.env
}
