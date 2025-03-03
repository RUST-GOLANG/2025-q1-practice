package models

import (
	"github.com/google/uuid"
)

// Product представляет собой продукт с его характеристиками
type Product struct {
	ID              uuid.UUID              `json:"id" db:"id"`                           // Уникальный идентификатор продукта
	Name            string                 `json:"name" db:"name"`                       // Название продукта
	Description     string                 `json:"description" db:"description"`         // Описание продукта
	Characteristics map[string]interface{} `json:"characteristics" db:"characteristics"` // Характеристики продукта
	Weight          float64                `json:"weight" db:"weight"`                   // Вес продукта
	Barcode         string                 `json:"barcode" db:"barcode"`                 // Штрих-код продукта
}

// NewProduct создает новый экземпляр Product с уникальным ID
func NewProduct(name, description string, characteristics map[string]interface{}, weight float64, barcode string) *Product {
	return &Product{
		ID:              uuid.New(),
		Name:            name,
		Description:     description,
		Characteristics: characteristics,
		Weight:          weight,
		Barcode:         barcode,
	}
}
