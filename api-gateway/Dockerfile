FROM golang:1.23.6 AS build

WORKDIR /app

COPY go.mod go.sum ./
# RUN go mod download
RUN go env -w GOPROXY=https://goproxy.io,direct && go mod download
# т.к была такая ошибка
# go: github.com/gabriel-vasile/mimetype@v1.4.3: Get "https://storage.googleapis.com/proxy-golang-org-prod/...": net/http: TLS handshake timeout

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o out/dist ./cmd;

FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/out/dist /app/dist

RUN chmod +x /app/dist

CMD ["/app/dist"]