FROM docker.arvancloud.ir/golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o transport ./cmd/transport/main.go


# FROM docker.arvancloud.ir/alpine:latest
FROM docker.arvancloud.ir/golang:1.23-alpine
WORKDIR /app/
# COPY --from=builder /app/ .
COPY --from=builder /app/transport .
COPY --from=builder /app/config.yml .
EXPOSE 8082
CMD ["./transport"]
# CMD ["sh"]