package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConnect(config *Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.DB,
		config.Charset,
	)
	gormConfig := &gorm.Config{}
	gormConfig.Logger = config.Logger
	db, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return db, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return db, err
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifeTime)
	return db, err
}
