package bd

import (
	"context"
	"time"

	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(cxt, t)
	if err != nil {
		return false, err
	}

	return true, err
}
