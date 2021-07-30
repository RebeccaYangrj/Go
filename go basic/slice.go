package main

import (
	"fmt"
)

func changeValue(myArray []int) {
	for _, value := range myArray {
		fmt.Println("Value = ", value)
	}
	//也可以通过slice的地址传递来改变value
	myArray[0] = 100
}

// 动态数组传递的是地址不是指针
func main() {
	myArray := []int{1, 2, 3, 4} //动态数组，slice

	changeValue(myArray)

	fmt.Println("------------")
	for _, value := range myArray {
		fmt.Println("Value = ", value)
	}

}
