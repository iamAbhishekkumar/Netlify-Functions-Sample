package main

import (
	"net/http"

	"github.com/iamAbhishekkumar/echoscript-lamda/functions/hello-world/handler"
)

func main() {
	http.HandleFunc("/", handler.HelloHandler)
	http.ListenAndServe(":9999", nil)
}
