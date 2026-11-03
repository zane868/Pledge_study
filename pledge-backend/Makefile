.PHONY: build test api sync chaincheck bindings
build:
	mkdir -p bin
	go build -o bin/pledge-api ./api
	go build -o bin/pledge-sync ./schedule
test:
	go test ./...
api:
	go run ./api
sync:
	go run ./schedule
chaincheck:
	go run ./cmd/chaincheck
bindings:
	python3 tools/generate-bindings.py
