package usecase

import (
	"github.com/martinjirku/zasobar/entity"
)

type StorageItemRepository interface {
	Create(storageItem entity.NewStorageItem) (entity.StorageItem, error)
	GetStorageItemById(storageItemId uint) (entity.StorageItem, error)
	GetStorageConsumptionById(storageItemId uint) ([]entity.StorageItemConsumption, error)
	AddStorageConsumption(storageConsumption entity.StorageItemConsumption) (entity.StorageItemConsumption, error)
	UpdateColumn(id uint, fieldName string, fieldValue interface{}) error
	List() ([]entity.StorageItem, error)
}

type StorageItemService struct {
	storageItemRepository StorageItemRepository
}

func NewStorageItemService(storageItemRepository StorageItemRepository) *StorageItemService {
	return &StorageItemService{storageItemRepository}
}

func (s *StorageItemService) Create(storageItem entity.NewStorageItem) (entity.StorageItem, error) {
	return s.storageItemRepository.Create(storageItem)
}

func (s *StorageItemService) UpdateField(storageItemId uint, fieldName string, value interface{}) error {
	return s.storageItemRepository.UpdateColumn(storageItemId, fieldName, value)
}

func (s *StorageItemService) Consumpt(storageItemId uint, amount float64, unit string) (entity.StorageItem, error) {
	storageItem, err := entity.LoadStorageItem(storageItemId, s.storageItemRepository)
	if err != nil {
		return storageItem, err
	}
	err = storageItem.Consumpt(amount, unit)
	if err != nil {
		return storageItem, err
	}
	idx := len(storageItem.Consumptions) - 1
	consumption, err := s.storageItemRepository.AddStorageConsumption(storageItem.Consumptions[idx])
	if err != nil {
		return storageItem, err
	}
	storageItem.Consumptions[idx].StorageItemConsumptionId = consumption.StorageItemId
	err = s.storageItemRepository.UpdateColumn(storageItemId, "currentAmount", storageItem.CurrentAmount)
	if err != nil {
		return storageItem, err
	}
	return storageItem, nil
}

func (s *StorageItemService) List() ([]entity.StorageItem, error) {
	result, err := s.storageItemRepository.List()
	if err != nil {
		return nil, err
	}
	return result, nil
}
