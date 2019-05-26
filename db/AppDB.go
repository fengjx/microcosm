package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"microcosm/conf"
)

var db *gorm.DB

type AppDB struct {
}

func (appDB *AppDB) Start(config *conf.Config) {
	var err error
	db, err = gorm.Open(config.GetString("db", "dialect"), config.GetString("db", "link"))
	if err != nil {
		log.Fatalln("AppDB start error")
		return
	}

	db.LogMode(config.IsDebug())
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	log.Info("AppDB started")
}

func (appDB *AppDB) Stop(config *conf.Config) {
	err := db.Close()
	if err != nil {
		log.Error("AppDB stop error", err)
	} else {
		log.Info("AppDB stoped")
	}
}

func New() *AppDB {
	db := new(AppDB)
	return db
}

func GetDB() *gorm.DB {
	return db
}
