package usecases

import (
	"context"

	"github.com/martinjirku/zasobar/domain"
)

type StorageItemRepository interface {
	Create(ctx context.Context, storageItem domain.NewStorageItem) (domain.StorageItem, error)
	GetStorageItemById(ctx context.Context, storageItemId uint) (domain.StorageItem, error)
	GetStorageConsumptionById(ctx context.Context, storageItemId uint) ([]domain.StorageItemConsumption, error)
	AddStorageConsumption(ctx context.Context, storageConsumption domain.StorageItemConsumption) (domain.StorageItemConsumption, error)
	List(ctx context.Context) ([]domain.StorageItem, error)
}

type StorageItemService struct {
	storageItemRepository StorageItemRepository
}

func NewStorageItemService(storageItemRepository StorageItemRepository) *StorageItemService {
	return &StorageItemService{storageItemRepository}
}

func (s *StorageItemService) Create(ctx context.Context, storageItem domain.NewStorageItem) (domain.StorageItem, error) {
	return s.storageItemRepository.Create(ctx, storageItem)
}

func (s *StorageItemService) Consumpt(ctx context.Context, storageItemId uint, amount float64, unit string) (domain.StorageItem, error) {
	storageItem, err := domain.LoadStorageItem(ctx, storageItemId, s.storageItemRepository)
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
	return storageItem, nil
}

func (s *StorageItemService) List(ctx context.Context) ([]domain.StorageItem, error) {
	result, err := s.storageItemRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
