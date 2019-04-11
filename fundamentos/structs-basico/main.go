package main

import "fmt"

type Imovel struct {
	X     int
	Y     int
	Nome  string
	Valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("Casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 76000}
	fmt.Printf("O Apartamento é %+v\r\n", apartamento)

	chacara := Imovel{
		Y:    85,
		Nome: "Chacara",
		X:    22,
	}
	fmt.Printf("A Chacara é: %+v\r\n", chacara)
	casa.Nome = "Lar Doce Lar"
	casa.Valor = 45000
	casa.X = 18
	casa.Y = 31
	fmt.Printf("A casa é: %+v\r\n", casa)
}
