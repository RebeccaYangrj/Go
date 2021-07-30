package main

import (
	"fmt"
	"reflect"
)

// function 判断类型的函数
func TypeofNum(arg interface{}) {
	fmt.Println("Type = ", reflect.TypeOf(arg))
	fmt.Println("Value = ", reflect.ValueOf(arg))
}

func main() {
	a := 1.2345
	TypeofNum(a)
}
