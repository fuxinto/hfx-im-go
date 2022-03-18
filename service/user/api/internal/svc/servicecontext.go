package svc

import (
	"HIMGo/service/user/api/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	//SqlConn sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//SqlConn: postgres.New(c.PostgreSQLConf.DataSource),
		Db: GormPgSql(c.PostgreSQLConf.DataSource),
	}
}
