package main

import (
	"fmt"
	"sort"
)

func main() {
	var nameList []string
	nameList = append(nameList, "张三")
	nameList = append(nameList, "李四")
	fmt.Println(nameList)

	nameList1 := make([]string, 0)
	fmt.Println(nameList1)

	var ints = []int{1, 4, 3}
	fmt.Println(ints)
	sort.Ints(ints) //升序
	fmt.Println(ints)
}
