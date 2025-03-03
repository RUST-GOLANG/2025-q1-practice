package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RUST-GOLANG/2025-q1-practice.git/config"
	"github.com/RUST-GOLANG/2025-q1-practice.git/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	// Регистрация маршрутов
	r := mux.NewRouter()
	routes.RegisterRoutes(r, nil) // Передаем nil вместо соединения с базой данных

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

	log.Println("Server exiting")
}
