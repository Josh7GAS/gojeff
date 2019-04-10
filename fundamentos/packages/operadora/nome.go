package operadora

import (
	"strconv"

	prefixo "github.com/Josh7GAS/gojeff/fundamentos/packages/prefixo"
)

//NomeOperadora representa o nome da operadora

var NomeOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora= prefixo mais o nome da operadora, que importa da pasta numeros
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
