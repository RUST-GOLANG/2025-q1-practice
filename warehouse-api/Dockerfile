# Этап сборки
FROM golang:1.21 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Этап запуска
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированное приложение из предыдущего этапа
COPY --from=builder /app/main .

# Указываем команду для запуска
CMD ["./main"]

# Открываем порт, если приложение его использует
EXPOSE 8080