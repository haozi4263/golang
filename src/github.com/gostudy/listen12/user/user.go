package user



type User struct {  //自定义静态包
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}



func NewUser(username string, sex string, age int, avater string) *User {   //自定义构照函数,返回user类型指针，不适用函数拷贝减小性能
	//user := &User{      //方法1.分配一个指向user类型的内存实例
	//	Username:	username, //初始化
	//	Sex:		sex,
	//	Age:		age,
	//	AvatrUrl:	avater,
	//}
	user := new(User)   //方法2
	user.Username = username
	user.Sex = sex
	user.Age = age
	user.AvatarUrl = avater
	return user
}
