run:
	@go run ./cmd/main.go

run-db:
	@sudo docker compose up

build:
	@go -o ./bin/ecom ./cmd/main.go

test:
	@go test -v ./...

migration:
	@migrate create -ext sql -dir db/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run db/migrate/migration.go up

migrate-down:
	@go run db/migrate/migration.go down

seed:
	@go run db/seed/seed.go

format:
	@go fmt ./...
