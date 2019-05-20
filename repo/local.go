package repo

import (
	"github.com/Josh7GAS/gojeff/avancado/MongoDB/model"
	"gopkg.in/mgo.v2/bson"
)

//ObtemLocal retorna um local do MongoDB
func ObtemLocal(codigoTelefone int) (local model.Local, err error) {
	sessao := SessaoMongo.Copy()
	defer sessao.Close()
	colecao := sessao.DB("cursodego").C("local")
	err = colecao.Find(bson.M{"telcode": codigoTelefone}).One(&local)
	return
}
