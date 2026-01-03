package routes

import (
	"foodos-backend/internal/modules/order"

	"github.com/gin-gonic/gin"
	"foodos-backend/internal/config"

)

func RegisterRoutes(router *gin.Engine) {

	redisClient := config.NewRedisClient()
	publisher := order.NewRedisPublisher(redisClient)

	db := config.NewPostgresDB()

	orderRepo := order.NewPostgresRepository(db)
	orderService := order.NewService(
		orderRepo,
		orderRepo,
		publisher,
	)


	orderHandler := order.NewHandler(orderService)

	orders := router.Group("/orders")
	{
		orders.POST("", orderHandler.CreateOrder)
		orders.GET("/:id", orderHandler.GetOrder)
		orders.PATCH("/:id/status", orderHandler.ChangeOrderStatus)

	}
}
