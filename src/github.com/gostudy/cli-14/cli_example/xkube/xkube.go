package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gostudy/cli-14/cli_example/xkube/cmd"
	//"github.com/gostudy/cli-14/cli_example/Args"
	"github.com/urfave/cli"
	"sort"
	"github.com/gostudy/cli-14/cli_example/xkube/global"
)

const (
	Name      = "xkube"
	UsageText = "xkube -n [global options] command [command options] [arguments...]"
	Version   = `
			 ___       ___       ___       ___
   /\__\     /\  \     /\__\     /\  \
  |::L__L   _\:\  \   |::L__L   _\:\  \
  |:::\__\ /\/::\__\  |:::\__\ /\/::\__\
  /:;;/__/ \::/\/__/  /:;;/__/ \::/\/__/
  \/__/     \:\__\    \/__/     \:\__\
             \/__/               \/__/   v0.1

`
	Auther = "Jude"
	Email  = "329672369@qq.com"
)



func main() {

	//app.Flags = args.Cmd()
	app := cli.NewApp()
	app.Name = Name
	app.Usage = UsageText
	app.Author = Auther
	app.Email = Email
	app.Version = Version
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Printf("[WARNING] Not Found Command: %s\n", command)
		fmt.Printf("[MESSAGE] Please Type: cmd --help")
	}



	app.Flags = global.Flags()
	//app.Action = func(c *cli.Context) error {
	//	name := "Nefertiti"
	//	if c.NArg() > 0 {
	//		name = c.Args().Get(1)
	//		fmt.Printf(" if name:%v\n", name)
	//	}
	//	return nil
	//}


	app.Commands = cmd.Get()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}



	app.Before = func(c *cli.Context) error {
		//app.Commands = cmd.Get()
		return nil
	}




	//app.Commands = commands.Commands()

	//app.Action = func(c *cli.Context) error {
	//	//fmt.Println(c.String("lang"), c.Int("port"))
	//	//fmt.Println(language)
	//	fmt.Printf("cmd run lang:%s, port:%v\n", language, port)
	//	return nil
	//}


	app.After = func(c *cli.Context) error {
		fmt.Println("app runing after")
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "help",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "print version, v",
		Usage: "print version",
	}

	//err := app.Run(os.Args)
	//if err != nil {
	//	log.Fatal(err)
	//}

}

