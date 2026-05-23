package main

import "fmt"

func main() {
	fmt.Println("请输入你的名字：")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)

	var n, error = fmt.Scan(&name)
	fmt.Println(n, error, name)

}
