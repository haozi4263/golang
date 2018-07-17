package main

import "fmt"


//质数判断，能被1和自身整出的数
func justify(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n;i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func example1()  {
	for i :=2; i < 100; i++ {
		if justify(i) == true {
			fmt.Printf("[%d] is prime\n",i)
		}
	}
}

func is_shuixianhua(n int ) bool  {

}

func example2()  {
	for i :=100; i < 1000; i++ {
		if is_shuixianhua(i) == true {
			fmt.Printf("[%d] is shuixiaohua\n",i)
		}
	}
}








func main()  {
	example1()
}