package mongo

import (
	"context"
	"net/url"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnect(ctx context.Context, config *Config) (*mongo.Client, error) {
	// connectionString: [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	connectionString := "mongodb://"
	if len(config.UserName) > 0 && len(config.Password) > 0 {
		connectionString = connectionString + config.UserName + ":" + url.QueryEscape(config.Password) + "@"
	}
	connectionString += strings.Join(config.Hosts, ",")
	if len(config.AuthSource) > 0 {
		connectionString += "/admin?authSource=" + config.AuthSource
	}
	clientOptions := options.Client().ApplyURI(connectionString)
	// connect to Mongodb
	timeOut := config.Timeout * time.Second
	clientOptions.ConnectTimeout = &timeOut
	session, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return session, err
	}
	if err = session.Ping(ctx, nil); err != nil {
		return session, err
	}
	return session, nil
}
