package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	listen, err1 := net.ListenTCP("tcp", addr)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("tcp listen:", addr.String())
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println(conn.RemoteAddr())
		conn.Write([]byte("HELLOWORLD"))
		time.Sleep(time.Second * 1)
		conn.Close()

	}

}
