created-b:
	docker compose up
migrate-up:
	@migrate -path internal/db/migration -database "postgresql://postgres:password@localhost:5432/todoapp?sslmode=disable" -verbose up
migrate-down:
	@migrate -path internal/db/migration -database "postgresql://postgres:password@localhost:5432/todoapp?sslmode=disable" -verbose down
migrate-force:
	@read -p "Enter migration version: " version; \
	migrate -path internal/db/migration -database "postgresql://postgres:password@localhost:5432/todoapp?sslmode=disable" -verbose force $$version
create-migrate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir internal/db/migration -seq $$name
sqlc:
	@sqlc generate
test:
	@go test -v -cover ./...

.PHONY: migrateup migratedown migrateforce sqlc createdb test
