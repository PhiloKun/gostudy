package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	//for {
	//	time.Sleep(1 * time.Second)
	//}
	var list = []string{"1", "2"}

	for index, item := range list {
		fmt.Println(index, item)
	}
}
