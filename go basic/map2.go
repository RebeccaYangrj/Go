package main

import (
	"fmt"
)

func printMap(mymap map[string]string) {
	for key, value := range mymap {
		fmt.Println("key is ", key)
		fmt.Println("value is ", value)
	}
}

// map和slice一样是可以自己扩充空间的一直存储，不过map是一种字典型，有key和value
func main() {
	var myMap1 map[string]string
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "C++"
	myMap1["two"] = "python"
	myMap1["three"] = "java"
	fmt.Println(myMap1) //string打印根据字典顺序

	//遍历
	for key, value := range myMap1 {
		fmt.Println("key is ", key)
		fmt.Println("value is ", value)
	}
	fmt.Println("-----------")
	//修改
	myMap1["one"] = "php"
	//删除
	delete(myMap1, "three")
	//传参
	printMap(myMap1)

}
