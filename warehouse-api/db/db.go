package db

import (
	"context"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
)

// Connect устанавливает соединение с базой данных PostgreSQL
func Connect(databaseURL string) *pgx.Conn {
	// Проверка и добавление sslmode=disable, если это необходимо
	if !containsSSLMode(databaseURL) {
		databaseURL += " sslmode=disable"
	}

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	return conn
}

// containsSSLMode проверяет, содержит ли строка подключения параметр sslmode
func containsSSLMode(databaseURL string) bool {
	return strings.Contains(databaseURL, "sslmode=")
}
