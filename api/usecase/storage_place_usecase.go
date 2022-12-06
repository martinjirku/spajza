package usecase

import (
	"github.com/martinjirku/zasobar/entity"
)

type StoragePlaceCruder interface {
	Create(storagePlace entity.StoragePlace) (entity.StoragePlace, error)
	Get(storagePlaceId uint) (entity.StoragePlace, error)
	List() ([]entity.StoragePlace, error)
	Update(storagePlace entity.StoragePlace) (entity.StoragePlace, error)
	Delete(storagePlaceId uint) error
}

type StoragePlaceUsecase struct {
	storagePlaceRepository StoragePlaceCruder
}

func NewStoragePlaceUsecase(storagePlaceRepository StoragePlaceCruder) *StoragePlaceUsecase {
	return &StoragePlaceUsecase{storagePlaceRepository}
}

func (s *StoragePlaceUsecase) Create(storagePlace entity.StoragePlace) (entity.StoragePlace, error) {
	return s.storagePlaceRepository.Create(storagePlace)
}

func (s *StoragePlaceUsecase) Get(storagePlaceId uint) (entity.StoragePlace, error) {
	return s.storagePlaceRepository.Get(storagePlaceId)
}

func (s *StoragePlaceUsecase) List() ([]entity.StoragePlace, error) {
	return s.storagePlaceRepository.List()
}

func (s *StoragePlaceUsecase) Update(storagePlace entity.StoragePlace) (entity.StoragePlace, error) {
	return s.storagePlaceRepository.Update(storagePlace)
}

func (s *StoragePlaceUsecase) Delete(storagePlaceId uint) error {
	return s.storagePlaceRepository.Delete(storagePlaceId)
}
