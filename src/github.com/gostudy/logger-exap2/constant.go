package logger

const (
	Debug = iota
	Traceing
	Info
	Warn
	Error
	Fatal
)

func GetLevelint(levle int) string  {
	switch levle {
	case Debug:
		return "DEBUG"
	case Traceing:
		return "TRACEING"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}
	return "UNKNOW"
}

func GetLevelstr(level string) int  {
	switch level {
	case "debug":
		return Debug
	case "traceing":
		return Traceing
	case "info":
		return Info
	case "warr":
		return Warn
	case "err":
		return Error
	case "fata":
		return Fatal
	}
	return Debug
}

