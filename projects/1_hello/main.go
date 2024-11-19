package main

import (
	"fmt"
	"net/http"
)

type Response struct {
	Result string `json:"result"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, web!"))
}

func main() {
	http.HandleFunc("/get", handler)
	fmt.Println("Server is listening on 127.0.0.1:8080")
	error := http.ListenAndServe("127.0.0.1:8080", nil)
	if error != nil {
		fmt.Println(error)
	}
}
