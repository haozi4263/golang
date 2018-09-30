package cmd

import "github.com/urfave/cli"
import (
	"github.com/gostudy/cli-14/cli_example/xkube/action"
	"github.com/gostudy/cli-14/cli_example/xkube/global"
	//"fmt"

)


var Pods string = ""

func Get_cmd() []cli.Command {
	return []cli.Command{
		{

			Name:   "pod",
			Usage:  "get pod",
			Category: "get",
			Action: action.Pod,
			Description: global.Pods,
		},

		{
			Name:   "svc",
			Usage:  "get svc",
			Category: "get",
			Action: action.Svc,
		},
	}
}

func Pod() (c *cli.Command, a *cli.Context) {
	c.Name = "pod"
	return
	if a.NArg() > 0 {
		c.Name = a.Args().Get(0)
	}
	c.Usage = "get pod"
	c.Category = "get"
	c.Action = action.Pod
	return

}

