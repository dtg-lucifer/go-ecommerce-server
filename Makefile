run:
	@go run ./cmd/main.go

run_db:
	@docker compose up -d

build:
	@go -o ./bin/ecom ./cmd/main.go

test:
	@go test -v ./...
