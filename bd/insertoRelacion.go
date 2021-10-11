package bd

import (
	"context"
	"time"

	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
)

func InsertoRelacion(t models.Relacion) (bool, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(cxt, t)
	if err != nil {
		return false, err
	}

	return true, err

}
