package repo

import (
	_ "github.com/go-sql-driver/mysql"
	/*
		gitub.com/go-sql-server/mysql não é usado diretamente pela aplicação*/
	"github.com/jmoiron/sqlx"
)

//Db armazena a conexão com o banco de dados
var Db *sqlx.DB

//Função que abre a conexão com o banco mysql
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root@tcp(40.76.49.248)/mundo-go/src/github.com/Josh7GAS/gojeff/avancado/banco_sql")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
