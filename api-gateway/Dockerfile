FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
# RUN go mod download
RUN go env -w GOPROXY=https://goproxy.io,direct && go mod download
# т.к была такая ошибка
# go: github.com/gabriel-vasile/mimetype@v1.4.3: Get "https://storage.googleapis.com/proxy-golang-org-prod/...": net/http: TLS handshake timeout

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o build ./cmd;

FROM alpine:3.19

#add timezone for correct work of time.Now() and create appuser
RUN apk --no-cache add ca-certificates tzdata && adduser -D -s /bin/sh appuser

WORKDIR /app

# бинарник + env
COPY --from=builder /app/build /app/program
COPY --from=builder /app/.env /app/.env

RUN chmod +x /app/program && chown -R appuser:appuser /app

USER appuser

CMD ["/app/program"]