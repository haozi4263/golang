package main

import (
	"fmt"
	"os"
	"log"

	"sort"
	"github.com/urfave/cli"
)

const (
	Name =		 "cmd"
	UsageText =		"cmd -p/-l [global options] command [command options] [arguments...]"
	Version   =   `
			 ___       ___       ___       ___
   /\__\     /\  \     /\__\     /\  \
  |::L__L   _\:\  \   |::L__L   _\:\  \
  |:::\__\ /\/::\__\  |:::\__\ /\/::\__\
  /:;;/__/ \::/\/__/  /:;;/__/ \::/\/__/
  \/__/     \:\__\    \/__/     \:\__\
             \/__/               \/__/   v0.1

`
  Auther =		"Jude"
  Email = 		"329672369@qq.com"


)

func main()  {
	var language string
	var port int

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

	app.Flags = []cli.Flag {
		cli.IntFlag{
			Name: "port, p",
			Value: 8080,
			Usage: "listen prot",
			Destination: &port,
		},
		cli.StringFlag{
			Name: "lang, l",
			Value: "china",
			Usage: "read from file",
			Destination: &language,
		},
	}
	app.Commands = []cli.Command {
		{
			Name: "add",
			Aliases: []string{"a"},
			Usage: "calc 1+1",
			Category: "calc",
			Action: func(c *cli.Context) error {
				fmt.Printf("cmd run lang:%s, port:%v\n", language, port)
				fmt.Println("1 + 1 =", 1 + 1)
				return nil
			},
		},
		{
			Name: "sub",
			Aliases: []string{"s"},
			Usage: "calc 5-3",
			Category: "calc",
			Action: func(c *cli.Context) error {
				fmt.Println(" 5 - 3 = ", 5 - 3)
				return nil
			},
		},
		{
			Name: "db",
			Usage: "databasre args",
			Category: "db",
			Subcommands: []cli.Command{
				{
					Name: "insert",
					Usage: "insert data",
					Action: func(c *cli.Context) error {
						fmt.Println(" insert in to values",)
							return nil
					},
				},
				{
					Name: "delete",
					Usage: "delete data",
					Action: func(c *cli.Context) error {
						fmt.Println(" delete in to values",)
						return nil
					},
				},
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		//fmt.Println(c.String("lang"), c.Int("port"))
		//fmt.Println(language)
		fmt.Printf("cmd run lang:%s, port:%v\n", language, port)
		return nil
	}

	app.Before = func(c *cli.Context) error {
		fmt.Println("app runing before")
		return nil
	}
	app.After = func(c *cli.Context) error {
		fmt.Println("app runing after")
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))


	cli.HelpFlag = cli.BoolFlag{
		Name: "help, h",
		Usage: "help",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name: "print version, v",
		Usage: "print version",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}




}

