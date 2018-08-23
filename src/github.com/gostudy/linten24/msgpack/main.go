package main

import (
	"fmt"
	"math/rand"
	"io/ioutil"
	"github.com/vmihailenco/msgpack"
)

type Person struct {
	Name string
	Age int
	Sex string
}

func writeJson(filename string) (err error) {
	var persons []*Person
	for i := 0; i < 10; i++ {
		p :=&Person{
			Name: fmt.Sprintf("name%d", i),
			Age: rand.Intn(100),
			Sex: "Man",
		}

		persons = append(persons, p)
	}


	data, err := msgpack.Marshal(persons)
	if err != nil {
		fmt.Printf("marshal failed ,err:%v", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
	return
}

func readJson(filename string) (err error)  {
	var person []*Person
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = msgpack.Unmarshal(data, &person)
	if err != nil {
		return
	}
	for _, v := range person {
		fmt.Printf("%#v\n", v)
	}
	return
}

func main()  {
	filename := "C:/tmp/person.data"
	err := writeJson(filename)
	if err != nil {
		fmt.Printf("write json faild ,err:%v\n", err)
	}

	err = readJson(filename)
	if err != nil {
		fmt.Printf("read json failed, err:%v\n", err)
		return
	}
}
