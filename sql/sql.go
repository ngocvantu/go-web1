package sql

const (
	CREATE_TERMS_TABLE = "CREATE TABLE IF NOT EXISTS TERMS (ID int NOT NULL AUTO_INCREMENT PRIMARY KEY, KEY_J varchar(20),KEY_E varchar(20),VAL_E varchar(50),VAL_J varchar(50),VAL_V varchar(50),CATEGORY int)"
	GET_ALL_TERMS      = "SELECT * from TERMS"
	ADD_TERMS          = "INSERT INTO TERMS ( KEY_J, KEY_E, VAL_E, VAL_J, VAL_V ) VALUES ( ?, ?, ?, ?, ? )"
)
