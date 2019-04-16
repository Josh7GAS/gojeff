package main

import "fmt"

func soma(A int, B int) int {
	return A + B
}

func main() {
	fmt.Println("Soma é %+v\r\n", soma(34, 88))
}
