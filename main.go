package main

import (
	"database/sql"
	"log"

	"github.com/DbWang888/simplebank/api"
	db "github.com/DbWang888/simplebank/db/sqlc"
	"github.com/DbWang888/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
