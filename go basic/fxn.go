package main

import "fmt"

func fool1(a string, b int) int {
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)
	c := 100
	return c

}

// return multiple return value but anonymity
func fool2(a string, b int) (int, int) {
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)

	return 66, 77

}

// return multiple return value with name
func fool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----fool3----")
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)
	fmt.Println("r1= ", r1)
	fmt.Println("r2= ", r2) //r1 r2 initial default number is 0
	// r1 r2 的作用空间是整个fool3函数体的内部空间{}
	r1 = 2000
	r2 = 3000
	return r1, r2
}

// combine the return type
//func fool3(a string, b int) (r1, r2 int) {}

func main() {
	c := fool1("abc", 123)
	fmt.Println("c= ", c)
	e, f := fool2("hhh", 888)
	fmt.Println("e= ", e, "f= ", f)
	g, h := fool3("xxx", 333)
	fmt.Println("g= ", g, "h= ", h)
}
