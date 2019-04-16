package model

//representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
}

//representa uma ave do tipo pato
type Pato interface {
	Quack() string
}

type Ave struct {
	Nome string
}

//Representa um som feito por uma galinha
func (a Ave) Cacareja() string {
	return "Cocorico..."
}

//Representa um som feito por um pato
func (a Ave) Quack() string {
	return "Quack, Quack..."
}
