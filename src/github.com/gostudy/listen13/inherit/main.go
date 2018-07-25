package main

import "fmt"

type Animal struct {
	Name string
	Sex  string
}

func (a *Animal) Talk()  {
	fmt.Printf("i talk,im is %s\n",a.Name)
}

type PuruAnimal struct {

}

func (p *PuruAnimal) Talk()  {
	fmt.Println("puru dongwu talk")
}



type Dog struct {
	Feet string
	//Animal     //继承Animal所有字段和方法，值类型基于拷贝
	*Animal //指针类型值拷贝指针地址大小,匿名字段继承
	*PuruAnimal  //talk方法冲突
}

func (d *Dog) Eat()  {
	fmt.Println("dog is est")
}

func (d *Dog) Talk()  {
	fmt.Println("dog is takling")
}


func main()  {
	var d  *Dog = &Dog{
		Feet: "four feet",
		Animal: &Animal{  //初始化
			Name:"dog",
			Sex:"xiong",
		},
		}
	d.Name = "dog"
	d.Sex  = "xiong"
	d.Eat()
	d.Talk()  //先使用当前实例内部方法，没有的话找匿名函数里方法
	d.Animal.Talk()   //冲突解决加上结构体方法

}
