package main

import "fmt"

func sayHello() {
	fmt.Println("hello world")
}

func params1(id string) {
	fmt.Println(id)
}
func params2(id, name string) {
	fmt.Println(id, name)
}
func params3(nameList ...string) {
	for _, name := range nameList {
		fmt.Println(name)
	}
}

func duo() bool {
	return true
}

func duo1() (string, int) {
	return "1", 2
}

// 匿名函数
var getName = func(name string) string {
	return name
}

func main() {
	sayHello()
	params1("123")
	params2("123", "123")
	params3("434343", "3232")
	duo()
	duo1()
	println(getName("1"))

}
