package main

import (
	"fmt"

	"github.com/Josh7GAS/gojeff/fundamentos/packages/operadora"
	"github.com/Josh7GAS/gojeff/fundamentos/packages/prefixos"
)

//NomedoUsusario do sistema
var NomedoUsusario = "Joshua"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomedoUsusario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixos.Capital)
	fmt.Printf("Prefixo da Capital: %s\r\n", operadora.NomeOperadora)

}
