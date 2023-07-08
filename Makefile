postgres:
	docker run --name gopg-db -p 5432:5432 -e POSTGRES_PASSWORD=pablo2846 -e POSTGRES_USER=root -d postgres:latest

createdb:
	docker exec -it gopg-db createdb --username=root --owner=root gopg-db

dropdb:
	docker exec -it gopg-db dropdb gopg-db

migrateup:
	migrate -path db/migration -database "postgresql://root:pablo2846@localhost:5432/gopg-db?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:pablo2846@localhost:5432/gopg-db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

PHONY: createdb dropdb migrateup migratedown postgres sqlc
