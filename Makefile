postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=scret -d postgres
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres dropdb simple_bank
migrateUp:
	migrate -path db/migration -database "postgresql://root:scret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateDown: 
	migrate -path db/migration -database "postgresql://root:scret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateUp migrateDown sqlc test