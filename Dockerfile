# --------------------
# BUILD STAGE
# --------------------
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# copy dependency files trước để cache
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api/main.go

# --------------------
# RUNTIME STAGE
# --------------------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

# Không chép .env vào image: biến môi trường được truyền lúc chạy (docker compose env_file / môi trường triển khai).

EXPOSE 8080

CMD ["./app"]
