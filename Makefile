.PHONY: sql build 

sql:
	@echo "Generating SQL code..."
	sqlc generate

build: sql
	@echo "Compiling Go binary..."
	go build -o bin/cbdb ./cmd/cli/main.go
