package main

import "fmt"

type Employer interface {
	CalcSalary() float32
}

type Programer struct {   //定义一个程序员类
	name string
	base float32    //基本工资
	extra float32   //奖金系数
}

func NewProgramer(name string, base float32,extra float32)  Programer {
	return Programer{
		name: name,
		base: base,
		extra: extra,
	}
}

func (p Programer) CalcSalary() float32  {   //计算程序员工资
	return p.base
}



type Sale struct {   //定义一个销售类
	name string
	base float32    //基本工资
	extra float32   //奖金系数
}

func NewSale(name string,base,extra float32) Sale  {   //初始化
	return Sale{
		name: name,
		base: base,
		extra: extra,
	}
}

func (p Sale)  CalcSalary() float32 {   //定义销售工资计算方式
	return p.base + p.extra*p.base*0.5
}

func calcALL( e []Employer) float32  {   //定义求总的工资数切片
	var cost float32
	for _, v := range e {
		cost = cost + v.CalcSalary()
	}
	return cost
}

func main()  {
	p1 := NewProgramer("jude1",2000.0,0)
	p2 := NewProgramer("jude2",2500.0,0)
	p3 := NewProgramer("jude3",3000.0,0)


	s1 := NewSale("销售1", 1000.0, 0.5)
	s2 := NewSale("销售2", 1500.0, 0.5)
	s3 := NewSale("销售3", 2000.0, 0.5)

	var employList []Employer
	employList = append(employList, p1)
	employList = append(employList, p2)
	employList = append(employList, p3)

	employList = append(employList, s1)
	employList = append(employList, s2)
	employList = append(employList, s3)

	cost := calcALL(employList)
	fmt.Printf("这个月人力成本:%f\n", cost)

}