package routes

import (
	"foodos-backend/internal/modules/order"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// ---- Order module wiring ----
	orderRepo := order.NewInMemoryRepository()
	orderService := order.NewService(orderRepo, orderRepo)
	orderHandler := order.NewHandler(orderService)


	orders := router.Group("/orders")
	{
		orders.POST("", orderHandler.CreateOrder)
		orders.PATCH("/:id/status", orderHandler.ChangeOrderStatus)
	}
}
