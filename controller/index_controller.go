package controller

import (
	"net/http"
	"html/template"
	"fmt"
)



func Index(w http.ResponseWriter, r *http.Request){
	var t,er = template.ParseFiles("web-content/pages/index.html")
	if er != nil {
		fmt.Println(er.Error())
	}
	t.Execute(w,nil)
}
