package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Josh7GAS/gojeff/avancado/webpost/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Jeff Prestes"

	conteudoEnviar, err := json.Marshall(usuario)
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

	request, err := http.NewRequest("Post", "https://enkgr2tmpa2ja.x.pipedream.net", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para a pagina do Google Brasil. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	resposta, err = cliente.Do(request)
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
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}
}
