build:
	go build -o cmd/server/main.go
run:
	go run cmd/server/main.go
test:
	go test -v ./...
