package model

import "database/sql"

//Local armazena os dados da localidade pelo seu codigo telefonico
type Local struct {
	Pais             string         `joson:"pais" db:"country" bson:"contry"`
	Cidade           sql.NullString `json:"cidade" db:"city" bson:"city"`
	CodigoTelefonico int            `json:"cod_telefone" db:"telcode" bson:"telcode"`
}
