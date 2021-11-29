package service

import "go_trainee/models"

type Orders interface {
	GetOrderByUID(string) *models.Order
	SaveOrderCache(*models.Order) error
	SaveOrderData(*models.Order) error
	LoadAllOrders() error
	UpdateCache([]byte)
	GetSaveOrderByUID(string) *models.OrderSave
	GetAllUID() []string
}

type Service interface {
	Orders
}
