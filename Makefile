postgres:
	@echo "Checking if container 'postgres12' exists..."
	@if [ "$$(docker ps -a -q -f name=^/postgres12$$)" = "" ]; then \
		echo "Container not found. Creating..."; \
		docker run --name postgres12 --network bank-network -p 5432:5432 \
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

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

build:
	docker build -t simplebank:latest .

run:
	docker run \
		--name simplebank \
		--network bank-network \
		-p 8080:8080 \
		-e GIN_MODE=release \
		-e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" \
		simplebank:latest

clean-image:
	@echo "🔍 Checking for container named 'simplebank'..."
	@if docker ps -a --format '{{.Names}}' | grep -q '^simplebank$$'; then \
		echo "🗑️  Removing container 'simplebank'..."; \
		docker rm -f simplebank; \
	else \
		echo "✅ No container named 'simplebank' found."; \
	fi

	@echo "🔍 Checking for image 'simplebank:latest'..."
	@if docker images --format '{{.Repository}}:{{.Tag}}' | grep -q '^simplebank:latest$$'; then \
		echo "🗑️  Removing image 'simplebank:latest'..."; \
		docker rmi -f simplebank:latest; \
	else \
		echo "✅ No image 'simplebank:latest' found."; \
	fi

.PHONY: createdb dropdb postgres migrateup migrateup1 migratedown migratedown1 sqlc test server mock build run clean-image