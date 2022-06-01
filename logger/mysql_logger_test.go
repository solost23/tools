package logger

import (
	"context"
	"github.com/solost23/tools/log"
	"testing"
)

func TestNewMysqlLogger(t *testing.T) {
	logger := log.NewLogger("/logs/")
	logger.LogMode(log.DebugLevel)
	ctx := context.Background()
	NewMysqlLogger().LogMode(Info).Info(ctx, "INSERT")
	NewMysqlLogger(WithCloseDDL(true)).LogMode(Info).Info(ctx, "INSERT")
	NewMysqlLogger(WithCloseDDL(true), WithCloseDQL(true)).LogMode(Info).Info(ctx, "INSERT")
	NewMysqlLogger(WithCloseDDL(true), WithCloseDML(true), WithCloseDQL(true)).LogMode(Info).Info(ctx, "INSERT")
}
