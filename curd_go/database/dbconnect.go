package database

import (
	"context"
	"curd/errorhandler"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var CrudDB *mongo.Database

func Start() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	Client, err := mongo.Connect(ctx)

	if err != nil {
		e := errorhandler.GetErorr(err, http.StatusInternalServerError, "Databse start issue.")
		log.Fatal(e)
	}
	CrudDB = Client.Database("crud")
	fmt.Println("Db connected!")

}
