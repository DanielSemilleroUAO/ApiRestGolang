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
	router.HandleFunc("/listaUsuarios", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.LeotTweetsSeguidores))).Methods("GET")
	//Tweets
	router.HandleFunc("/tweet", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/subirAvatar", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/subirBanner", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.ObtenerBanner))).Methods("GET")
	//Relaciones
	router.HandleFunc("/altaRelacion", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlewares.ChequeoDB(middlewares.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
