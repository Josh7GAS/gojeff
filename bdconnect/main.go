package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

func main() {
	//credencial := cfgArquivo() Array vinda da função
	credencial := cfgArquivo()

	//variaveis que alimentam as index da array credencial
	// host := credencial[0]
	// port := credencial[1]
	user := credencial[2]
	password := credencial[3]

	//conctando o banco de dados
	db, _ := sql.Open("mysql", user+":"+password+"@/testedb")
	if db != nil {
		//caso bem sucedido ira fazer o globalstatus e armazenar em um slice

		//trabalhe aqui com (foor range loop), (Next()), (Ponteiros)
		mostraGlobalStatus, _ := db.Query("SHOW GLOBAL STATUS;")
		if mostraGlobalStatus != nil {
			slice1 := make([]*sql.Rows, 20)
			slice1 = append(slice1, mostraGlobalStatus)
			for _, v := range slice1 {
				fmt.Println(v)
			}
		}

	}

	defer db.Close()

}

//função para avbrir o arquivo .my.cnf e procurar as informações das variaveis, elas são armazenadas na array credencial []string and
//e são usadad mais tarde quando essa função devolver a array
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
