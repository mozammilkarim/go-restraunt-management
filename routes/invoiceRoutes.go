package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/syed/go-restraunt-management/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controllers.GetInvoices())
	incomingRoutes.GET("/invoice/invoice_id", controllers.GetInvoice())
	incomingRoutes.POST("/invoice", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoice/invoice_id", controllers.UpdateInvoice())
}
