include .env

up:
	@echo "Starting Docker containers..."
	docker-compose up --build -d --remove-orphans

down: 
	@echo "Stopping Docker containers..."
	docker-compose down

build:
	@echo "Building Go application..."
	go build -o ${BINARY} ./cmd/api/

start:
	@echo "Starting Go application..."
	./${BINARY}

restart:
	@echo "Restarting Go application..."
	$(MAKE) build && $(MAKE) start
