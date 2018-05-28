package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var t2, er2 = template.ParseFiles("web-content/pages/index.html")

func Index(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if er2 != nil {
		fmt.Println(er2.Error())
	}
	t2.Execute(w, nil)
	fmt.Println("index take:", time.Since(start))
}
