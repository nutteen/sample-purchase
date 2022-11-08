package model

import "github.com/shopspring/decimal"

type GetPurchaseOrderResponse struct {
	ID	string `json:"id"`
	TotalAmount decimal.Decimal `json:"totalAmount"`
	PurchaseOrderItems []PurchaseOrderItemDto `json:"purchaseOrderItems"`
}

type PurchaseOrderItemDto struct {
	ID string	`json:"id"`
	PurchaseOrderID	string	`json:"purchaseOrderId"`
	ItemID		string `json:"itemId"`
	ItemPrice	decimal.Decimal `json:"itemPrice"`
	Quantity	int	`json:"quantity"`
	Amount		decimal.Decimal	`json:"amount"`
}