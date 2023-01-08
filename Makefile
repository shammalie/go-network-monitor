run:
	go run ./cmd/main.go

run-race:
	go run -race ./cmd/main.go

build-binary:
	go build -o go-network-service ./cmd/main.go

build-docker:
	docker build -f build/docker/Dockerfile -t shammalie/go-network-service:0.0.1 .