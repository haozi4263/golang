package logger

import (
	"os"
	"fmt"
)

type ConsoleLogger struct {
	level int
}

func (c *ConsoleLogger) SetLevel(level int)  {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	c.level = LogLevelDebug
}

func NewConsoleLog(config map[string]string ) (log LogInterface, err error) {  //实例化
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level")
		return
	}
	level := getLogLevel(logLevel)
	log =  &ConsoleLogger{
		level:level,
	}
	return
}




func (c *ConsoleLogger) Debug(format string, args ...interface{})  {   //接口方法
	if (c.level > LogLevelDebug) {
		return
	}
	writeLog(os.Stdout, LogLevelDebug, format, args...)

}

func (c *ConsoleLogger) Trace(format string, args ...interface{})  {   //接口方法
	if (c.level > LogLevelTrace) {
		return
	}
	writeLog(os.Stdout, LogLevelTrace, format, args...)
}

func (c *ConsoleLogger) Info(format string, args ...interface{})  {   //接口方法
	if (c.level > LogLevelInfo) {
		return
	}
	writeLog(os.Stdout, LogLevelInfo, format, args...)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{})  {   //接口方法
	if (c.level > LogLevelWarn) {
		return
	}
	writeLog(os.Stdout, LogLevelWarn, format, args...)
}

func (c *ConsoleLogger) Error(format string, args ...interface{})  {   //接口方法
	if (c.level > LogLevelError) {
		return
	}
	writeLog(os.Stdout, LogLevelError, format, args...)
}

func (f *ConsoleLogger) Fatal(format string, args ...interface{})  {   //接口方法
	if (f.level > LogLevelFatal) {
		return
	}
	writeLog(os.Stdout, LogLevelFatal, format, args...)
}

func (c *ConsoleLogger) Close() {

}
func (c *ConsoleLogger) Init() {

}
