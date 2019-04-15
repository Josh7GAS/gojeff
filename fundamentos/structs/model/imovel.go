package model

///Imovel Representa informações de um imóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
