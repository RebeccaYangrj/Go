package main

import (
	"fmt"
)

//两个goroutine之间用channel来返回值
//携程之间的消息消息异步

//无缓冲的chanel会造成阻塞
func main() {
	//创建一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("Sub goroutine ended")
		fmt.Println("Sub goroutine is running")
		//将值传入到channel里面
		c <- 666
	}()
	//将channel里面的值传入到nums里面去

	nums := <-c
	fmt.Println(nums)
	fmt.Println("Main goroutine ended")

}
