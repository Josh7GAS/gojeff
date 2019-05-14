package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

func main() {

	db, err := sql.Open("mysql", "root:senha@/testedb")
	if err != nil {
		fmt.Println("Deu Erro",
			err.Error())
		return

	} else {
		fmt.Println("Conxão Estabelecida!!!!")
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Não conseguiu pingar", err.Error())
	}

	arquivo, err := ini.Load(os.Getenv("HOME") + "/.my.cnf")
	if err != nil {
		fmt.Printf("Não foi possível ler o aquivo: %v", err)
		os.Exit(1)
	} else {

		// lendoArquivo, err := arquivo.Section("client").Key("hosts").Validate(func(in string) string {
		// 	if len(in) == 0 {
		// 		return "default"
		// 	}
		// 	return in
		// })

		fmt.Println("host:", arquivo.Section("client").Key("host").String())
		fmt.Println("user:", arquivo.Section("client").Key("port").String())
		fmt.Println("password:", arquivo.Section("client").Key("password").String())

		var []string = arquivo.Section("client").Key("host")
	}
}

// [client]
// host=127.0.0.1
// port=3306
// user=root
// password=senha
