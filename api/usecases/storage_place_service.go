package usecases

import (
	"context"

	"github.com/martinjirku/zasobar/domain"
)

type StoragePlaceRepository interface {
	Create(ctx context.Context, storagePlace domain.StoragePlace) (domain.StoragePlace, error)
	Get(ctx context.Context, storagePlaceId uint) (domain.StoragePlace, error)
	List(ctx context.Context) ([]domain.StoragePlace, error)
	Update(ctx context.Context, storagePlace domain.StoragePlace) (domain.StoragePlace, error)
	Delete(ctx context.Context, storagePlaceId uint) error
}

type StoragePlaceService struct {
	storagePlaceRepository StoragePlaceRepository
}

func NewStoragePlaceService(storagePlaceRepository StoragePlaceRepository) *StoragePlaceService {
	return &StoragePlaceService{storagePlaceRepository}
}

func (s *StoragePlaceService) Create(ctx context.Context, storagePlace domain.StoragePlace) (domain.StoragePlace, error) {
	return s.storagePlaceRepository.Create(ctx, storagePlace)
}

func (s *StoragePlaceService) Get(ctx context.Context, storagePlaceId uint) (domain.StoragePlace, error) {
	return s.storagePlaceRepository.Get(ctx, storagePlaceId)
}

func (s *StoragePlaceService) List(ctx context.Context) ([]domain.StoragePlace, error) {
	return s.storagePlaceRepository.List(ctx)
}

func (s *StoragePlaceService) Update(ctx context.Context, storagePlace domain.StoragePlace) (domain.StoragePlace, error) {
	return s.storagePlaceRepository.Update(ctx, storagePlace)
}

func (s *StoragePlaceService) Delete(ctx context.Context, storagePlaceId uint) error {
	return s.storagePlaceRepository.Delete(ctx, storagePlaceId)
}
