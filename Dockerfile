FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download

# Собираем приложение
RUN go build -o test-medods cmd/main.go

# Запускаем приложение
CMD ["./test-medods"]