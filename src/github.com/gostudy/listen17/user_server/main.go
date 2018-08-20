package main

import (
	//"github.com/gostudy/listen17/log"

	"github.com/gostudy/logger"
	"time"
)

//var log logger.LogInterface

func initLogger(logPath, logName string, level string)  {
	//log = logger.NewFileLogger(level, logPath, logName)
	m :=make(map[string]string, 8)
	m["log_path"] = logPath
	m["log_name"] = "user_server"
	m["log_level"] = level
	err := logger.InitLogger("console",m)
	if err != nil {
		return
	}
	logger.Debug("init logger success")
	return
}

func Run()  {
	for {
		logger.Debug("user server is running")
		logger.Info("-----------------------info日志-------------------------")
		time.Sleep(time.Second)
	}
}

func main()  {
	initLogger("c:/logs/", "user_server", "debug")
	Run()
	return
}























