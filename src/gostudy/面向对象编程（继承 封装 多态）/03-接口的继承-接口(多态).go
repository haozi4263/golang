package main

import "fmt"

type Humaner interface {
	sayhi()//接口方法
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

//定义一个普通变量，函数的参数为接口类型
//只有一个函数，可以有不同表现，多态
func WhoSayHi(i Humaner)  {
	i.sayhi()
}
func main()  {
	s := &Student{"jude", 11}
	t := &Teacher{"wuhan", "golang"}
	var str Mystr = "hello jude"

	//调用同一函数，不同表现，多态，多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	since := make([]Humaner, 3)
	since[0] = s
	since[1] = t
	since[2] = &str

	for _, i := range since {
		i.sayhi()
	}

}