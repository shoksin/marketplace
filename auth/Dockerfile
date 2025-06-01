FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o build ./cmd;


FROM alpine:3.19

#add timezone for correct work of time.Now() and create appuser
RUN apk --no-cache add ca-certificates tzdata && adduser -D -s /bin/sh appuser

WORKDIR /app

# app + env + migrations
COPY --from=builder /app/build   /app/program
COPY --from=builder /app/.env       /app/.env
COPY --from=builder /app/migrations /app/migrations

RUN chmod +x /app/program && chown -R appuser:appuser /app

USER appuser

CMD ["/app/program"]