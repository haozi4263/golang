package iniconfig

import (
	"testing"
	"fmt"
	"io/ioutil"

)

type Config struct {
	ServerConf	ServerConfig  `ini:"server"`
	MysqlConf	MysqlConfig   `ini:"mysql"`

}

type ServerConfig struct {
	Ip string		`ini:"ip"`  //tag映射到配置文件中小写
	Port int		`:"port"`
}

type MysqlConfig struct {
	Username string		`ini:"username"`
	Passwd 	 string		`ini:"passwd"`
	Database string		`ini:"database"`
	Host     string		`ini:"host"`
	Port     int		`ini:"port"`
}

func TestIniConfig(t *testing.T)  {
	fmt.Println("hello")
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Errorf("read file failed")
	}
	var conf Config
	err = UnMarshal(data,&conf)  //传入指针地址,因为结构体是值传递的
	if err != nil {
		t.Errorf("unmarshal failed ,err:%v", err)
	}

	t.Logf("unmarshal success, conf:%#v", conf)

}