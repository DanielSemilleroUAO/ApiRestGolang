package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(t models.Relacion) (bool, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := col.FindOne(cxt, condicion).Decode(&resultado)
	fmt.Println(resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
