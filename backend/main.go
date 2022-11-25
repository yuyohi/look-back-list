package main

import (
	"io"
	"log"
	"net/http"
)

func postIssueHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "create issue\n")
}

func main() {
	http.HandleFunc("/issue", postIssueHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}