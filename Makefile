#!make
include .env.development

createdb:
	docker exec -it ${CONTAINER_NAME} createdb --username=${DATABASE_USERNAME} --owner=${DATABASE_USERNAME} ${DATABASE_NAME}

dropdb:
	docker exec -it ${CONTAINER_NAME} dropdb ${DATABASE_NAME}

migrateup:
	migrate --path db/migration -database "postgresql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@localhost:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate --path db/migration -database "postgresql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@localhost:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

dcupdev:
	docker-compose --env-file .env.development up -d

dcdowndev:
	docker-compose --env-file .env.development down

.PHONY: createdb dropdb migrateup migratedown sqlc test dcupdev dcdowndev