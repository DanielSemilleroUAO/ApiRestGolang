package routers

import (
	"net/http"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Error", 400)
		return
	}

	var t models.Relacion
	t.UsuarioID = IdUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Error", 400)
		return
	}

	if status == false {
		http.Error(w, "Error", 400)
		return
	}

	w.WriteHeader(201)
}
