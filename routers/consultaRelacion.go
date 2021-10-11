package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IdUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-type", "applcation/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)

}
