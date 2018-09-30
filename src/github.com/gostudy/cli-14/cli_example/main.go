package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

const (
	usage = "golang cmd "
	version = "1.0"
)

func main()  {
	app := cli.NewApp()  //生成一个app
	app.Name = "gg"		//app名字
	app.Version = version
	app.Usage = usage

	app.Flags = []cli.Flag {
		cli.BoolFlag{
			Name: "msg",
			Usage: "打印消息",
		},
		cli.StringFlag{
			Name: "log",
			Value: "/dev/null",
			Usage:"设置log日志存储路径",
		},
	}

	app.Commands= []cli.Command{
		createCommand
	}


	//app.Action = func(c *cli.Context) error {  //执行动作
	//	var cmd string
	//	if c.NArg() > 0 {  //判断用户传递过来参数个数
	//		cmd = c.Args()[0]
	//	}
	//	fmt.Println("hello golang cmd:", cmd)
	//	return nil
	//}

	app.Run(os.Args)  //执行调用函数
}



var createCommand = cli.Command{
	Name: "create",
	Usage: "create a xxx",
	ArgsUsage: "xxx",
	Flags: []cli.Flag {
		cli.StringFlag{
			Name: "bundle, b",
			Value: "",
			Usage: "path to root or director",
		},
	},
	Action: func(c *cli.Context) err {
		fmt.Printf("testing")
		return nil
	},
}
