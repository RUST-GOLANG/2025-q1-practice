package routes

import (
	"github.com/RUST-GOLANG/2025-q1-practice.git/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

// RegisterRoutes регистрирует маршруты для API и HTML
func RegisterRoutes(r *mux.Router, db *pgx.Conn) {
	// Маршруты для работы со складами
	r.HandleFunc("/api/warehouses", handlers.CreateWarehouse(db)).Methods("POST")
	r.HandleFunc("/api/warehouses", handlers.GetWarehouses(db)).Methods("GET")

	// Маршруты для работы с продуктами
	r.HandleFunc("/api/products", handlers.CreateProduct(db)).Methods("POST")
	r.HandleFunc("/api/products", handlers.GetProducts(db)).Methods("GET")

	// Маршруты для работы с инвентарем
	r.HandleFunc("/api/inventory", handlers.CreateInventory(db)).Methods("POST")
	r.HandleFunc("/api/inventory/update", handlers.UpdateInventoryQuantity(db)).Methods("PUT")

	// Маршруты для аналитики
	r.HandleFunc("/api/analytics/warehouse", handlers.GetAnalyticsByWarehouse(db)).Methods("GET")
	r.HandleFunc("/api/analytics/top-warehouses", handlers.GetTopWarehouses(db)).Methods("GET")

	// Маршрут для главной страницы
	r.HandleFunc("/", handlers.HTMLHandler).Methods("GET")
}
