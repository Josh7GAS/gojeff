package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/Josh7GAS/gojeff/intermediario/lerarquivos/model"
) 

func main() {
	arquivo, err := os.Open("cidades.cvs")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir ao arquivo. Erro: ", err.Error())
		return
	}

	// scanner := bufio.NewScanner(arquivo)
	// for scanner.Scan() {
	// 	linha := scanner.Text
	// 	fmt.Println("O Conteúdo da linha é: ", linha())
	// }

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com leitor CSV. Error: ",
			err.Error())
		return
	}

	arquivoJson, err := os.Create("cidades.json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro: ", err.ERRO())
		return
	}


	escritor := bufio.NewWriter(arquivoJson)
	escritor.WriteString("[\r\n")

	for _, linha := range conteudo {
		for indiceItem, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome := dados[0]
			cidade.Estado := dados[1]
			fmt.Printf("Cidade: %+v\r\n", cidade)

			cidadeJSON:=json.Marshal(cidade)

			if err != nil {
				fmt.Println("[main] Houve um erro ao gerar o json do item",item, ".Erro: ", err.Error() )
				
			}
			escritor.WriteString(" " + string(cidadeJSON))
			if (indiceItem+1) < len(linha){
				escritor.WriteString(",\r\n")
			}
		 
		}
	}
	escritor.WriteString("\r\n]")
	escritor.Flush()
	arquivo.JSON.Close()

	arquivo.Close()
}
