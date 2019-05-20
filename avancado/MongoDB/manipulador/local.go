package manipulador

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Josh7GAS/gojeff/avancado/banco_sql/model"
	"github.com/Josh7GAS/gojeff/repo"
)

//Local é o manipulador da requisição de rota /local/
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um numero válido.",
			http.StatusBadRequest)
		fmt.Println("[local] erro ao converter o numero enviado", err.Error())
		return
	}
	sql := "select county, city, telcode from place where telcode = ?"
	linha, err := repo.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possivel pesquisar esse número.", http.StatusInternalServerError)
		fmt.Println("[local] não foi possivel executar a query: ", sql, "Erro: ",
			err.Error())
		return
	}
	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possivel pesquisar esse número.", http.StatusInternalServerError)
			fmt.Println("[local] não foi possivel fazer o binding dos dados do banco na struct: ", sql, "Erro: ",
				err.Error())
			return
		}
	}

	localMongo, err := repo.ObtemLocal(codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possivel pesquisar esse número.", http.StatusInternalServerError)
		fmt.Println("[local] não foi possivel pesquisar a Query no MongoDB. Erro:",
			err.Error())

		return
	}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", localMongo); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execção do modelo: ", err.Error())
	}

}
