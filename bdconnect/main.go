package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Josh7GAS/gojeff/bdconnect/cfg"
	"github.com/Josh7GAS/gojeff/bdconnect/model"
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
	var (
		valores []string
	)
	valores[0] = keys
	valores[1] = values

	db, _ := sql.Open("mysql", user+":"+password+"@tcp"+"("+host+":"+port+")/testedb")

	mostraGlobalStatus, _ := db.Query("SHOW GLOBAL STATUS;")

	for mostraGlobalStatus.Next() {
		mostraGlobalStatus.Scan(&keys, &values)
		fmt.Println(keys, values)
	}

	GlobalStatusJSON, _ := os.Create("globalstatus.json")
	if GlobalStatusJSON != nil {
		escrevendoGlobalStatusJSON := bufio.NewWriter(GlobalStatusJSON)
		escrevendoGlobalStatusJSON.WriteString("[\r\n")
		for _, rows := range valores {
			for indiceValores, item := range valores {
				dados := strings.Split(item, "/")
				primeiraLinha := model.Cabecalio{}
				primeiraLinha.Keys = dados[0]
				primeiraLinha.Values = dados[1]
				fmt.Printf("Cabeçalio: %+v\r\n", primeiraLinha)

				cabecalioJSON, _ := json.Marshal(primeiraLinha)
				if cabecalioJSON != nil {
					escrevendoGlobalStatusJSON.WriteString(" " + string(cabecalioJSON))
					if (indiceValores + 1) < len(rows) {
						escrevendoGlobalStatusJSON.WriteString(",\r\n")

					}
					return
				}
			}

		}
		escrevendoGlobalStatusJSON.WriteString("\r\n]")
		escrevendoGlobalStatusJSON.Flush()
		GlobalStatusJSON.Close()
		db.Close()
	}

}
