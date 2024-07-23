package db_helper

import (
	"fmt"
	"os"
	"time"
	"xorm.io/xorm/log"

	"growth/conf"
	xlog "log"
	"xorm.io/xorm"
)

var dbEngine *xorm.Engine

func initDb() {
	if dbEngine != nil {
		return
	}

	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		conf.GlobalConfig.Db.Username,
		conf.GlobalConfig.Db.Password,
		conf.GlobalConfig.Db.Host,
		conf.GlobalConfig.Db.Port,
		conf.GlobalConfig.Db.Database,
		conf.GlobalConfig.Db.Charset,
	)

	if engin, err := xorm.NewEngine(conf.GlobalConfig.Db.Engine, sourceName); err != nil {
		xlog.Fatalf("dbhelper initDb(%s) error%s\n", sourceName, err.Error())
		return
	} else {
		dbEngine = engin
	}

	// sql log write to stdout
	logger := log.NewSimpleLogger(os.Stdout)
	logger.ShowSQL(conf.GlobalConfig.Db.ShowSql)
	dbEngine.SetLogger(logger)

	if conf.GlobalConfig.Db.ShowSql {
		dbEngine.SetLogLevel(log.DEFAULT_LOG_LEVEL)
	} else {
		dbEngine.SetLogLevel(log.LOG_ERR)
	}

	// more database config
	if conf.GlobalConfig.Db.MaxIdleConns > 0 {
		dbEngine.SetMaxIdleConns(conf.GlobalConfig.Db.MaxIdleConns)
	}

	if conf.GlobalConfig.Db.MaxOpenConns > 0 {
		dbEngine.SetMaxOpenConns(conf.GlobalConfig.Db.MaxOpenConns)
	}
	if conf.GlobalConfig.Db.CoonMaxLifetime > 0 {
		dbEngine.SetConnMaxLifetime(time.Minute * time.Duration(conf.GlobalConfig.Db.CoonMaxLifetime))
	}

}

func GetDb() *xorm.Engine {
	return dbEngine
}
