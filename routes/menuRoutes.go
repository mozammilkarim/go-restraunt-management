package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/syed/go-restraunt-management/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controllers.GetMenus())
	incomingRoutes.GET("/menu/menu_id", controllers.GetMenu())
	incomingRoutes.POST("/menu", controllers.CreateMenu())
	incomingRoutes.PATCH("/menu/menu_id", controllers.UpdateMenu())
}
