package main

import (
	"fmt"

	"github.com/Josh7GAS/gojeff/fundamentos/structs/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Println("Sim, e o que achei foi:", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apartamento Azul"
	apto.X = 19
	apto.Y = 26
	apto.SetValor(70000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens a no cache?", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n\r", chave, imovel)
	}
	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Quantos itens há no cache?", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n\r", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Quantos itens há no cache", len(cache))
}
