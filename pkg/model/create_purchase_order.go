package model

import "github.com/shopspring/decimal"

type CreatePurchaseOrderRequest struct {
	PurchaseOrderItems []CreatePurchaseOrderItem `json:"purchaseOrderItems" validate:"required,min=1"`
}

type CreatePurchaseOrderItem struct {
	ItemID	string `json:"itemId" validate:"required"`
	ItemPrice	decimal.Decimal `json:"itemPrice" validate:"required"`
	Quantity	int `json:"quantity" validate:"required"`
}

type CreatePurchaseOrderResponse struct {
	PurchaseOrderId	string `json:"purchaseOrderId"`
}