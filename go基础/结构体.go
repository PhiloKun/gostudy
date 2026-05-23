package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Study() {
	fmt.Printf("%s study", u.Name)
}
func main() {
	fmt.Println(User{})
	s1 := User{Name: "1", Age: 20}
	s1.Study()
}
