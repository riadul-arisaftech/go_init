package main

import (
	"go_sample/api"
	"go_sample/core/config"
	"go_sample/core/db"
	"go_sample/database"
	"log"
	"time"
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

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(70)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// DB migrations
	if err := database.Migrate(db); err != nil {
		log.Println(err)
		panic(err.Error())
	}

	server := api.NewServer(conf, database.NewStore(db))
	server.Run()
}
