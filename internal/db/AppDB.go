package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"microcosm/internal/config"
)

var db *gorm.DB

type AppDB struct {
	dbConfig *DbConfig
}

type DbConfig struct {
	Dialect string
	Link    string
}

func (appDB *AppDB) Start(config *config.Config) {
	appDB.initConfig(config)
	dbConfig := appDB.dbConfig
	var err error
	db, err = gorm.Open(dbConfig.Dialect, dbConfig.Link)
	if err != nil {
		log.Fatalln("AppDB start error", err)
		return
	}

	db.LogMode(config.IsDebug())
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	log.Info("AppDB started")
}

func (appDB *AppDB) Stop(config *config.Config) {
	err := db.Close()
	if err != nil {
		log.Error("AppDB stop error", err)
	} else {
		log.Info("AppDB stoped")
	}
}

func (appDB *AppDB) initConfig(config *config.Config) {
	dbConfig := new(DbConfig)
	dbConfig.Dialect = "mysql"
	err := config.GetCfg().Section("db").MapTo(dbConfig)
	if err != nil {
		log.Fatalln("get db config error", err)
	}
	appDB.dbConfig = dbConfig
}

func New() *AppDB {
	appDB := new(AppDB)
	return appDB
}

func GetOrm() *gorm.DB {
	if db == nil {
		log.Fatalln("AppDB is not start")
	}
	return db
}
