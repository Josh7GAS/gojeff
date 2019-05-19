package main

import (
	"database/sql"
	"fmt"

	"github.com/Josh7GAS/gojeff/bdconnect/cfg"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	credencial := cfg.Credencial
	host := credencial[0]
	port := credencial[1]
	user := credencial[2]
	password := credencial[3]
	var (
		keys, values string
	)

	db, _ := sql.Open("mysql", user+":"+password+"@tcp"+"("+host+":"+port+")/testedb")

	mostraGlobalStatus, _ := db.Query("SHOW GLOBAL STATUS;")

	for mostraGlobalStatus.Next() {
		mostraGlobalStatus.Scan(&keys, &values)
		fmt.Println(keys, values)
	}

}
