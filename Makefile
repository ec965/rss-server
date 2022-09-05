all: server

run:
	go run cmd/server/main.go
.PHONY: run

migrate:
	go run cmd/migrate/main.go
.PHONY: migrate

server:
	go build -o server cmd/server/main.go
.PHONY: server

clean:
	rm server
