package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IdUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen "+err.Error(), 400)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), 400)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IdUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IdUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la bd "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)

}
