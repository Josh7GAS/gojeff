package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWrite, r *http.Request) {
		fmt.Fprintln(w, "Olá mundo")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", mainpulador.Ola)
	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServer(":8181", nil)
}
