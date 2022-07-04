postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=4524 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank 

migrateup:
	migrate -path db/migration -database "postgresql://root:4524@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:4524@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc-generate:
	docker run --rm -v "d:\ProgramPro\workspace\GoProject\simplebank:/src" -w /src kjconroy/sqlc generate

sqlc-init:
	docker run --rm -v "d:\ProgramPro\workspace\GoProject\simplebank:/src" -w /src kjconroy/sqlc init

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown showgcc