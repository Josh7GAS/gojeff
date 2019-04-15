package model

import "errors"

///Imovel Representa informações de um imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_Y"`
	Nome  string `json:"nome"`
	valor int
}

/*
ErrValorDeImovelInvalido é um erro que é apresentado quando é atribuido a imovel em um valor muito baixo
*/

/*
ErrValorDeImovelMuitoAlto é um erro que é apresentado quando é atribuido ao valor muito alto, fora do mercado
*/

var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto")

var ErrValorDeImovelInvalido = errors.New("O valor informado não é valido para um imóvel")

func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
