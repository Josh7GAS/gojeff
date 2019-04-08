package main

import "fmt"

var (
	endereco string
	telefone  string
	quantidade int
	comprou bool
	valor float64
	palavras rune
)

func main(){
	teste := "Valor de Teste"
	fmt.Printf("endereco: %s\r\n", endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("quantidade: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("teste: %v\r\n ", teste)
}
