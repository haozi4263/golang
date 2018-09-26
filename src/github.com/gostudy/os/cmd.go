package main

import (
"fmt"
"os/exec"
)

func ping()  {
	ping := `ls /
			pwd
			ls -l
            docker ps
		`
	cmd := exec.Command("/bin/bash", "-c", ping)
	out, _ :=cmd.Output()
	fmt.Println(string(out))
}


func main()  {
	ping()
}