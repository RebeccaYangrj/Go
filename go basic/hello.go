package main

//import "fmt"
import (
	"fmt"
	"time"
)

func main() { //{ must be in the same line as the fxn name
	fmt.Println("Hello Go")
	time.Sleep(1 * time.Second)
	fmt.Println("Bye Go")
}
