package routers

import (
	"net/http"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", 400)
		return
	}

	err := bd.BorroTweet(ID, IdUsuario)
	if err != nil {
		http.Error(w, "Debe enviar el parametro ID "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
}
