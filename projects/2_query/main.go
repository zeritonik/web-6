package main

import (
	"fmt"
	"net/http"
)

type Response struct {
	Result string `json:"result"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "No name passed", http.StatusBadRequest)
		return
	}
	w.Write([]byte("Hello, " + name + "!"))
}

func main() {
	http.HandleFunc("/api/user", handler)
	fmt.Println("Server is listening on 127.0.0.1:8080")
	error := http.ListenAndServe("127.0.0.1:8080", nil)
	if error != nil {
		fmt.Println(error)
	}
}
