postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

cleantestcache:
	go clean -testcache

test1: 
	cd db && export "GO111MODULE=on" && go mod tidy && go test -v ./sqlc -cover

dockerconnect:
	docker exec -it postgres12 psql -U root -d simple_bank

isolationtestconnect:
	docker exec -it mysql8 mysql -uroot -psecret simple_bank
# go test -v ./sqlc -cover

server: 
	export "GO111MODULE=on" && go mod tidy && go run main.go
	
.PHONY: postgres createdb dropdb sqlc test dockerconnect server cleantestcache