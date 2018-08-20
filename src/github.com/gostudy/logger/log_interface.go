package logger


type LogInterface interface {  //定义日志规范和方法
	Init()
	SetLevel(level int)
	Debug(format string, args ...interface{})  //接受一个可变的参数
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}
