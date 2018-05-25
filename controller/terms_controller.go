package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"go-web1/dbmng"
	"go-web1/sql"
	sql2 "database/sql"
	"time"
)

type TermsObj struct {
	Id int
	Key_J sql2.NullString
	Key_E sql2.NullString
	VAL_E sql2.NullString
	VAL_J sql2.NullString
	VAL_V sql2.NullString
	CATEGORY int
}


var db = dbmng.DB_MANAGER
var t,er = template.ParseFiles("web-content/pages/terms.html")

func Terms(w http.ResponseWriter, r *http.Request){
	start := time.Now()
	//createTermsTable()
	getTerms()
	if er != nil {
		fmt.Println(er.Error())
	}
	t.Execute(w,getTerms())
	fmt.Println("term take:", time.Since(start))
}

func createTermsTable(){
	db.Exec(sql.CREATE_TERMS_TABLE)
}

func getTerms() []TermsObj{
	result := []TermsObj{}
	listTerms, err := db.Query(sql.GET_ALL_TERMS)
	defer listTerms.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	term := &TermsObj{}
	for listTerms.Next()  {
		listTerms.Scan(&term.Id, &term.Key_J, &term.Key_E, &term.VAL_E, &term.VAL_J, &term.VAL_V, &term.CATEGORY)
		result = append(result, *term)
	}

	return result
}
