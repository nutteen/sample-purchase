package service

import (
	"context"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/nutteen/sample-purchase/pkg/domain"
	"github.com/nutteen/sample-purchase/pkg/model"
	"github.com/nutteen/sample-purchase/pkg/repository"
	"github.com/shopspring/decimal"
)

type PurchaseOrderService struct {
	purchaseOrderRepository repository.PurchaseOrderRepository
}

func NewPurchaseOrderService(purchaseOrderRepository repository.PurchaseOrderRepository) *PurchaseOrderService{
	return &PurchaseOrderService{purchaseOrderRepository: purchaseOrderRepository}
}

func (service PurchaseOrderService) CreatePurchaseOrder(ctx context.Context, request *model.CreatePurchaseOrderRequest) (*model.CreatePurchaseOrderResponse, error) {
	var items []domain.PurchaseOrderItem
	for _,itemDto := range request.PurchaseOrderItems {
		itemEntity := domain.NewPurchaseOrderItem(utils.UUID(), itemDto.ItemID, itemDto.ItemPrice, itemDto.Quantity, itemDto.ItemPrice.Mul(decimal.NewFromInt32(int32(itemDto.Quantity))))
		items = append(items, *itemEntity)
	}
	purchaseOrder := domain.NewPurchaseOrder(utils.UUID(), items)

	err := service.purchaseOrderRepository.Create(purchaseOrder)

	if err != nil {
		return nil, err
	}

	response := &model.CreatePurchaseOrderResponse{
		PurchaseOrderId: purchaseOrder.ID,
	}

	return response, nil
}