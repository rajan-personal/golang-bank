package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rajan-personal/golang-bank/api"
	db "github.com/rajan-personal/golang-bank/db/sqlc"
	"github.com/rajan-personal/golang-bank/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// connect to database
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		panic(err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
