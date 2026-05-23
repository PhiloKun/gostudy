package main

import (
	"fmt"
	"gostudy/version"
)

func hello() {
	fmt.Println(quanju)
}

// 全局变量
var quanju string = "全局变量"

// 常量
const Version string = "1.0.0"

func main() {

	fmt.Println(version.Version)
	//先声明
	var name string
	//在赋值
	name = "philokun"
	fmt.Println(name)

	//声明并赋值
	var name1 string = "philokun1"
	fmt.Println(name1)

	//省略类型
	var name2 = "philokun2"
	fmt.Println(name2)

	//声明并赋值,短声明符号象牙
	name4 := "philokun3"
	fmt.Println(name4)

	hello()

	//定义多个变量
	var a1, a2, a3 = 1, 2, 3
	fmt.Println(a1, a2, a3)

	var (
		str1 = "1"
		str2 = "2"
	)
	fmt.Println(str1, str2)
}
