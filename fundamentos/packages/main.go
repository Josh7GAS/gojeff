package main

import (
	"fmt"

	prefixo "github.com/Josh7GAS/gojeff/fundamentos/packages/prefixo"

	"github.com/Josh7GAS/gojeff/fundamentos/packages/operadora"
)

var NomeDoUsuario = "Joshua"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome de operadora: %s\r\n", operadora.NomeOperadora)
	//fmt.Printf("Valor de teste: %s\r\n", prefixo.teste)
	fmt.Printf("Teste Com Prefixo: %s\r\n", prefixo.TesteComPrefixo)
}
