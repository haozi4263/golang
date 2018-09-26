package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var language string
	var recusive bool
	app := cli.NewApp()
	app.Flags = []cli.Flag{   //切片
		cli.StringFlag{ //字符选项
			Name:        "lang, l",  //--lang长选项，-l短选项
			Value:       "english",  //默认值
			Usage:       "select language",
			Destination: &language,
		},
		cli.BoolFlag{
			Name:        "recusive, r",
			Usage:       "recusive for the greeting",
			Destination: &recusive,
		},
	}

	app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args()[0]
			fmt.Println("cmd is ", cmd)
		}
		fmt.Println("recusive is ", recusive)
		fmt.Println("language is ", language)
		return nil
		//return process(c)  //调用函数执行业务逻辑
	}
	app.Run(os.Args)
}

//func process(c *cli.Context) (err error)  {
//
//}