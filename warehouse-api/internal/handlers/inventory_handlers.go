package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RUST-GOLANG/2025-q1-practice.git/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// CreateInventory создает новую запись инвентаря
func CreateInventory(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inventory models.Inventory
		if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
			http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Валидация входных данных
		if inventory.Quantity < 0 {
			http.Error(w, "Quantity cannot be negative", http.StatusBadRequest)
			return
		}

		inventory.ID = uuid.New()

		result, err := db.Exec(r.Context(), "INSERT INTO inventory (id, product_id, warehouse_id, quantity, price, discount) VALUES ($1, $2, $3, $4, $5, $6)",
			inventory.ID, inventory.ProductID, inventory.WarehouseID, inventory.Quantity, inventory.Price, inventory.Discount)
		if err != nil {
			http.Error(w, "Error inserting inventory record: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Проверка, была ли вставлена хотя бы одна запись
		if result.RowsAffected() == 0 {
			http.Error(w, "No rows affected", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(inventory); err != nil {
			http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// UpdateInventoryQuantity обновляет количество инвентаря
func UpdateInventoryQuantity(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inventory models.Inventory
		if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
			http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Валидация входных данных
		if inventory.Quantity < 0 {
			http.Error(w, "Quantity cannot be negative", http.StatusBadRequest)
			return
		}

		result, err := db.Exec(r.Context(), "UPDATE inventory SET quantity = quantity + $1 WHERE product_id = $2 AND warehouse_id = $3",
			inventory.Quantity, inventory.ProductID, inventory.WarehouseID)
		if err != nil {
			http.Error(w, "Error updating inventory: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Проверка, была ли обновлена хотя бы одна запись
		if result.RowsAffected() == 0 {
			http.Error(w, "Inventory record not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
