package main

import (
	"net/http"
	"go-web1/controller"
)

func main() {
	registerHandler()

	http.ListenAndServe(":8070", nil)
}

func registerHandler() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/terms/", controller.Terms)

	http.Handle("/web-content/", http.StripPrefix("/web-content/", http.FileServer(http.Dir("web-content"))))
}
