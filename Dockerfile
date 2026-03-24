# --------------------
# BUILD STAGE
# --------------------
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# copy dependency files trước để cache
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# --------------------
# RUNTIME STAGE
# --------------------
FROM alpine:latest

WORKDIR /app

# copy binary từ build stage
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]