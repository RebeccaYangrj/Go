package main

import (
	"fmt"
	"time"
)

//子goroutine
func ChildGoRoutine() {
	i := 0
	for {
		i++
		fmt.Printf("Child %d", i)
		time.Sleep(1 * time.Second)
	}

}

//主goroutine
//子 主goroutine会并发运行
func main() {
	//创建一个go程去执行ChildGoRoutine()
	go ChildGoRoutine()
	//如果主程序的goroutine不进行的话那么子goroutine也会被杀死
	fmt.Println("main goroutine exit")

	/* 选中sft+opt+a全部注释
	i := 0
	for {
		i++
		fmt.Printf("Parent %d", i)
		time.Sleep(1 * time.Second)
	} */

}
