package repository

import "assignment2/model"

type IOrderRepository interface {
	GetAll() ([]model.Order, error)
	CreateOrder(model.Order) (model.Order, error)
	UpdateOrder(string, model.Order) (model.Order, error)
	DeleteOrder(string) error
}
