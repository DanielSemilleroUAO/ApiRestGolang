package middlewares

import (
	"net/http"

	"github.com/DanielSemilleroUAO/ApiRestGolang/bd"
)

func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(rw, "Conexión Perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
