tests:
	go test -v ./... 

.PHONY: dev-docker
dev-docker:
	@echo "Starting Application Docker"
	docker-compose -f docker-compose.yml -f docker-compose-build.yml up --build

.PHONY: down-docker
down-docker:
	docker-compose -f docker-compose.yml -f docker-compose-build.yml down

.PHONY: dev
dev:
	@echo "Starting application Local"
	go mod download
	docker-compose up -d
	go run cmd/api/*.go

.PHONY: down
down:
	docker-compose down
