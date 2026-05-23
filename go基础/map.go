package main

import "fmt"

func main() {
	var userMap map[int]string = map[int]string{1: "A", 2: "B", 3: "C"}
	fmt.Println(userMap)

	fmt.Println(userMap[1])
	value, ok := userMap[2]
	fmt.Println(value, ok)

	userMap[1] = "Hello"
	fmt.Println(userMap)
	delete(userMap, 2)
	fmt.Println(userMap)
}
