package main

import (
	"fmt"

	"github.com/Josh7GAS/gojeff/tree/master/packages/prefixo"
)

var NomeDoUsuario = "Joshua"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
}
