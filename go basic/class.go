package main

import (
	"fmt"
)

type Hero struct {
	Name string
	Size int
	Id   int
}

/* 传递的是值，可以读取但是无法写入
func (this Hero) Show() {
	fmt.Println("Name: ", this.Name)
	fmt.Println("Size: ", this.Size)
	fmt.Println("ID: ", this.Id)

}
func (this Hero) GetName() string {
	return this.Name

}
func (this Hero) ChangeName(newName string) {
	this.Name = newName
*/

//给struct绑定方法一点要用*
//类名，属性名，方法名首字母大写就是外界可以访问
func (this *Hero) Show() {
	fmt.Println("Name: ", this.Name)
	fmt.Println("Size: ", this.Size)
	fmt.Println("ID: ", this.Id)

}
func (this *Hero) GetName() string {
	return this.Name

}
func (this *Hero) ChangeName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "Zhang3", Size: 160, Id: 1}
	hero.Show()
	hero.GetName()
	hero.ChangeName("Li4")
	hero.Show()
}
