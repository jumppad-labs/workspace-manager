package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		panic(err)
	}
}
