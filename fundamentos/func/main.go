package main

import (
	"fmt"

	"github.com/Josh7GAS/gojeff/fundamentos/func/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 90, 3)
	fmt.Printf("O resultaddo da divisão é: %d\r\n", resultado)
}

func multiplicador(x int, y int) int {

	return x * y
}
