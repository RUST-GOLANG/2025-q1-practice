package models

import (
	"github.com/google/uuid"
)

// Warehouse представляет собой склад с его адресом
type Warehouse struct {
	ID      uuid.UUID `json:"id" db:"id"`           // Уникальный идентификатор склада
	Address string    `json:"address" db:"address"` // Адрес склада
}

// NewWarehouse создает новый экземпляр Warehouse с уникальным ID
func NewWarehouse(address string) *Warehouse {
	return &Warehouse{
		ID:      uuid.New(),
		Address: address,
	}
}
