package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/DanielSemilleroUAO/ApiRestGolang/middlewares"
	"github.com/DanielSemilleroUAO/ApiRestGolang/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoDB(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
