package main

import (
	"fmt"
)

// 用select实现fibonacii
func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x: //如果c是可以写的话，那么case就会进来
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

func main() {
	// make two channel
	c := make(chan int)
	quit := make(chan int)

	// sub go
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		//如果quit不写的话那么就会阻塞在casec那个地方因为没有东西继续写入了所以无法到达return那个地方
		quit <- 0

	}()

	// main go
	fibonacii(c, quit)

}
