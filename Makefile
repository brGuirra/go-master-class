#!make
include app.env

createdb:
	docker exec -it ${CONTAINER_NAME} createdb --username=${DATABASE_USERNAME} --owner=${DATABASE_USERNAME} ${DATABASE_NAME}

dropdb:
	docker exec -it ${CONTAINER_NAME} dropdb ${DATABASE_NAME}

migrateup:
	migrate --path db/migration -database "postgresql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@localhost:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" -verbose up

migrateup1:
	migrate --path db/migration -database "postgresql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@localhost:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" -verbose up 1

migratedown:
	migrate --path db/migration -database "postgresql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@localhost:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" -verbose down

migratedown1:
	migrate --path db/migration -database "postgresql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@localhost:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

dcupdev:
	docker-compose --env-file app.env up -d

dcdowndev:
	docker-compose --env-file app.env down

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/brGuirra/simple-bank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown sqlc test dcupdev dcdowndev server mock migrateup1 migratedown1
