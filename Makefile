postgres:
	docker run --name banktransf -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine
createdb:
	docker exec -it banktransf createdb --username=root --owner=root bank_transf

dropdb:
	docker exec -it banktransf dropdb bank_transf

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_transf?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_transf?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/BuiNhatTruong99/Go-Simple-Bank-/db/sqlc Store

.PHONY: posgres createdb dropdb migrateup migratedown sqlc test server mock
