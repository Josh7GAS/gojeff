package manipulador

import (
	"net/http"
	"time"
)

//olá é o manipulador da requisição a rota/ola
func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")
	pagina := model.pagina{}
	pagina.Hora = hora

	if err := Modelos.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro na rendereização da página.", http.StatusInternalServerError)
		fmt.Prinln("[Ola] erro na execução do modelo ", err.Error())
	}
}
