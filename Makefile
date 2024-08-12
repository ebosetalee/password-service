build:
	@go build -o bin/password-service cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/password-service