package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielSemilleroUAO/ApiRestGolang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSkip((page - 1) * 20)

	query := bson.M{
		"nombre": bson.M{
			"$regex": `(?i)` + search,
		},
	}

	cur, err := col.Find(cxt, query, findOptions)
	if err != nil {
		fmt.Println(results)
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(cxt) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()
		incluir = false
		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		if tipo == "follow" && encontrado == true {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
		}

		results = append(results, &s)
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	return results, true
}
