.PHONY: dev migrate clean build

dev:
	docker compose up --build

migrate:
	docker compose --profile migrate run --rm migrate

clean:
	docker compose down -v

build:
	go build -o bin/calm-map ./cmd/server
