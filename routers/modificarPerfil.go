package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IdUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Ocurrio un error al intentar modificar el registro "+err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}
