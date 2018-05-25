package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

var t,er = template.ParseFiles("web-content/pages/terms.html")

func Terms(w http.ResponseWriter, r *http.Request){

	if er != nil {
		fmt.Println(er.Error())
	}
	t.Execute(w,nil)
}
