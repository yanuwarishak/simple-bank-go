postgres:
	@echo "Checking if container 'postgres12' exists..."
	@if [ "$$(docker ps -a -q -f name=^/postgres12$$)" = "" ]; then \
		echo "Container not found. Creating..."; \
		docker run --name postgres12 -p 5432:5432 \
			-e POSTGRES_USER=root \
			-e POSTGRES_PASSWORD=secret \
			-d postgres:12-alpine; \
	elif [ "$$(docker ps -q -f name=^/postgres12$$)" = "" ]; then \
		echo "Container exists but not running. Starting..."; \
		docker start postgres12; \
	else \
		echo "Container 'postgres12' is already running."; \
	fi

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server