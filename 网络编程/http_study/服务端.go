package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path, r.UserAgent())
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("服务启动在: http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
