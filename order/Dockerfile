FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o out/dist ./cmd;

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/out/dist /app/dist

RUN chmod +x /app/dist

CMD ["/app/dist"]
