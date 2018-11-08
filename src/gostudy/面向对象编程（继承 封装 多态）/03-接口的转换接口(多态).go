package main

import "fmt"

type Humaner interface {  //子集
	sayhi()//接口方法
}
type Personer interface {  //超集
	Humaner //匿名字段，继承了sayhi
	sing(lrc string)
}
type Student struct {
	name string
	id int
}
//Student实现了此方法
func (s *Student) sayhi()  {
	fmt.Printf("Student:[%s, %d] sayhi\n", s.name, s.id)
}

func (s *Student) sing (lrc string)  {
	fmt.Println("Student在唱着:", lrc)
}

func main()  {
	var i Personer
	s := &Student{"jude", 11}
	i = s

	i.sayhi()
	i.sing("死了都要爱")//继承过来的方法


}