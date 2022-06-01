package log

import "testing"

func TestNewLogger(t *testing.T) {
	log := NewLogger("/logs/")
	log.LogMode(DebugLevel)
	log.Info("这是一条普通信息日志")
	log.Warn("这是一条警告日志")

	LOGGER.Info("这是第二条普通日志")
}
