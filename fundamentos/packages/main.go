package main

import (
	"fmt"

	Prefixo "github.com/Josh7GAS/gojeff/fundamentos/packages/prefixo"
)

var NomeDoUsuario = "Joshua"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", Prefixo.Capital)
}
