package main

import (
	"net/http"
	"go-web1/controller"
	"database/sql"
	"fmt"

)

var DB, errglobal = sql.Open("mysql", "root:@tcp(localhost:3306)/webapp")

func main() {
	if errglobal != nil {
		fmt.Println(errglobal.Error())
	}
	registerHandler()
	http.ListenAndServe(":8070", nil)
}

func registerHandler() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/terms/", controller.Terms)

	http.Handle("/web-content/", http.StripPrefix("/web-content/", http.FileServer(http.Dir("web-content"))))
}
