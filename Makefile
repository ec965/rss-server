all: server

run:
	go run main.go
.PHONY: run

migrate:
	go run cmd/migrate/main.go
.PHONY: migrate

clean:
	rm rss-server
