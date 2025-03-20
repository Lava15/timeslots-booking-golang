build:
	go build -o cmd/server/main.go
run:
	go run cmd/server/main.go
test:
	go test -v ./...
air-install:
	@command -v air >/dev/null 2>&1 || { echo "Installing air..."; go install github.com/air-verse/air@latest; }
