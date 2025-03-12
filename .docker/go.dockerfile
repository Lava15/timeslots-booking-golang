FROM golang:1.24-alpine
WORKDIR /app

RUN apk add --no-cache wget && \
  wget -O /wait-for https://raw.githubusercontent.com/eficode/wait-for/master/wait-for && \
  chmod +x /wait-for

COPY . .

CMD ["sh", "-c", "/wait-for db:5432 -- go run ./cmd/server/main.go"]
