package domain

import "github.com/shopspring/decimal"

type PurchaseOrder struct {
	ID	string `gorm:"primaryKey"`
	TotalAmount	decimal.Decimal
	Items	[]PurchaseOrderItem `gorm:"foreignKey:PurchaseOrderID"`
}

type PurchaseOrderItem struct {
	ID string	`gorm:"primaryKey"`
	PurchaseOrderID	string
	ItemID		string
	ItemPrice	decimal.Decimal
	Quantity	int
	Amount		decimal.Decimal
}

func NewPurchaseOrder(id string, items []PurchaseOrderItem) *PurchaseOrder {
	totalAmount := decimal.Decimal{}
	for _, item := range items {
		item.PurchaseOrderID = id
		totalAmount = totalAmount.Add(item.Amount)
	}
	return &PurchaseOrder{
		ID: id,
		TotalAmount: totalAmount,
		Items: items,
	}
}

func NewPurchaseOrderItem(id string, itemId string, itemPrice decimal.Decimal, quantity int, amount decimal.Decimal) *PurchaseOrderItem {
	return &PurchaseOrderItem{
		ID: id,
		ItemID: itemId,
		ItemPrice: itemPrice,
		Quantity: quantity,
		Amount: amount,
	}
}