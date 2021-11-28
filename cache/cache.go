package cache

import "go_trainee/models"

type Orders interface {
	GetOrderByUID(string) *models.Order
	SaveOrder(*models.Order) error
	Change(map[string]*models.Order)
	GetAllOrders() []*models.Order
}

type DatabaseCache struct {
	Orders
	order map[string]*models.Order
}
