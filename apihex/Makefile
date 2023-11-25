.PHONY: api

api:
	go run main.go
alot:
	TRANSFER_URL=http://localhost:8080/transfer go run cmd/alot/main.go

pid := $(shell lsof -ti:8080)
sigterm:
	kill -SIGTERM $(pid)

build:
	docker build -t todo:latest -f Dockerfile .
run:
	docker run -p:8081:8081 --env-file ./.env --name mytodo todo:latest