package action

import (
	"fmt"
	"github.com/urfave/cli"
)

func Add(c *cli.Context) {

		fmt.Printf("cmd run add to Commands func port:%v\n")
		fmt.Println("1 + 1 =", 1+1)
		//fmt.Printf("language:%v\n",language)

}

func Sub(c *cli.Context) {
	fmt.Printf("cmd run sub to Commands func\n")
	fmt.Println("5 - 3 =", 5-3)

}

func Insert(c *cli.Context) {
	fmt.Printf("cmd run insert to Commands func\n")
	fmt.Println(" insert in to values")

}

func Delete(c *cli.Context) {
	fmt.Printf("cmd run Delete to Commands func\n")
	fmt.Println(" Delete in to values")

}
