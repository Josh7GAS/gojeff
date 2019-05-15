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

	fmt.Println(cfgArquivo)
	db, _ := sql.Open("mysql", cfgArquivo(user)+":"+cfgArquivo(password)+"@/testedb")
	if db == nil {
		fmt.Println("Conexão Estabelecida")
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Não conseguiu pingar", err.Error())
	} else {
		fmt.Println("Pong")
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
	if arquivo == nil {

		host = arquivo.Section("client").Key("host").Validate(func(in string) string {
			if len(in) == 0 {
				return user
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
		port = arquivo.Section("client").Key("port").Validate(func(in string) string {
			if len(in) == 0 {
				return "default"
			}
			return in
		})
		credencial = append(credencial, port)

	}
	return credencial
}

// db, err := sql.Open("mysql", user+":"+ password+"@/testedb")
// if err != nil {
// 	fmt.Println("Deu Erro",
// 		err.Error())
// 	return

// } else {
// 	fmt.Println("Conxão Estabelecida!!!!")
// }

// defer db.Close()

// err = db.Ping()
// if err != nil {
// 	fmt.Println("Não conseguiu pingar", err.Error())
// } else {
// 	fmt.Println("Pong")
// }
//fmt.Println(cfgArquivo(host, port, user, password))
