init:
	go mod download

test:
	go test -race ./pkg/...

lint:
	staticcheck ./...

build:
	go build -o ./dist/main ./cmd/*.go

start: build
	./dist/main
