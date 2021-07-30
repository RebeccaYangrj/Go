package main

import "fmt"

// constant use for enumeration
const (
	// inside const() we can use iota
	// iota can only used in const()
	BEIJING   = 10 * iota // iota = 0
	SHANGHAI  = 10 * iota // iota = 1
	GUANGZHOU = 10 * iota // iota  = 2
)
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func main() {
	//constant number cann't be edit
	const length = 100
	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println(a, b, c, d, e, f)
}
