package log

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

type Level uint32

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

const (
	YEARMONTHDAY                 = "2006-01-02"
	YEARMONTHDAYHOURMINUTESECOND = "2006-01-02 15:04:05"
)

type Logger struct {
	logrus.Logger
}

var LOGGER *Logger

func NewLogger(logFilePath string) *Logger {
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + logFilePath
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		log.Println(err.Error())
	}
	logFileName := fmt.Sprintf("%s%s", time.Now().Format(YEARMONTHDAY), ".log")
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(logFileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
		}
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err.Error())
	}
	logger := &Logger{}
	logger.Out = src
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: YEARMONTHDAYHOURMINUTESECOND,
	})
	LOGGER = logger
	return logger
}

func (l *Logger) LogMode(level Level) {
	l.Level = logrus.Level(level)
	LOGGER.Level = logrus.Level(level)
}

func Error(ctx context.Context, logStr string) {
	LOGGER.Error(ctx, logStr)
}

func Info(ctx context.Context, logStr string) {
	LOGGER.Info(ctx, logStr)
}

func Warn(ctx context.Context, logStr string) {
	LOGGER.Warn(ctx, logStr)
}

func Trace(args ...interface{}) {
	LOGGER.Trace(args...)
}
