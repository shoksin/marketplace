FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.24.3
RUN go build -o out/dist ./cmd;

FROM alpine:3.19

WORKDIR /app

# бинарник + env + migtations
COPY --from=builder /app/out/dist /app/dist
COPY --from=builder /app/.env /app/.env
COPY --from=builder /app/migrations /app/migrations
# goose
COPY --from=builder /go/bin/goose /usr/local/bin/goose

RUN chmod +x /app/dist /usr/local/bin/goose

CMD ["sh", "-c", "goose -dir ./migrations postgres \"$DATABASE_CONFIG\" up && /app/dist" ]