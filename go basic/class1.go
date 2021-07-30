package main

// OOD的继承
import (
	"fmt"
)

//父类
type Human struct {
	Name string
	Age  int
}

//父类的行为
func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (this *Human) Run() {
	fmt.Println("Human.Run()...")
}

//子类
type Superman struct {
	Human //从父类进行继承
	level int
}

//修改父类的行为
func (this *Superman) Eat() {
	fmt.Println("Superman.Eat()...")
}

//添加子类的行为
func (this *Superman) Fly() {
	fmt.Println("Superman.Fly()...")
}

//子类打印
func (this *Superman) Print() {
	fmt.Println("Name: ", this.Name)
	fmt.Println("Age: ", this.Age)
	fmt.Println("Level: ", this.level)
}

func main() {
	h := Human{"Zhang3", 11}
	h.Eat()
	h.Run()

	//s := Superman{Human{"Li4", 12}, 10}
	var s Superman
	s.Name = "Zhao5"
	s.Age = 14
	s.level = 100
	s.Print()
	s.Eat()
	s.Run()
	s.Fly()

}
