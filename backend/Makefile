MYSQL_URL="mysql://myapp:myapppassword@tcp(localhost:3306)/myapp"

.PHONY: migrate-up migrate-down

migrate-up:
	migrate -database ${MYSQL_URL} -path internal/db/migrations up

migrate-down:
	migrate -database ${MYSQL_URL} -path internal/db/migrations down 