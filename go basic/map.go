package main

import (
	"fmt"
)

// map和slice一样是可以自己扩充空间的一直存储，不过map是一种字典型，有key和value
func main() {
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is empty")
	}
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "C++"
	myMap1["two"] = "python"
	myMap1["three"] = "java"
	fmt.Println(myMap1) //string打印根据字典顺序

	myMap2 := make(map[int]string, 10)
	myMap2[1] = "C++"
	myMap2[2] = "python"
	myMap2[3] = "java"
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"one":   "C++",
		"two":   "python",
		"three": "java", //最后一行也要加，
	}
	fmt.Println(myMap3)
}
