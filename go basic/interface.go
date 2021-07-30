package main

import (
	"fmt"
)

// interface是一种万能类型不管传入的数据是什么类型的都可以实现
func myFunc(arg interface{}) {
	//??????fmt.Fprintln("myFunc() is call....")
	//./interface.go:9:15: cannot use "myFunc() is call...." (type string) as type io.Writer in argument to fmt.Fprintln:
	//string does not implement io.Writer (missing Write method)
	fmt.Println(arg)
	//区分引用的数据类型
	//interface有一种断言的机制
	//必须要对空接口来判断
	value, ok := arg.(string)
	if !ok {
		fmt.Println("is not string\n")
	} else {
		fmt.Println("is string with value\n", value)
		fmt.Printf("Value type = %T\n", value)
	}
}

type Book struct {
	name string
}

func main() {
	book := Book{"Golang"}
	myFunc(book)
	myFunc("Hello go")
	myFunc(123)
	myFunc(123.4444)
}
