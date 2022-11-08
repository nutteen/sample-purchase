package model

import "github.com/shopspring/decimal"

type CreatePurchaseOrderRequest struct {
	PurchaseOrderItems []CreatePurchaseOrderItem `json:"purchaseOrderItems"`
}

type CreatePurchaseOrderItem struct {
	ItemID	string `json:"itemId"`
	ItemPrice	decimal.Decimal `json:"itemPrice"`
	Quantity	int `json:"quantity"`
}

type CreatePurchaseOrderResponse struct {
	PurchaseOrderId	string `json:"purchaseOrderId"`
}