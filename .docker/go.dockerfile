FROM golang:1.24-alpine
WORKDIR /app

ENV GOPROXY=direct
ENV GO111MODULE=on

RUN apk update && \
  apk upgrade && \
  apk add --no-cache make wget git bash && \
  wget -O /wait-for https://raw.githubusercontent.com/eficode/wait-for/master/wait-for && \
  chmod +x /wait-for && \
  rm -rf /var/cache/apk/*

COPY . .

ADD https://github.com/pressly/goose/releases/download/v3.24.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

ARG UID=1000
ARG GID=1000
RUN addgroup -g $GID -S appgroup && adduser -u $UID -S appuser -G appgroup
USER appuser

CMD ["sh", "-c", "/wait-for db:5432 -- goose -dir /app/migrations postgres \"postgres://postgres:postgres@db:5432/booking?sslmode=disable\" up && go run ./cmd/server/main.go"]
