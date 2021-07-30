package main

import "fmt"

func printArray(myArray [4]int) {
	for idx, val := range myArray {
		fmt.Println("index = ", idx, "value = ", val)
	}

}

func main() {
	//固定长度数组
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < 10; i++ {
		fmt.Println(myArray1[i])
	}
	for idx, val := range myArray2 {
		fmt.Println("index = ", idx, "value = ", val)
	}

	//查看数组的类型
	fmt.Printf("The type of myArray1: %T\n", myArray1)
	fmt.Printf("The type of myArray2: %T\n", myArray2)
	fmt.Printf("The type of myArray2: %T\n", myArray3)

}
