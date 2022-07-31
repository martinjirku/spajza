package usecases

import (
	"context"

	"github.com/martinjirku/zasobar/domain"
)

type StorageItemRepository interface {
	Create(ctx context.Context, storageItem domain.NewStorageItemRequest) (domain.StorageItemResponse, error)
}

type StorageItemService struct {
	StorageItemRepository StorageItemRepository
}

func NewStorageItemService(storageItemRepository StorageItemRepository) *StorageItemService {
	return &StorageItemService{storageItemRepository}
}

func (s *StorageItemService) Create(ctx context.Context, storageItem domain.NewStorageItemRequest) (domain.StorageItemResponse, error) {
	return s.StorageItemRepository.Create(ctx, storageItem)
}
