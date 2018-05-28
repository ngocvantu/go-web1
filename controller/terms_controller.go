package controller

import (
	sql2 "database/sql"
	"fmt"
	"go-web1/dbmng"
	"go-web1/sql"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type TermsObj struct {
	Id       int
	Key_J    sql2.NullString
	Key_E    sql2.NullString
	VAL_E    sql2.NullString
	VAL_J    sql2.NullString
	VAL_V    sql2.NullString
	CATEGORY int
}

var db = dbmng.DB_MANAGER
var t, er = template.ParseFiles("web-content/pages/terms.html")

func Terms(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		start := time.Now()
		r.ParseForm()
		key_jp := r.FormValue("key-jp")
		key_en := r.FormValue("key-en")
		val_en := r.FormValue("val-en")
		val_jp := r.FormValue("val-jp")
		val_vi := r.FormValue("val-vi")
		fmt.Println("key_en:", key_en)
		err := AddTerms(&key_jp, &key_en, &val_en, &val_jp, &val_vi)
		if err == nil {
			http.Redirect(w, r, "/terms/", http.StatusFound)
		} else {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("add term take:", time.Since(start))
	} else {
		start := time.Now()
		createTermsTable()
		//getTerms()
		if er != nil {
			fmt.Println(er.Error())
		}
		t.Execute(w, getTerms())
		fmt.Println("term take:", time.Since(start))
	}
}
func AddTerms(key_jp *string, key_en *string, val_en *string, val_jp *string, val_vi *string) error {
	if strings.Trim(*key_jp, " ") == "" && strings.Trim(*val_en, " ") == "" {
		return nil
	}
	_, err := db.Exec(sql.ADD_TERMS, *key_jp, *key_en, *val_en, *val_jp, *val_vi)
	return err
}

func createTermsTable() {
	db.Exec(sql.CREATE_TERMS_TABLE)
}

func getTerms() []TermsObj {
	result := []TermsObj{}
	listTerms, err := db.Query(sql.GET_ALL_TERMS)
	defer listTerms.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	term := &TermsObj{}
	for listTerms.Next() {
		listTerms.Scan(&term.Id, &term.Key_J, &term.Key_E, &term.VAL_E,
			&term.VAL_J, &term.VAL_V, &term.CATEGORY)
		result = append(result, *term)
	}

	return result
}
