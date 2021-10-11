package bd

import (
	"context"
	"time"

	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}

	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}

	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	}

	_, err := col.UpdateOne(cxt, filtro, updateString)

	if err != nil {
		return false, err
	}

	return true, nil

}
