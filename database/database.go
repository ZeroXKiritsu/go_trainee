package database

import "go_trainee/models"

type Orders interface {
	DatabaseSaveOrder(*models.Order) error
	GetAllOrders() ([]models.OrderSave, error)
	CheckOrder(*models.Order) (bool, error)
	RunService(string, func([]byte)) error
	Close()
}

type Database struct {
	Orders
}