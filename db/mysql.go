package db

import (
	"demo/conf"
	"demo/util"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"os"
)

var Db *gorm.DB

func init()  {
	var err error
	var dbConfig  = conf.Conf.Db
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)

	Db, err = gorm.Open(conf.Conf.Db.Dialects, url)

	if err != nil {
		panic(err)
	}

	if Db.Error != nil {
		panic(Db.Error)
	}

	Db.DB().SetMaxIdleConns(dbConfig.MaxIdle)
	Db.DB().SetMaxOpenConns(dbConfig.MaxOpen)
	logger := log()
	Db.SetLogger(logger)
	Db.LogMode(true)
	logger.Info("mysql connect success")

}

func log() *logrus.Logger {
	logger := util.Log()
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{TimestampFormat:"2006-01-02 15:04:05"}
	return logger
}

