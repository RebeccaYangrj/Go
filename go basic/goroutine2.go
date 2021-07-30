package main

import (
	"fmt"
	"time"
)

func main() {
	// 用匿名函数来运行
	//用go承载一个匿名函数
	/* 无参数的匿名函数
	go func() {
		defer fmt.Println("A defer")

		//提前加return可以提前退出
		//return
		func() {
			defer fmt.Println("B defer")
			//提前加return只可以提前退出当前子函数
			//return
			//退出go函数
			runtime.Goexit()
			fmt.Println("B")
		}() //结尾加上（）定义了一个子函数然后进行调用

		fmt.Println("A")
	}() */
	//有参数的匿名函数

	//此处不能 returnvalue := 因为go是一个并行操作
	//需要用channel
	go func(a int, b int) bool {
		fmt.Println(a, b)
		return true
	}(10, 20)
	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
