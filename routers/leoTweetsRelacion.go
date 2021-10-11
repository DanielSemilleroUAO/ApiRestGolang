package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
)

func LeotTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Error", 400)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Error", 400)
		return
	}

	respuesta, correcto := bd.LeotTweetsSeguidores(IdUsuario, pagina)

	if correcto == false {
		http.Error(w, "Error", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuesta)
}
