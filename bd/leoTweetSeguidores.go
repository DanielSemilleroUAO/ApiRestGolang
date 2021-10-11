package bd

import (
	"context"
	"time"

	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeotTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})

	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})

	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})

	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})
	cursor, err := col.Aggregate(cxt, condiciones)

	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(cxt, &result)

	if err != nil {
		return result, false
	}

	return result, true

}
