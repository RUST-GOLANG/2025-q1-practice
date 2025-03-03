package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RUST-GOLANG/2025-q1-practice.git/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// CreateProduct создает новый продукт
func CreateProduct(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
			return
		}
		product.ID = uuid.New()

		_, err := db.Exec(r.Context(), "INSERT INTO products (id, name, description, characteristics, weight, barcode) VALUES ($1, $2, $3, $4, $5, $6)",
			product.ID, product.Name, product.Description, product.Characteristics, product.Weight, product.Barcode)
		if err != nil {
			http.Error(w, "Error inserting product: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// GetProducts возвращает список всех продуктов
func GetProducts(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(r.Context(), "SELECT id, name, description, characteristics, weight, barcode FROM products")
		if err != nil {
			http.Error(w, "Error fetching products: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []models.Product
		for rows.Next() {
			var product models.Product
			if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Characteristics, &product.Weight, &product.Barcode); err != nil {
				http.Error(w, "Error scanning product: "+err.Error(), http.StatusInternalServerError)
				return
			}
			products = append(products, product)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Error iterating over products: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
