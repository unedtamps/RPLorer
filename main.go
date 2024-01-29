package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Testhandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Testhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("testhelo world"))
}
