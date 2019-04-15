package main

import "fmt"

func main() {
	meses := 0
	situacao := true
	cidade := "Teste"

	if meses <= 6 {
		fmt.Println("Esse Credor deve a pouco tempo.")

	}

	if situacao {
		fmt.Println("Ele está devendo")
	}

	if !situacao {
		fmt.Println("Ele está adimplente")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente está no estado teste")
	}
	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual é a situação do cliente?", descricao)
		return
	}

	fmt.Println("Obrigado por nos Consultar")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses >= 1 {
		status = true
		descricao = "O Cliente está devendo."
		return
	}
	descricao = "O Cliente está em dia!"
	return
}
