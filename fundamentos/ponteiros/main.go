package main

import "fmt"

//Imovel é uma struct que armazena dados de um imovel

type Imovel struct {
	X     int
	Y     int
	Nome  string
	Valor int
}

func main() {

	casa := new(Imovel)
	fmt.Printf("Casa é: %p - %+v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "Chacara Linda", 28000}
	fmt.Printf("chacara é: %p - %+v\r\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("chacara é: %p - %+v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
