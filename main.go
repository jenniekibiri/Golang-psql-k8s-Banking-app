package main

import (
	"database/sql"
	"log"

	"github.com/jenniekibiri/Golang-psql-k8s-Banking-app/api"
	db "github.com/jenniekibiri/Golang-psql-k8s-Banking-app/db/sqlc"
	"github.com/jenniekibiri/Golang-psql-k8s-Banking-app/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	println(config.DBDriver)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
