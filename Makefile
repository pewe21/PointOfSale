db_up:
	migrate -database postgres://pos:passpos@localhost:5432/pointOfSale?sslmode=disable -path internal/database/migration up

db_up_1:
	migrate -database postgres://pos:passpos@localhost:5432/pointOfSale?sslmode=disable -path internal/database/migration up 1

db_down:
	migrate -database postgres://pos:passpos@localhost:5432/pointOfSale?sslmode=disable -path internal/database/migration down

db_down_1:
	migrate -database postgres://pos:passpos@localhost:5432/pointOfSale?sslmode=disable -path internal/database/migration down 1

run:
	go run cmd/main.go

#create migration
# migrate create -ext sql -dir internal/database/migration "create_customers_table"