package main

import "fmt"

func main() {
	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1
	fmt.Println("Qual é a capacidade deste array", len(teste))

	cantores := [2]string{"Michael Jackson", "Bruno Mars"}
	fmt.Printf("O que a neste array? \n\r%v\r\n", cantores)

	capitais := [...]string{"Lisboa", "Brasilia", "Luanda", "Maputo"}
	fmt.Println("Qual a capacidade deste array?", len(capitais))
	for indice, cidade := range capitais {
		fmt.Printf("Capital[%d] é  %s", indice, cidade)
	}
}
