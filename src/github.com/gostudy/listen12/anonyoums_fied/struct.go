package main   //嵌套匿名结构体
import "fmt"

type Address struct {
	Province	string
	City		string   //字段小写，在函数外部无法访问
	CreateTime 	string
}

type Email struct {
	account string
	CreateTime string
}

type User struct {
	Username string
	Sex string
	*Address  //匿名嵌套结构体
}

type User01 struct {
	City string
	Username string
	Sex string
	*Address  //存在冲突匿名结构体,先访问自身结构体中字段，没有的话去访问匿名结构体中字段，如匿名结构体未初始化，会报错
	*Email
}




func test1()  {   //嵌套匿名结构
	var user User
	user.Username = "jude"
	user.Sex = "man"
	user.Address = &Address{   //第一种指针方式
		Province:"beijing",
		City:	"bj",

	}
	//第二种方式，访问Province如果主结构体种没有，如果有匿名结构是否存在，有的话会使用匿名结构体中赋值
	user.Province = "beijing"
	user.City = "bj01"

	fmt.Printf("user=%#v addr=%#v\n", user, user.Address )
}

func test02() {
	var user User01
	user.Username = "jude01"
	user.Sex = "man"
	user.City = "bj"
	//user.Address = &Address{   //第一种指针方式
	user.Address = new(Address) //初始化匿名结构体积
	user.Email = new(Email)

	fmt.Printf("user=%#v\n", user)
	user.Address.City = "Bj01"
	fmt.Printf("user=%#v city of address:%s\n",user, user.Address.City)



	user.Address.CreateTime = "002"  //如果冲突的话需要指定匿名结构体加字段访问
	user.Address.CreateTime = "001"  //如果2个匿名结构体中存在相同字段，访问需要加上访问具体某个结构体中字段

	fmt.Printf("user=%#v creaTime=%s city of address=%s\n", user, user.Address.CreateTime, user.Address.City)

}

func main()  {
	//test1()
	test02()
}