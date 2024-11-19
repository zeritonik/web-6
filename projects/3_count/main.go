package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Body struct {
	Count int `json:"count"`
}

var count int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var body Body
		e := json.NewDecoder(r.Body).Decode(&body)
		if e != nil {
			http.Error(w, "Это не число!!", http.StatusBadRequest)
			return
		}
		count += body.Count
		return
	} else if r.Method == "GET" {
		w.Write([]byte(fmt.Sprintf("Count: %d", count)))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/count", handler)
	fmt.Println("Server is listening on 127.0.0.1:3333")
	error := http.ListenAndServe("127.0.0.1:3333", nil)
	if error != nil {
		fmt.Println(error)
	}
}
