package action

import (
	"fmt"
	"github.com/urfave/cli"
	"os/exec"
	"github.com/gostudy/cli-14/cli_example/xkube/global"
	//"github.com/gostudy/cli-14/cli_example/xkube/cmd"
)



func Pod(c *cli.Context) {
	name := "Nefertiti"
	if c.NArg() > 0 {
		name = c.Args().Get(1)
		
		fmt.Printf(" if name:%v\n", name)
	}


	K8s_get_pod(c.Args().Get(0),c.Args().Get(1) )
	if c.NumFlags() > 0 {
		a := c.Command.Description
		fmt.Printf("a:%v\n", a)
	}


}


func Svc(c *cli.Context) {

	fmt.Printf("xkube app")
	app := `kubectl get svc`
	cmd := exec.Command(app)
	out, _ := cmd.Output()
	fmt.Println(string(out))

}

func K8s_get_pod(app, podname string)  {
	fmt.Printf("command run Ns:%v\n", global.Pods)

	app = `get  -n im-dev ` + app + "|grep "  + podname
	fmt.Printf("cmd:%v\n", app)
	cmd := exec.Command("kubectl.exe", app)
	out, _ := cmd.Output()
	fmt.Println(string(out))
}