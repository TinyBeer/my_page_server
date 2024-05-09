package svc

import (
	"fmt"
	"personal_page/zero/model/user"
	"personal_page/zero/user/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.GormUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%ds",
		mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.DbName, mysql.Timeout)
	conn := sqlx.NewSqlConn("mysql", dsn)
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewGormUserModel(conn),
	}
}
