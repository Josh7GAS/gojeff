package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Josh7GAS/gojeff/avancado/jason_marshall/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("https://www.google.com.br/")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do Google Brasil. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler a pagina do Google Brasil. Erro: ", err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	request, err := http.NewRequest("Get", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o servidor. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do Servidor. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler a pagina do Google Brasil. Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		post := model.BlogPost
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter o retorno JSON do Servidor.Erro: ", err.Error())
			return
		}
		fmt.Printf("Post é: %+v\r\n", post)
	}
}
