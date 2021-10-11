package middlewares

import (
	"net/http"

	"github.com/DanielSemilleroUAO/ApiRestGolang/routers"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(rw, "Error en el Token "+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(rw, r)
	}
}
