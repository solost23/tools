package mongo

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestNewMongoConnect(t *testing.T) {
	config := &Config{
		Hosts:      []string{"localhost"},
		AuthSource: "",
		UserName:   "",
		Password:   "",
		Database:   "mongo-test",
		Timeout:    30,
	}
	ctx := context.Background()
	mongoClient, err := NewMongoConnect(ctx, config)
	if err != nil {
		t.Error(err.Error())
	}
	// 插入一条数据查看是否成功
	_, err = mongoClient.Database(config.Database).Collection("user").InsertOne(ctx, bson.M{
		"user_name": "ty",
		"age":       20,
	})
	if err != nil {
		t.Error(err.Error())
	}
}
