package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path, r.UserAgent())
}

func main() {
	http.HandleFunc("/", Index)
	fmt.Println("running:     http://127.0.0.1:8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
