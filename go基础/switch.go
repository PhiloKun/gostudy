package main

import "fmt"

func main() {
	var age int = 30

	switch {
	case age <= 30:
		fmt.Println(30)
	case age > 30:
		fmt.Println(31)
		fallthrough
	default:
		fmt.Println("default")
	}
}
