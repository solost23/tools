package mysql

import (
	"testing"
)

type User struct {
	UserName string
}

func TestNewMysqlConnect(t *testing.T) {
	config := &Config{
		UserName: "root",
		Password: "123",
		Host:     "localhost",
		Port:     3306,
		DB:       "mysql-test",
		Charset:  "utf8mb4",
	}
	db, err := NewMysqlConnect(config)
	if err != nil {
		t.Errorf("err: %v \n", err.Error())
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		t.Errorf("err: %v \n", err.Error())
	}
}
