package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/syed/go-restraunt-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/order/order_id", controllers.GetOrder())
	incomingRoutes.POST("/order", controllers.CreateOrder())
	incomingRoutes.PATCH("/order/order_id", controllers.UpdateOrder())
}
