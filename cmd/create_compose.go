package main

import (
	"fmt"
	"go_sample/core/config"
	"os"
)

func main() {
	conf := config.All()

	content := fmt.Sprintf(`version: '3'

services:
  postgres:
    container_name: %s
    image: %s
    environment:
      - POSTGRES_USER=%s
      - POSTGRES_PASSWORD=%s
    ports:
      - '5432:%s'
`, conf.Docker.Postgres.Container, conf.Docker.Postgres.Image, conf.Database.Username, conf.Database.Password, conf.Database.Port)

	// Create the migration file
	fileName := fmt.Sprintf("%s/docker-compose.yaml", "./")
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating docker compose file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("file created: %s\n", fileName)
}
