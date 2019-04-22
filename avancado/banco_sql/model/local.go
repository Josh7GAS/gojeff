package model

import "database/sql"

//local armazena os dados da localidade pelo seu codigo telefonico

type Local struct {
	Pais             string         `joson:"pais" db:"country"`
	Cidade           sql.NullString `json:"cidade" db:"city"`
	CodigoTelefonico int            `json:"cod_telefone" db:"telcode"`
}
