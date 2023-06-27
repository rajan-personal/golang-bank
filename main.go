package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rajan-personal/golang-bank/api"
	db "github.com/rajan-personal/golang-bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	// connect to database
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
