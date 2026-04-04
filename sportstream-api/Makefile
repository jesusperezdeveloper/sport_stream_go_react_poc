.PHONY: run build test lint docker-build docker-run

run:
	go run ./cmd/server/main.go

build:
	CGO_ENABLED=0 go build -o bin/sportstream-api ./cmd/server/main.go

test:
	go test ./... -v -race

lint:
	golangci-lint run ./...

docker-build:
	docker build -t sportstream-api -f deployments/Dockerfile .

docker-run:
	docker run -p 8080:8080 sportstream-api
