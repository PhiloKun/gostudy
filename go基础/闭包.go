package main

import "fmt"

func nameList(a int) func(...int) {
	return func(a ...int) {
		fmt.Println(a)
	}
}

func main() {
	nameList(1)(1)
}
