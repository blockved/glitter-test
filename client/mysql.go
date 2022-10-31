package client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"otherpay-test/common"
)

var DB *sql.DB
var DBIndex3 *sql.DB

func MysqlClient() *sql.DB {
	if DB == nil {
		config := common.SerConf.MysqlCfg
		DB, _ = GetDB(config)
	}
	return DB
}

func MysqlClientIndex3() *sql.DB {
	if DBIndex3 == nil {
		config := common.SerConf.MysqlIndex3Cfg
		DBIndex3, _ = GetDB(config)
	}
	return DBIndex3
}

func GetDB(config common.MysqlConfig) (*sql.DB, error) {
	dsn := config.Dsn
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
