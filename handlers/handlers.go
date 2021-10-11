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

	//Usuarios
	router.HandleFunc("/registro", middlewares.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	//Tweets
	router.HandleFunc("/tweet", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
