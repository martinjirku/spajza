package usecases

import (
	"context"

	"github.com/martinjirku/zasobar/entity"
)

type StoragePlaceRepository interface {
	Create(ctx context.Context, storagePlace entity.StoragePlace) (entity.StoragePlace, error)
	Get(ctx context.Context, storagePlaceId uint) (entity.StoragePlace, error)
	List(ctx context.Context) ([]entity.StoragePlace, error)
	Update(ctx context.Context, storagePlace entity.StoragePlace) (entity.StoragePlace, error)
	Delete(ctx context.Context, storagePlaceId uint) error
}

type StoragePlaceService struct {
	storagePlaceRepository StoragePlaceRepository
}

func NewStoragePlaceService(storagePlaceRepository StoragePlaceRepository) *StoragePlaceService {
	return &StoragePlaceService{storagePlaceRepository}
}

func (s *StoragePlaceService) Create(ctx context.Context, storagePlace entity.StoragePlace) (entity.StoragePlace, error) {
	return s.storagePlaceRepository.Create(ctx, storagePlace)
}

func (s *StoragePlaceService) Get(ctx context.Context, storagePlaceId uint) (entity.StoragePlace, error) {
	return s.storagePlaceRepository.Get(ctx, storagePlaceId)
}

func (s *StoragePlaceService) List(ctx context.Context) ([]entity.StoragePlace, error) {
	return s.storagePlaceRepository.List(ctx)
}

func (s *StoragePlaceService) Update(ctx context.Context, storagePlace entity.StoragePlace) (entity.StoragePlace, error) {
	return s.storagePlaceRepository.Update(ctx, storagePlace)
}

func (s *StoragePlaceService) Delete(ctx context.Context, storagePlaceId uint) error {
	return s.storagePlaceRepository.Delete(ctx, storagePlaceId)
}
