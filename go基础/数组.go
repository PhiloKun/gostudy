package main

import "fmt"

func main() {
	var nameList [3]string = [3]string{"zhangsan", "lisi", "wangwu"}
	fmt.Println(nameList)
	fmt.Println(nameList[0], nameList[1], nameList[2])
	//fmt.Println(nameList[-1]) go语言不支持

	fmt.Println(nameList[len(nameList)-1])
	nameList[0] = "didi"
	fmt.Println(nameList)
}
