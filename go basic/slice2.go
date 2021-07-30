package main

import (
	"fmt"
)

// 4种slice的声明方式，如果不赋值，初始值是0
func main() {
	slice1 := []int{1, 2, 3, 4}

	var slice2 []int
	slice2 = make([]int, 3)

	var slice3 []int = make([]int, 5)

	slice4 := make([]int, 6)

	fmt.Printf("len of slice1 =%d , slice = %v\n ", len(slice1), slice1)
	fmt.Printf("len of slice2 =%d , slice = %v\n ", len(slice2), slice2)
	fmt.Printf("len of slice3 =%d , slice = %v\n ", len(slice3), slice3)
	fmt.Printf("len of slice4 =%d , slice = %v\n ", len(slice4), slice4)

	if slice2 == nil {
		fmt.Println("slice2 is an empty slice")
	} else {
		fmt.Println("slice2 has space")
	}

}
