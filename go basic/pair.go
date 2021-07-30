package main

import (
	"fmt"
)

//接口1 read
type Reader interface {
	Readbook()
}

//接口2 write
type Writer interface {
	Writerbook()
}

//class book类
type Book struct {
}

//book类readfunc
func (this *Book) Readbook() {
	fmt.Println("Reading a book ...")
}

//book类writefunc
func (this *Book) Writerbook() {
	fmt.Println("writing a book ...")
}

func main() {
	// b pair: type: Book, value book address
	b := &Book{}

	// r pair: type and value not determined
	var r Reader
	// r pair: type: Book, value book address
	r = b
	r.Readbook()

	// w pair: type and value not determined
	var w Writer
	// w pair: type: Book, value book address
	w = r.(Writer) //此处断言成功是因为：w和r是一个type的
	w.Writerbook()
}
