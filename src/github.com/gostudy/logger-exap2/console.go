package logger

import "os"

type ConsoleLogger struct {
	level int
}

func NewConsole(level int) LoggerInterfacel {
	return &ConsoleLogger{
		level:level,
	}
}

func (c *ConsoleLogger) SetLevel(level int)  {
	if level < Debug || level > Fatal {
		level = Debug
	}
	c.level = level
}



func (c *ConsoleLogger) Debug(loging string, args ...interface{})  {
	if (c.level > Debug)  {
		return
	}
	writeLogger(os.Stdout, Debug,loging, args...)
}

func (c *ConsoleLogger) Traceing(loging string, args ...interface{})  {
	if (c.level > Traceing)  {
		return
	}
	writeLogger(os.Stdout, Traceing,loging, args...)
}

func (c *ConsoleLogger) Info(loging string, args ...interface{})  {
	if (c.level > Info)  {
		return
	}
	writeLogger(os.Stdout, Info,loging, args...)
}

func (c *ConsoleLogger) Warn(loging string, args ...interface{})  {
	if (c.level > Warn)  {
		return
	}
	writeLogger(os.Stdout, Warn,loging, args...)
}

func (c *ConsoleLogger) Error(loging string, args ...interface{})  {
	if (c.level > Error)  {
		return
	}
	writeLogger(os.Stdout, Error,loging, args...)
}

func (c *ConsoleLogger) Fatal(loging string, args ...interface{})  {
	if (c.level > Fatal)  {
		return
	}
	writeLogger(os.Stdout, Fatal,loging, args...)
}


func ( *ConsoleLogger) Close() {

}


