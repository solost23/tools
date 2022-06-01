package logger

import (
	"context"
	"fmt"
	"strings"
	"time"

	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"github.com/solost23/tools/log"
)

const (
	Silent gormLogger.LogLevel = iota + 1
	Error
	Warn
	Info
	Debug
)

const (
	CLOSEDDLANDDMLANDDQL uint8 = 7
	CLOSEDDL                   = 1
	CLOSEDML                   = 2
	CLOSEDQL                   = 4
	CLOSEDDLANDCLOSEDML        = 3
	CLOSEDDLANDDQL             = 5
	CLOSEDMLANDCLOSEDQL        = 6
)

// 选择模式:实现兼容多代码版本
var (
	DEFAULT_MESSAGE = Message{CloseDDL: 0, CloseDML: 0, CloseDQL: 0}

	DDL = []string{"CREATE", "DROP", "ALTER"}
	DML = []string{"INSERT", "DELETE", "UPDATE"}
	DQL = []string{"SELECT"}
)

type Message struct {
	CloseDDL, CloseDML, CloseDQL uint8
}

type Option func(*Message)

func WithCloseDDL(closeDDL bool) Option {
	return func(msg *Message) {
		if closeDDL {
			msg.CloseDDL = 1
		}
	}
}

func WithCloseDML(closeDML bool) Option {
	return func(msg *Message) {
		if closeDML {
			msg.CloseDML = 2
		}
	}
}

func WithCloseDQL(closeDQL bool) Option {
	return func(msg *Message) {
		if closeDQL {
			msg.CloseDQL = 4
		}
	}
}

type MysqlLogger struct {
	logLevel                            gormLogger.LogLevel
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
	Message
}

func NewMysqlLogger(opts ...Option) gormLogger.Interface {
	var (
		infoStr      = "%v "
		warnStr      = "%v "
		errStr       = "%v "
		traceStr     = "%v [%.3fms] [rows:%v] %v"
		traceWarnStr = "%v %v [%.3fms] [rows:%v] %v"
		traceErrStr  = "%v %v [%.3fms] [rows:%v] %v"
	)
	mysqlLogger := &MysqlLogger{
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
	mysqlLogger.Message = DEFAULT_MESSAGE
	for _, o := range opts {
		o(&mysqlLogger.Message)
	}
	return mysqlLogger
}

func (l *MysqlLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *l
	newLogger.logLevel = level
	return &newLogger
}

func (l *MysqlLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if !l.check(l.CloseDDL|l.CloseDML|l.CloseDQL, msg) {
		return
	}
	if l.logLevel <= Info {
		logStr := fmt.Sprintf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		log.Info(ctx, logStr)
	}
}

func (l *MysqlLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if !l.check(l.CloseDDL|l.CloseDML|l.CloseDQL, msg) {
		return
	}
	if l.logLevel >= Warn {
		logStr := fmt.Sprintf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		log.Warn(ctx, logStr)
	}
}

func (l *MysqlLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if !l.check(l.CloseDDL|l.CloseDML|l.CloseDQL, msg) {
		return
	}
	if l.logLevel >= Error {
		logStr := fmt.Sprintf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		log.Error(ctx, logStr)
	}
}

func (l *MysqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.logLevel <= Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, _ := fc()
	if !l.check(l.CloseDDL|l.CloseDML|l.CloseDQL, sql) {
		return
	}
	log.Trace(ctx, l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
	return
}

// 判断是什么操作
func (l *MysqlLogger) check(operator uint8, msg string) bool {
	action := strings.Trim(strings.Split(msg, " ")[0], " ")
	if operator == CLOSEDDLANDDMLANDDQL {
		return false
	}
	if operator == CLOSEDDL {
		var DmlAndDql []string
		DmlAndDql = append(DML, DQL...)
		for _, item := range DmlAndDql {
			if item == action {
				return true
			}
		}
		return false
	}
	if operator == CLOSEDML {
		var DdlAndDql []string
		DdlAndDql = append(DDL, DQL...)
		for _, item := range DdlAndDql {
			if item == action {
				return true
			}
		}
		return false
	}
	if operator == CLOSEDQL {
		var DdlAndDml []string
		DdlAndDml = append(DDL, DML...)
		for _, item := range DdlAndDml {
			if item == action {
				return true
			}
		}
		return false
	}
	if operator == CLOSEDDLANDCLOSEDML {
		for _, item := range DQL {
			if item == action {
				return true
			}
		}
		return false
	}
	if operator == CLOSEDDLANDDQL {
		for _, item := range DML {
			if item == action {
				return true
			}
		}
		return false
	}
	if operator == CLOSEDMLANDCLOSEDQL {
		for _, item := range DDL {
			if item == action {
				return true
			}
		}
		return false
	}
	return true
}
