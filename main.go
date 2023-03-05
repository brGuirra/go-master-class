package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/brGuirra/simple-bank/api"
	db "github.com/brGuirra/simple-bank/db/sqlc"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

const environment = ".env.development"

func main() {
	err := godotenv.Load(environment)
	if err != nil {
		log.Fatalln("cannot load env vars: ", err)
	}

	conn, err := sql.Open(os.Getenv("DATABASE_DRIVER"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(os.Getenv("SERVER_BASE_URL"))
	if err != nil {
		log.Fatalln("cannot start server: ", err)
	}
}
