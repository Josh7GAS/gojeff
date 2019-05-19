package operadora

import (
	"strconv"

	"github.com/Josh7GAS/gojeff/fundamentos/packages/prefixos"
)

//NomeOperadora é o nome da operadora
var NomeOperadora = "XPTO Telecom"

//PrefixoOperadora prefixo mais o nome da operadora
var PrefixoOperadora = strconv.Itoa(prefixos.Capital) + "" + NomeOperadora
