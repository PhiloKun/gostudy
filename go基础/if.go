package main

import "fmt"

func main() {
	var age int = 30

	if age < 18 {
		fmt.Println(age)
	} else {
		fmt.Println("大于30")
	}
}
