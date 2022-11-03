package main

import (
	"io"
	"log"
	"net/http"
)

func postTaskHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "create task\n")
}

func main() {
	http.HandleFunc("/task", postTaskHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}