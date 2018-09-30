package commands

import (
	"github.com/gostudy/cli-14/cli_example/action"
	"github.com/urfave/cli"
	"fmt"
)

func Commands(language string) []cli.Command {

	if language == "china" {
		fmt.Printf("language:%v\n",language)
		return []cli.Command{
			{
				Name:     "add",
				Aliases:  []string{"a"},
				Usage:    "calc 1+1",
				Category: "calc",
				Action:   action.Add,
			},
			{
				Name:     "sub",
				Aliases:  []string{"s"},
				Usage:    "calc 5-3",
				Category: "calc",
				Action:   action.Sub,
			},
			{
				Name:        "db",
				Usage:       "databasre args",
				Category:    "db",
				Subcommands: Db_cmd(),
			},
		}
	}
	return nil

}
