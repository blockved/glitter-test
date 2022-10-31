package common

import (
	"github.com/BurntSushi/toml"
	"time"
)

type ServiceConfig struct {
	MysqlCfg       MysqlConfig `toml:"mysql"`
	MysqlIndex3Cfg MysqlConfig `toml:"mysql_index3"`
}

type MysqlConfig struct {
	Dsn         string        `toml:"dsn"`                       // data source name
	DriverName  string        `toml:"driver_name"`               // data source driver name
	Retry       int           `toml:"retry"`                     // retry time
	MaxIdle     int           `toml:"db_conn_pool_max_idle"`     // zero means defaultMaxIdleConns; negative means 0
	MaxOpen     int           `toml:"db_conn_pool_max_open"`     // <= 0 means unlimited
	MaxLifetime time.Duration `toml:"db_conn_pool_max_lifetime"` // maximum amount of time a connection may be reused
}

var SerConf ServiceConfig
var confPath string = "/home/work/otherpay-test/conf/service.conf"

func Load(confFile string) (err error) {
	_, err = toml.DecodeFile(confFile, &SerConf)
	return
}

func init() {
	Load(confPath)
}
