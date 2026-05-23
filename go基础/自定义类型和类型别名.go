package main

import "fmt"

type MyCode int        //自定义类型
type MyAliasCode = int //类型别名

func (m MyCode) Code() {

}

const myCode MyCode = 1
const myAliasCode MyAliasCode = 1

func main() {

	fmt.Printf("%v,%T\n", myCode, myCode)
	fmt.Printf("%v,%T\n", myAliasCode, myAliasCode)
}
