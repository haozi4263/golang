package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func main()  {
	app := cli.NewApp()  //生成一个app
	app.Name = "gg"		//app名字

	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {  //执行动作
		var cmd string
		if c.NArg() > 0 {  //判断用户传递过来参数个数
			cmd = c.Args()[0]
		}
		fmt.Println("hello golang cmd:", cmd)
		return nil
	}

	app.Run(os.Args)  //执行调用函数
}
