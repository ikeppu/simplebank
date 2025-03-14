package main

import (
	"context"
	"log"

	"github.com/ikeppu/simplebank/api"
	db "github.com/ikeppu/simplebank/db/sqlc"
	"github.com/ikeppu/simplebank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Load configuration
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
