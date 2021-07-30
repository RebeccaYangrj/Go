package main

import "fmt"

func changeValue1(p int) {
	p = 10
}

//如果是形参的话就是没法直接传入，他传入changeValue里面p = a =1
//在changeValue里面p又重新被赋值为10，但是a的地址存储的东西并没有发生改变还是1，所以输出的是1
func changeValue2(p *int) {
	*p = 10
}

//如果是*int表示的是 p是一种指针类型，传入的&b是b的地址，让p的内存里面存储的是b变量的地址
//在func cV2中*p =10代表的是将p储存地址位置的数值改变为10
func main() {
	var a int = 1
	changeValue1(a)
	fmt.Println("a= ", a)
	var b int = 1
	changeValue2(&b)
	fmt.Println("b= ", b)
}
