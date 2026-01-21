include .env
export

service-run:
	@export conn_string=postgres://postgres:admin@localhost:5432/postgres && \
	go run main.go

migrate-up:
	migrate -path migrations -database ${CONN_STRING} up

migrate-down:
	migrate -path migrations -database ${CONN_STRING} down