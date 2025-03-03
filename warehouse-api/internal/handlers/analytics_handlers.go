package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// AnalyticsRecord представляет собой запись аналитики по складам
type AnalyticsRecord struct {
	ProductID   uuid.UUID `json:"product_id"`
	TotalSold   int       `json:"total_sold"`
	TotalAmount float64   `json:"total_amount"`
}

// TopWarehouseRecord представляет собой запись топ-складов
type TopWarehouseRecord struct {
	WarehouseID  uuid.UUID `json:"warehouse_id"`
	TotalRevenue float64   `json:"total_revenue"`
}

// GetAnalyticsByWarehouse возвращает аналитику по складам
func GetAnalyticsByWarehouse(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		analytics, err := fetchAnalytics(db, r.Context())
		if err != nil {
			http.Error(w, "Error fetching analytics: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(analytics); err != nil {
			http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// fetchAnalytics выполняет запрос к базе данных и возвращает аналитику
func fetchAnalytics(db *pgx.Conn, ctx context.Context) ([]AnalyticsRecord, error) {
	rows, err := db.Query(ctx, `
		SELECT product_id, SUM(sold_quantity) AS total_sold, SUM(total_amount) AS total_amount 
		FROM analytics 
		GROUP BY product_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var analytics []AnalyticsRecord
	for rows.Next() {
		var record AnalyticsRecord
		if err := rows.Scan(&record.ProductID, &record.TotalSold, &record.TotalAmount); err != nil {
			return nil, err
		}
		analytics = append(analytics, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return analytics, nil
}

// GetTopWarehouses возвращает топ-10 складов по выручке
func GetTopWarehouses(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		topWarehouses, err := fetchTopWarehouses(db, r.Context())
		if err != nil {
			http.Error(w, "Error fetching top warehouses: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(topWarehouses); err != nil {
			http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// fetchTopWarehouses выполняет запрос к базе данных и возвращает топ-склады
func fetchTopWarehouses(db *pgx.Conn, ctx context.Context) ([]TopWarehouseRecord, error) {
	rows, err := db.Query(ctx, `
		SELECT warehouse_id, SUM(total_amount) AS total_revenue 
		FROM analytics 
		GROUP BY warehouse_id 
		ORDER BY total_revenue DESC 
		LIMIT 10
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topWarehouses []TopWarehouseRecord
	for rows.Next() {
		var record TopWarehouseRecord
		if err := rows.Scan(&record.WarehouseID, &record.TotalRevenue); err != nil {
			return nil, err
		}
		topWarehouses = append(topWarehouses, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return topWarehouses, nil
}
