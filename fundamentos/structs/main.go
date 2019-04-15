package main

import (
	"encoding/json"
	"fmt"

	"github.com/Josh7GAS/gojeff/fundamentos/structs/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	fmt.Printf("O valor da casa é: %d\r\n", casa.GetValor())
	objJson, _ := json.Marshal(casa)

	fmt.Println("A casa em JSON: ", string(objJson))
}
