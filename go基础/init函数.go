package main

import "fmt"

// 谁先在前面先加载谁，可以多个init 函数
func init() {
	fmt.Println(1)
}
func init() {
	fmt.Println(2)
}
func init() {
	fmt.Println(3)
}

func main() {

}
