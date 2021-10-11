package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IdUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Error en la petición "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Error en la petición", 400)
		return
	}

	w.WriteHeader(200)
}
