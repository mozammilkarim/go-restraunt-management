package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/syed/go-restraunt-management/database"
	"github.com/syed/go-restraunt-management/middleware"
	"github.com/syed/go-restraunt-management/routes"
	"go.mongodb.org/mongo-driver/mongo"
)
var foodCollection *mongo.Collection:= database.OpenCollection(database.client,"food")

func main(){
	port:= os.Getenv("PORT")
	if( port==""){
		port="8000"
	}
	router:=gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())
	router.UserRoutes(router)
	router.FoodRoutes(router)
	router.TableRoutes(router)
	router.OrderRoutes(router)
	router.OrderItemRoutes(router)
	router.InvoiceRoutes(router)


	log.Fatal(router.Run(":"+port))
}