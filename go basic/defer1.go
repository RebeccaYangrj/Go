package main

import "fmt"

func main() {

	defer fmt.Println("main: end1")
	defer fmt.Println("main: end2")

	fmt.Println("main: hello go1")
	fmt.Println("main: hello go2")
}

// defer 是在运行玩所有main函数里面其他的内容之后最后再运行
// defer 是stack类型的函数，FILO
