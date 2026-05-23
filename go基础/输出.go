package main

import "fmt"

func main() {
	fmt.Println("helloworld", "你好世界")
	//字符
	fmt.Printf("你好 %s", "世界")
	//整数
	fmt.Printf("%d\n", 123)
	//浮点数
	fmt.Printf("%.2f\n", 123.0415)

	//类型
	fmt.Printf("%T  %T", "123", 123)
	/*
	   %v 表示以变量的默认格式输出。
	   对于基本类型（如整数、浮点数等），直接输出其值。
	   对于结构体，输出字段值。
	*/
	fmt.Printf("%v  %v\n", 123, "string")
	fmt.Printf("%#v\n", "")
	fmt.Printf("%v\n", "")
	//格式化内容赋给一个变量
	var fmt1 = fmt.Sprintf("%.2f\n", 123.0415)
	fmt.Println(fmt1)

}
