
setup: 
	go install github.com/cosmtrek/air@latest

new-migration:
	docker run --rm -v $(CURDIR)/database/migrations:/migrations \
	migrate/migrate create -ext sql -dir migrations -seq $(NAME)

migrate:
	docker run --rm -v $(CURDIR)/database/migrations:/migrations --network host migrate/migrate \
  	-path=/migrations/ -database postgres://monban:monban@localhost:5432/monban?sslmode=disable up

dump-schema:
	docker exec -e PGPASSWORD=monban monban-db-1 pg_dump --schema-only -U monban monban > db_schema.sql

sqlc-generate:
	docker run --rm -v $(CURDIR):/src -w /src/sqlc sqlc/sqlc generate

