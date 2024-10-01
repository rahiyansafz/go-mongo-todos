include .env

PROTO_DIR = pb
GO_OUT_DIR = pb

.PHONY: proto
proto:
	protoc --proto_path=$(PROTO_DIR) --go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto

.PHONY: build
build:
	go build -o ${BINARY} ./cmd/server

.PHONY: run
run:
	go run ./cmd/server

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: restart
restart: down up