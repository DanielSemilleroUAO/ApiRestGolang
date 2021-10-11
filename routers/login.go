package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
	"github.com/DanielSemilleroUAO/ApiRestGolang/jwt"
	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(rw, "Usuario y/o contraseña inválidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(rw, "El email del usuario es requerido "+err.Error(), 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(rw, "Usuario no registrado"+err.Error(), 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(rw, "Usuario no registrado"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(resp)

	/*Cookie desde el backend*/
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
