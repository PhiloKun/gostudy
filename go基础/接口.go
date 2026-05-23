package main

import "fmt"

type SingInter interface {
	Sing()
}

type Cat struct {
	Name string
}

func (c Cat) Sing() {
	fmt.Println(c.Name + "在唱歌")
}

type EmptyInter interface { //空接口，任何类型都实现了空接口
}

func Print(val EmptyInter) {
	fmt.Println(val)
}

func Print1(val interface{}) { //空接口方式2    或者any =interface{}
	fmt.Println(val)
}
func Print2(val any) { //空接口方式2    或者any =interface{}
	fmt.Println(val)
}
func main() {

	c := Cat{
		Name: "Mimi",
	}
	c.Sing()
	Print(c)
	Print1(c)
	Print2(c)

}
