package usecases

import (
	"context"

	"github.com/martinjirku/zasobar/entity"
)

type StorageItemRepository interface {
	Create(ctx context.Context, storageItem entity.NewStorageItem) (entity.StorageItem, error)
	GetStorageItemById(ctx context.Context, storageItemId uint) (entity.StorageItem, error)
	GetStorageConsumptionById(ctx context.Context, storageItemId uint) ([]entity.StorageItemConsumption, error)
	AddStorageConsumption(ctx context.Context, storageConsumption entity.StorageItemConsumption) (entity.StorageItemConsumption, error)
	UpdateColumn(ctx context.Context, id uint, fieldName string, fieldValue interface{}) error
	List(ctx context.Context) ([]entity.StorageItem, error)
}

type StorageItemService struct {
	storageItemRepository StorageItemRepository
}

func NewStorageItemService(storageItemRepository StorageItemRepository) *StorageItemService {
	return &StorageItemService{storageItemRepository}
}

func (s *StorageItemService) Create(ctx context.Context, storageItem entity.NewStorageItem) (entity.StorageItem, error) {
	return s.storageItemRepository.Create(ctx, storageItem)
}

func (s *StorageItemService) UpdateField(ctx context.Context, storageItemId uint, fieldName string, value interface{}) error {
	return s.storageItemRepository.UpdateColumn(ctx, storageItemId, fieldName, value)
}

func (s *StorageItemService) Consumpt(ctx context.Context, storageItemId uint, amount float64, unit string) (entity.StorageItem, error) {
	storageItem, err := entity.LoadStorageItem(ctx, storageItemId, s.storageItemRepository)
	if err != nil {
		return storageItem, err
	}
	err = storageItem.Consumpt(amount, unit)
	if err != nil {
		return storageItem, err
	}
	idx := len(storageItem.Consumptions) - 1
	consumption, err := s.storageItemRepository.AddStorageConsumption(ctx, storageItem.Consumptions[idx])
	if err != nil {
		return storageItem, err
	}
	storageItem.Consumptions[idx].StorageItemConsumptionId = consumption.StorageItemId
	err = s.storageItemRepository.UpdateColumn(ctx, storageItemId, "currentAmount", storageItem.CurrentAmount)
	if err != nil {
		return storageItem, err
	}
	return storageItem, nil
}

func (s *StorageItemService) List(ctx context.Context) ([]entity.StorageItem, error) {
	result, err := s.storageItemRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
