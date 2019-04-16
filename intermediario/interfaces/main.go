package main

import (
	"fmt"

	"github.com/Josh7GAS/gojeff/intermediario/interfaces/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	QueroAcordarComUmCacarejo(jojo)
	QueroOuvirUmaPataNoLago(jojo)

}

func QueroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func QueroOuvirUmaPataNoLago(g model.Pato) {
	fmt.Println(g.Quack())
}
