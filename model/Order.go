package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	OrderID      string    `json:"order_id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name"`
	OrderAt      time.Time `json:"order_at"`
	Items        []Item    `json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.OrderID = uuid.NewString()
	o.OrderAt = time.Now()
	return nil
}

func (o *Order) BeforeUpdate(tx *gorm.DB) error {
	o.OrderAt = time.Now()
	return nil
}
