run:
	@docker-compose up -d
	@go run ./cmd/main.go

build:
	@go -o ./bin/ecom ./cmd/main.go

test:
	@go test -v ./...
