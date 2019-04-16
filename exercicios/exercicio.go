package main

import "fmt"

func soma(A int, B int) int {
	return A + B
}

//le numero e devolve
func lernumero() float64 {
	fmt.Println("Entre um número: ")
	var x float64
	fmt.Scanf("%f", &x)
	if x < 0 {
		x = x * (-1)
	}
	return x
}

func main() {
	fmt.Println("Soma é ", soma(34, 88))
	fmt.Println("Numero é ", lernumero())
}
