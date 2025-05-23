FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Сборка приложения
RUN go build -o service_quote ./cmd/main.go

CMD ["/app/service_quote"]
