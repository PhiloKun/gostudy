package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"` //忽略空值
	Sex  int    `json:"-"`             //不转化
}

func main() {
	stu := Student{Name: "zhangsan"}
	fmt.Println(stu.Name)
	jsondata, _ := json.Marshal(stu)
	fmt.Println(string(jsondata))
}
