package main

import (
	"fmt"
)

type Hunaner interface {  //子集
	syahi()
}

type Personer interface {
	Hunaner  //匿名字段，继承了sayhi()
	sing(lrc string)
}

type Student struct {
	name		string
	id 			int
}

func (tmp *Student) sayhi()  {
	fmt.Printf("student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string)  {
	fmt.Println("student在唱歌:", lrc)
}


func main()  {
	//超集可以转换成子集，反过来不可以
	var iPro Personer  //超集
	iPro = &Student{"jude", 111}
	var i Hunaner      //子集

	//iPro = i  //error
	i = iPro
	i.syahi()

}