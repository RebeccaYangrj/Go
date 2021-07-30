package main

import (
	"encoding/json"
	"fmt"
)

//建立一个结构体
type Movie struct {
	Name  string   `json: "title"`
	Year  int      `json: "time"`
	Actor []string `json: "actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, []string{"周星驰", "张柏芝"}}
	//编码过程 结构体 --> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json error", err)
		return
	}
	fmt.Println("jsonStr = ", string(jsonStr))
	//解码过程 json --> 结构体
	my_movie := Movie{}
	err1 := json.Unmarshal(jsonStr, &my_movie)
	if err1 != nil {
		fmt.Println("Error", err1)
		return
	}
	fmt.Printf("%v\n", my_movie)
}
