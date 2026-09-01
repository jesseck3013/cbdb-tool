.PHONY: sql test build 

sql:
	@echo "Generating SQL code..."
	sqlc generate
	@echo "====================="

test: 
	@echo "Running tests..."
	go test ./... -v
	@echo "====================="

build: sql test
	@echo "Compiling Go binary..."
	go build -o bin/cbdb ./cmd/cli/main.go

build-frontend:
	cd frontend && npm install && npm run build
