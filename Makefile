build:
	@go build -o bin/password-service cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/password-service
	
# Target to create a new migration
# @read -p "Enter migration name: " name; 
migration:
	@read -p "Enter migration name: " name; 
	touch cmd/migrate/migrations/`date +%s`_$$name.up.sql cmd/migrate/migrations/`date +%s`_$$name.down.sql; \
	echo "Migration files created: cmd/migrate/migrations/`date +%s`_$$name.up.sql and .down.sql"

# Target to run the migrations up
migrate-up:
	go run cmd/migrate/main.go up

# Target to run the migrations down
migrate-down:
	go run cmd/migrate/main.go down

# migration:
# 	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

# migrate-up:
# 	@go run cmd/migrate/main.go up

# migrate-down:
# 	@go run cmd/migrate/main.go down
