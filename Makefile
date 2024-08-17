db_up:
	migrate -database postgres://postgres@localhost:5432/pointOfSale?sslmode=disable -path internal/database/migration up

db_down:
	migrate -database postgres://postgres@localhost:5432/pointOfSale?sslmode=disable -path internal/database/migration down

run:
	go run cmd/main.go