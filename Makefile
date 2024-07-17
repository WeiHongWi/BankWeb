postgres:
	docker run --name postgres_ct -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=fighting -d postgres:16-alpine
createdb:
	docker exec -it postgres_ct createdb --username=root --owner=root bank
dropdb:
	docker exec -it postgres_ct dropdb bank
migrateup:
	migrate -path db/migration -database "postgresql://root:fighting@localhost:5432/bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:fighting@localhost:5432/bank?sslmode=disable" -verbose down
test:
	go test -v -cover ./...
.PHONY: createdb,dropdb,postgres,migrateup,migratedownp