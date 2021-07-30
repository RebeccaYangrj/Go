package main

import (
	"fmt"
	"reflect"
)

//先建立一个user的type，以及他对应的一些action
type User struct {
	Name string
	ID   int
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called ..")
	fmt.Printf("%v\n", this)
}

func main() {
	user1 := User{"Rebecca", 1, 26}

	DoFieldandMethod(user1)

}

// fxn 通过reflect来写出来如何调用user的field和method

func DoFieldandMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)

	//获取input的value
	inputValue := reflect.ValueOf(input)

	//通过type获取字段：
	//1. 获取interface的reflecttype，通过type.Numfield进行遍历， field就是一共有几个properties
	//2. 得到每个field的数据类型
	//3. 通过field里的一个interface的方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v = %v\n", field.Name, field.Type, value)
	}
	//通过type 获得里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)

	}
}
