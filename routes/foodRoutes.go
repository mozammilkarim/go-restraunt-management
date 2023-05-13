package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/syed/go-restraunt-management/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/food/food_id", controllers.GetFood())
	incomingRoutes.POST("/food", controllers.CreateFood())
	incomingRoutes.PATCH("/food/food_id", controllers.UpdateFood())
}
