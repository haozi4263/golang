package logger

import (
"testing"
)

func TestFileLogger(t *testing.T)  {   //测试函数，单元测试
	logger := NewLogger(0, "C:/log", "test")  //初始化


	logger.Debug("user id[%d] is come from chian", 32442)
	logger.Warn("test warn log")
	logger.Error("test Error log")
	logger.Fatal("test fatal log")
	logger.Close()
}