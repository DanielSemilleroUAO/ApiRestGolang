package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
)

func Registro(rw http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(rw, "Fallo de conexión"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(rw, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(rw, "La contraseña no es valida", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(rw, "El usuario ya existe", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(rw, "Ocurrió un error inesperado en el registro", 400)
		return
	}

	if status == false {
		http.Error(rw, "Ocurrió un error inesperado en el registro", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
