package logger

type LoggerInterfacel interface {
	SetLevel(level int)
	Debug(loging string, args ...interface{})
	Traceing(loging string, args ...interface{})
	Info(loging string, args ...interface{})
	Warn(login string, args ...interface{})
	Error(loging string, args ...interface{})
	Fatal(loging string, args ...interface{})
	Close()
}