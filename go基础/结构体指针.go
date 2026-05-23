package main

import "fmt"

type UserInfo struct {
	Name string `json:"name"`
}

func (u UserInfo) SetName(name string) {
	u.Name = name
}

func (u *UserInfo) SetName1(name string) {
	u.Name = name

}
func main() {
	//值传递
	user := UserInfo{Name: "zhangsan"}
	user.SetName("lisi")
	fmt.Println(user.Name)

	//引用传递
	user1 := UserInfo{Name: "zhangsan"}
	user1.SetName1("lisi")
	fmt.Println(user1.Name)
}
