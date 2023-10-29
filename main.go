package main

import (
	"github.com/go_sample/api"
	"github.com/go_sample/core/config"
	"github.com/go_sample/core/db"
	"github.com/go_sample/database/repository"
)

// @title Go Sample Backend
// @version 1.0
// @description This is a sample server for using Swagger with Echo.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	conf := config.All()

	dbConfig := db.NewConfiguration(conf)
	db := db.InitDB(dbConfig)
	store := repository.NewStore(db)

	server := api.NewServer(conf, store)
	server.Run()
}
