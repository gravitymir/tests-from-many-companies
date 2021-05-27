.PHONY: build
build:
		go build -o ./cmd/job/bergen-tech -v ./cmd/job
		./cmd/job/bergen-tech

.PHONY: run
run:
		docker-compose up --remove-orphans --build 

.PHONY: docker-build
docker-build:
		docker build -t bergen-tech:latest -f Dockerfile.bergen-tech .

.PHONY: docker-run
docker-run:
		docker run -p 127.0.0.1:9090:9090/tcp bergen-tech:latest

.PHONY: docker-build-run
docker-build-run:
		docker build -t bergen-tech:latest -f Dockerfile.bergen-tech .
		docker run -p 127.0.0.1:9090:9090/tcp bergen-tech:latest

.PHONY: test
test:
#go test -v -race -timeout 30s ./...
		go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build