package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/kytruong0712/getgo/api/internal/appconfig/db/pg"
	"github.com/kytruong0712/getgo/api/internal/controller/products"
	"github.com/kytruong0712/getgo/api/internal/httpserver"
	"github.com/kytruong0712/getgo/api/internal/repository"
	"github.com/kytruong0712/getgo/api/internal/repository/generator"
	"os"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	fmt.Println("PROJECT_NAME: ", os.Getenv("PROJECT_NAME"))

	ctx := context.Background()

	conn, err := pg.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	generator.InitSnowflakeGenerators()

	rtr, err := initRouter(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	httpserver.Start(httpserver.Handler(ctx, rtr.routes))
}

func initRouter(_ context.Context, db *sql.DB) (router, error) {
	repo := repository.New(db)

	productCtrl := products.New(repo)

	return router{
		productCtrl: productCtrl,
	}, nil
}
