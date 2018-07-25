package main

import "fmt"


type People struct {
	Name string
	Country string
}

func (p People) Printn()  {    //p相当于是p1的拷贝
	fmt.Printf("name=%s country=%s\n",p.Name,p.Country)
}

func (p People) Set(name string, country string)  {
	p.Name = name
	p.Country = country
}

func (p *People)  SetV2(name string, country string) {   //指针p,p是存放people实例地址,通过指针类型来修改
	p.Country = country
	p.Name = name
}



func main()  {
	var p1 People = People{  //实例化
		Name:     "jude",
		Country:  "wuhan",
	}
	p1.Printn()  //p1传递给Printn()函数
	p1.Set("lj","usa")
	p1.Printn()  //无法对结构体修改

	//(&p1).SetV2("Lj","usa")    //取指针类型值
	p1.SetV2("P2","C2")  //编译器发现Setv2发现需要一个地址会自动转换&p1
	p1.Printn()   //通过指针类型修改实例值
}