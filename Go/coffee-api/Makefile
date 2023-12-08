DB_DOCKER_CONTAINER=coffee-api-container
BINARY_NAME=coffeeapi
DB_NAME=coffeedb
USER=root
PW=password
DB_URL="postgresql://${USER}:${PW}@localhost:5432/${DB_NAME}?sslmode=disable"

# Creating the container with Postgres software
postgres:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PW} -d postgres:12-alpine

# Creating the coffee database inside the postgres container
createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

# Stop other Docker containers
stop_containers:
	@echo "Stopping other Docker containers";
	if [ "$$(docker ps -q)" ]; then \
		echo "Found and stopped containers..."; \
		docker stop $$(docker ps -q); \
	else \
		echo "No active containers found."; \
	fi

# Start Docker container
start_docker:
	docker start ${DB_DOCKER_CONTAINER}

create_migrations:
	sqlx migrate add -r init

migrate_up:
	sqlx migrate run --database-url=${DB_URL}

migrate_down:
	sqlx migrate revert --database-url=${DB_URL}
	
run:
	go run cmd/server/main.go