package dbmng

import ("database/sql"
	_ "github.com/go-sql-driver/mysql")

var DB_MANAGER, errglobal = sql.Open("mysql", "root:@tcp(localhost:3306)/webapp")
