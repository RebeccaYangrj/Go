package main

import (
	"fmt"
)

//声明一个数据，或者声明一个结构体
type myint int //其实是int的一个别名
//定义一个class 结构体里面是可以声明其他变量的
type Book struct {
	title  string
	author string
}

func printBook(book Book) {
	fmt.Printf("%v\n", book)
	book.author = "zhao5"
}
func changeBook(book *Book) {
	book.author = "li4"

}

func main() {
	var a myint = 10
	fmt.Println("myint = ", a)
	fmt.Printf("myint type = %T\n", a)
	fmt.Println("------------------\n")

	var book1 Book
	book1.title = "golang"
	book1.author = "zhang3"
	printBook(book1)
	fmt.Println("------------------\n")

	changeBook(&book1)
	printBook(book1)

}
