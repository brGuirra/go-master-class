package main

import (
	"database/sql"
	"log"

	"github.com/brGuirra/simple-bank/api"
	db "github.com/brGuirra/simple-bank/db/sqlc"
	"github.com/brGuirra/simple-bank/utils"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("cannot start server: ", err)
	}
}
