package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RUST-GOLANG/2025-q1-practice.git/config"
	"github.com/RUST-GOLANG/2025-q1-practice.git/internal/routes"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5" // Импортируем pgx для работы с PostgreSQL
)

func main() {
	cfg := config.LoadConfig()

	// Замените 'your_password' и 'your_database' на ваши реальные данные
	databaseURL := "user=postgres password=Ishtaev73 dbname=postgres sslmode=disable"

	// Подключение к базе данных
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background()) // Передаем контекст в Close()

	fmt.Println("Successfully connected to the database!")

	// Регистрация маршрутов
	r := mux.NewRouter()
	routes.RegisterRoutes(r, conn)

	srv := &http.Server{
		Handler: r,
		Addr:    cfg.ServerAddress,
	}

	go func() {
		log.Println("Starting server on", cfg.ServerAddress)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
