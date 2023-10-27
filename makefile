###==============================> start docker and database initialization
docker:
	go run ./cmd/create_docker.go

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
###==============================> end docker and database initialization



###==============================> start custom cmd CLI
seed:
	# refresh / flush
	go run ./cmd/make_seed.go

migrate:
	# make migrate name=migration_name
	go run ./cmd/create_migrate.go -name ${name}

module:
	# make module name=demo
	go run ./cmd/create_module.go -name ${name}

swagger:
	swag init
###==============================> end custom cmd CLI



###==============================> start project with build and watch
start:
	CompileDaemon -command="./go_sample"


.PHOY: docker docker_up docker_down db_up db_down seed migrate module swagger start