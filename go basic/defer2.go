package main

import "fmt"

func firstdefer() int {
	fmt.Println("Defer func call...")
	return 0

}

func firstreturn() int {
	fmt.Println("Return func call...")
	return 0

}
func returnanddefer() int {
	defer firstdefer()
	return firstreturn()
}

func main() {
	returnanddefer()
}
