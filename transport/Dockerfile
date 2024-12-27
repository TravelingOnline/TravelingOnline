FROM docker.arvancloud.ir/golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN GO111MODULE=on GOPROXY=https://goproxy.cn,direct go mod download
COPY . .
RUN GO111MODULE=on GOPROXY=https://goproxy.cn,direct go build -o vehicle ./cmd/vehicle/main.go

FROM docker.arvancloud.ir/alpine:latest
WORKDIR /app/
COPY --from=builder /app/vehicle .
COPY --from=builder /app/config.yml .
EXPOSE 8081
CMD ["./vehicle"]