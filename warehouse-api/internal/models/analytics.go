package models

import (
	"github.com/google/uuid"
)

// Analytics представляет собой запись аналитики по складам
type Analytics struct {
	ID           uuid.UUID `json:"id" db:"id"`                       // Уникальный идентификатор записи
	WarehouseID  uuid.UUID `json:"warehouse_id" db:"warehouse_id"`   // Идентификатор склада
	ProductID    uuid.UUID `json:"product_id" db:"product_id"`       // Идентификатор продукта
	SoldQuantity int       `json:"sold_quantity" db:"sold_quantity"` // Количество проданных единиц
	TotalAmount  float64   `json:"total_amount" db:"total_amount"`   // Общая сумма продаж
}

// NewAnalytics создает новый экземпляр Analytics с уникальным ID
func NewAnalytics(warehouseID, productID uuid.UUID, soldQuantity int, totalAmount float64) *Analytics {
	return &Analytics{
		ID:           uuid.New(),
		WarehouseID:  warehouseID,
		ProductID:    productID,
		SoldQuantity: soldQuantity,
		TotalAmount:  totalAmount,
	}
}
