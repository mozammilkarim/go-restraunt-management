package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/syed/go-restraunt-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItem/orderItem_id", controllers.GetOrderitem())
	incomingRoutes.GET("/orderItems-order/order_id", controllers.ItemsByOrderId())
	incomingRoutes.POST("/order", controllers.CreateOrderitem())
	incomingRoutes.PATCH("/order/orderItem_id", controllers.UpdateOrderitem())
}
