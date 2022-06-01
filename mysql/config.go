package mysql

import (
	"time"

	"gorm.io/gorm/logger"
)

type Config struct {
	UserName string
	Password string
	Host     string
	Port     int
	DB       string
	Charset  string

	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration

	Logger logger.Interface
}
