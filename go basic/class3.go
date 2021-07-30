package main

import (
	"fmt"
)

//一个指针接口
type AnimalIF interface {
	Sleep()
	GetColor() string
	Gettype() string
}

//具体的类
type Cat struct {
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}
func (this *Cat) GetColor() string {
	return this.Color
}
func (this *Cat) Gettype() string {
	return "Cat"

}

func main() {
	var animal1 AnimalIF //接口数据类型
	animal1 = &Cat{"Yellow"}
	animal1.Sleep()
	fmt.Println(animal1.GetColor())
	fmt.Println(animal1.Gettype())

}
