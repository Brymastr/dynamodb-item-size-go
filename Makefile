init:
	go mod download

test:
	go test -race ./pkg/...

lint:
	staticcheck ./...
