package gormrepository

import (
	"github.com/nutteen/sample-purchase/pkg/domain"
	"gorm.io/gorm"
)

type GormPurchaseOrderRepository struct {
	db *gorm.DB
}

func NewGormPurchaseOrderRepository(gormDB *gorm.DB) *GormPurchaseOrderRepository {
	return &GormPurchaseOrderRepository{db: gormDB}
}

func (repo GormPurchaseOrderRepository) GetById(id string) (*domain.PurchaseOrder, error) {
	purchaseOrder := new(domain.PurchaseOrder)
	result := repo.db.First(purchaseOrder, id)
	return purchaseOrder, result.Error
}

func (repo GormPurchaseOrderRepository) Create(purchaseOrder *domain.PurchaseOrder) error {
	result := repo.db.Create(purchaseOrder)
	return result.Error
}