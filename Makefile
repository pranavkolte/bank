postgres:
	docker run --name go-bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

createdb:
	docker exec -it go-bank createdb --username=root --owner=root bank

dropdb:
	docker exec -it go-bank dropdb bank

migrateup:
	migrate -path db/migration  -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration  -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrateup migratedown sqlc