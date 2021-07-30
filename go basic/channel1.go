package main

import (
	"fmt"
	"time"
)

func main() {
	// build a channel
	c := make(chan int, 3) //带有缓冲的channel
	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))

	// child goroutine to store elements in channel
	// 没有go就会阻塞，用go才是多线程进行了
	go func() {
		defer fmt.Println("Child goroutine has finished")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("input number i: ", i, "cap(c)", cap(c), "len(c)", len(c))
		}
	}()
	time.Sleep(2 * time.Second)
	// main gorotine to get elements in channel
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("output number num: ", num)
	}
	fmt.Println("Main goroutine has finished")

}
