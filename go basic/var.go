package main

import (
	"fmt"
)

func main() {
	//method 1, default number is 0
	var a int
	//method 2 def a number and give a number
	var b int = 100
	var c = 200

	fmt.Println(a, b, c)
	fmt.Printf("type of a: %T, b: %T, c: %T", a, b, c)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb is %T", bb, bb)
	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc is %T", cc, cc)

	// method4 only inside the fxn cannot make a global variable
	e := 100
	fmt.Printf("e = %T, type of e is %T", e, e)

	// method5 indicate multiple varibles
	var xx, yy, zz int = 100, 200, 300
	var kk, ll = 200, "string"
	fmt.Println(xx, yy, zz, kk, ll)

	// constant number
	const ccc int = 1
}
