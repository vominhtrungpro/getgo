package main

import (
	"context"
	"log"

	"github.com/kytruong0712/getgo/api/internal/appconfig/db/pg"
	"github.com/kytruong0712/getgo/api/internal/httpserver"
)

func main() {
	ctx := context.Background()

	conn, err := pg.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	rtr, err := initRouter(ctx)
	if err != nil {
		log.Fatal(err)
	}

	httpserver.Start(httpserver.Handler(ctx, rtr.routes))
}

func initRouter(_ context.Context) (router, error) {
	return router{}, nil
}
