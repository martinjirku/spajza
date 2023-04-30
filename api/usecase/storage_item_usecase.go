package usecase

import (
	"github.com/martinjirku/zasobar/entity"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . StorageItemRepository

type StorageItemRepository interface {
	Create(storageItem entity.StorageItem) (entity.StorageItem, error)
	ById(storageItemId uint) (entity.StorageItem, error)
	Update(storageItem entity.StorageItem) error
	List() ([]entity.StorageItem, error)
}

type StorageItemUsecase struct {
	repo StorageItemRepository
}

func NewStorageItemUsecase(repo StorageItemRepository) *StorageItemUsecase {
	return &StorageItemUsecase{repo}
}

func (s *StorageItemUsecase) Create(storageItem entity.StorageItem) (entity.StorageItem, error) {
	return s.repo.Create(storageItem)
}

func (s *StorageItemUsecase) UpdateField(id uint, fieldName string, value interface{}) (entity.StorageItem, error) {
	item, err := s.repo.ById(id)
	if err != nil {
		return item, err
	}
	if fieldName == "storagePlaceId" {
		parsedValue, ok := value.(float64)
		if !ok {
			return item, entity.ErrInvalidParameter
		}
		item.StoragePlaceId = int32(parsedValue)
	}
	if fieldName == "title" {
		parsedValue, ok := value.(string)
		if !ok {
			return item, entity.ErrInvalidParameter
		}
		item.Title = parsedValue
	}
	return item, s.repo.Update(item)
}

func (s *StorageItemUsecase) Consumpt(id uint, amount float64, unit entity.UnitName) (entity.StorageItem, error) {
	item, err := s.repo.ById(id)
	if err != nil {
		return item, err
	}
	err = item.Consumpt(amount, unit)
	if err != nil {
		return item, err
	}
	err = s.repo.Update(item)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *StorageItemUsecase) List() ([]entity.StorageItem, error) {
	result, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return result, nil
}
