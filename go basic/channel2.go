package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 4)

	go func() {
		for i := 0; i < 4; i++ {
			c <- i
		}
		//如果不close那么会出现fatal error: all goroutines are asleep - deadlock!
		//因为写入数据的时候已经是写完了，当读完了c中所有的数据之后不会在有人往c里写入数据了而c并没有关闭所以会堵塞在上面
		close(c)
	}()
	/* for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	} */
	//可以使用range来迭代不断操作的channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main finished")

}
