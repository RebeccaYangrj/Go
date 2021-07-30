package main

import (
	"fmt"
)

func swap(a int, b int) {
	var tmp int
	tmp = a
	a = b
	b = tmp
}

// a,b传入以后再func swap里面换了值了，但是在main函数里面没有huan值
func swap2(pa *int, pb *int) {
	var tmp int
	tmp = *pa
	*pa = *pb
	*pb = tmp
}

// a,b传入的是main函数中a,b的地址 swap2里面换了main函数中a,b的值，所以打印出来也是main函数里的值换了位置

func main() {
	var a int = 10
	var b int = 20
	swap(a, b)
	fmt.Println("a = ", a, "b = ", b)
	swap2(&a, &b)
	fmt.Println("a = ", a, "b = ", b)
}
