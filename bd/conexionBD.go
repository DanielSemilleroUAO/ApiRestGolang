package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la bd*/
var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.zdcsx.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conectar la BD*/
func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Printf("Conexión Exitosa con la BD")
	return client
}

/*ChequeoConnection es el Ping a la bd*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
