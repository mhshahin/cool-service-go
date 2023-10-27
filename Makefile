serve:
	go run ./main.go serve

config:
	cp config.yaml.example config.yaml

# Change the database connection string according to your own credentials
migrate-up:
	migrate -path ./database/migrations -database "postgres://user:password@localhost:5432/database?sslmode=disable" up

migrate-down:
	migrate -path ./database/migrations -database "postgres://user:password@localhost:5432/database?sslmode=disable" down

generate-secret:
	openssl rand -hex 32