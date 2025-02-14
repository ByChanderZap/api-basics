build:
	@go build -o bin/econ cmd/main.go

run: build
	@./bin/econ

migrate-up:
	@goose -dir cmd/migrate/migrations postgres "postgres://alex:password@localhost:5432/mydatabase" up

migrate-down:
	@goose -dir cmd/migrate/migrations postgres "postgres://alex:password@localhost:5432/mydatabase" down
