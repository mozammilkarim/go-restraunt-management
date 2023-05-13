package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/syed/go-restraunt-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/table/table_id", controllers.GetTable())
	incomingRoutes.POST("/table", controllers.CreateTable())
	incomingRoutes.PATCH("/table/table_id", controllers.UpdateTable())
}
