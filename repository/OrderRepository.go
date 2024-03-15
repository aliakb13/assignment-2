package repository

import (
	"assignment2/model"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) GetAll() ([]model.Order, error) {
	orders := []model.Order{}
	tx := or.db.Preload("Items").Find(&orders)
	return orders, tx.Error
}

func (or *orderRepository) CreateOrder(order model.Order) (model.Order, error) {
	tx := or.db.Create(&order)
	return order, tx.Error
}

func (or *orderRepository) UpdateOrder(id string, order model.Order) (model.Order, error) {
	err := or.db.Model(&model.Item{}).Where("order_id = ?", id).Delete(&model.Item{}).Error

	if err != nil {
		return model.Order{}, err
	}

	err = or.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error
	return order, err
}

func (or *orderRepository) DeleteOrder(id string) error {
	order := model.Order{}
	tx := or.db.Where("order_id = ?", id).Delete(&order)
	return tx.Error
}
