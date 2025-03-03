package models

import (
	"github.com/google/uuid"
)

// Inventory представляет собой запись инвентаря
type Inventory struct {
	ID          uuid.UUID `json:"id" db:"id"`                     // Уникальный идентификатор записи
	ProductID   uuid.UUID `json:"product_id" db:"product_id"`     // Идентификатор продукта
	WarehouseID uuid.UUID `json:"warehouse_id" db:"warehouse_id"` // Идентификатор склада
	Quantity    int       `json:"quantity" db:"quantity"`         // Количество на складе
	Price       float64   `json:"price" db:"price"`               // Цена продукта
	Discount    float64   `json:"discount" db:"discount"`         // Скидка на продукт
}

// NewInventory создает новый экземпляр Inventory с уникальным ID
func NewInventory(productID, warehouseID uuid.UUID, quantity int, price, discount float64) *Inventory {
	return &Inventory{
		ID:          uuid.New(),
		ProductID:   productID,
		WarehouseID: warehouseID,
		Quantity:    quantity,
		Price:       price,
		Discount:    discount,
	}
}
