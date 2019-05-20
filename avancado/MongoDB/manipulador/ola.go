package manipulador

import (
	"net/http"
	"time"

	"github.com/Josh7GAS/gojeff/avancado/servidorweb/model"
)

//Ola é um manipulador da requesição a rota /ola
func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:05:07")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := ModeloLocal.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
	}
}
