package log


type LogInterface interface {   //定义log接口规范
	LogDebug(msg string)   //规范，无论打印到文件或者终端都需要LogDebug方法
	LogWarn(msg string)
}
