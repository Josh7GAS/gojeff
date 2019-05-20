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
	if err := casa.SetValor(11000); err != nil {
		fmt.Println("[main] houve um erro na atribuição de valor a casa: ", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor, Por favor verifique a sua avaliação")
		}
		return
	}
	fmt.Printf("O valor da casa é: %d\r\n", casa.GetValor())

	objJson, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] houve um erro na geração do objeto JSON: ", err.Error())

		return
	}
	fmt.Println("A casa em JSON: ", string(objJson))
}
