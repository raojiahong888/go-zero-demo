package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	Mysql    MysqlConfig  // mysql configuration
}

type MysqlConfig struct {
	Host         string
	Port         int
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	MaxConnLifetime time.Duration
	SqlLog       bool
	Timezone    string
}
