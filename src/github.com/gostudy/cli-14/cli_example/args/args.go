package args

import (
	"github.com/urfave/cli"
)




func Cmd() []cli.Flag {

		var language string
		var port int

		cli.IntFlag{
		Name:        "port, p",
		Value:       8080,
		Usage:       "listen prot",
		Destination: &port,
		},
		cli.StringFlag{
		Name:        "lang, l",
		Value:       "china",
		Usage:       "read from file",
		Destination: &language,
		},
}

