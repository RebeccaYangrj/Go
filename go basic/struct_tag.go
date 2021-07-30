package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `info:"Name" doc:"MyName"` //Tag结构体的标签，keyvalue随便写也可以写多个，中间用空格隔开，不是逗号
	Sex  string `info:"Sex"`               //在其他packeage调用我们这个属性的时候来判断说明
}

func FindTag(input interface{}) {

	t := reflect.TypeOf(input).Elem() //得到当前结构体全部的元素
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		fmt.Println("Tag: ", taginfo)
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("Tag: ", tagdoc)
	}

}

func main() {
	//可以用reflect来解析结构体里的tag
	var rm Resume
	FindTag(&rm)
	/*传入的是地址才能进行解析，不然会出现错误
	panic: reflect: Elem of invalid type main.Resume
	goroutine 1 [running]:
	reflect.(*rtype).Elem(0x10b7580, 0xc000092e30, 0x10152d7)
	/usr/local/go/src/reflect/type.go:915 +0x1a5
	*/
}
