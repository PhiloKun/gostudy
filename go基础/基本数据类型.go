package main

import "fmt"

func main() {
	//var u8 uint8 = 255
	var a byte = 'a'
	fmt.Printf("%c %d\n", a, a)
	var a1 uint8 = 97
	fmt.Printf("%c %d\n", a1, a1)

	var z rune = '中'
	fmt.Printf("%c %d\n", z, z)

	//零值
	var s string
	var bo bool
	var b int
	fmt.Printf("%v %v %v\n", s, bo, b)
}
