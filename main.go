package main

import (
	"bank/api"
	db "bank/db/sqlc"
	"bank/util"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	fmt.Println("Connecting to DB")
	if err != nil {
		log.Fatal("cannot connect to Database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server")
	}

}
