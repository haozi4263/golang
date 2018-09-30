package cmd

import (
	"github.com/urfave/cli"
	"github.com/gostudy/cli-14/cli_example/xkube/action"
	//"fmt"

	//"os/exec"

)

var Ns string = "im-dev"

func Get() []cli.Command {
		return []cli.Command{
			{
				Name:     "get",
				Aliases:  []string{"g"},
				Usage:    "get",
				Category: "get",
				//Subcommands: Get_cmd(),
				Action: action.Pod,
				Description: Ns,
			},

		}

}
