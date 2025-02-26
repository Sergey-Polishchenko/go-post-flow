FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /server ./cmd/server/server.go

FROM alpine:3.18

WORKDIR /

COPY .env .
COPY --from=builder /server /server

EXPOSE 8080
ENTRYPOINT ["/server"]
