package main

import (
	"fmt"
	"net/http"

	"github.com/Josh7GAS/gojeff/avancado/MongoDB/manipulador"
	"github.com/Josh7GAS/gojeff/avancado/servidorweb/manipulador"
	"github.com/Josh7GAS/gojeff/repo"
)

func init() {
	fmt.Println("vamos começar a subir o servidor...")
}

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados:",
			err.Error())
		return
	}

	err = repo.AbreSessaoComMongo()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados:",
			err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Estou escutando comandante....")
	http.ListenAndServe(":8181", nil)
}
