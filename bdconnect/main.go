package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

func main() {

	credencial := cfgArquivo()
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

func cfgArquivo() []string {
	var (
		credencial []string
		host       string
		port       string
		user       string
		password   string
	)
	arquivo, _ := ini.Load(os.Getenv("HOME") + "/.my.cnf")
	if arquivo != nil {

		host = arquivo.Section("client").Key("host").Validate(func(in string) string {
			if len(in) == 0 {
				return "default"
			}

			return in
		})
		credencial = append(credencial, host)

		port = arquivo.Section("client").Key("port").Validate(func(in string) string {
			if len(in) == 0 {
				return "default"
			}
			return in
		})
		credencial = append(credencial, port)

		user = arquivo.Section("client").Key("user").Validate(func(in string) string {
			if len(in) == 0 {
				return "default"
			}
			return in
		})
		credencial = append(credencial, user)

		password = arquivo.Section("client").Key("password").Validate(func(in string) string {
			if len(in) == 0 {
				return "default"
			}
			return in
		})
		credencial = append(credencial, password)

	}

	return credencial
}
