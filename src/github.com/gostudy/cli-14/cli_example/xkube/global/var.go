package global

import (
	//"fmt"
	"github.com/urfave/cli"

)

//var Pods string = "pod"
var Ns string  = "im-dev"
var Pods string = ""






func Flags() ([]cli.Flag) {
	return []cli.Flag {
		cli.StringFlag{
			Name:        "ns",
			Value:       "im-dev",
			Usage:       "xkube -n im-dev",
			Destination: &Ns,
			EnvVar: "im-dev",
		},
	}
	//fmt.Printf("ns:%v\n", Ns)
}

