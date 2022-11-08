package router

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/nutteen/sample-purchase/pkg/api/v1"
	"github.com/nutteen/sample-purchase/pkg/service"
)

func RegisterRoutes(app *fiber.App, purchaseOrderService *service.PurchaseOrderService){
	v1PurchaseOrder := v1.NewPurchaseOrderHandler(purchaseOrderService)
	routes := app.Group("/purchase-orders")
	routes.Post("/", v1PurchaseOrder.CreatePurchaseOrder)
	routes.Get("/:id", v1PurchaseOrder.GetById)
}