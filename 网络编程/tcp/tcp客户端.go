package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var bytedata = make([]byte, 1024)
		n, err := conn.Read(bytedata)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(bytedata[:n]))
	}

}
