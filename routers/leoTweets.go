package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
)

func LeoTweets(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el id", 400)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(rw, "Debe enviar el parámetro página", 400)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(rw, "Debe enviar el parámetro página", 400)
		return
	}

	pag := int64(pagina)
	respuesta, corecto := bd.LeoTweet(ID, pag)
	if corecto == false {
		http.Error(rw, "Error al leer los tweets", 400)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(respuesta)
}
