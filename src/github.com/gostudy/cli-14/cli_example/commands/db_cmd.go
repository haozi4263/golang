package commands

import (
	"github.com/gostudy/cli-14/cli_example/action"
	"github.com/urfave/cli"
)

func Db_cmd() []cli.Command {
	return []cli.Command{
		{
			Name:   "insert",
			Usage:  "insert data",
			Action: action.Insert,
		},
		{
			Name:   "delete",
			Usage:  "delete data",
			Action: action.Delete,
		},
	}
}
