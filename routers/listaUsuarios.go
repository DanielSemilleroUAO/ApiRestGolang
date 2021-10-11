package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
)

func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Error", 400)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.LeoUsuariosTodos(IdUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)
}
