postgres:
	wsl sudo docker run --name go-bank -p 127.0.0.1:5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

createdb:
	wsl sudo docker exec -it go-bank createdb --username=root --owner=root bank

dropdb:
	wsl docker exec -it go-bank dropdb bank

migrateup:
	migrate -path db/migration  -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration  -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration  -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration  -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock