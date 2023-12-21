package main

import (
	"database/sql"
	"github.com/BuiNhatTruong99/Go-Simple-Bank-/api"
	db "github.com/BuiNhatTruong99/Go-Simple-Bank-/db/sqlc"
	"github.com/BuiNhatTruong99/Go-Simple-Bank-/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	log.Fatal("cannot start server: ", err)
}
