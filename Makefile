# Makefile
# Directory you want to create
DIR := $(HOME)/plak3-data

# Default action is to start all services
all: up

# Start all services
start:
	@[ -d "$(DIR)" ] || mkdir -p "$(DIR)" && cd docker && cd postgres && docker-compose up -d

# Stop all services
stop:
	cd docker && cd postgres && docker-compose down

# Rebuild and start all services
rebuild:
	cd docker && cd postgres && docker-compose up -d --build

# View logs
logs:
	cd docker && cd postgres && docker-compose logs

# Enter the PostgreSQL container
psql:
	cd docker && cd postgres && docker-compose exec db psql -U exampleuser exampledb

# Stop and remove all containers, networks, and volumes
clean:
	cd docker && cd postgres && docker-compose down -v

swagger-docs:
	swag init -g cmd/app/main.go -o cmd/docs 

migrate:
	goose -dir db/migrations postgres "host=10.0.0.78 user=postgres password=nopassword dbname=fms sslmode=disable" up

# Stop and remove all containers, networks, and volumes
migrate-delete:
	goose -dir db/migrations postgres "host=10.0.0.78 user=postgres password=nopassword dbname=fms sslmode=disable" down


.PHONY: start stop rebuild logs psql clean swagger-docs migrate-d migrate-u



	cd migrations && goose create $(PARAM)