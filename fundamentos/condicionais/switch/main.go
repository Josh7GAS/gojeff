package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numero := 3
	fmt.Print("O numero ", numero, "se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	}

	fmt.Println("Você é da Familia do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim")
	default:
		fmt.Println("Deixa disso")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Pode dormir até mais tarde")
	default:
		fmt.Println("Vamos lá que é dia de Trabalho")
	}

	numero = 10
	fmt.Println("este numero cabe dois digito?", numero)
	switch {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve dois dígitos....")
	case numero >= 100:
		fmt.Println("Não, dá um numero muito grande")

	}

}
