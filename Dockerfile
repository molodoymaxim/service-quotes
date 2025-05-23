# Сборка
FROM golang:1.22
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o quotes-service ./cmd/main.go

# Запуск
FROM alpine:3.20.3
WORKDIR /app
COPY --from=builder /app/service-quotes .
EXPOSE 8080
ENTRYPOINT ["./quotes-service"]