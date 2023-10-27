seed:
	# refresh / flush
	go run ./cmd/make_seed.go

migrate:
	# make migrate name=migration_name
	go run ./cmd/create_migrate.go -name ${name}

module:
	# make module name=demo
	go run ./cmd/create_module.go -name ${name}

compose:
	go run ./cmd/create_compose.go

swagger:
	swag init

start:
	CompileDaemon -command="./go_sample"

docker_up:
	# postgres up - create postgres server
	docker-compose up -d

docker_down:
	# postgres down - delete postgres server
	docker-compose down

db_up:
	docker exec -it sample-go createdb --username=postgres --owner=postgres simple_db

db_down:
	docker exec -it sample-go dropdb --username=postgres simple_db


.PHOY: seed migrate module compose swagger start docker_up docker_down db_up db_down