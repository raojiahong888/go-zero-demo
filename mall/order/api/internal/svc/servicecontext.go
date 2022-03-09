package svc

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/mall/order/api/internal/config"
	"go-zero-demo/mall/order/api/internal/middleware"
	"go-zero-demo/mall/order/api/internal/types"
	//"go-zero-demo/mall/user/rpc/user"
	"net/url"
	"os"
	"time"

	//"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	//UserRpc user.User
	UserMiddleware rest.Middleware
	Db *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	timezone := url.QueryEscape(c.Mysql.Timezone)
	resource := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=%s", c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.Database, c.Mysql.Charset, timezone)
	db, err := gorm.Open("mysql", resource)
	defer db.Close()
	if err != nil {
		fmt.Printf("mysql connect error:%s\r\n", err)
		os.Exit(-1)
	}
	db.SingularTable(true)                                     //Global table name cannot be plural.
	db.AutoMigrate(&types.Order{})
	db.DB().SetMaxIdleConns(c.Mysql.MaxIdleConns) //Maximum number of connections when idle
	db.DB().SetMaxOpenConns(c.Mysql.MaxOpenConns) //Maximum connections
	db.DB().SetConnMaxLifetime(time.Second*c.Mysql.MaxConnLifetime) // default lift time
	if c.Mysql.SqlLog {
		db.LogMode(true)
	}
	return &ServiceContext{
		Config: c,
		//UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserMiddleware: middleware.NewUserMiddleware().Handle,
		Db: db,
	}
}
