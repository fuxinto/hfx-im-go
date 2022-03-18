package svc

import (
	"HIMGo/service/user/model"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func pgSqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Auth{},
	)
	if err != nil {
		log.Fatal("创建表失败")
	}
	logx.Info("register table success")
}

func GormPgSql(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,

		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		log.Panic("数据库打开失败")
	}
	pgSqlTables(db)
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return config.Database_table_prefix + defaultTableName;
	//}
	////// 启用Logger，显示详细日志
	//Db.LogMode(true)
	logx.Info("数据库打开成功")
	return db
}
