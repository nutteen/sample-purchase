package repository

import "github.com/nutteen/sample-purchase/pkg/domain"

type PurchaseOrderRepository interface {
	GetById(id string) (*domain.PurchaseOrder, error)
	Create(purchaseOrder *domain.PurchaseOrder) error
}