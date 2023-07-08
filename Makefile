postgres:
	docker run --name gopg-db -p 5432:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:latest

createdb:
	docker exec -it gopg-db createdb --username=root --owner=root gopg-db

dropdb:
	docker exec -it gopg-db dropdb gopg-db

createtestdb:
	docker exec -it gopg-db createdb --username=root --owner=root gopg-db-test

droptestdb:
	docker exec -it gopg-db dropdb gopg-db-test

migrateuptest:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gopg-db-test?sslmode=disable" -verbose up

migratedowntest:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gopg-db-test?sslmode=disable" -verbose down

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gopg-db?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gopg-db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

testdb:
	go test -v -cover ./... -coverprofile=coverage.out 

coverhtml:
	go tool cover -html=coverage.out

PHONY: createdb dropdb migrateup migratedown postgres sqlc testdb createtestdb droptestdb migrateuptest migratedowntest coverhtml
