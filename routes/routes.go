package routes

import (
	"go_trainee/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{service: s}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With")
	corsConfig.AddAllowMethods("GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS")

	r.Use(
		cors.New(corsConfig),
	)

	api := r.Group("/orders")
	{
		api.GET("", handler.getOrders)
		api.GET("/order", handler.getOrderByUID)
	}

	return r
}
