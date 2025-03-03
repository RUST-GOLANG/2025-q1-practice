package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RUST-GOLANG/2025-q1-practice.git/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// CreateWarehouse создает новый склад
func CreateWarehouse(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warehouse models.Warehouse
		if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
			http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
			return
		}
		warehouse.ID = uuid.New()

		_, err := db.Exec(r.Context(), "INSERT INTO warehouses (id, address) VALUES ($1, $2)", warehouse.ID, warehouse.Address)
		if err != nil {
			http.Error(w, "Error inserting warehouse: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(warehouse); err != nil {
			http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// GetWarehouses возвращает список всех складов
func GetWarehouses(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(r.Context(), "SELECT id, address FROM warehouses")
		if err != nil {
			http.Error(w, "Error fetching warehouses: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var warehouses []models.Warehouse
		for rows.Next() {
			var warehouse models.Warehouse
			if err := rows.Scan(&warehouse.ID, &warehouse.Address); err != nil {
				http.Error(w, "Error scanning warehouse: "+err.Error(), http.StatusInternalServerError)
				return
			}
			warehouses = append(warehouses, warehouse)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Error iterating over warehouses: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(warehouses); err != nil {
			http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
