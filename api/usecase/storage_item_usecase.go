package usecase

import (
	"github.com/martinjirku/zasobar/entity"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . StorageItemRepository

type StorageItemRepository interface {
	Create(storageItem entity.StorageItem) (entity.StorageItem, error)
	ById(storageItemId int32) (entity.StorageItem, error)
	Update(storageItem entity.StorageItem) error
	List(pagination entity.Pagination) ([]entity.StorageItem, error)
	Count(pagination entity.Pagination) (int64, error)
	CountAll() (int64, error)
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

func (s *StorageItemUsecase) UpdateField(id int32, fieldName string, value interface{}) (entity.StorageItem, error) {
	item, err := s.repo.ById(id)
	if err != nil {
		return item, err
	}
	if fieldName == "storagePlaceId" {
		parsedValue, ok := value.(float64)
		if !ok {
			return item, entity.ErrInvalidParameter
		}
		item.StoragePlaceID = int32(parsedValue)
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

func (s *StorageItemUsecase) Consumpt(id int32, amount float64, unit entity.UnitName) (entity.StorageItem, error) {
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

func (s *StorageItemUsecase) List(pagination entity.Pagination) (entity.StorageItemList, error) {
	count, err := s.repo.CountAll()
	if err != nil {
		return entity.StorageItemList{}, err
	}
	data, err := s.repo.List(pagination)
	if err != nil {
		return entity.StorageItemList{}, err
	}
	return entity.StorageItemList{Data: data, Count: count}, nil
}
