package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/nsqio/go-nsq"
)


var producer *nsq.Producer

func initProducer(str string) error {
	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)

	if err != nil {
		return err
	}
	return nil
}


func main()  {
	nsqAddress := "192.168.10.106:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init produceer failed, err:%v\n", err)
		return
	}

	//duqu kongzhitai shuru
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string failed, err:%v]n", err)
			continue
		}

		data = strings.TrimSpace(data)
		if data == "stop" {
			break
		}

		err = producer.Publish("order_queue", []byte(data))
		if err != nil {
			fmt.Printf("publish message failed, err:%v\n", err)
			continue
		}
		fmt.Printf("publish data:%s success\n", data)
	}
}

