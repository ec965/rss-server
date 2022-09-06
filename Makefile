all: server

run:
	go run main.go
.PHONY: run

migrate:
	go run cmd/migrate/main.go
.PHONY: migrate

docker-build:
	docker build --tag rss-server .
.PHONY: docker-build

docker-run:
	docker run --name rss-server -p 127.0.0.1:3000:3000/tcp rss-server 
.PHONY: docker-run

clean:
	rm rss-server
