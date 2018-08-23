package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Servers struct {
	Name    xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Encoding string	 `xml:"encoding,attr"`
	Servers []Server `xml:"server"`
}

type Server struct {
	SrcName string `xml:"serverName"`
	SrcIP   string `xml:"serverIP"`
}

func main() {
	data, err := ioutil.ReadFile("./config.xml")
	if err != nil {
		fmt.Printf("read config.xml failed, err:%v\n", err)
		return
	}
	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("xml: %#v\n", servers)
}