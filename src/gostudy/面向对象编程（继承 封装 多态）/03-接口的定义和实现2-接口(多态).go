package main

import "fmt"

type Humaner interface {
	sayhi()
}
type Student struct {
	name string
	id int
}
//Student实现了此方法
func (s *Student) sayhi()  {
	fmt.Printf("Student:[%s, %d] sayhi\n", s.name, s.id)
}

//Tercher实现此方法
type Teacher struct {
	addr string
	group string
}
func (t *Teacher) sayhi() {
	fmt.Printf("Teacher:[%s, %s] sayhi\n", t.addr, t.group)
}

//Mystr实现此方法
type Mystr string

func (m *Mystr) sayhi()  {
	fmt.Printf("Mystr:[%s] sayhi\n", *m)
}

func main()  {
	//定义接口类型
	var i Humaner

	//只要实现了此接口方法的类型，那么这个类型的变量（接受者类型）就可以给i赋值
	s := &Student{"jude", 11}
	i = s
	i.sayhi()

	t := &Teacher{"wuhan", "golang"}
	i = t
	i.sayhi()

	var str Mystr = "hello jude"
	i = &str
	i.sayhi()
}